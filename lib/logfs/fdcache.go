package logfs

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

// fdCache is a bounded, reference-counted cache of read-only volume fds so
// GetData avoids an open()/close() syscall pair per read. Volumes are append-
// only and never deleted, so a cached fd stays valid indefinitely; entries are
// closed only on eviction (when unreferenced) or closeAll. os.File.ReadAt uses
// pread and is concurrency-safe, so many readers share one fd. An in-use fd
// (refs>0) is never evicted, so a read never races a Close. Disabled when max<=0
// (falls back to a transient open/close per read).
type fdEntry struct {
	f    *os.File
	refs int
}

type fdCache struct {
	mu      sync.Mutex
	m       map[uint64]*fdEntry
	max     int
	basedir string
}

func newFdCache(basedir string, max int) *fdCache {
	return &fdCache{m: make(map[uint64]*fdEntry), max: max, basedir: basedir}
}

func (c *fdCache) path(idx uint64) string {
	return filepath.Join(c.basedir, fmt.Sprintf("%d.vol", idx))
}

// acquire returns a shared read-only fd for volume idx plus a release func that
// MUST be called (typically deferred); the fd must not be used after release.
func (c *fdCache) acquire(idx uint64) (*os.File, func(), error) {
	if c == nil || c.max <= 0 {
		f, err := os.OpenFile(c.path(idx), os.O_RDONLY, os.ModePerm)
		if err != nil {
			return nil, nil, err
		}
		return f, func() { f.Close() }, nil
	}

	c.mu.Lock()
	if e, ok := c.m[idx]; ok {
		e.refs++
		c.mu.Unlock()
		return e.f, func() { c.release(idx) }, nil
	}
	// evict one unreferenced entry if at capacity (close outside the lock)
	var toClose *os.File
	if len(c.m) >= c.max {
		for k, e := range c.m {
			if e.refs == 0 {
				toClose = e.f
				delete(c.m, k)
				break
			}
		}
	}
	c.mu.Unlock()
	if toClose != nil {
		toClose.Close()
	}

	f, err := os.OpenFile(c.path(idx), os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil, nil, err
	}

	c.mu.Lock()
	if e, ok := c.m[idx]; ok { // lost a race: reuse the existing, drop ours
		e.refs++
		c.mu.Unlock()
		f.Close()
		return e.f, func() { c.release(idx) }, nil
	}
	if len(c.m) >= c.max { // still saturated with in-use fds: hand out a transient
		c.mu.Unlock()
		return f, func() { f.Close() }, nil
	}
	c.m[idx] = &fdEntry{f: f, refs: 1}
	c.mu.Unlock()
	return f, func() { c.release(idx) }, nil
}

func (c *fdCache) release(idx uint64) {
	c.mu.Lock()
	if e, ok := c.m[idx]; ok && e.refs > 0 {
		e.refs--
	}
	c.mu.Unlock()
}

func (c *fdCache) closeAll() {
	if c == nil {
		return
	}
	c.mu.Lock()
	for k, e := range c.m {
		e.f.Close()
		delete(c.m, k)
	}
	c.mu.Unlock()
}
