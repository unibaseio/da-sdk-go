package hub

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"

	"github.com/unibaseio/da-sdk-go/lib/env"
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

// env helpers now live in lib/env (Int64/Float/Int); HUB_* keys stay local
// string literals since they're hub-specific.

// authBypassPaths are exempted from AuthMiddleware. Keep this list tiny.
var authBypassPaths = map[string]bool{
	"/api/info": true,
}

// ----------------------------------------------------------------------------
// Body size limit middleware
// ----------------------------------------------------------------------------

// MaxBodySize wraps r.Body with http.MaxBytesReader using a per-route cap.
// Multipart file-upload routes (/uploadData, /seal) get the larger cap;
// everything else gets the JSON cap.
func MaxBodySize() gin.HandlerFunc {
	jsonCap := env.Int64("HUB_MAX_JSON_BYTES", defaultMaxJSONBytes)
	multipartCap := env.Int64("HUB_MAX_MULTIPART_BYTES", defaultMaxMultipartBytes)

	return func(c *gin.Context) {
		var capBytes int64 = jsonCap
		switch c.Request.URL.Path {
		case "/api/uploadData", "/api/seal":
			// both stream a (potentially large) file part — must not be
			// truncated by the small JSON cap.
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
// recoverSigner verifies an Authorization header and returns the recovered
// lowercased signer address, or an error. Shared by AuthMiddleware (write group)
// and the read-side enumeration guards (RequireOwnerForList / RequireAuthenticated)
// so identical signature + freshness rules apply everywhere.
func recoverSigner(authStr string, drift int64) (string, error) {
	au, err := sdk.DecodeAuth(authStr)
	if err != nil {
		return "", fmt.Errorf("decode auth: %w", err)
	}

	// signature + freshness against the timestamp bound into the signature,
	// and a SIWE message only if it was issued for this hub's domain
	if err := sdk.VerifyAuthFreshDomains(au, drift, siweDomains()); err != nil {
		return "", err
	}
	return strings.ToLower(au.Addr.Hex()), nil
}

// siweDomains lists the domains a SIWE sign-in must be issued for to be
// accepted here (HUB_SIWE_DOMAINS, comma-separated, e.g. the frontends that
// talk to this hub). Unset accepts any domain — set it in production.
func siweDomains() []string {
	v := strings.TrimSpace(env.Str("HUB_SIWE_DOMAINS", ""))
	if v == "" {
		return nil
	}
	return strings.Split(v, ",")
}

func AuthMiddleware() gin.HandlerFunc {
	drift := env.Int64("HUB_AUTH_DRIFT_SEC", defaultAuthDriftSec)

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

		addr, err := recoverSigner(authStr, drift)
		if err != nil {
			abortWithAuthError(c, err)
			return
		}

		c.Set(ctxAuthAddrKey, addr)
		c.Next()
	}
}

// RequireOwnerForList guards owner-scoped enumeration (list buckets/needles/
// volumes/objects). Unlike ResolveOwnerForList (used by public point-reads) it
// REQUIRES a valid signed request and scopes the result to the signer: an
// anonymous caller can no longer enumerate any owner's namespace, and a signer
// can only list their own. This is the "hybrid" read model — enumeration is
// gated while exact-key point reads and aggregate stats stay public. It reads
// the Authorization header directly (independent of group middleware), so the
// public read group and its point-read handlers are untouched.
func RequireOwnerForList(c *gin.Context, owner string) (string, bool) {
	authStr := c.GetHeader("Authorization")
	if authStr == "" {
		abortWithAuthError(c, fmt.Errorf("listing requires a signed request (namespace enumeration is not public)"))
		return "", false
	}
	signer, err := recoverSigner(authStr, env.Int64("HUB_AUTH_DRIFT_SEC", defaultAuthDriftSec))
	if err != nil {
		abortWithAuthError(c, err)
		return "", false
	}
	// A privileged read-only "reader" identity (e.g. the block explorer) may
	// enumerate ANY owner — metadata only; stored content stays client-encrypted.
	// Writes are unaffected (RequireOwnerMatch still demands owner==signer).
	if isReader(signer) {
		return readerScope(c, owner)
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

// isReader reports whether addr is in the HUB_READER_ADDRS allowlist — a set of
// addresses permitted to enumerate any owner for reads only. Empty list (default)
// means no readers, so behaviour is identical to a plain owner==signer hub.
func isReader(addr string) bool {
	list := env.Str("HUB_READER_ADDRS", "")
	if list == "" || addr == "" {
		return false
	}
	for _, a := range strings.Split(list, ",") {
		if strings.EqualFold(strings.TrimSpace(a), addr) {
			return true
		}
	}
	return false
}

// readerScope resolves the requested owner for a reader: empty = list-all, else
// the requested owner canonicalised (any owner, not just the signer's own).
func readerScope(c *gin.Context, owner string) (string, bool) {
	if owner == "" {
		return "", true
	}
	if !common.IsHexAddress(owner) {
		abortWithBadRequest(c, fmt.Errorf("owner must be a 0x-prefixed Ethereum address"))
		return "", false
	}
	return CanonOwner(owner), true
}

// RequireAuthenticated guards global (not owner-scoped) enumeration such as the
// account registry: any valid signed request passes, anonymous is rejected — so
// the full owner list can't be dumped anonymously.
func RequireAuthenticated(c *gin.Context) bool {
	authStr := c.GetHeader("Authorization")
	if authStr == "" {
		abortWithAuthError(c, fmt.Errorf("this listing requires a signed request"))
		return false
	}
	if _, err := recoverSigner(authStr, env.Int64("HUB_AUTH_DRIFT_SEC", defaultAuthDriftSec)); err != nil {
		abortWithAuthError(c, err)
		return false
	}
	return true
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

	// Reader identity: enumerate any owner (read-only), see RequireOwnerForList.
	if isReader(signer) {
		return readerScope(c, owner)
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
