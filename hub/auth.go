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

	// rate limit defaults
	defaultIPReqPerSec    = 10.0
	defaultIPBurst        = 30
	defaultOwnerReqPerSec = 5.0
	defaultOwnerBurst     = 15
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

		// freshness: |now - au.Time| <= drift
		now := time.Now().Unix()
		delta := now - au.Time
		if delta < 0 {
			delta = -delta
		}
		if delta > drift {
			abortWithAuthError(c, fmt.Errorf("auth timestamp out of window (delta=%ds, max=%ds)", delta, drift))
			return
		}

		if err := sdk.VerifyAuth(au); err != nil {
			abortWithAuthError(c, fmt.Errorf("verify auth: %w", err))
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

// ----------------------------------------------------------------------------
// Owner ownership check (signer must own the namespace they're touching)
// ----------------------------------------------------------------------------

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
// Behavior:
//   - if owner is empty, default to the signer's own address
//   - if owner is provided, it must equal the signer (and be a valid ETH addr)
//
// Returns (resolvedOwner, ok). When ok is false, the response has already
// been written and the handler must return.
func ResolveOwnerForList(c *gin.Context, owner string) (string, bool) {
	signer := CtxAuthAddr(c)
	if signer == "" {
		abortWithAuthError(c, fmt.Errorf("no signer in context"))
		return "", false
	}

	if owner == "" {
		return signer, true
	}

	if !common.IsHexAddress(owner) {
		abortWithAuthError(c, fmt.Errorf("owner must be a 0x-prefixed Ethereum address"))
		return "", false
	}

	if !strings.EqualFold(owner, signer) {
		abortWithAuthError(c, fmt.Errorf("owner %s does not match signer %s", owner, signer))
		return "", false
	}

	return strings.ToLower(owner), true
}

// ----------------------------------------------------------------------------
// Rate limiting (two-tier: per-IP and per-owner)
// ----------------------------------------------------------------------------

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

// RateLimit produces a middleware that enforces:
//   - per-IP rate (DOS protection, applies before auth-recovered identity)
//   - per-owner rate (applied when AuthMiddleware has set ctxAuthAddrKey)
//
// Returns 429 on either tier exceedance.
func RateLimit() gin.HandlerFunc {
	ipRPS := envFloat("HUB_RATE_IP_RPS", defaultIPReqPerSec)
	ipBurst := envInt("HUB_RATE_IP_BURST", defaultIPBurst)
	ownerRPS := envFloat("HUB_RATE_OWNER_RPS", defaultOwnerReqPerSec)
	ownerBurst := envInt("HUB_RATE_OWNER_BURST", defaultOwnerBurst)

	ipReg := newLimiterRegistry(ipRPS, ipBurst)
	ownerReg := newLimiterRegistry(ownerRPS, ownerBurst)

	return func(c *gin.Context) {
		// per-IP first
		ip := c.ClientIP()
		if ip != "" && !ipReg.get(ip).Allow() {
			logger.Warnf("rate limit hit (ip) from %s %s", ip, c.Request.URL.Path)
			c.AbortWithStatusJSON(http.StatusTooManyRequests, lerror.ToAPIError("hub", fmt.Errorf("rate limit exceeded (ip)")))
			return
		}

		// per-owner if available (set by AuthMiddleware in the chain before us)
		if addr := CtxAuthAddr(c); addr != "" {
			if !ownerReg.get(addr).Allow() {
				logger.Warnf("rate limit hit (owner) from %s %s", addr, c.Request.URL.Path)
				c.AbortWithStatusJSON(http.StatusTooManyRequests, lerror.ToAPIError("hub", fmt.Errorf("rate limit exceeded (owner)")))
				return
			}
		}

		c.Next()
	}
}
