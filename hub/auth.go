package hub

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"

	lerror "github.com/unibaseio/da-sdk-go/lib/error"
	"github.com/unibaseio/da-sdk-go/sdk"
)

// ----------------------------------------------------------------------------
// Tunables (env-overridable)
// ----------------------------------------------------------------------------

const (
	ctxAuthAddrKey = "auth_addr" // recovered ETH address (lowercased, 0x...)

	// signature freshness window, both sides
	defaultAuthDriftSec int64 = 600 // 10 min

	// body size caps
	defaultMaxJSONBytes      int64 = 4 << 20  // 4 MB for /upload (JSON message)
	defaultMaxMultipartBytes int64 = 64 << 20 // 64 MB for /uploadData (file)

	// rate limit defaults. Deliberately generous: (1) legitimate explorer
	// traffic all arrives from the explorer's reverse-proxy IP (one IP, many
	// users); (2) there is no batch/stream read API, so a client syncing N
	// records issues N separate small GET /download calls — a few thousand
	// objects must not trip the limiter. The negative cache (not this limiter)
	// is the primary absorber of non-existent-key floods, so a high cap here is
	// safe; the real ceiling is single-instance read throughput.
	//
	// burst = 2x rps so a one-shot batch of up to `burst` requests clears
	// immediately, then sustains at `rps`. Tune per deployment via
	// HUB_RATE_IP_RPS / _BURST and HUB_RATE_OWNER_RPS / _BURST.
	defaultIPReqPerSec    = 1000.0
	defaultIPBurst        = 2000
	defaultOwnerReqPerSec = 1000.0
	defaultOwnerBurst     = 2000
)

func envInt64(key string, fallback int64) int64 {
	if v := os.Getenv(key); v != "" {
		n, err := strconv.ParseInt(v, 10, 64)
		if err == nil {
			return n
		}
	}
	return fallback
}

func envFloat(key string, fallback float64) float64 {
	if v := os.Getenv(key); v != "" {
		f, err := strconv.ParseFloat(v, 64)
		if err == nil {
			return f
		}
	}
	return fallback
}

func envInt(key string, fallback int) int {
	if v := os.Getenv(key); v != "" {
		n, err := strconv.Atoi(v)
		if err == nil {
			return n
		}
	}
	return fallback
}

// authBypassPaths are exempted from AuthMiddleware. Keep this list tiny.
var authBypassPaths = map[string]bool{
	"/api/info": true,
}

// ----------------------------------------------------------------------------
// Body size limit middleware
// ----------------------------------------------------------------------------

// MaxBodySize wraps r.Body with http.MaxBytesReader using a per-route cap.
// /uploadData (multipart) gets the larger cap, everything else gets the JSON cap.
func MaxBodySize() gin.HandlerFunc {
	jsonCap := envInt64("HUB_MAX_JSON_BYTES", defaultMaxJSONBytes)
	multipartCap := envInt64("HUB_MAX_MULTIPART_BYTES", defaultMaxMultipartBytes)

	return func(c *gin.Context) {
		var capBytes int64 = jsonCap
		if c.Request.URL.Path == "/api/uploadData" {
			capBytes = multipartCap
		}
		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, capBytes)
		c.Next()
	}
}

// ----------------------------------------------------------------------------
// Auth middleware
// ----------------------------------------------------------------------------

// AuthMiddleware parses Authorization header, verifies signature, enforces
// freshness, and stores the recovered ETH address (lowercased) in the gin
// context under ctxAuthAddrKey.
//
// Any /api/* request that isn't in authBypassPaths must carry a valid signature.
func AuthMiddleware() gin.HandlerFunc {
	drift := envInt64("HUB_AUTH_DRIFT_SEC", defaultAuthDriftSec)

	return func(c *gin.Context) {
		if authBypassPaths[c.Request.URL.Path] {
			c.Next()
			return
		}

		authStr := c.GetHeader("Authorization")
		if authStr == "" {
			abortWithAuthError(c, fmt.Errorf("missing Authorization header"))
			return
		}

		au, err := sdk.DecodeAuth(authStr)
		if err != nil {
			abortWithAuthError(c, fmt.Errorf("decode auth: %w", err))
			return
		}

		// Verify the signature first, recovering the timestamp BOUND INTO the
		// signature: au.Time for the legacy Hash||be64(Time) bytes, or the
		// "Issued At" embedded in the SIWE message. Freshness is then checked
		// against that bound timestamp, so a tampered envelope can't widen the
		// window.
		var signedAt int64
		if len(au.Msg) > 0 {
			// EIP-4361 / SIWE human-readable message (new clients)
			signedAt, err = sdk.VerifySIWE(au)
			if err != nil {
				abortWithAuthError(c, fmt.Errorf("verify auth: %w", err))
				return
			}
		} else {
			// legacy: signature over Hash(label) || be64(Time)
			signedAt = au.Time
			if err := sdk.VerifyAuth(au); err != nil {
				abortWithAuthError(c, fmt.Errorf("verify auth: %w", err))
				return
			}
		}

		// freshness: |now - signedAt| <= drift
		now := time.Now().Unix()
		delta := now - signedAt
		if delta < 0 {
			delta = -delta
		}
		if delta > drift {
			abortWithAuthError(c, fmt.Errorf("auth timestamp out of window (delta=%ds, max=%ds)", delta, drift))
			return
		}

		c.Set(ctxAuthAddrKey, strings.ToLower(au.Addr.Hex()))
		c.Next()
	}
}

