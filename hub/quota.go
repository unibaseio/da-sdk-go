package hub

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/unibaseio/da-sdk-go/lib/env"
	lerror "github.com/unibaseio/da-sdk-go/lib/error"
)

// Per-signer write quota for writes the hub pays for (review 2026-10-04 M).
// Every object PUT/POST is drained to DA by the hub (it pays gas + the piece
// bond), and so is a seal with register=hub / hub_attributed. Without a quota
// one key could make the hub spend its whole balance.
//
//   - HUB_WRITE_QUOTA_BYTES (default 8 GiB): bytes one signer may write per
//     window; 0 turns the quota off.
//   - HUB_WRITE_QUOTA_WINDOW_SEC (default 86400): the rolling window. The quota
//     is a token bucket of QUOTA bytes refilling at QUOTA/WINDOW, so a signer
//     can burst the whole quota and then sustains QUOTA per WINDOW.
//   - HUB_WRITE_QUOTA_EXEMPT: comma-separated signer addresses with no quota
//     (e.g. an operator's service key).
//
// The default is deliberately generous (several times the largest legitimate
// daily writer we expect on testnet) — it caps abuse, not normal agents. State
// is in-memory per hub: with owner sharding every signer's writes land on its
// home shard, so the count is per signer; a restart resets it.
const (
	defaultWriteQuotaBytes     int64 = 8 << 30
	defaultWriteQuotaWindowSec int64 = 86400
)

type writeQuota struct {
	reg    *limiterRegistry
	limit  int64
	exempt map[string]bool
}

// newWriteQuota returns nil when the quota is disabled.
func newWriteQuota() *writeQuota {
	limit := env.Int64("HUB_WRITE_QUOTA_BYTES", defaultWriteQuotaBytes)
	window := env.Int64("HUB_WRITE_QUOTA_WINDOW_SEC", defaultWriteQuotaWindowSec)
	if limit <= 0 || window <= 0 {
		return nil
	}
	q := &writeQuota{
		reg:    newLimiterRegistry(float64(limit)/float64(window), int(limit)),
		limit:  limit,
		exempt: map[string]bool{},
	}
	for _, a := range strings.Split(env.Str("HUB_WRITE_QUOTA_EXEMPT", ""), ",") {
		if a = strings.ToLower(strings.TrimSpace(a)); a != "" {
			q.exempt[a] = true
		}
	}
	return q
}

// charge takes n bytes from signer's quota, or returns an error (nothing taken)
// when they don't fit.
func (q *writeQuota) charge(signer string, n int64) error {
	if q == nil || n <= 0 {
		return nil
	}
	signer = strings.ToLower(signer)
	if q.exempt[signer] {
		return nil
	}
	if n > q.limit {
		return fmt.Errorf("write of %d bytes exceeds the per-signer quota (%d bytes)", n, q.limit)
	}
	if !q.reg.get(signer).AllowN(time.Now(), int(n)) {
		return fmt.Errorf("write quota exceeded for %s (%d bytes per window); retry later", signer, q.limit)
	}
	return nil
}

// chargeWrite charges a hub-paid write of n bytes to the request's signer. On
// refusal it has written a 429 and returns false.
func (s *Server) chargeWrite(c *gin.Context, n int64) bool {
	if err := s.quota.charge(CtxAuthAddr(c), n); err != nil {
		logger.Warnf("write quota: %v", err)
		c.AbortWithStatusJSON(http.StatusTooManyRequests, lerror.ToAPIError("hub", err))
		return false
	}
	return true
}
