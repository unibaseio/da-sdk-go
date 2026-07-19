package hub

import (
	"hash/crc32"
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
}

// shardFwdHeader marks a request already forwarded once by a shard peer. The
// receiving hub must answer locally and never re-proxy: home resolution is
// deterministic, so a second hop only happens under topology misconfig (e.g.
// two hubs disagreeing on TOTAL) — without this guard that would ping-pong
// forever. One hop max, then serve or fail locally.
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
	logger.Infof("owner sharding enabled: index=%d/%d peers=%v", index, total, parts)
	return sr
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
		if c.GetHeader(shardFwdHeader) != "" {
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
		c.Request.Header.Set(shardFwdHeader, "1")
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
	if c.GetHeader(shardFwdHeader) != "" {
		return false // one hop max — answer (or 404) locally
	}
	home := sr.shardOf(owner)
	if home == sr.index || sr.proxies[home] == nil {
		return false // we ARE home (or no proxy): nothing better than local
	}
	sr.readProxied.Add(1)
	logger.Debugf("shard: proxying read for owner %s to home shard %d (%s)", owner, home, sr.peers[home])
	c.Request.Header.Set(shardFwdHeader, "1")
	sr.proxies[home].ServeHTTP(c.Writer, c.Request)
	c.Abort()
	return true
}
