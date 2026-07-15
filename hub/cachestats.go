package hub

import (
	"net/http"

	"github.com/unibaseio/da-sdk-go/lib/logfs"

	"github.com/gin-gonic/gin"
)

// kvStatSource is implemented by the Badger metastore (lib/kv) to expose
// write-stall telemetry without coupling the hub to a concrete KV type.
type kvStatSource interface {
	KVStats() (puts, slowPuts, maxPutMs int64, l0Tables int)
}

// v1CacheStats surfaces the Phase-Q cache/FS effectiveness counters so an
// operator can measure the single-node ceiling before deciding whether to
// escalate to P3/P4 (S3/PG/Redis). Read-only, no auth, no sensitive data —
// just aggregate hit/miss rates across the four layers touched by Phase Q:
//
//   - fs:        per-owner LogFS get-or-create (sync.Map hit vs logfs.New)
//   - readcache: hot-object byte LRU (read-through)
//   - fdcache:   read-only volume fd reuse, summed across all loaded owners
//   - download:  DA-reconstruct fallbacks and how many were coalesced by
//     singleflight (shared), plus the optional concurrency cap
//
// GET /v1/cachestats
func (s *Server) v1CacheStats(c *gin.Context) {
	rcHit, rcMiss := s.readCache.Stats()
	l2Hit, l2On := s.readCache.L2Stats()

	var fdHit, fdMiss int64
	var ownersLoaded int
	s.lfs.Range(func(_, v any) bool {
		h, m := v.(*logfs.LogFS).FdStats()
		fdHit += h
		fdMiss += m
		ownersLoaded++
		return true
	})

	// Badger write-stall telemetry (P3-DB2): a high slow-put ratio or L0 table
	// count is the signal to migrate high-churn state to Postgres (P3 escalation).
	badger := gin.H{"available": false}
	if s.rp != nil {
		if src, ok := s.rp.MetaStore().(kvStatSource); ok {
			puts, slow, maxMs, l0 := src.KVStats()
			badger = gin.H{
				"available":  true,
				"puts":       puts,
				"slow_puts":  slow, // Put latency over BADGER_SLOW_PUT_MS
				"max_put_ms": maxMs,
				"l0_tables":  l0, // Badger stalls writes when L0 grows too large
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"fs": gin.H{
			"hit":           s.fsHit.Load(),    // lock-free sync.Map hit
			"create":        s.fsCreate.Load(), // on-demand logfs.New
			"owners_loaded": ownersLoaded,
		},
		"readcache": gin.H{
			"enabled":    s.readCache != nil, // L1 in-proc LRU
			"hit":        rcHit,              // L1 hit
			"miss":       rcMiss,             // both tiers miss
			"l2_enabled": l2On,               // shared Redis tier
			"l2_hit":     l2Hit,              // L1 miss, L2 hit
		},
		"fdcache": gin.H{
			"hit":  fdHit,
			"miss": fdMiss,
		},
		"badger": badger,
		"download": gin.H{
			"total":             s.dlTotal.Load(),  // cold-key DA fallbacks
			"shared":            s.dlShared.Load(), // coalesced by singleflight
			"concurrency_limit": cap(s.dlSem),      // 0 = unlimited
		},
	})
}
