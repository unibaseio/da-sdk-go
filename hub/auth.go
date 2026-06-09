package hub

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"

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

// abortWithBadRequest is for client errors on public (unauthenticated) reads —
// e.g. a missing or malformed owner — where a 401 would be misleading.
func abortWithBadRequest(c *gin.Context, err error) {
	c.AbortWithStatusJSON(http.StatusBadRequest, lerror.ToAPIError("hub", err))
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
// IMPORTANT: the owner is returned VERBATIM — its case is never altered. The
// storage layer keys LogFS instances, needle metadata and MetaStore entries on
// the exact owner string used at write time, which is usually an EIP-55
// mixed-case address (e.g. 0x6370eF2f...). Lower-casing here would look up a
// different key than was written, breaking read-after-write and making all
// existing mixed-case-owner data un-downloadable.
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
		return owner, true
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

	return owner, true
}
