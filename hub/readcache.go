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

	// l2 is the optional shared tier (P4; Redis in prod). nil = L1-only. On an L1
	// miss we consult L2 and, on a hit, promote the bytes back into L1.
	l2 l2cache

	hits   atomic.Int64 // L1 hit
	l2hits atomic.Int64 // L1 miss, L2 hit
	misses atomic.Int64 // both miss
}

// Stats reports L1 cache effectiveness (hits, misses) for observability.
func (c *readCache) Stats() (hits, misses int64) {
	if c == nil {
		return 0, 0
	}
	return c.hits.Load(), c.misses.Load()
}

// L2Stats reports the shared-tier hit count and whether L2 is configured.
func (c *readCache) L2Stats() (hits int64, enabled bool) {
	if c == nil || c.l2 == nil {
		return 0, false
	}
	return c.l2hits.Load(), true
}

const (
	defaultReadCacheMB   int64 = 256
	readCacheMaxItemByte int64 = 4 << 20 // don't cache objects larger than 4MiB
)

func newReadCache() *readCache {
	mb := env.Int64("HUB_READCACHE_MB", defaultReadCacheMB)
	if mb <= 0 {
		return nil // both tiers disabled
	}
	c := &readCache{
		ll:       list.New(),
		m:        make(map[string]*list.Element),
		maxBytes: mb << 20,
		maxItem:  readCacheMaxItemByte,
	}
	// Attach the shared L2 only when configured; assigning a typed-nil *redisL2 to
	// the interface field would make c.l2 != nil wrongly true.
	if l2 := newRedisL2(); l2 != nil {
		c.l2 = l2
	}
	return c
}

// get returns a cached copy-safe reference for (owner,name). The returned slice
// must not be mutated (callers only write it to a Writer). ok=false on miss.
// L1 (in-proc) is checked first; on a miss the shared L2 (Redis) is consulted
// and, on a hit, promoted back into L1. L2 network I/O never holds the L1 lock.
func (c *readCache) get(owner, name string) ([]byte, bool) {
	if c == nil {
		return nil, false
	}
	k := missKey(owner, name)

	c.mu.Lock()
	if el, ok := c.m[k]; ok {
		c.ll.MoveToFront(el)
		val := el.Value.(*rcEntry).val
		c.mu.Unlock()
		c.hits.Add(1)
		return val, true
	}
	c.mu.Unlock()

	if c.l2 != nil {
		if v, ok := c.l2.get(k); ok {
			c.l2hits.Add(1)
			c.putL1(k, v) // promote to L1; don't re-write L2
			return v, true
		}
	}
	c.misses.Add(1)
	return nil, false
}

// put caches (owner,name)->val in both tiers if it fits the per-item cap. val
// must be immutable after this call (the cache retains the reference).
func (c *readCache) put(owner, name string, val []byte) {
	if c == nil || int64(len(val)) > c.maxItem {
		return
	}
	k := missKey(owner, name)
	if c.l2 != nil {
		c.l2.set(k, val) // outside the L1 lock (network)
	}
	c.putL1(k, val)
}

// putL1 inserts/refreshes an L1 entry and evicts the LRU tail to budget.
func (c *readCache) putL1(k string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if el, ok := c.m[k]; ok {
		old := el.Value.(*rcEntry)
		c.curBytes += int64(len(val)) - int64(len(old.val))
		old.val = val
		c.ll.MoveToFront(el)
	} else {
		el := c.ll.PushFront(&rcEntry{key: k, val: val})
		c.m[k] = el
		c.curBytes += int64(len(val))
	}
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

// del drops any cached entry for (owner,name) in both tiers — call after an
// overwrite so a freshly written value is never shadowed by a stale cache hit.
func (c *readCache) del(owner, name string) {
	if c == nil {
		return
	}
	k := missKey(owner, name)
	if c.l2 != nil {
		c.l2.del(k) // outside the L1 lock (network)
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	if el, ok := c.m[k]; ok {
		e := el.Value.(*rcEntry)
		c.ll.Remove(el)
		delete(c.m, k)
		c.curBytes -= int64(len(e.val))
	}
}
