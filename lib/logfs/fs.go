package logfs

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/alecthomas/units"
	"github.com/fxamacker/cbor/v2"
	"github.com/unibaseio/da-sdk-go/lib/log"
	"github.com/unibaseio/da-sdk-go/lib/types"
	"github.com/unibaseio/da-sdk-go/lib/utils"
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
	// mu guards the offset-reservation critical section only (curSize/curIndex/
	// curFi/full). The actual WriteAt + KV index Put happen OUTSIDE it, so writes
	// to distinct reserved offsets on one owner proceed concurrently (POSIX
	// pwrite is safe at distinct offsets). Crash-consistency: data is written
	// before its index, so a crash leaves at worst orphan bytes (reclaimable
	// space), never an index entry pointing at unwritten data.
	sync.RWMutex
	ds       types.IKVStore
	addr     string
	curSize  int64
	curIndex uint64
	curFi    *os.File
	basedir  string
	openedAt time.Time // when the current (open) volume got its first byte
	full     bool      // current volume passed MaxSize; roll on next reserve
	// wg tracks in-flight WriteAt calls on the CURRENT volume; forward() drains
	// it before closing the fd. Swapped for a fresh one on each roll.
	wg *sync.WaitGroup
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
		wg:      &sync.WaitGroup{},
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
	if sf.curSize > 0 {
		// existing partial volume: approximate its age from now (we don't persist
		// the original first-write time). The time-flush will commit it promptly.
		sf.openedAt = time.Now()
	}

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

// forward rolls to the next volume. MUST be called with sf.Lock held. It first
// drains in-flight WriteAt calls on the current volume (they hold no lock and
// finish independently), then closes the fd and opens the next one. New
// reservations are blocked meanwhile because they too need sf.Lock.
func (sf *LogFS) forward() error {
	sf.wg.Wait()          // let in-flight writes on the current volume finish
	sf.wg = &sync.WaitGroup{}

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
	sf.openedAt = time.Time{} // reset; set on next first write
	sf.full = false
	logger.Infof("logfs %s forward to: %d", sf.addr, sf.curIndex)
	return nil
}

// Pending reports the unflushed bytes in the current (open) volume and how long
// ago its first byte was written. Used by the hub to time-flush small writes:
// commit a partial volume once it ages past a threshold (size trigger is MaxSize).
func (sf *LogFS) Pending() (int64, time.Duration) {
	sf.Lock()
	defer sf.Unlock()
	if sf.curSize == 0 {
		return 0, 0
	}
	return sf.curSize, time.Since(sf.openedAt)
}

// Roll force-closes the current (open) volume into a completed one so it becomes
// uploadable, even below MaxSize. No-op if the volume is empty.
func (sf *LogFS) Roll() error {
	sf.Lock()
	defer sf.Unlock()
	if sf.curSize == 0 {
		return nil
	}
	return sf.forward()
}

func (sf *LogFS) Put(key, val []byte) error {
	sum := sha256.Sum256(val)
	if len(key) == 0 {
		key = sum[:]
	}
	dskey := types.NewKey(types.DsLogFS, sf.addr, key)

	if has, herr := sf.ds.Has(dskey); herr == nil && has {
		logger.Infof("%s overwrite key: %s", sf.addr, string(key))
	}

	// padded on-disk length (records are 31-byte aligned)
	n := len(val)
	padded := int64(n)
	if n%31 != 0 {
		padded += int64(31 - n%31)
	}

	// --- reserve an offset (short critical section) ---
	sf.Lock()
	// roll a full volume before reserving on it (deferred from the write that
	// filled it, so forward's wg.Wait never races that same write).
	if sf.full && sf.curSize > 0 {
		if err := sf.forward(); err != nil {
			sf.Unlock()
			return err
		}
	}
	if sf.curSize == 0 {
		sf.openedAt = time.Now() // first byte of a fresh volume
	}
	start := sf.curSize
	volIndex := sf.curIndex
	fi := sf.curFi
	wg := sf.wg
	sf.curSize += padded
	if int(sf.curSize) > MaxSize {
		sf.full = true // roll on the next reserve
	}
	wg.Add(1)
	sf.Unlock()

	// --- write payload + padding + index OUTSIDE the lock ---
	defer wg.Done()

	if _, err := fi.WriteAt(val, start); err != nil {
		return err
	}
	if pad := padded - int64(n); pad > 0 {
		if _, err := fi.WriteAt(make([]byte, pad), start+int64(n)); err != nil {
			return err
		}
	}
	logger.Debugf("logfs write at: %s %d %d %d", sf.addr, volIndex, start, n)

	lm := LogMeta{
		Index: volIndex,
		Start: uint64(start),
		Size:  uint64(n),
		Hash:  sum[:],
		Name:  key,
	}
	lmv, err := lm.Serialize()
	if err != nil {
		return err
	}
	// data-before-index: only now is the object safely indexed.
	return sf.ds.Put(dskey, lmv)
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
	sf.wg.Wait() // let in-flight writes finish before closing the fd
	return sf.curFi.Close()
}
