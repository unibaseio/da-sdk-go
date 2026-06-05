package logfs

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/unibaseio/da-sdk-go/lib/log"
	"github.com/unibaseio/da-sdk-go/lib/types"
	"github.com/unibaseio/da-sdk-go/lib/utils"
	"github.com/alecthomas/units"
	"github.com/fxamacker/cbor/v2"
)

var logger = log.Logger("logfs")

var _ types.IFileStore = (*LogFS)(nil)

var MaxSize = 31 * 1024 * 1024

func init() {
	hs := os.Getenv("MAX_SIZE")
	bb, err := units.ParseBase2Bytes(hs)
	if err == nil && int(bb) > 0 {
		MaxSize = int(bb)
	}
}

type LogMeta struct {
	Index uint64 // which volume
	Start uint64
	Size  uint64
	Hash  []byte // 32byte
	Name  []byte
}

func (lm *LogMeta) Serialize() ([]byte, error) {
	return cbor.Marshal(lm)
}

func (lm *LogMeta) Deserialize(b []byte) error {
	return cbor.Unmarshal(b, lm)
}

type LogFS struct {
	sync.RWMutex
	ds       types.IKVStore
	addr     string
	curSize  int64
	curIndex uint64
	curFi    *os.File
	basedir  string
}

// todo: each one has its own maxsize
func New(ds types.IKVStore, dir string, local, addr string) (*LogFS, error) {
	//log.SetLogLevel("debug")
	dir = filepath.Join(dir, addr)
	logger.Infof("logfs start at: %s with maxsize: %d", dir, MaxSize)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return nil, err
	}

	sf := &LogFS{
		basedir: dir,
		ds:      ds,
		addr:    addr,
	}

	dsKey := types.NewKey(types.DsLogFS, addr)
	val, err := sf.ds.Get(dsKey)
	if err == nil && len(val) == 8 {
		sf.curIndex = binary.BigEndian.Uint64(val)
	} else {
		sf.curIndex = GetIndex(local, addr)
	}

	curlog := filepath.Join(sf.basedir, fmt.Sprintf("%d.vol", sf.curIndex))
	sf.curFi, err = os.OpenFile(curlog, os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		return nil, err
	}
	fi, err := sf.curFi.Stat()
	if err != nil {
		return nil, err
	}
	sf.curSize = fi.Size()

	logger.Infof("logfs started at: %s %d %d", dir, sf.curIndex, sf.curSize)
	return sf, nil
}

func GetIndex(local, addr string) uint64 {
	h := sha256.New()
	h.Write([]byte(local))
	h.Write([]byte(addr))
	res := h.Sum(nil)
	st := binary.BigEndian.Uint32(res[:4])
	return uint64(st)
}

func (sf *LogFS) forward() error {
	err := sf.curFi.Close()
	if err != nil {
		return err
	}
	sf.curIndex++

	dsKey := types.NewKey(types.DsLogFS, sf.addr)
	val := make([]byte, 8)
	binary.BigEndian.PutUint64(val, sf.curIndex)
	err = sf.ds.Put(dsKey, val)
	if err != nil {
		return err
	}

	curlog := filepath.Join(sf.basedir, fmt.Sprintf("%d.vol", sf.curIndex))
	fi, err := os.OpenFile(curlog, os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		return err
	}
	sf.curFi = fi
	sf.curSize = 0
	logger.Infof("logfs %s forward to: %d", sf.addr, sf.curIndex)
	return nil
}

func (sf *LogFS) Put(key, val []byte) error {
	sum := sha256.Sum256(val)
	if len(key) == 0 {
		key = sum[:]
	}
	dskey := types.NewKey(types.DsLogFS, sf.addr, key)

	sf.Lock()
	defer sf.Unlock()

	has, err := sf.ds.Has(dskey)
	if err == nil && has {
		logger.Infof("%s overwrite key: %s", sf.addr, string(key))
	}

	n, err := sf.curFi.WriteAt(val, sf.curSize)
	if err != nil {
		return err
	}

	logger.Debugf("logfs write at: %s %d %d %d %d", sf.addr, sf.curIndex, sf.curSize, n, len(val))

	if n%31 != 0 {
		data := make([]byte, 31-n%31)
		pn, err := sf.curFi.WriteAt(data, sf.curSize+int64(n))
		if err != nil {
			return err
		}
		n += pn
		logger.Debugf("logfs write padding: %d", pn)
	}

	lm := LogMeta{
		Index: sf.curIndex,
		Start: uint64(sf.curSize),
		Size:  uint64(len(val)),
		Hash:  sum[:],
		Name:  key,
	}

	sf.curSize += int64(n)

	lmv, err := lm.Serialize()
	if err != nil {
		return err
	}
	err = sf.ds.Put(dskey, lmv)
	if err != nil {
		return err
	}

	if int(sf.curSize) > MaxSize {
		err := sf.forward()
		if err != nil {
			return err
		}
	}

	return nil
}

func (sf *LogFS) Get(key []byte, opts ...int) ([]byte, error) {
	lm, err := sf.GetMeta(key)
	if err != nil {
		return nil, err
	}

	return sf.GetData(lm, opts...)
}

func (sf *LogFS) GetMeta(key []byte) (*LogMeta, error) {
	dskey := types.NewKey(types.DsLogFS, sf.addr, key)
	val, err := sf.ds.Get(dskey)
	if err != nil {
		return nil, err
	}

	lm := new(LogMeta)
	err = lm.Deserialize(val)
	if err != nil {
		return nil, err
	}
	return lm, nil
}

func (sf *LogFS) GetData(lm *LogMeta, opts ...int) ([]byte, error) {
	logger.Infof("logfs read at: %s %d %d %d", sf.addr, lm.Index, lm.Start, lm.Size)
	curlog := filepath.Join(sf.basedir, fmt.Sprintf("%d.vol", lm.Index))
	fi, err := os.OpenFile(curlog, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer fi.Close()

	res := make([]byte, lm.Size)
	n, err := fi.ReadAt(res, int64(lm.Start))
	if err != nil {
		return nil, err
	}
	if n != int(lm.Size) {
		return nil, fmt.Errorf("unequal size")
	}

	sum := sha256.Sum256(res)
	if !bytes.Equal(sum[:], lm.Hash) {
		return nil, fmt.Errorf("unequal content")
	}

	return res, nil
}

func (sf *LogFS) Has(key []byte) (bool, error) {
	return false, nil
}

func (sf *LogFS) Delete(key []byte) error {
	return nil
}

func (sf *LogFS) Size() types.DiskStats {
	ds, _ := utils.GetDiskStatus(sf.basedir)
	return ds
}

func (sf *LogFS) Close() error {
	sf.Lock()
	defer sf.Unlock()
	return sf.curFi.Close()
}
