package hub

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

func TestHTTPServerTimeouts(t *testing.T) {
	srv := newHTTPServer(":0", http.NotFoundHandler())
	if srv.ReadHeaderTimeout != 10*time.Second || srv.IdleTimeout != 120*time.Second ||
		srv.ReadTimeout != 10*time.Minute {
		t.Fatalf("unexpected defaults: %+v", srv)
	}
	// no write deadline by default: large downloads and seal must not be cut off
	if srv.WriteTimeout != 0 {
		t.Fatalf("WriteTimeout = %s, want 0", srv.WriteTimeout)
	}
	t.Setenv("HUB_HTTP_WRITE_TIMEOUT_SEC", "30")
	if srv := newHTTPServer(":0", nil); srv.WriteTimeout != 30*time.Second {
		t.Fatalf("env override ignored: %s", srv.WriteTimeout)
	}
}

// A peer that accepts the connection but never answers no longer pins the
// forwarded request: the shard transport gives up after the header timeout.
func TestShardProxyHungPeerTimesOut(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Setenv("HUB_SHARD_PROXY_HEADER_TIMEOUT_SEC", "1")
	release := make(chan struct{})
	peer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		<-release
	}))
	defer peer.Close()
	defer close(release)

	sr := testShardRouter(t, 0, 2, "http://127.0.0.1:1,"+peer.URL)
	if sr == nil {
		t.Fatal("expected enabled router")
	}
	s := &Server{shard: sr}
	r := gin.New()
	owner := ownerHomedAt(t, sr, 1)
	r.Use(func(c *gin.Context) { c.Set(ctxAuthAddrKey, owner); c.Next() })
	r.Use(s.shardWrite())
	r.PUT("/w", func(c *gin.Context) { c.Status(http.StatusOK) })

	front := httptest.NewServer(r)
	defer front.Close()

	start := time.Now()
	req, _ := http.NewRequest("PUT", front.URL+"/w", nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	resp.Body.Close()
	if resp.StatusCode != http.StatusBadGateway {
		t.Fatalf("got %d, want 502 from the timed-out proxy", resp.StatusCode)
	}
	if d := time.Since(start); d > 5*time.Second {
		t.Fatalf("proxy waited %s on a hung peer", d)
	}
}

// ownerHomedAt finds an owner address whose home shard is idx.
func ownerHomedAt(t *testing.T, sr *shardRouter, idx int) string {
	t.Helper()
	for i := 0; i < 10000; i++ {
		o := "0x" + time.Duration(i).String() + "aa"
		if sr.shardOf(o) == idx {
			return o
		}
	}
	t.Fatal("no owner found")
	return ""
}