func abortWithAuthError(c *gin.Context, err error) {
	logger.Warnf("auth reject from %s %s %s: %v", c.ClientIP(), c.Request.Method, c.Request.URL.Path, err)
	c.AbortWithStatusJSON(http.StatusUnauthorized, lerror.ToAPIError("hub", err))
}

// abortWithBadRequest is for client errors on public (unauthenticated) reads —
// e.g. a missing or malformed owner — where a 401 would be misleading.
func abortWithBadRequest(c *gin.Context, err error) {
	c.AbortWithStatusJSON(http.StatusBadRequest, lerror.ToAPIError("hub", err))
}

// ----------------------------------------------------------------------------
// Owner ownership check (signer must own the namespace they're touching)
// ----------------------------------------------------------------------------

// CanonOwner normalizes an owner to the canonical lowercase storage form.
// Ethereum addresses are case-insensitive — EIP-55 mixed case is only a UI
// checksum — so keying storage on lowercase stops one wallet from splitting
// into separate mixed-case vs lowercase namespaces. Non-address owners
// (legacy string ids) are lowercased too, which is harmless for matching.
func CanonOwner(owner string) string {
	return strings.ToLower(owner)
}

// ownerCandidates returns the owner forms to try when reading, newest-scheme
// first: the canonical lowercase form (how we store going forward), then the
// EIP-55 checksum form (how legacy data was stored). Deduped.
func ownerCandidates(owner string) []string {
	lc := strings.ToLower(owner)
	out := []string{lc}
	if common.IsHexAddress(owner) {
		if cs := common.HexToAddress(owner).Hex(); cs != lc {
			out = append(out, cs)
		}
	} else if owner != lc {
		out = append(out, owner)
	}
	return out
}

// CtxAuthAddr returns the lowercased 0x... address stored by AuthMiddleware,
// or "" if auth wasn't applied (e.g. on bypass routes).
func CtxAuthAddr(c *gin.Context) string {
	if v, ok := c.Get(ctxAuthAddrKey); ok {
		if s, ok2 := v.(string); ok2 {
			return s
		}
	}
	return ""
}

// RequireOwnerMatch enforces:
//  1. owner is a valid 0x-prefixed Ethereum address
//  2. owner (lowercased) equals the signer address recovered by AuthMiddleware
//
// On mismatch it writes a 401 + APIError and returns false; the caller MUST
// return immediately when this returns false.
func RequireOwnerMatch(c *gin.Context, owner string) bool {
	signer := CtxAuthAddr(c)
	if signer == "" {
		abortWithAuthError(c, fmt.Errorf("no signer in context"))
		return false
	}

	if owner == "" {
		abortWithAuthError(c, fmt.Errorf("owner is required"))
		return false
	}

	if !common.IsHexAddress(owner) {
		abortWithAuthError(c, fmt.Errorf("owner must be a 0x-prefixed Ethereum address"))
		return false
	}

	if !strings.EqualFold(owner, signer) {
		abortWithAuthError(c, fmt.Errorf("owner %s does not match signer %s", owner, signer))
		return false
	}

	return true
}

