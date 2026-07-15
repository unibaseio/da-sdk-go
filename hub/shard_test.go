package hub

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
)

func testShardRouter(t *testing.T, index, total int, peers string) *shardRouter {
	t.Helper()
	t.Setenv("HUB_SHARD_TOTAL", strconv.Itoa(total))
	t.Setenv("HUB_SHARD_INDEX", strconv.Itoa(index))
	t.Setenv("HUB_SHARD_PEERS", peers)
	return newShardRouter()
}

// TestShardRouterDisabledByDefault: TOTAL<=1 or misconfig => nil (single-node).
func TestShardRouterDisabledByDefault(t *testing.T) {
	if sr := testShardRouter(t, 0, 1, ""); sr != nil {
		t.Fatal("TOTAL=1 must disable sharding")
	}
	if sr := testShardRouter(t, 5, 2, "http://a,http://b"); sr != nil {
		t.Fatal("out-of-range index must disable sharding")
	}
	if sr := testShardRouter(t, 0, 2, "http://only-one"); sr != nil {
		t.Fatal("peer count != total must disable sharding")
	}
	if sr := testShardRouter(t, 0, 2, "bad-url,http://b"); sr != nil {
		t.Fatal("invalid peer url must disable sharding")
	}
}

// TestShardOfStableAndInRange: shardOf is deterministic, case-insensitive, and
// always within [0,total).
func TestShardOfStableAndInRange(t *testing.T) {
	sr := testShardRouter(t, 0, 3, "http://a,http://b,http://c")
	if sr == nil {
		t.Fatal("expected an enabled router")
	}
	owner := "0xAbC123"
	first := sr.shardOf(owner)
	for i := 0; i < 100; i++ {
		if sr.shardOf(owner) != first {
			t.Fatal("shardOf not deterministic")
		}
	}
	if sr.shardOf("0xabc123") != first {
		t.Fatal("shardOf must be case-insensitive")
	}
	for _, o := range []string{"0x1", "0x2", "0xdeadbeef", "0xFEED"} {
		if s := sr.shardOf(o); s < 0 || s >= 3 {
			t.Fatalf("shardOf(%s)=%d out of range", o, s)
		}
	}
}

// TestShardWriteProxiesNonHome: a write whose owner isn't homed here is
// transparently proxied to the peer; a home-owner write runs the local handler.
// Both hubs are real httptest servers so the reverse proxy exercises real HTTP.
func TestShardWriteProxiesNonHome(t *testing.T) {
	gin.SetMode(gin.TestMode)

	var peerHit bool
	peer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		peerHit = true
		w.WriteHeader(http.StatusTeapot) // distinctive
		_, _ = w.Write([]byte("from-peer"))
	}))
	defer peer.Close()

	// this hub is shard 0; peer is shard 1
	sr := testShardRouter(t, 0, 2, "http://127.0.0.1:1,"+peer.URL)
	if sr == nil {
		t.Fatal("expected enabled router")
	}
	s := &Server{shard: sr}

	// owner comes from a test header (simulating what AuthMiddleware would set)
	var localHit bool
	r := gin.New()
	r.POST("/v1/x",
		func(c *gin.Context) { c.Set(ctxAuthAddrKey, c.GetHeader("X-Test-Owner")) },
		s.shardWrite(),
		func(c *gin.Context) { localHit = true; c.String(http.StatusOK, "from-local") },
	)
	local := httptest.NewServer(r)
	defer local.Close()

	// pick owners that hash to shard 0 (home) and shard 1 (away)
	var homeOwner, awayOwner string
	for _, o := range []string{"0x01", "0x02", "0x03", "0x04", "0x05", "0x06", "0x07", "0x08"} {
		switch sr.shardOf(o) {
		case 0:
			if homeOwner == "" {
				homeOwner = o
			}
		case 1:
			if awayOwner == "" {
				awayOwner = o
			}
		}
	}
	if homeOwner == "" || awayOwner == "" {
		t.Fatal("could not find owners for both shards")
	}

	post := func(owner string) *http.Response {
		req, _ := http.NewRequest(http.MethodPost, local.URL+"/v1/x", nil)
		req.Header.Set("X-Test-Owner", owner)
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("post: %v", err)
		}
		return resp
	}

	// away → proxied to peer
	resp := post(awayOwner)
	resp.Body.Close()
	if !peerHit || localHit {
		t.Fatalf("away write should proxy (peerHit=%v localHit=%v)", peerHit, localHit)
	}
	if resp.StatusCode != http.StatusTeapot {
		t.Fatalf("expected peer status echoed, got %d", resp.StatusCode)
	}
	if sr.proxied.Load() != 1 {
		t.Fatalf("proxied = %d, want 1", sr.proxied.Load())
	}

	// home → local handler
	peerHit, localHit = false, false
	resp = post(homeOwner)
	resp.Body.Close()
	if peerHit || !localHit {
		t.Fatalf("home write should run local (peerHit=%v localHit=%v)", peerHit, localHit)
	}
	if sr.local.Load() != 1 {
		t.Fatalf("local = %d, want 1", sr.local.Load())
	}
}
