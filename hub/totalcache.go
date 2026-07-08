package hub

import (
	"context"
	"sync"
	"time"

	"github.com/unibaseio/da-sdk-go/lib/env"
	"github.com/unibaseio/da-sdk-go/lib/types"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
)

// totalCache caches the expensive COUNT(*) results behind the /v1 lists'
// withTotal option. The needles table is tens of millions of rows, so an exact
// count costs seconds on Postgres; pagers tolerate a slightly stale total.
//
// TTL via HUB_TOTAL_CACHE_SEC (default 60; 0 disables caching). Lookups within
// the TTL are served from memory; the first lookup after expiry recomputes it,
// deduplicated by singleflight so a burst of first-page loads runs ONE count.
// On recompute error a stale value keeps being served rather than failing the
// list request.
type totalCache struct {
	mu sync.Mutex
	m  map[string]totalEntry
	sf singleflight.Group
}

type totalEntry struct {
	val int64
	at  time.Time
}

func newTotalCache() *totalCache { return &totalCache{m: make(map[string]totalEntry)} }

func totalCacheTTL() time.Duration {
	return time.Duration(env.Int64("HUB_TOTAL_CACHE_SEC", 60)) * time.Second
}

func (t *totalCache) get(key string, count func() (int64, error)) (int64, error) {
	ttl := totalCacheTTL()
	if ttl <= 0 {
		return count()
	}
	t.mu.Lock()
	e, ok := t.m[key]
	t.mu.Unlock()
	if ok && time.Since(e.at) < ttl {
		return e.val, nil
	}
	v, err, _ := t.sf.Do(key, func() (interface{}, error) {
		n, err := count()
		if err != nil {
			return int64(0), err
		}
		t.mu.Lock()
		t.m[key] = totalEntry{val: n, at: time.Now()}
		t.mu.Unlock()
		return n, nil
	})
	if err != nil {
		if ok { // serve stale over failing the whole list request
			return e.val, nil
		}
		return 0, err
	}
	return v.(int64), nil
}

// cachedCount runs (or serves from cache) a COUNT over model with the given
// filter, keyed by the caller-provided scope string.
func (s *Server) cachedCount(key string, model interface{}, filter func(*gorm.DB) *gorm.DB) (int64, error) {
	return s.totals.get(key, func() (int64, error) {
		var n int64
		err := filter(s.gdb.Model(model)).Count(&n).Error
		return n, err
	})
}

// startTotalRefresh keeps the hottest keys warm in the background so even the
// first request after TTL expiry never pays the multi-second count: the global
// needles total (explorer Memory page) and the global buckets total (Agents
// page). Owner-/bucket-scoped totals stay on-demand (long tail, cheaper).
func (s *Server) startTotalRefresh(ctx context.Context) {
	ttl := totalCacheTTL()
	if ttl <= 0 {
		return
	}
	refresh := func() {
		var n int64
		if err := s.gdb.Model(&types.Needle{}).Count(&n).Error; err == nil {
			s.totals.mu.Lock()
			s.totals.m["needles||"] = totalEntry{val: n, at: time.Now()}
			s.totals.mu.Unlock()
		}
		if err := s.gdb.Model(&types.Bucket{}).Count(&n).Error; err == nil {
			s.totals.mu.Lock()
			s.totals.m["buckets||"] = totalEntry{val: n, at: time.Now()}
			s.totals.mu.Unlock()
		}
	}
	go func() {
		refresh() // warm immediately at startup
		ticker := time.NewTicker(ttl)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-s.shutdownChan:
				return
			case <-ticker.C:
				refresh()
			}
		}
	}()
}
