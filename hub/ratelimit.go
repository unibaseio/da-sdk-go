package hub

import (
	"fmt"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"

	"github.com/unibaseio/da-sdk-go/lib/env"
	lerror "github.com/unibaseio/da-sdk-go/lib/error"
)

// ----------------------------------------------------------------------------
// Client IP / trusted proxies
// ----------------------------------------------------------------------------
//
// gin trusts EVERY proxy by default, so c.ClientIP() returned whatever the
// client wrote into X-Forwarded-For and a forged header dodged the per-IP limit
// (review 2026-10-04). The hub now trusts only the proxies listed in
// HUB_TRUSTED_PROXIES (comma-separated CIDRs or IPs):
//
//   - unset (default): private + loopback ranges. The testnet/mainnet hubs sit
//     behind an AWS ALB inside a private VPC (10.0.0.0/16), and shard peers
//     reverse-proxy each other over private IPs, so their XFF must be honored —
//     with trust-none every request would appear to come from the ALB and the
//     per-IP limiter would throttle all users together. A hub exposed directly
//     to the internet is still safe under this default: a public client's
//     RemoteAddr is untrusted, so its XFF is ignored.
//   - "none": trust nobody; ClientIP is the TCP peer (RemoteAddr).
//   - a list: trust exactly those ranges (e.g. "10.0.0.0/16" for the ALB VPC).
//
// Only X-Forwarded-For is consulted (not X-Real-IP): the ALB sets XFF, and
// gin walks it right-to-left skipping trusted hops, so a client-supplied prefix
// is never taken as the client address.
var defaultTrustedProxies = []string{
	"10.0.0.0/8", "172.16.0.0/12", "192.168.0.0/16", // RFC 1918
	"127.0.0.0/8", "::1/128", // loopback
	"fc00::/7", // IPv6 unique-local
}

// trustedProxies resolves HUB_TRUSTED_PROXIES; nil means trust none.
func trustedProxies() []string {
	v := strings.TrimSpace(env.Str("HUB_TRUSTED_PROXIES", ""))
	switch strings.ToLower(v) {
	case "":
		return defaultTrustedProxies
	case "none", "off", "-":
		return nil
	}
	var out []string
	for _, p := range strings.Split(v, ",") {
		if p = strings.TrimSpace(p); p != "" {
			out = append(out, p)
		}
	}
	return out
}

// configureClientIP applies the trusted-proxy policy to a gin engine. An
// invalid HUB_TRUSTED_PROXIES is a startup error rather than a silent fallback
// to some other trust set.
func configureClientIP(r *gin.Engine) error {
	r.RemoteIPHeaders = []string{"X-Forwarded-For"}
	r.ForwardedByClientIP = true
	if err := r.SetTrustedProxies(trustedProxies()); err != nil {
		return fmt.Errorf("HUB_TRUSTED_PROXIES: %w", err)
	}
	return nil
}

// ----------------------------------------------------------------------------
// Rate limiting (per-IP + per-owner)
// ----------------------------------------------------------------------------
//
// Two tiers. The per-IP tier runs BEFORE signature verification (the expensive
// part of a request), so a flood of junk-signed requests is cut off cheaply;
// the per-owner tier needs the recovered signer and therefore runs after
// AuthMiddleware. Limits are deliberately generous — see the defaults in
// auth.go — and the download negative cache stays the primary flood absorber.

const (
	// how often a registry sweeps idle entries, and the entry cap beyond which
	// new keys share a single overflow bucket (bounds memory under key churn,
	// e.g. an attacker rotating IPv6 addresses or fresh signer keys).
	limiterSweepEvery = time.Minute
	limiterMinIdle    = time.Minute
	limiterMaxEntries = 100_000
)

type limiterEntry struct {
	l    *rate.Limiter
	seen time.Time
}

// limiterRegistry is a keyed set of token buckets with idle eviction. An entry
// idle for longer than burst/rate has refilled completely, so dropping it and
// recreating it on the next request is indistinguishable from keeping it —
// eviction never loosens a limit.
type limiterRegistry struct {
	mu        sync.Mutex
	limiters  map[string]*limiterEntry
	r         rate.Limit
	burst     int
	idle      time.Duration
	max       int
	lastSweep time.Time
	overflow  *rate.Limiter
	now       func() time.Time
}

