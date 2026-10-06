package hub

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync/atomic"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"

	com "github.com/unibaseio/da-sdk-go/contract/common"
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

	// body size cap of the /v1 write group (object bodies, multipart batches, seal)
	defaultMaxMultipartBytes int64 = 64 << 20 // 64 MB

	// rate limit defaults. Deliberately generous: (1) legitimate explorer
	// traffic all arrives from the explorer's reverse-proxy IP (one IP, many
	// users); (2) there is no batch/stream read API, so a client syncing N
	// records issues N separate small object reads — a few thousand objects
	// must not trip the limiter. The real ceiling is single-instance read
	// throughput.
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

// ----------------------------------------------------------------------------
// Auth middleware
// ----------------------------------------------------------------------------

// recoverSigner verifies an Authorization header and returns the recovered
// lowercased signer address, or an error. Shared by AuthMiddleware (write group)
// and the read-side enumeration guard (RequireOwnerForList) so identical
// signature + freshness rules apply everywhere.
func recoverSigner(authStr string, drift int64) (string, error) {
	au, err := sdk.DecodeAuth(authStr)
	if err != nil {
		return "", fmt.Errorf("decode auth: %w", err)
	}

	// signature + freshness against the timestamp bound into the signature,
	// and a sign-in message only if it is for this hub's chain and (when
	// HUB_SIWE_DOMAINS is set) issued for, and pointing at, an allowed
	// domain. Its Nonce is not checked: see sdk.SIWEPolicy (deferred, needs
	// a client protocol change).
	if err := sdk.VerifyAuthFreshPolicy(au, drift, sdk.SIWEPolicy{Domains: siweDomains(), ChainIDs: siweChainIDs()}); err != nil {
		return "", err
	}
	return strings.ToLower(au.Addr.Hex()), nil
}

// hubChainID is the EIP-155 id of the chain this hub is configured for (set
// by NewServer; 0 = unknown).
var hubChainID atomic.Int64

// siweChainIDs lists the chains a sign-in message may name: HUB_SIWE_CHAIN_IDS
// (comma-separated, for frontends whose wallets sign on another chain), else
// the hub's own chain. An unparsable entry is ignored with a warning; if none
// is left, the hub's chain applies.
func siweChainIDs() []int64 {
	var ids []int64
	for _, f := range strings.Split(env.Str("HUB_SIWE_CHAIN_IDS", ""), ",") {
		if f = strings.TrimSpace(f); f == "" {
			continue
		}
		id, err := strconv.ParseInt(f, 10, 64)
		if err != nil || id <= 0 {
			logger.Warnf("HUB_SIWE_CHAIN_IDS: ignoring %q", f)
			continue
		}
		ids = append(ids, id)
	}
	if len(ids) == 0 {
		if id := hubChainID.Load(); id != 0 {
			ids = []int64{id}
		}
	}
	return ids
}

// chainIDOf maps a configured chain type to its EIP-155 id (0 if unknown).
func chainIDOf(chainType string) int64 {
	switch chainType {
	case com.BaseSepolia:
		return com.BaseSepoliaChainID
	case com.BaseMainnet:
		return com.BaseMainnetChainID
	case com.BSCMainnet:
		return com.BSCMainnetChainID
	case com.ETHMainnet:
		return com.ETHMainnetChainID
	case com.BNBTestnetV2:
		return int64(com.BNBTestnetChainID)
	case com.BNBTestnetDAO:
		return int64(com.BNBTestnetDAOChainID)
	case com.LocalAnvil:
		return int64(com.LocalAnvilChainID)
	}
	return 0
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

// AuthMiddleware parses the Authorization header, verifies the signature,
// enforces freshness, and stores the recovered ETH address (lowercased) in the
// gin context under ctxAuthAddrKey. Every request through it must be signed.
func AuthMiddleware() gin.HandlerFunc {
	drift := env.Int64("HUB_AUTH_DRIFT_SEC", defaultAuthDriftSec)

	return func(c *gin.Context) {
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

// CtxAuthAddr returns the lowercased 0x... address stored by AuthMiddleware,
// or "" if auth wasn't applied (the public read group).
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

// ResolveOwnerForList is the read-side variant used by point-read endpoints.
// These run on the public (unauthenticated) /v1 group, so the common case
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
// (the gorm queries use LOWER(owner)=?; content is then read under the owner
// form the matched row stores). This keeps a single wallet from splitting
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