// ResolveOwnerForList is the read-side variant used by list/get endpoints.
// These run on the public (unauthenticated) /api group, so the common case
// has no signer. Behavior:
//   - no signer (public read): owner is OPTIONAL. Empty means "no filter —
//     list everything", which the explorer's global /agents and /memory
//     browse views rely on. A provided owner must be a valid address and
//     simply scopes the listing to that owner.
//   - signer present (e.g. if a route is ever moved to the authed group):
//     empty owner defaults to the signer; an explicit owner must match it.
//
// Stored content is client-encrypted, so an unscoped listing exposes only
// ciphertext plus public metadata (bucket / needle names) — the same data
// a block explorer shows.
//
// The owner is returned in canonical lowercase form (CanonOwner). Ethereum
// addresses are case-insensitive, so callers must match it case-insensitively
// (the gorm queries use LOWER(owner)=?, and logFSRead also tries the EIP-55
// checksum form for legacy data). This keeps a single wallet from splitting
// into mixed-case vs lowercase namespaces regardless of what case the client
// sent.
//
// Returns (resolvedOwner, ok). When ok is false, the response has already
// been written and the handler must return. An empty resolvedOwner with
// ok==true means "list all" — the gorm queries omit a zero-value owner from
// the WHERE clause, so this is the correct unscoped query.
func ResolveOwnerForList(c *gin.Context, owner string) (string, bool) {
	signer := CtxAuthAddr(c)

	if signer == "" {
		if owner == "" {
			return "", true
		}
		if !common.IsHexAddress(owner) {
			abortWithBadRequest(c, fmt.Errorf("owner must be a 0x-prefixed Ethereum address"))
			return "", false
		}
		return CanonOwner(owner), true
	}

	if owner == "" {
		return signer, true
	}

	if !common.IsHexAddress(owner) {
		abortWithBadRequest(c, fmt.Errorf("owner must be a 0x-prefixed Ethereum address"))
		return "", false
	}

	if !strings.EqualFold(owner, signer) {
		abortWithAuthError(c, fmt.Errorf("owner %s does not match signer %s", owner, signer))
		return "", false
	}

	return CanonOwner(owner), true
}

// ----------------------------------------------------------------------------
// Rate limiting (per-IP + per-owner)
// ----------------------------------------------------------------------------
//
// Reinstated as a per-IP guard. It works correctly when abusive hosts connect
// from their own IP (the case we observed: a bot hammering /api/download from
// a dedicated IP). Legitimate explorer traffic arrives from the explorer's
// reverse-proxy IP — one IP shared by many users — so the per-IP cap is set
// generously (see defaults above) to clear that aggregate while still cutting
// off a single host doing hundreds of req/s.
//
// NOTE: behind a proxy, c.ClientIP() reflects the proxy IP unless gin is told
// to trust it (SetTrustedProxies) and the proxy forwards X-Forwarded-For. We
// deliberately do NOT trust-all + read XFF here, because that lets any client
// spoof its IP and dodge the limit. So this throttles per *connecting* IP,
// which is exactly right for direct abusers and safe (just coarse) for proxied
// traffic. The download negative cache is the primary flood absorber; this is
// defense-in-depth.

type limiterRegistry struct {
	mu       sync.Mutex
	limiters map[string]*rate.Limiter
	r        rate.Limit
	burst    int
}

func newLimiterRegistry(rps float64, burst int) *limiterRegistry {
	return &limiterRegistry{
		limiters: make(map[string]*rate.Limiter),
		r:        rate.Limit(rps),
		burst:    burst,
	}
}

func (lr *limiterRegistry) get(key string) *rate.Limiter {
	lr.mu.Lock()
	defer lr.mu.Unlock()
	if l, ok := lr.limiters[key]; ok {
		return l
	}
	l := rate.NewLimiter(lr.r, lr.burst)
	lr.limiters[key] = l
	return l
}

// RateLimit enforces a per-IP token bucket, plus a per-owner bucket once
// AuthMiddleware has recovered a signer. Returns 429 on either tier.
func RateLimit() gin.HandlerFunc {
	ipRPS := envFloat("HUB_RATE_IP_RPS", defaultIPReqPerSec)
	ipBurst := envInt("HUB_RATE_IP_BURST", defaultIPBurst)
	ownerRPS := envFloat("HUB_RATE_OWNER_RPS", defaultOwnerReqPerSec)
	ownerBurst := envInt("HUB_RATE_OWNER_BURST", defaultOwnerBurst)

	ipReg := newLimiterRegistry(ipRPS, ipBurst)
	ownerReg := newLimiterRegistry(ownerRPS, ownerBurst)

	return func(c *gin.Context) {
		ip := c.ClientIP()
		if ip != "" && !ipReg.get(ip).Allow() {
			logger.Warnf("rate limit hit (ip) from %s %s", ip, c.Request.URL.Path)
			c.AbortWithStatusJSON(http.StatusTooManyRequests, lerror.ToAPIError("hub", fmt.Errorf("rate limit exceeded")))
			return
		}

		if addr := CtxAuthAddr(c); addr != "" {
			if !ownerReg.get(addr).Allow() {
				logger.Warnf("rate limit hit (owner) from %s %s", addr, c.Request.URL.Path)
				c.AbortWithStatusJSON(http.StatusTooManyRequests, lerror.ToAPIError("hub", fmt.Errorf("rate limit exceeded")))
				return
			}
		}

		c.Next()
	}
}
