package hub

import "testing"

func TestMissCache_AddHasDel(t *testing.T) {
	m := newMissCache()
	o, n := "0xabc", "key-1"

	if m.has(o, n) {
		t.Fatal("fresh cache should not contain key")
	}
	m.add(o, n)
	if !m.has(o, n) {
		t.Fatal("key should be present after add")
	}
	// different key must not collide
	if m.has(o, "key-2") {
		t.Fatal("unrelated key should be absent")
	}
	// owner scoping: same name, different owner is a different entry
	if m.has("0xdef", n) {
		t.Fatal("same name under a different owner should be absent")
	}

	m.del(o, n)
	if m.has(o, n) {
		t.Fatal("key should be gone after del")
	}
}

func TestMissCache_ExpiredEntryEvicted(t *testing.T) {
	m := newMissCache()
	o, n := "0xabc", "key-1"
	m.add(o, n)
	// force expiry into the past
	m.mu.Lock()
	m.entries[missKey(o, n)] = 1 // unix nanos in 1970
	m.mu.Unlock()
	if m.has(o, n) {
		t.Fatal("expired entry must read as absent")
	}
	m.mu.Lock()
	_, still := m.entries[missKey(o, n)]
	m.mu.Unlock()
	if still {
		t.Fatal("expired entry should be lazily deleted on has()")
	}
}

func TestMissCache_DisabledWhenTTLZero(t *testing.T) {
	t.Setenv("HUB_DOWNLOAD_MISS_TTL_SEC", "0")
	m := newMissCache()
	m.add("0xabc", "key-1")
	if m.has("0xabc", "key-1") {
		t.Fatal("ttl<=0 must disable the cache (never reports present)")
	}
}

func TestMissCache_NilSafe(t *testing.T) {
	var m *missCache // e.g. a Server constructed without init
	if m.has("o", "n") {
		t.Fatal("nil cache has() must be false")
	}
	m.add("o", "n") // must not panic
	m.del("o", "n") // must not panic
}
