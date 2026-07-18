package hub

import (
	"bytes"
	"container/list"
	"sync"
	"testing"
)

// fakeL2 is an in-memory l2cache for testing the two-tier logic without Redis.
type fakeL2 struct {
	mu               sync.Mutex
	m                map[string][]byte
	gets, sets, dels int
}

func newFakeL2() *fakeL2 { return &fakeL2{m: map[string][]byte{}} }

func (f *fakeL2) get(k string) ([]byte, bool) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.gets++
	v, ok := f.m[k]
	return v, ok
}
func (f *fakeL2) set(k string, v []byte) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.sets++
	f.m[k] = append([]byte(nil), v...)
}
func (f *fakeL2) del(k string) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.dels++
	delete(f.m, k)
}
func (f *fakeL2) ping() error { return nil }
func (f *fakeL2) close()      {}

func newTestReadCache(maxBytes, maxItem int64) *readCache {
	return &readCache{
		ll:       list.New(),
		m:        make(map[string]*list.Element),
		maxBytes: maxBytes,
		maxItem:  maxItem,
	}
}

func TestReadCacheHitPutDel(t *testing.T) {
	c := newTestReadCache(250, 4<<20)
	blob := func(b byte) []byte { return bytes.Repeat([]byte{b}, 100) }

	c.put("o", "a", blob(1))
	c.put("o", "b", blob(2))
	if v, ok := c.get("o", "a"); !ok || !bytes.Equal(v, blob(1)) {
		t.Fatal("a should hit")
	}

	// third item (total 300 > 250) evicts the LRU tail. get(a) above moved a to
	// front, so b is the tail and gets evicted.
	c.put("o", "c", blob(3))
	if _, ok := c.get("o", "b"); ok {
		t.Fatal("b should have been evicted")
	}
	if _, ok := c.get("o", "a"); !ok {
		t.Fatal("a should survive")
	}
	if _, ok := c.get("o", "c"); !ok {
		t.Fatal("c should be present")
	}

	// del drops the entry
	c.del("o", "a")
	if _, ok := c.get("o", "a"); ok {
		t.Fatal("a should be deleted")
	}
}

func TestReadCacheMaxItemGuard(t *testing.T) {
	c := newTestReadCache(1<<20, 1024)
	c.put("o", "big", make([]byte, 2048)) // > maxItem
	if _, ok := c.get("o", "big"); ok {
		t.Fatal("oversized object must not be cached")
	}
}

func TestReadCacheNilSafe(t *testing.T) {
	var c *readCache
	c.put("o", "a", []byte{1})
	if _, ok := c.get("o", "a"); ok {
		t.Fatal("nil cache must miss")
	}
	c.del("o", "a") // must not panic
}

// TestReadCacheL2Promote: an L1 miss that hits L2 returns the value, promotes it
// into L1 (next get is an L1 hit), and counts an l2 hit.
func TestReadCacheL2Promote(t *testing.T) {
	c := newTestReadCache(1<<20, 4<<20)
	f := newFakeL2()
	c.l2 = f
	// seed only L2 (as if another replica cached it)
	f.set(missKey("o", "x"), bytes.Repeat([]byte{9}, 50))

	v, ok := c.get("o", "x")
	if !ok || len(v) != 50 {
		t.Fatalf("expected L2 hit, ok=%v len=%d", ok, len(v))
	}
	if h, on := c.L2Stats(); !on || h != 1 {
		t.Fatalf("L2Stats = (%d,%v), want (1,true)", h, on)
	}
	// promoted: now an L1 hit, no further L2 get
	getsBefore := f.gets
	if _, ok := c.get("o", "x"); !ok {
		t.Fatal("promoted entry should hit L1")
	}
	if f.gets != getsBefore {
		t.Fatal("L1 hit must not consult L2")
	}
	if c.hits.Load() != 1 {
		t.Fatalf("L1 hits = %d, want 1", c.hits.Load())
	}
}

// TestReadCacheL2WriteThroughAndDel: put writes both tiers; del clears both.
func TestReadCacheL2WriteThroughAndDel(t *testing.T) {
	c := newTestReadCache(1<<20, 4<<20)
	f := newFakeL2()
	c.l2 = f

	c.put("o", "k", bytes.Repeat([]byte{1}, 64))
	if _, ok := f.get(missKey("o", "k")); !ok {
		t.Fatal("put must write through to L2")
	}
	c.del("o", "k")
	if _, ok := f.get(missKey("o", "k")); ok {
		t.Fatal("del must clear L2")
	}
	if _, ok := c.get("o", "k"); ok {
		t.Fatal("del must clear L1")
	}
}

// TestReadCacheL2BothMiss: L1+L2 miss counts a single miss and returns false.
func TestReadCacheL2BothMiss(t *testing.T) {
	c := newTestReadCache(1<<20, 4<<20)
	c.l2 = newFakeL2()
	if _, ok := c.get("o", "none"); ok {
		t.Fatal("both-miss must return false")
	}
	if c.misses.Load() != 1 {
		t.Fatalf("misses = %d, want 1", c.misses.Load())
	}
}

func TestReadCacheOverwriteRefresh(t *testing.T) {
	c := newTestReadCache(1<<20, 4<<20)
	c.put("o", "a", bytes.Repeat([]byte{1}, 100))
	c.put("o", "a", bytes.Repeat([]byte{2}, 200)) // overwrite, larger
	v, ok := c.get("o", "a")
	if !ok || len(v) != 200 || v[0] != 2 {
		t.Fatalf("overwrite not reflected: ok=%v len=%d", ok, len(v))
	}
	if c.curBytes != 200 {
		t.Fatalf("curBytes accounting off: %d", c.curBytes)
	}
}
