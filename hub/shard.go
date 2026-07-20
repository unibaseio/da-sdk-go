package hub

import (
	"crypto/subtle"
	"hash/crc32"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"sync/atomic"

	"github.com/gin-gonic/gin"

	"github.com/unibaseio/da-sdk-go/lib/env"
)

// shardRouter implements owner-sharded sticky WRITE routing (P4). Writes for one
// owner MUST be single-homed: each hub owns its own local logfs + Badger offset
// state per owner, so two hubs writing the same owner would build divergent
// volumes. We map hash(owner) -> exactly one hub; a write that lands on the
// wrong hub is transparently reverse-proxied to the owner's home hub. Reads are
// unrestricted (any hub serves them via L2 Redis / DA reconstruct).
//
// Disabled (nil) unless HUB_SHARD_TOTAL>1 with a valid HUB_SHARD_INDEX and a
// HUB_SHARD_PEERS list of exactly TOTAL comma-separated base URLs (index-order).
// Default single-node behavior is unchanged.
type shardRouter struct {
	index   int
	total   int
	peers   []*url.URL
	proxies []*httputil.ReverseProxy

	proxied     atomic.Int64 // writes forwarded to a peer
	local       atomic.Int64 // writes served locally (home shard)
	readProxied atomic.Int64 // reads forwarded to the owner's home shard

	// fwdSecret is the shared value a peer stamps on shardFwdHeader when it
	// forwards. When set (HUB_SHARD_FWD_SECRET), the receiver honors the
	// forwarded-once marker ONLY if it matches — so an external client can't
	// forge the header to misroute its own writes or suppress the read
	// fallback. Empty = legacy behavior (any non-empty value counts).
	fwdSecret string
}

// shardFwdHeader marks a request already forwarded once by a shard peer. The
// receiving hub must answer locally and never re-proxy: home resolution is
// deterministic, so a second hop only happens under topology misconfig (e.g.
// two hubs disagreeing on TOTAL) — without this guard that would ping-pong
// forever. One hop max, then serve or fail locally. The value is a shared
// secret (fwdSecret) so the marker cannot be forged from outside the fleet.
const shardFwdHeader = "X-Hub-Shard-Fwd"

func newShardRouter() *shardRouter {
	total := env.Int("HUB_SHARD_TOTAL", 1)
	if total <= 1 {
		return nil // single-node = unchanged
	}
	index := env.Int("HUB_SHARD_INDEX", -1)
	if index < 0 || index >= total {
		logger.Errorf("HUB_SHARD_TOTAL=%d but HUB_SHARD_INDEX=%d is out of range; sharding disabled", total, index)
		return nil
	}
	parts := strings.Split(env.Str("HUB_SHARD_PEERS", ""), ",")
	if len(parts) != total {
		logger.Errorf("HUB_SHARD_PEERS must list %d urls (got %d); sharding disabled", total, len(parts))
		return nil
	}

	sr := &shardRouter{
		index:   index,
		total:   total,
		peers:   make([]*url.URL, total),
		proxies: make([]*httputil.ReverseProxy, total),
	}
	for i, p := range parts {
		u, err := url.Parse(strings.TrimSpace(p))
		if err != nil || u.Host == "" || u.Scheme == "" {
			logger.Errorf("HUB_SHARD_PEERS[%d]=%q invalid; sharding disabled", i, p)
			return nil
		}
		sr.peers[i] = u
		if i != index {
			sr.proxies[i] = httputil.NewSingleHostReverseProxy(u)
		}
	}
	sr.fwdSecret = strings.TrimSpace(env.Str("HUB_SHARD_FWD_SECRET", ""))
	if sr.fwdSecret == "" {
		logger.Warnf("HUB_SHARD_FWD_SECRET unset: the X-Hub-Shard-Fwd loop/fallback marker is forgeable by external clients (can misroute their own writes or suppress the read fallback). Set a shared secret across all shards.")
	}
	logger.Infof("owner sharding enabled: index=%d/%d peers=%v", index, total, parts)
	return sr
}

// markForwarded stamps the forwarded-once marker on a request about to be
// reverse-proxied to a peer (the shared secret when configured, else "1").
func (sr *shardRouter) markForwarded(req *http.Request) {
	v := sr.fwdSecret
	if v == "" {
		v = "1"
	}
	req.Header.Set(shardFwdHeader, v)
}

