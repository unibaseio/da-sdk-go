package hub

import (
	"bytes"
	"container/list"
	"testing"
)

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
