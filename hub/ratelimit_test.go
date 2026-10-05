package hub

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

// newIPEngine mounts IPRateLimit (burst 2, ~no refill) on a gin engine with the
// hub's trusted-proxy policy and returns a helper that reports the status for a
// request from remote with an optional X-Forwarded-For.
func newIPEngine(t *testing.T) func(remote, xff string) int {
	t.Helper()
	t.Setenv("HUB_RATE_IP_RPS", "0.001")
	t.Setenv("HUB_RATE_IP_BURST", "2")
	gin.SetMode(gin.TestMode)
	r := gin.New()
	if err := configureClientIP(r); err != nil {
		t.Fatal(err)
	}
	r.Use(IPRateLimit())
	r.GET("/x", func(c *gin.Context) { c.String(http.StatusOK, c.ClientIP()) })
	return func(remote, xff string) int {
		req := httptest.NewRequest("GET", "/x", nil)
		req.RemoteAddr = remote
		if xff != "" {
			req.Header.Set("X-Forwarded-For", xff)
		}
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		return w.Code
	}
}

// A client talking to the hub directly cannot mint fresh buckets by forging
// X-Forwarded-For: its public RemoteAddr is not a trusted proxy.
func TestIPRateLimitIgnoresForgedXFF(t *testing.T) {
	hit := newIPEngine(t)
	codes := []int{
		hit("203.0.113.7:1000", "1.1.1.1"),
		hit("203.0.113.7:1000", "2.2.2.2"),
		hit("203.0.113.7:1000", "3.3.3.3"),
	}
	if codes[2] != http.StatusTooManyRequests {
		t.Fatalf("forged XFF bypassed the per-IP limit: %v", codes)
	}
}

// Behind the ALB (a private peer) the real client from XFF is used, so two
// users behind the same ALB get separate buckets — and a client-supplied XFF
// prefix is skipped in favor of the hop the ALB appended.
func TestIPRateLimitHonorsTrustedProxy(t *testing.T) {
	hit := newIPEngine(t)
	alb := "10.0.21.5:443"
	for i := 0; i < 2; i++ {
		if c := hit(alb, "198.51.100.1"); c != http.StatusOK {
			t.Fatalf("user A request %d: %d", i, c)
		}
	}
	if c := hit(alb, "198.51.100.1"); c != http.StatusTooManyRequests {
		t.Fatalf("user A not limited: %d", c)
	}
	// user B behind the same ALB is unaffected
	if c := hit(alb, "198.51.100.2"); c != http.StatusOK {
		t.Fatalf("user B throttled with A: %d", c)
	}
	// user A forging a prefix still lands in A's bucket (ALB appends the real IP)
	if c := hit(alb, "9.9.9.9, 198.51.100.1"); c != http.StatusTooManyRequests {
		t.Fatalf("forged prefix escaped the limit: %d", c)
	}
}

func TestTrustedProxiesNone(t *testing.T) {
	t.Setenv("HUB_TRUSTED_PROXIES", "none")
	hit := newIPEngine(t)
	hit("10.0.21.5:443", "198.51.100.1")
	hit("10.0.21.5:443", "198.51.100.2")
	if c := hit("10.0.21.5:443", "198.51.100.3"); c != http.StatusTooManyRequests {
		t.Fatalf("trust-none still read XFF: %d", c)
	}
}

func TestTrustedProxiesInvalid(t *testing.T) {
	t.Setenv("HUB_TRUSTED_PROXIES", "not-a-cidr")
	if err := configureClientIP(gin.New()); err == nil {
		t.Fatal("invalid HUB_TRUSTED_PROXIES accepted")
	}
}

// The per-IP tier on the write group runs before signature verification: a
// flood of junk-signed writes gets 429, not a signature check each.
func TestV1WriteIPLimitBeforeAuth(t *testing.T) {
	t.Setenv("HUB_RATE_IP_RPS", "0.001")
	t.Setenv("HUB_RATE_IP_BURST", "1")
	s := newV1TestServer(t)
	codes := map[int]int{}
	for i := 0; i < 3; i++ {
		req := httptest.NewRequest("PUT", "/v1/buckets/b", nil)
		req.RemoteAddr = "203.0.113.9:1"
		req.Header.Set("Authorization", "junk")
		w := httptest.NewRecorder()
		s.Router.ServeHTTP(w, req)
		codes[w.Code]++
	}
	if codes[http.StatusUnauthorized] != 1 || codes[http.StatusTooManyRequests] != 2 {
		t.Fatalf("want 1x401 then 429s, got %v", codes)
	}
}

func TestLimiterRegistryEvictsIdle(t *testing.T) {
	lr := newLimiterRegistry(10, 20) // full refill in 2s → idle floor of 1 min
	now := time.Unix(1000, 0)
	lr.now = func() time.Time { return now }
	for _, k := range []string{"a", "b", "c"} {
		lr.get(k)
	}
	now = now.Add(30 * time.Second)
	lr.get("a") // touch a
	now = now.Add(45 * time.Second)
	lr.get("d") // triggers a sweep: b, c idle 75s > 60s; a idle 45s
	if n := lr.size(); n != 2 {
		t.Fatalf("want a+d after sweep, got %d entries", n)
	}
}

func TestLimiterRegistryBounded(t *testing.T) {
	lr := newLimiterRegistry(1, 1)
	lr.max = 3
	for _, k := range []string{"a", "b", "c", "d", "e"} {
		lr.get(k)
	}
	if n := lr.size(); n != 3 {
		t.Fatalf("registry grew past its cap: %d", n)
	}
	if lr.get("e") != lr.overflow {
		t.Fatal("new key past the cap should share the overflow bucket")
	}
}

func TestIPKeyIPv6Slash64(t *testing.T) {
	if ipKey("2001:db8:1:2::1") != ipKey("2001:db8:1:2:ffff::9") {
		t.Fatal("same /64 should share a key")
	}
	if ipKey("2001:db8:1:2::1") == ipKey("2001:db8:1:3::1") {
		t.Fatal("different /64s should not share a key")
	}
	if ipKey("203.0.113.7") != "203.0.113.7" {
		t.Fatal("IPv4 key changed")
	}
}
