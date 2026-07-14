package hub

import (
	"container/list"
	"sync"
	"sync/atomic"

	"github.com/unibaseio/da-sdk-go/lib/env"
)

// readCache is a bounded, byte-sized LRU of small hot objects, keyed by
// owner/name. It turns the hub's negative-only missCache into a full
// read-through cache for the common agent-memory pattern (read-heavy, small,
// high fan-out). Object content is content-addressed and immutable, so cache
// entries never go stale — they're evicted only by capacity, and proactively on
// an overwrite of the same key (logFSWrite calls del). Large objects
// (model/dataset passthrough) are never cached (maxItem guard) to avoid
// polluting the small-object cache. Disabled (nil) when HUB_READCACHE_MB=0.
type rcEntry struct {
	key string
	val []byte
}

type readCache struct {
	mu       sync.Mutex
	ll       *list.List
	m        map[string]*list.Element
	curBytes int64
	maxBytes int64
	maxItem  int64

	hits   atomic.Int64
	misses atomic.Int64
}

// Stats reports cache effectiveness (hits, misses) for observability.
func (c *readCache) Stats() (hits, misses int64) {
	if c == nil {
		return 0, 0
	}
	return c.hits.Load(), c.misses.Load()
}

const (
	defaultReadCacheMB   int64 = 256
	readCacheMaxItemByte int64 = 4 << 20 // don't cache objects larger than 4MiB
)

func newReadCache() *readCache {
	mb := env.Int64("HUB_READCACHE_MB", defaultReadCacheMB)
	if mb <= 0 {
		return nil // disabled
	}
	return &readCache{
		ll:       list.New(),
		m:        make(map[string]*list.Element),
		maxBytes: mb << 20,
		maxItem:  readCacheMaxItemByte,
	}
}

// get returns a cached copy-safe reference for (owner,name). The returned slice
// must not be mutated (callers only write it to a Writer). ok=false on miss.
func (c *readCache) get(owner, name string) ([]byte, bool) {
	if c == nil {
		return nil, false
	}
	k := missKey(owner, name)
	c.mu.Lock()
	defer c.mu.Unlock()
	if el, ok := c.m[k]; ok {
		c.ll.MoveToFront(el)
		c.hits.Add(1)
		return el.Value.(*rcEntry).val, true
	}
	c.misses.Add(1)
	return nil, false
}

// put caches (owner,name)->val if it fits the per-item cap. val must be
// immutable after this call (the cache retains the reference).
func (c *readCache) put(owner, name string, val []byte) {
	if c == nil || int64(len(val)) > c.maxItem {
		return
	}
	k := missKey(owner, name)
	c.mu.Lock()
	defer c.mu.Unlock()
	if el, ok := c.m[k]; ok {
		// refresh existing
		old := el.Value.(*rcEntry)
		c.curBytes += int64(len(val)) - int64(len(old.val))
		old.val = val
		c.ll.MoveToFront(el)
	} else {
		el := c.ll.PushFront(&rcEntry{key: k, val: val})
		c.m[k] = el
		c.curBytes += int64(len(val))
	}
	// evict LRU tail until within budget
	for c.curBytes > c.maxBytes {
		tail := c.ll.Back()
		if tail == nil {
			break
		}
		e := tail.Value.(*rcEntry)
		c.ll.Remove(tail)
		delete(c.m, e.key)
		c.curBytes -= int64(len(e.val))
	}
}

// del drops any cached entry for (owner,name) — call after an overwrite so a
// freshly written value is never shadowed by a stale cache hit.
func (c *readCache) del(owner, name string) {
	if c == nil {
		return
	}
	k := missKey(owner, name)
	c.mu.Lock()
	defer c.mu.Unlock()
	if el, ok := c.m[k]; ok {
		e := el.Value.(*rcEntry)
		c.ll.Remove(el)
		delete(c.m, k)
		c.curBytes -= int64(len(e.val))
	}
}
