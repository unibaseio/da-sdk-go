package hub

import (
	"bytes"
	"os"
	"testing"
)

// TestRedisL2Integration exercises the real redisL2 against a live Redis. Skipped
// unless HUB_REDIS_TEST_ADDR is set, so `go test` stays hermetic by default.
//
//	docker run -d -p 6379:6379 redis
//	HUB_REDIS_TEST_ADDR=localhost:6379 go test ./hub/ -run RedisL2Integration -v
func TestRedisL2Integration(t *testing.T) {
	addr := os.Getenv("HUB_REDIS_TEST_ADDR")
	if addr == "" {
		t.Skip("set HUB_REDIS_TEST_ADDR to run the Redis L2 integration test")
	}
	t.Setenv("HUB_REDIS_ADDR", addr)

	r := newRedisL2()
	if r == nil {
		t.Fatal("newRedisL2 returned nil with HUB_REDIS_ADDR set")
	}
	defer r.close()
	if err := r.ping(); err != nil {
		t.Fatalf("ping: %v", err)
	}

	k := "itest/key"
	val := bytes.Repeat([]byte{0x3c}, 512)

	// miss before set
	if _, ok := r.get(k); ok {
		r.del(k) // clean stale
	}
	r.set(k, val)
	got, ok := r.get(k)
	if !ok || !bytes.Equal(got, val) {
		t.Fatalf("get after set: ok=%v len=%d", ok, len(got))
	}
	r.del(k)
	if _, ok := r.get(k); ok {
		t.Fatal("get after del should miss")
	}
}

// TestRedisL2FailOpen: an unreachable Redis never errors — get misses, set/del
// are silent no-ops, so a read/write path degrades to L1-only.
func TestRedisL2FailOpen(t *testing.T) {
	t.Setenv("HUB_REDIS_ADDR", "127.0.0.1:1") // nothing listening
	r := newRedisL2()
	if r == nil {
		t.Fatal("expected a client")
	}
	defer r.close()
	if _, ok := r.get("k"); ok {
		t.Fatal("unreachable redis must miss, not hit")
	}
	r.set("k", []byte("v")) // must not panic/block beyond timeout
	r.del("k")
}