// isForwarded reports whether this request was already forwarded once by a peer.
// With a secret configured the marker must match it (constant-time) — a forged
// or absent header is treated as not-forwarded, so the request routes normally.
// Without a secret, any non-empty value counts (legacy, forgeable).
func (sr *shardRouter) isForwarded(c *gin.Context) bool {
	got := c.GetHeader(shardFwdHeader)
	if got == "" {
		return false
	}
	if sr.fwdSecret == "" {
		return true
	}
	return subtle.ConstantTimeCompare([]byte(got), []byte(sr.fwdSecret)) == 1
}

// shardStats reports the sharding topology + write-routing counters for
// /v1/cachestats. enabled=false on a single-node hub.
func (s *Server) shardStats() gin.H {
	sr := s.shard
	if sr == nil {
		return gin.H{"enabled": false}
	}
	return gin.H{
		"enabled":        true,
		"index":          sr.index,
		"total":          sr.total,
		"local_writes":   sr.local.Load(),
		"proxied_writes": sr.proxied.Load(),
		"proxied_reads":  sr.readProxied.Load(),
	}
}

// shardOf maps an owner to its home shard. Stable + case-insensitive (owner is a
// hex address). crc32 is fine for fixed TOTAL; re-sharding on a topology change
// moves keys but writes are idempotent by da_cid and recoverable from S3+DA.
func (sr *shardRouter) shardOf(owner string) int {
	h := crc32.ChecksumIEEE([]byte(strings.ToLower(owner)))
	return int(h % uint32(sr.total))
}

// shardWrite is the write-path middleware. It MUST run after AuthMiddleware so
// the owner == authed signer. A write for a non-home owner is reverse-proxied to
// that owner's hub and the local chain aborts. No-op when sharding is disabled.
func (s *Server) shardWrite() gin.HandlerFunc {
	sr := s.shard
	return func(c *gin.Context) {
		if sr == nil {
			c.Next()
			return
		}
		owner := CtxAuthAddr(c)
		if owner == "" {
			c.Next() // unauthenticated — let the handler reject
			return
		}
		// Already forwarded once by a peer: answer locally, never re-proxy
		// (see shardFwdHeader — breaks the ping-pong on topology misconfig).
		if sr.isForwarded(c) {
			sr.local.Add(1)
			c.Next()
			return
		}
		home := sr.shardOf(owner)
		if home == sr.index || sr.proxies[home] == nil {
			sr.local.Add(1)
			c.Next()
			return
		}
		sr.proxied.Add(1)
		logger.Debugf("shard: proxying %s write for owner %s to shard %d (%s)", c.Request.Method, owner, home, sr.peers[home])
		sr.markForwarded(c.Request)
		sr.proxies[home].ServeHTTP(c.Writer, c.Request)
		c.Abort()
	}
}

// shardReadProxy forwards a READ to the owner's home shard and returns true if
// it did (the response has been written; the caller must return). Reads are
// normally served anywhere (L1/L2 cache, DA reconstruct) — but a STAGED
// object's bytes exist ONLY in the home shard's local logfs (the needle index
// is shared via PG; blobs are not, and the shared L2 is filled on read, not on
// write). Without this, a read-after-write that lands on a non-home shard 404s
// until the object is either read once on its home shard or committed to DA.
// On a local miss the content handler calls this as the fallback: single hop
// (forwarded requests are answered locally), and the home shard's own read
// populates the shared L2, so subsequent reads hit the cache on any shard.
func (s *Server) shardReadProxy(c *gin.Context, owner string) bool {
	sr := s.shard
	if sr == nil || owner == "" {
		return false
	}
	if sr.isForwarded(c) {
		return false // one hop max — answer (or 404) locally
	}
	home := sr.shardOf(owner)
	if home == sr.index || sr.proxies[home] == nil {
		return false // we ARE home (or no proxy): nothing better than local
	}
	sr.readProxied.Add(1)
	logger.Debugf("shard: proxying read for owner %s to home shard %d (%s)", owner, home, sr.peers[home])
	sr.markForwarded(c.Request)
	sr.proxies[home].ServeHTTP(c.Writer, c.Request)
	c.Abort()
	return true
}
