package logfs

import (
	"bytes"
	"fmt"
	"path/filepath"
	"sync"
	"testing"

	"github.com/unibaseio/da-sdk-go/lib/kv"
)

func newTestFS(t *testing.T) *LogFS {
	t.Helper()
	dir := t.TempDir()
	ds, err := kv.NewBadgerStore(filepath.Join(dir, "kv"), &kv.DefaultOptions)
	if err != nil {
		t.Fatalf("badger: %v", err)
	}
	fs, err := New(ds, filepath.Join(dir, "data"), "0xlocal", "0xowner")
	if err != nil {
		t.Fatalf("logfs new: %v", err)
	}
	return fs
}

// TestGroupFsyncConcurrent: group-commit fsync loop must not race concurrent
// writes/rolls, and all data must remain correct.
func TestGroupFsyncConcurrent(t *testing.T) {
	t.Setenv("LOGFS_FSYNC", "group")
	t.Setenv("LOGFS_FSYNC_INTERVAL_MS", "3") // aggressive tick to overlap writes/rolls
	old := MaxSize
	MaxSize = 4096
	defer func() { MaxSize = old }()

	fs := newTestFS(t)
	defer fs.Close()
	if !fs.fsyncGroup {
		t.Fatal("group fsync should be enabled")
	}

	const n = 300
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := []byte(fmt.Sprintf("fs-%05d", i))
			if err := fs.Put(key, bytes.Repeat([]byte{byte(i)}, 200)); err != nil {
				t.Errorf("put %d: %v", i, err)
			}
		}(i)
	}
	wg.Wait()

	for i := 0; i < n; i++ {
		got, err := fs.Get([]byte(fmt.Sprintf("fs-%05d", i)))
		if err != nil {
			t.Fatalf("get %d: %v", i, err)
		}
		if !bytes.Equal(got, bytes.Repeat([]byte{byte(i)}, 200)) {
			t.Fatalf("mismatch %d", i)
		}
	}
}

// TestConcurrentPutSameOwner: many concurrent Put on one owner (the intra-owner
// path that T2 unblocked) — every object must round-trip intact, proving offset
// reservations don't overlap and writes-outside-the-lock are safe.
func TestConcurrentPutSameOwner(t *testing.T) {
	fs := newTestFS(t)
	const n = 300
	keys := make([][]byte, n)
	vals := make([][]byte, n)
	for i := 0; i < n; i++ {
		keys[i] = []byte(fmt.Sprintf("key-%05d", i))
		vals[i] = bytes.Repeat([]byte{byte(i), byte(i >> 8)}, 20+i) // distinct content+len
	}

	var wg sync.WaitGroup
	errs := make([]error, n)
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			errs[i] = fs.Put(keys[i], vals[i])
		}(i)
	}
	wg.Wait()

	for i := 0; i < n; i++ {
		if errs[i] != nil {
			t.Fatalf("put %d: %v", i, errs[i])
		}
		got, err := fs.Get(keys[i])
		if err != nil {
			t.Fatalf("get %d: %v", i, err)
		}
		if !bytes.Equal(got, vals[i]) {
			t.Fatalf("mismatch %d: got %d bytes want %d", i, len(got), len(vals[i]))
		}
	}
}

// TestConcurrentPutWithRolls: force many volume rolls (tiny MaxSize) under
// concurrent writes — exercises forward()'s in-flight drain vs the writes that
// filled the volume. All objects must survive.
func TestConcurrentPutWithRolls(t *testing.T) {
	old := MaxSize
	MaxSize = 4096
	defer func() { MaxSize = old }()

	fs := newTestFS(t)
	const n = 400
	keys := make([][]byte, n)
	vals := make([][]byte, n)
	for i := 0; i < n; i++ {
		keys[i] = []byte(fmt.Sprintf("k-%05d", i))
		vals[i] = bytes.Repeat([]byte{byte(i)}, 200) // ~200B ⇒ ~20 records/volume ⇒ many rolls
	}

	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			if err := fs.Put(keys[i], vals[i]); err != nil {
				t.Errorf("put %d: %v", i, err)
			}
		}(i)
	}
	wg.Wait()

	for i := 0; i < n; i++ {
		got, err := fs.Get(keys[i])
		if err != nil {
			t.Fatalf("get %d: %v", i, err)
		}
		if !bytes.Equal(got, vals[i]) {
			t.Fatalf("mismatch %d", i)
		}
	}
}

// TestConcurrentPutAndGet: readers hammer keys while writers add more — no
// panic/corruption; every completed write is readable.
func TestConcurrentPutAndGet(t *testing.T) {
	fs := newTestFS(t)
	const n = 200
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := []byte(fmt.Sprintf("rw-%05d", i))
			val := bytes.Repeat([]byte{byte(i)}, 50+i)
			if err := fs.Put(key, val); err != nil {
				t.Errorf("put %d: %v", i, err)
				return
			}
			got, err := fs.Get(key)
			if err != nil {
				t.Errorf("get %d: %v", i, err)
				return
			}
			if !bytes.Equal(got, val) {
				t.Errorf("mismatch %d", i)
			}
		}(i)
	}
	wg.Wait()
}

// TestConcurrentOverwriteSameKey: concurrent writes to one key are last-writer-
// wins and never corrupt — the final read must equal one of the written values.
func TestConcurrentOverwriteSameKey(t *testing.T) {
	fs := newTestFS(t)
	key := []byte("hot-key")
	const n = 100
	vals := make([][]byte, n)
	for i := 0; i < n; i++ {
		vals[i] = bytes.Repeat([]byte{byte(i)}, 64)
	}
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			if err := fs.Put(key, vals[i]); err != nil {
				t.Errorf("put %d: %v", i, err)
			}
		}(i)
	}
	wg.Wait()

	got, err := fs.Get(key)
	if err != nil {
		t.Fatalf("get: %v", err)
	}
	ok := false
	for i := 0; i < n; i++ {
		if bytes.Equal(got, vals[i]) {
			ok = true
			break
		}
	}
	if !ok {
		t.Fatalf("final value matches none of the writes (len=%d)", len(got))
	}
}
