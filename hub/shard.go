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

	proxied atomic.Int64 // writes forwarded to a peer
	local   atomic.Int64 // writes served locally (home shard)
}

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
		home := sr.shardOf(owner)
		if home == sr.index || sr.proxies[home] == nil {
			sr.local.Add(1)
			c.Next()
			return
		}
		sr.proxied.Add(1)
		logger.Debugf("shard: proxying %s write for owner %s to shard %d (%s)", c.Request.Method, owner, home, sr.peers[home])
		sr.proxies[home].ServeHTTP(c.Writer, c.Request)
		c.Abort()
	}
}
