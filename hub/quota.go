package hub

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"sync"
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

// Signer state is bounded. A bucket that has refilled to the full quota is
// indistinguishable from no bucket, so such buckets are dropped (at most every
// quotaSweepEvery, or every quotaFullSweepEvery while the table is full) and
// dropping one never gives a signer more than it would have had. Only signers
// still below their full quota (they wrote within the last window, in
// proportion to what they wrote) occupy the table.
//
// If quotaMaxSigners of those exist at once, a NEW signer is refused (429)
// until some refill. It used to be given one overflow bucket shared by all
// newcomers, which an attacker could drain to lock out every new signer for
// a window, or use as extra budget. Refusing is the fail-safe side: it spends
// nothing. Evicting a signer that is still in debt (LRU) was rejected because
// it hands that signer a fresh quota: an attacker cycling keys could reset its
// own. Filling the table is not cheap either: keeping a bucket below full for
// a time T takes T*QUOTA/WINDOW bytes actually written, i.e. at the defaults
// ~6 MB per signer per minute, ~600 GB a minute for 100k signers.
const (
	quotaMaxSigners     = 100_000
	quotaSweepEvery     = time.Minute
	quotaFullSweepEvery = time.Second
)

// errQuotaTableFull is returned to a new signer while every slot is taken by
// a signer that is still using its quota.
var errQuotaTableFull = errors.New("too many signers are using their write quota right now; retry later")

type quotaBucket struct {
	tokens float64 // bytes left
	last   time.Time
}

type writeQuota struct {
	mu        sync.Mutex
	buckets   map[string]*quotaBucket
	rate      float64 // refill, bytes per second
	limit     int64   // bucket size: the quota per window
	max       int
	lastSweep time.Time
	now       func() time.Time
	exempt    map[string]bool
}

// newWriteQuota returns nil when the quota is disabled.
func newWriteQuota() *writeQuota {
	limit := env.Int64("HUB_WRITE_QUOTA_BYTES", defaultWriteQuotaBytes)
	window := env.Int64("HUB_WRITE_QUOTA_WINDOW_SEC", defaultWriteQuotaWindowSec)
	if limit <= 0 || window <= 0 {
		return nil
	}
	q := &writeQuota{
		buckets: map[string]*quotaBucket{},
		rate:    float64(limit) / float64(window),
		limit:   limit,
		max:     quotaMaxSigners,
		now:     time.Now,
		exempt:  map[string]bool{},
	}
	for _, a := range strings.Split(env.Str("HUB_WRITE_QUOTA_EXEMPT", ""), ",") {
		if a = strings.ToLower(strings.TrimSpace(a)); a != "" {
			q.exempt[a] = true
		}
	}
	return q
}

// refill brings b up to now. Called with q.mu held.
func (q *writeQuota) refill(b *quotaBucket, now time.Time) {
	if dt := now.Sub(b.last).Seconds(); dt > 0 {
		b.tokens += dt * q.rate
		if b.tokens > float64(q.limit) {
			b.tokens = float64(q.limit)
		}
	}
	b.last = now
}

// sweepLocked drops buckets that have refilled completely (lossless).
func (q *writeQuota) sweepLocked(now time.Time) {
	q.lastSweep = now
	for k, b := range q.buckets {
		q.refill(b, now)
		if b.tokens >= float64(q.limit) {
			delete(q.buckets, k)
		}
	}
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
	q.mu.Lock()
	defer q.mu.Unlock()
	now := q.now()
	if now.Sub(q.lastSweep) >= quotaSweepEvery {
		q.sweepLocked(now)
	}
	b, ok := q.buckets[signer]
	if !ok {
		if len(q.buckets) >= q.max && now.Sub(q.lastSweep) >= quotaFullSweepEvery {
			q.sweepLocked(now)
		}
		if len(q.buckets) >= q.max {
			return errQuotaTableFull
		}
		b = &quotaBucket{tokens: float64(q.limit), last: now}
		q.buckets[signer] = b
	} else {
		q.refill(b, now)
	}
	if b.tokens < float64(n) {
		return fmt.Errorf("write quota exceeded for %s (%d bytes per window); retry later", signer, q.limit)
	}
	b.tokens -= float64(n)
	return nil
}

// refund gives back n bytes charged to signer for a write the hub did not
// carry out (capped at the full quota). A signer whose bucket has since been
// dropped is already at the full quota.
func (q *writeQuota) refund(signer string, n int64) {
	if q == nil || n <= 0 {
		return
	}
	signer = strings.ToLower(signer)
	if q.exempt[signer] {
		return
	}
	q.mu.Lock()
	defer q.mu.Unlock()
	b, ok := q.buckets[signer]
	if !ok {
		return
	}
	q.refill(b, q.now())
	b.tokens += float64(n)
	if b.tokens > float64(q.limit) {
		b.tokens = float64(q.limit)
	}
}

// size is the number of signers holding a bucket.
func (q *writeQuota) size() int {
	q.mu.Lock()
	defer q.mu.Unlock()
	return len(q.buckets)
}

// chargeWrite charges a hub-paid write of n bytes to the request's signer. On
// refusal it has written a 429 and returns false.
func (s *Server) chargeWrite(c *gin.Context, n int64) bool {
	if err := s.quota.charge(CtxAuthAddr(c), n); err != nil {
		logger.Warnf("write quota: %v", err)
		c.Header("Retry-After", "60")
		c.AbortWithStatusJSON(http.StatusTooManyRequests, lerror.ToAPIError("hub", err))
		return false
	}
	return true
}

// refundWrite returns n bytes of a charged write that failed before the hub
// spent anything on it.
func (s *Server) refundWrite(c *gin.Context, n int64) {
	s.quota.refund(CtxAuthAddr(c), n)
}
