package hub

import (
	"net/http"

	"github.com/unibaseio/da-sdk-go/lib/logfs"

	"github.com/gin-gonic/gin"
)

// v1CacheStats surfaces the Phase-Q cache/FS effectiveness counters so an
// operator can measure the single-node ceiling before deciding whether to
// escalate to P3/P4 (S3/PG/Redis). Read-only, no auth, no sensitive data —
// just aggregate hit/miss rates across the four layers touched by Phase Q:
//
//   - fs:        per-owner LogFS get-or-create (sync.Map hit vs logfs.New)
//   - readcache: hot-object byte LRU (read-through)
//   - fdcache:   read-only volume fd reuse, summed across all loaded owners
//   - download:  DA-reconstruct fallbacks and how many were coalesced by
//                singleflight (shared), plus the optional concurrency cap
//
// GET /v1/cachestats
func (s *Server) v1CacheStats(c *gin.Context) {
	rcHit, rcMiss := s.readCache.Stats()

	var fdHit, fdMiss int64
	var ownersLoaded int
	s.lfs.Range(func(_, v any) bool {
		h, m := v.(*logfs.LogFS).FdStats()
		fdHit += h
		fdMiss += m
		ownersLoaded++
		return true
	})

	c.JSON(http.StatusOK, gin.H{
		"fs": gin.H{
			"hit":           s.fsHit.Load(),    // lock-free sync.Map hit
			"create":        s.fsCreate.Load(), // on-demand logfs.New
			"owners_loaded": ownersLoaded,
		},
		"readcache": gin.H{
			"enabled": s.readCache != nil,
			"hit":     rcHit,
			"miss":    rcMiss,
		},
		"fdcache": gin.H{
			"hit":  fdHit,
			"miss": fdMiss,
		},
		"download": gin.H{
			"total":             s.dlTotal.Load(),  // cold-key DA fallbacks
			"shared":            s.dlShared.Load(), // coalesced by singleflight
			"concurrency_limit": cap(s.dlSem),      // 0 = unlimited
		},
	})
}
