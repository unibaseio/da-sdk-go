package hub

import (
	"sync"
	"time"

	"github.com/unibaseio/da-sdk-go/lib/env"
)

// missCache is a small TTL set of "this download key is definitely not here"
// markers. It absorbs download floods: a client hammering /api/download for a
// key that isn't stored would otherwise trigger the full fallback chain —
// including remote round-trips to the gateway — on every single request
// (~30-50ms each), which can saturate the hub. Once a key is confirmed
// missing, subsequent lookups short-circuit to a cheap 404 until the TTL
// expires or the key is actually written.
//
// Only the EXPENSIVE remote fallback is skipped; the cheap local LogFS lookup
// still runs first on every request, so a key that gets written stays
// immediately readable (and logFSWrite proactively evicts it anyway).
type missCache struct {
	mu      sync.Mutex
	entries map[string]int64 // key -> expiry (unix nanos)
	ttl     time.Duration
	max     int
}

const (
	defaultMissTTLSec   int64 = 30
	defaultMissMaxItems       = 50000
)

func newMissCache() *missCache {
	return &missCache{
		entries: make(map[string]int64),
		ttl:     time.Duration(env.Int64("HUB_DOWNLOAD_MISS_TTL_SEC", defaultMissTTLSec)) * time.Second,
		max:     defaultMissMaxItems,
	}
}

func missKey(owner, name string) string { return owner + "/" + name }

// has reports whether (owner,name) was recently confirmed missing.
func (m *missCache) has(owner, name string) bool {
	if m == nil || m.ttl <= 0 {
		return false
	}
	k := missKey(owner, name)
	now := time.Now().UnixNano()
	m.mu.Lock()
	defer m.mu.Unlock()
	exp, ok := m.entries[k]
	if !ok {
		return false
	}
	if now >= exp {
		delete(m.entries, k)
		return false
	}
	return true
}

// add records (owner,name) as missing for one TTL window.
func (m *missCache) add(owner, name string) {
	if m == nil || m.ttl <= 0 {
		return
	}
	k := missKey(owner, name)
	now := time.Now().UnixNano()
	m.mu.Lock()
	defer m.mu.Unlock()
	if len(m.entries) >= m.max {
		// bounded: drop expired entries; if still full, reset (a flood hits the
		// same few keys, so the cache re-populates cheaply).
		for ek, ee := range m.entries {
			if now >= ee {
				delete(m.entries, ek)
			}
		}
		if len(m.entries) >= m.max {
			m.entries = make(map[string]int64)
		}
	}
	m.entries[k] = now + int64(m.ttl)
}

// del removes any miss marker for (owner,name) — call after a successful write
// so a freshly stored key is never shadowed by a stale negative.
func (m *missCache) del(owner, name string) {
	if m == nil {
		return
	}
	k := missKey(owner, name)
	m.mu.Lock()
	delete(m.entries, k)
	m.mu.Unlock()
}
