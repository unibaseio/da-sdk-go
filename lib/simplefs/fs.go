package simplefs

import (
	"encoding/binary"
	"errors"
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
			os.Remove(fi.Name())
			continue
		}

		fii, err := fi.Info()
		if err != nil {
			continue
		}

		sf.size += fii.Size()
	}
}

func (sf *SimpleFs) getPath(key []byte) string {
	pbase := path.Base(string(key))
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

	os.MkdirAll(dir, 0755)
	return path.Join(dir, pbase)
}

func (sf *SimpleFs) Put(key, val []byte) error {
	fn := sf.getPath(key)
	info, err := os.Stat(fn)
	if err == nil && !info.IsDir() {
		err := os.Remove(fn)
		if err != nil {
			return err
		}

		sf.Lock()
		sf.size -= info.Size()
		sf.Unlock()
	}

	// write then rename
	tmpfn := fn + ".tmp"
	err = os.WriteFile(tmpfn, val, 0644)
	if err != nil {
		return err
	}
	err = os.Rename(tmpfn, fn)
	if err != nil {
		return err
	}

	sf.Lock()
	sf.size += int64(len(val))
	sf.writeSize()
	sf.Unlock()

	return nil
}

func (sf *SimpleFs) Get(key []byte, opts ...int) ([]byte, error) {
	fn := sf.getPath(key)

	if len(opts) == 2 {
		start := opts[0]
		length := opts[1]
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
	fn := sf.getPath(key)
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
	fn := sf.getPath(key)
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