func newLimiterRegistry(rps float64, burst int) *limiterRegistry {
	idle := limiterMinIdle
	if rps > 0 {
		if full := time.Duration(float64(burst) / rps * float64(time.Second)); full > idle {
			idle = full
		}
	}
	return &limiterRegistry{
		limiters: make(map[string]*limiterEntry),
		r:        rate.Limit(rps),
		burst:    burst,
		idle:     idle,
		max:      limiterMaxEntries,
		overflow: rate.NewLimiter(rate.Limit(rps), burst),
		now:      time.Now,
	}
}

func (lr *limiterRegistry) get(key string) *rate.Limiter {
	lr.mu.Lock()
	defer lr.mu.Unlock()
	now := lr.now()
	if e, ok := lr.limiters[key]; ok {
		e.seen = now
		return e.l
	}
	if now.Sub(lr.lastSweep) >= limiterSweepEvery || len(lr.limiters) >= lr.max {
		lr.sweepLocked(now)
	}
	if len(lr.limiters) >= lr.max {
		// still full of live keys: new keys share one bucket rather than
		// growing the map without bound
		return lr.overflow
	}
	l := rate.NewLimiter(lr.r, lr.burst)
	lr.limiters[key] = &limiterEntry{l: l, seen: now}
	return l
}

func (lr *limiterRegistry) sweepLocked(now time.Time) {
	lr.lastSweep = now
	for k, e := range lr.limiters {
		if now.Sub(e.seen) > lr.idle {
			delete(lr.limiters, k)
		}
	}
}

func (lr *limiterRegistry) size() int {
	lr.mu.Lock()
	defer lr.mu.Unlock()
	return len(lr.limiters)
}

// ipKey is the limiter key for a client IP. IPv6 clients are keyed by their
// /64: a single host typically owns a whole /64 and could otherwise rotate
// through it to get a fresh bucket per request.
func ipKey(ip string) string {
	p := net.ParseIP(ip)
	if p == nil || p.To4() != nil {
		return ip
	}
	return p.Mask(net.CIDRMask(64, 128)).String() + "/64"
}

func abortRateLimited(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusTooManyRequests, lerror.ToAPIError("hub", fmt.Errorf("rate limit exceeded")))
}

// ipAllowed / ownerAllowed run one tier; on a hit they write the 429 and
// return false.
func ipAllowed(c *gin.Context, reg *limiterRegistry) bool {
	ip := c.ClientIP()
	if ip != "" && !reg.get(ipKey(ip)).Allow() {
		logger.Warnf("rate limit hit (ip) from %s %s", ip, c.Request.URL.Path)
		abortRateLimited(c)
		return false
	}
	return true
}

func ownerAllowed(c *gin.Context, reg *limiterRegistry) bool {
	if addr := CtxAuthAddr(c); addr != "" && !reg.get(addr).Allow() {
		logger.Warnf("rate limit hit (owner) from %s %s", addr, c.Request.URL.Path)
		abortRateLimited(c)
		return false
	}
	return true
}

func newIPRegistry() *limiterRegistry {
	return newLimiterRegistry(env.Float("HUB_RATE_IP_RPS", defaultIPReqPerSec), env.Int("HUB_RATE_IP_BURST", defaultIPBurst))
}

func newOwnerRegistry() *limiterRegistry {
	return newLimiterRegistry(env.Float("HUB_RATE_OWNER_RPS", defaultOwnerReqPerSec), env.Int("HUB_RATE_OWNER_BURST", defaultOwnerBurst))
}

// IPRateLimit is the per-client-IP tier. Mount it before AuthMiddleware.
func IPRateLimit() gin.HandlerFunc {
	reg := newIPRegistry()
	return func(c *gin.Context) {
		if ipAllowed(c, reg) {
			c.Next()
		}
	}
}

// OwnerRateLimit is the per-signer tier; a no-op for requests without a
// recovered signer. Mount it after AuthMiddleware.
func OwnerRateLimit() gin.HandlerFunc {
	reg := newOwnerRegistry()
	return func(c *gin.Context) {
		if ownerAllowed(c, reg) {
			c.Next()
		}
	}
}

// RateLimit applies both tiers in one middleware (per-IP, then per-owner when a
// signer is already in context). Kept for callers that mount a single limiter;
// the /v1 write group splits them around AuthMiddleware instead.
func RateLimit() gin.HandlerFunc {
	ipReg, ownerReg := newIPRegistry(), newOwnerRegistry()
	return func(c *gin.Context) {
		if ipAllowed(c, ipReg) && ownerAllowed(c, ownerReg) {
			c.Next()
		}
	}
}
