package hub

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/unibaseio/da-sdk-go/lib/env"
)

// l2cache is the shared second-tier read cache behind the in-proc L1 LRU.
// redisL2 is the production impl; tests inject an in-memory fake. All ops are
// fail-open (see redisL2).
type l2cache interface {
	get(k string) ([]byte, bool)
	set(k string, v []byte)
	del(k string)
	ping() error
	close()
}

var _ l2cache = (*redisL2)(nil)

// redisL2 is the shared L2 read cache backing the in-proc L1 LRU (P4). It lets
// N stateless hub replicas share hot small-object bytes instead of each
// reconstructing from DA. Content is content-addressed and immutable, so entries
// never go stale — they are dropped only by Redis eviction (configure
// maxmemory-policy=allkeys-lru) or an explicit del on overwrite, or an optional
// TTL safety net (HUB_REDIS_TTL_SEC).
//
// Every op is FAIL-OPEN: a Redis error/timeout is treated as a miss (get) or
// silently ignored (set/del), so an unavailable Redis degrades to L1-only and
// never blocks or errors a read/write. nil when HUB_REDIS_ADDR is unset.
type redisL2 struct {
	rdb   *redis.Client
	ttl   time.Duration
	getTO time.Duration
	setTO time.Duration
}

func newRedisL2() *redisL2 {
	addr := env.Str("HUB_REDIS_ADDR", "")
	if addr == "" {
		return nil
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: env.Str("HUB_REDIS_PASSWORD", ""),
		DB:       env.Int("HUB_REDIS_DB", 0),
	})
	return &redisL2{
		rdb:   rdb,
		ttl:   time.Duration(env.Int("HUB_REDIS_TTL_SEC", 0)) * time.Second, // 0 = rely on maxmemory-policy
		getTO: 200 * time.Millisecond,
		setTO: 500 * time.Millisecond,
	}
}

func (r *redisL2) rkey(k string) string { return "rc:" + k }

// get returns the cached bytes for k, or (nil,false) on miss OR any error.
func (r *redisL2) get(k string) ([]byte, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), r.getTO)
	defer cancel()
	b, err := r.rdb.Get(ctx, r.rkey(k)).Bytes()
	if err != nil {
		return nil, false // redis.Nil (miss) or a transport error — both fail-open
	}
	return b, true
}

// set stores k->v (best-effort). Value must be immutable after this call.
func (r *redisL2) set(k string, v []byte) {
	ctx, cancel := context.WithTimeout(context.Background(), r.setTO)
	defer cancel()
	_ = r.rdb.Set(ctx, r.rkey(k), v, r.ttl).Err()
}

// del drops k (best-effort) — call on overwrite so a stale value can't shadow.
func (r *redisL2) del(k string) {
	ctx, cancel := context.WithTimeout(context.Background(), r.setTO)
	defer cancel()
	_ = r.rdb.Del(ctx, r.rkey(k)).Err()
}

func (r *redisL2) ping() error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	return r.rdb.Ping(ctx).Err()
}

func (r *redisL2) close() {
	if r != nil && r.rdb != nil {
		_ = r.rdb.Close()
	}
}
