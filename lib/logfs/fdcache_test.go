package logfs

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"testing"
)

// writes a tiny volume file so acquire has something to open (matches
// fdCache.path naming: "%d.vol").
func touchVol(t *testing.T, dir string, idx int) {
	t.Helper()
	p := filepath.Join(dir, fmt.Sprintf("%d.vol", idx))
	if err := os.WriteFile(p, []byte("hello-volume"), 0644); err != nil {
		t.Fatal(err)
	}
}

// TestFdCacheConcurrentAcquireEvict: concurrent acquire/release across more
// volumes than the cap forces eviction while other fds are in use — an in-use
// fd (refs>0) must never be closed, so reads never error.
func TestFdCacheConcurrentAcquireEvict(t *testing.T) {
	dir := t.TempDir()
	const vols = 16
	for i := 0; i < vols; i++ {
		touchVol(t, dir, i)
	}
	c := newFdCache(dir, 4) // cap 4, far below vols → constant eviction

	var wg sync.WaitGroup
	for g := 0; g < 32; g++ {
		wg.Add(1)
		go func(g int) {
			defer wg.Done()
			for r := 0; r < 200; r++ {
				idx := uint64((g + r) % vols)
				f, release, err := c.acquire(idx)
				if err != nil {
					t.Errorf("acquire %d: %v", idx, err)
					return
				}
				buf := make([]byte, 5)
				if _, err := f.ReadAt(buf, 0); err != nil {
					t.Errorf("readat %d: %v", idx, err)
				}
				release()
			}
		}(g)
	}
	wg.Wait()
	c.closeAll()
}

// TestFdCacheDisabled: max<=0 falls back to transient open/close per read.
func TestFdCacheDisabled(t *testing.T) {
	dir := t.TempDir()
	touchVol(t, dir, 0)
	c := newFdCache(dir, 0)
	f, release, err := c.acquire(0)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := f.ReadAt(make([]byte, 5), 0); err != nil {
		t.Fatal(err)
	}
	release() // closes the transient fd
}
