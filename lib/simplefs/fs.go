package simplefs

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"

	"github.com/unibaseio/da-sdk-go/lib/log"
	"github.com/unibaseio/da-sdk-go/lib/types"
	"github.com/unibaseio/da-sdk-go/lib/utils"
)

var logger = log.Logger("simplefs")

var _ types.IFileStore = (*SimpleFs)(nil)

const usage = "usage"

type SimpleFs struct {
	sync.RWMutex
	size    int64
	basedir string
}

func New(dir string) (*SimpleFs, error) {
	logger.Infof("simplefs start at: %s", dir)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return nil, err
	}

	sf := &SimpleFs{
		basedir: dir,
	}

	ub, err := os.ReadFile(filepath.Join(dir, usage))
	if err == nil && len(ub) >= 8 {
		sf.size = int64(binary.BigEndian.Uint64(ub))
	} else {
		logger.Infof("calculate usage by walking simplefs: %s", dir)
		sf.walk(dir)
		sf.writeSize()
	}

	logger.Infof("simplefs started at: %s %d", dir, sf.size)
	return sf, nil
}

func (sf *SimpleFs) walk(baseDir string) {
	rd, err := os.ReadDir(baseDir)
	if err != nil {
		return
	}

	for _, fi := range rd {
		if fi.IsDir() {
			sf.walk(path.Join(baseDir, fi.Name()))
			continue
		}

		if strings.Contains(fi.Name(), usage) {
			continue
		}

		if strings.HasSuffix(fi.Name(), ".tmp") {
			// a write interrupted before its rename; the name is relative to
			// this directory, not to the process's working directory
			os.Remove(path.Join(baseDir, fi.Name()))
			continue
		}

		fii, err := fi.Info()
		if err != nil {
			continue
		}

		sf.size += fii.Size()
	}
}

var errBadKey = errors.New("simplefs: key has no usable file name")

// getPath maps a key to its file: the key's last path element, sharded by its
// last 4 characters. A key whose file would land outside the base directory
// is refused: besides "." and "..", the shard directories come from the
// name's own characters, so a name ending in ".." would climb out of it.
func (sf *SimpleFs) getPath(key []byte) (string, error) {
	pbase := path.Base(string(key))
	if pbase == "." || pbase == ".." || pbase == "/" || strings.ContainsRune(pbase, 0) {
		return "", errBadKey
	}
	dir := "mo/ck"
	plen := len(pbase)
	if plen >= 2 {
		if plen >= 4 {
			dir = pbase[plen-2:] + "/" + pbase[plen-4:plen-2] // last 4
		} else {
			dir = pbase[plen-2:] + "/nt"
		}
	}
	dir = path.Join(sf.basedir, dir)
	fn := path.Join(dir, pbase)
	if rel, err := filepath.Rel(sf.basedir, fn); err != nil || rel == ".." || strings.HasPrefix(rel, "../") {
		return "", errBadKey
	}

	os.MkdirAll(dir, 0755)
	return fn, nil
}

func (sf *SimpleFs) Put(key, val []byte) error {
	fn, err := sf.getPath(key)
	if err != nil {
		return err
	}
	var old int64
	if info, err := os.Stat(fn); err == nil && !info.IsDir() {
		old = info.Size()
	}

	// write then rename: the rename replaces an existing file atomically, so
	// a crash never leaves the key without data
	tmpfn := fn + ".tmp"
	err = os.WriteFile(tmpfn, val, 0644)
	if err != nil {
		return err
	}
	err = os.Rename(tmpfn, fn)
	if err != nil {
		os.Remove(tmpfn)
		return err
	}

	sf.Lock()
	sf.size -= old
	sf.Unlock()

	sf.Lock()
	sf.size += int64(len(val))
	sf.writeSize()
	sf.Unlock()

	return nil
}

func (sf *SimpleFs) Get(key []byte, opts ...int) ([]byte, error) {
	fn, err := sf.getPath(key)
	if err != nil {
		return nil, err
	}

	if len(opts) == 2 {
		start := opts[0]
		length := opts[1]
		if start < 0 || length < 0 {
			return nil, fmt.Errorf("simplefs: bad range %d+%d", start, length)
		}
		osf, err := os.OpenFile(fn, os.O_RDONLY, os.ModePerm)
		if err != nil {
			return nil, err
		}
		defer osf.Close()

		res := make([]byte, length)
		n, err := osf.ReadAt(res, int64(start))
		if err != nil && err != io.EOF {

			return nil, err
		}
		return res[:n], nil
	}

	data, err := os.ReadFile(fn)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (sf *SimpleFs) Has(key []byte) (bool, error) {
	fn, err := sf.getPath(key)
	if err != nil {
		return false, err
	}
	info, err := os.Stat(fn)
	if err != nil || os.IsNotExist(err) {
		return false, err
	}

	if info.IsDir() {
		return false, errors.New("is a dir")
	}

	return true, nil
}

func (sf *SimpleFs) Delete(key []byte) error {
	fn, err := sf.getPath(key)
	if err != nil {
		return err
	}
	fi, err := os.Stat(fn)
	if err != nil {
		return err
	}

	err = os.Remove(fn)
	if err != nil {
		return err
	}

	if !fi.IsDir() {
		sf.Lock()
		sf.size -= fi.Size()
		sf.writeSize()
		sf.Unlock()
	}

	return nil
}

func (sf *SimpleFs) Size() types.DiskStats {
	ds, _ := utils.GetDiskStatus(sf.basedir)
	if sf.size >= 0 {
		ds.Used = uint64(sf.size)
	}

	return ds
}

func (sf *SimpleFs) writeSize() error {
	ub := make([]byte, 8)
	binary.BigEndian.PutUint64(ub, uint64(sf.size))

	return os.WriteFile(filepath.Join(sf.basedir, usage), ub, 0644)
}

func (sf *SimpleFs) Close() error {
	sf.Lock()
	defer sf.Unlock()
	return sf.writeSize()
}
