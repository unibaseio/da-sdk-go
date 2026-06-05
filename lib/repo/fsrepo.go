package repo

import (
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"sync"

	"github.com/unibaseio/da-sdk-go/lib/config"
	"github.com/unibaseio/da-sdk-go/lib/key"
	"github.com/unibaseio/da-sdk-go/lib/kv"
	"github.com/unibaseio/da-sdk-go/lib/simplefs"
	"github.com/unibaseio/da-sdk-go/lib/types"
	lockfile "github.com/ipfs/go-fs-lock"
	"github.com/mitchellh/go-homedir"
)

const (
	configFilename     = "config.json"
	tempConfigFilename = ".config.json.temp"
	lockFile           = "repo.lock"

	keyStorePath = "keystore" // $Path/keystore
	metaPath     = "meta"     // $Path/meta
	dataPath     = "data"     // $Path/data
)

type FSRepo struct {
	path string

	lk sync.RWMutex

	cfg *config.Config
	key *key.Key

	metaDs types.IKVStore
	fileDs types.IFileStore

	lockfile io.Closer
}

var _ Repo = (*FSRepo)(nil)

func NewFSRepo(dir string, cfg *config.Config) (*FSRepo, error) {
	if dir == "" {
		dir = "./"
	}

	rp, err := homedir.Expand(dir)
	if err != nil {
		return nil, err
	}

	err = ensureWritableDirectory(rp)
	if err != nil {
		return nil, fmt.Errorf("no writable directory %w", err)
	}

	info, err := os.Stat(rp)
	if err != nil {
		return nil, fmt.Errorf("failed to stat repo %s %s", rp, err)
	}

	var actualPath string
	if info.IsDir() {
		actualPath = rp
	} else {
		actualPath, err = os.Readlink(rp)
		if err != nil {
			return nil, fmt.Errorf("failed to follow repo symlink %s %w", rp, err)
		}
	}

	err = initFSRepo(actualPath, cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to init repo %s %w", rp, err)
	}

	r := &FSRepo{path: actualPath}

	r.lockfile, err = lockfile.Lock(r.path, lockFile)
	if err != nil {
		return nil, fmt.Errorf("failed to take repo lock %w", err)
	}

	err = r.loadFromDisk()
	if err != nil {
		_ = r.lockfile.Close()
		return nil, err
	}

	return r, nil
}

func initFSRepo(dir string, cfg *config.Config) error {
	configFile := filepath.Join(dir, configFilename)
	if cfg != nil {
		err := cfg.WriteFile(configFile)
		if err != nil {
			return err
		}
	}

	return nil
}

func Exists(repoPath string) bool {
	repoPath, _ = homedir.Expand(repoPath)

	configFile := filepath.Join(repoPath, configFilename)
	_, err := os.Stat(configFile)
	if err != nil {
		notExist := os.IsNotExist(err)
		return !notExist
	}

	cfg, err := config.ReadFile(configFile)
	if err != nil {
		return false
	}
	if cfg.Wallet.Address != "" {
		return true
	}

	return false
}

func (r *FSRepo) loadFromDisk() error {
	err := r.loadConfig()
	if err != nil {
		return fmt.Errorf("failed to load config %w", err)
	}

	err = r.openKey()
	if err != nil {
		return fmt.Errorf("failed to open keystore %w", err)
	}

	err = r.openMetaStore()
	if err != nil {
		return fmt.Errorf("failed to open metastore %w", err)
	}

	err = r.openFileStore()
	if err != nil {
		return fmt.Errorf("failed to open datastore %w", err)
	}

	return nil
}

// ReplaceConfig replaces the current config with the newly passed in one.
func (r *FSRepo) ReplaceConfig(cfg *config.Config) error {
	r.lk.Lock()
	defer r.lk.Unlock()

	r.cfg = cfg
	tmp := filepath.Join(r.path, tempConfigFilename)
	err := os.RemoveAll(tmp)
	if err != nil {
		return err
	}
	err = r.cfg.WriteFile(tmp)
	if err != nil {
		return err
	}
	return os.Rename(tmp, filepath.Join(r.path, configFilename))
}

func (r *FSRepo) Config() *config.Config {
	r.lk.RLock()
	defer r.lk.RUnlock()

	return r.cfg
}

func (r *FSRepo) Key() *key.Key {
	return r.key
}

func (r *FSRepo) MetaStore() types.IKVStore {
	return r.metaDs
}

func (r *FSRepo) DataStore() types.IFileStore {
	return r.fileDs
}

// Close closes the repo.
func (r *FSRepo) Close() error {
	err := r.fileDs.Close()
	if err != nil {
		return fmt.Errorf("failed to close file store %w", err)
	}

	err = r.metaDs.Close()
	if err != nil {
		return fmt.Errorf("failed to close meta store %w", err)
	}

	return r.lockfile.Close()
}

func (r *FSRepo) loadConfig() error {
	configFile := filepath.Join(r.path, configFilename)
	cfg, err := config.ReadFile(configFile)
	if err != nil {
		return fmt.Errorf("failed to read config file at %q %w", configFile, err)
	}

	r.cfg = cfg
	return nil
}
func (r *FSRepo) openKey() error {
	ksp := filepath.Join(r.path, keyStorePath)
	r.key = key.New(ksp)

	return nil
}

func (r *FSRepo) openMetaStore() error {
	mpath := path.Join(r.path, metaPath)
	opt := kv.DefaultOptions

	ds, err := kv.NewBadgerStore(mpath, &opt)
	if err != nil {
		return err
	}

	r.metaDs = ds

	return nil
}

func (r *FSRepo) openFileStore() error {

	dpath := path.Join(r.path, dataPath)

	ds, err := simplefs.New(dpath)
	if err != nil {
		return err
	}

	r.fileDs = ds

	return nil
}

func ensureWritableDirectory(path string) error {
	err := os.Mkdir(path, 0775)
	if err == nil {
		return nil
	} else if !os.IsExist(err) {
		return fmt.Errorf("failed to create directory %s %w", path, err)
	}

	stat, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("failed to stat path %s %w", path, err)
	}
	if !stat.IsDir() {
		return fmt.Errorf("%s is not a directory", path)
	}
	if (stat.Mode() & 0600) != 0600 {
		return fmt.Errorf("insufficient permissions for path %s, got %04o need %04o", path, stat.Mode(), 0600)
	}
	return nil
}

func (r *FSRepo) Path() string {
	return r.path
}

func (r *FSRepo) Repo() Repo {
	return r
}
