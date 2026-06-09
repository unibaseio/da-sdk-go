package hub

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"crypto/sha256"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
)

// Set CHAIN_TYPE so the indirect import of sdk (which init()s on it) doesn't
// panic. We never touch the chain in these tests.
func init() {
	if os.Getenv("CHAIN_TYPE") == "" {
		os.Setenv("CHAIN_TYPE", "bnb-testnet-dao")
	}
	gin.SetMode(gin.TestMode)
}

// buildHeader is the test-side mirror of sdk/lib/key.BuildAuth. It produces the
// exact Authorization header string that AuthMiddleware expects.
func buildHeader(t *testing.T, skHex string, label string, ts int64) string {
	t.Helper()
	sk, err := crypto.HexToECDSA(strings.TrimPrefix(skHex, "0x"))
	if err != nil {
		t.Fatalf("HexToECDSA: %v", err)
	}
	addr := crypto.PubkeyToAddress(sk.PublicKey)

	h := sha256.New()
	h.Write([]byte(label))
	tsb := make([]byte, 8)
	binary.BigEndian.PutUint64(tsb, uint64(ts))
	h.Write(tsb)
	digest := h.Sum(nil)

	sig, err := crypto.Sign(digest, sk)
	if err != nil {
		t.Fatalf("Sign: %v", err)
	}

	payload := map[string]any{
		"Type": "",
		"Addr": addr.Hex(),
		"Time": ts,
		"Hash": "0x" + fmt.Sprintf("%x", []byte(label)),
		"Sign": "0x" + fmt.Sprintf("%x", sig),
	}
	b, err := json.Marshal(payload)
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}
	return string(b)
}

func newTestRouter() *gin.Engine {
	r := gin.New()
	g := r.Group("/api")
	g.Use(MaxBodySize())
	g.Use(AuthMiddleware())
	g.Use(RateLimit())

	// fake info endpoint (bypassed)
	g.GET("/info", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"ok": true})
	})

	// fake upload that exercises RequireOwnerMatch
	g.POST("/upload", func(c *gin.Context) {
		var body struct {
			Owner   string `json:"owner"`
			ID      string `json:"id"`
			Message string `json:"message"`
		}
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(599, gin.H{"err": err.Error()})
			return
		}
		if !RequireOwnerMatch(c, body.Owner) {
			return
		}
		c.JSON(http.StatusOK, gin.H{"ok": true, "owner": body.Owner})
	})

	// fake download that exercises ResolveOwnerForList
	g.POST("/download", func(c *gin.Context) {
		owner, ok := ResolveOwnerForList(c, c.PostForm("owner"))
		if !ok {
			return
		}
		c.JSON(http.StatusOK, gin.H{"ok": true, "owner": owner})
	})
	return r
}

const testSK = "0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f20"
const testAddr = "0x6370eF2f4Db3611D657b90667De398a2Cc2a370C"

func TestAuthMiddleware_RejectsNoHeader(t *testing.T) {
	r := newTestRouter()
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/upload",
		strings.NewReader(`{"owner":"`+testAddr+`","id":"x","message":"y"}`))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	if w.Code != http.StatusUnauthorized {
		t.Fatalf("want 401, got %d body=%s", w.Code, w.Body.String())
	}
	if !strings.Contains(w.Body.String(), "missing Authorization") {
		t.Errorf("want 'missing Authorization' in body, got %s", w.Body.String())
	}
}

func TestAuthMiddleware_AcceptsValidSignature(t *testing.T) {
	r := newTestRouter()
	hdr := buildHeader(t, testSK, "upload", time.Now().Unix())

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/upload",
		strings.NewReader(`{"owner":"`+testAddr+`","id":"x","message":"y"}`))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", hdr)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("want 200, got %d body=%s", w.Code, w.Body.String())
	}
}

func TestAuthMiddleware_RejectsStaleTimestamp(t *testing.T) {
	r := newTestRouter()
	// 1 hour ago — outside the default 10 min window
	hdr := buildHeader(t, testSK, "upload", time.Now().Add(-1*time.Hour).Unix())

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/upload",
		strings.NewReader(`{"owner":"`+testAddr+`","id":"x","message":"y"}`))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", hdr)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Fatalf("want 401, got %d body=%s", w.Code, w.Body.String())
	}
	if !strings.Contains(w.Body.String(), "out of window") {
		t.Errorf("want 'out of window', got %s", w.Body.String())
	}
}

func TestOwnerMatch_RejectsNonHexOwner(t *testing.T) {
	r := newTestRouter()
	hdr := buildHeader(t, testSK, "upload", time.Now().Unix())

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/upload",
		strings.NewReader(`{"owner":"noah-2026","id":"x","message":"y"}`))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", hdr)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Fatalf("want 401, got %d body=%s", w.Code, w.Body.String())
	}
	if !strings.Contains(w.Body.String(), "Ethereum address") {
		t.Errorf("want 'Ethereum address' err, got %s", w.Body.String())
	}
}

func TestOwnerMatch_RejectsOtherAddress(t *testing.T) {
	r := newTestRouter()
	hdr := buildHeader(t, testSK, "upload", time.Now().Unix())

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/upload",
		strings.NewReader(`{"owner":"0x0000000000000000000000000000000000000001","id":"x","message":"y"}`))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", hdr)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Fatalf("want 401, got %d body=%s", w.Code, w.Body.String())
	}
	if !strings.Contains(w.Body.String(), "does not match signer") {
		t.Errorf("want 'does not match signer', got %s", w.Body.String())
	}
}

func TestResolveOwnerForList_DefaultsToSigner(t *testing.T) {
	r := newTestRouter()
	hdr := buildHeader(t, testSK, "download", time.Now().Unix())

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/download",
		strings.NewReader("")) // no owner field
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", hdr)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("want 200, got %d body=%s", w.Code, w.Body.String())
	}
	if !strings.Contains(strings.ToLower(w.Body.String()), strings.ToLower(testAddr)) {
		t.Errorf("expected signer addr in response, got %s", w.Body.String())
	}
}

// newPublicTestRouter mirrors the public, read-only /api group in server.go:
// body-size cap + rate limit, but NO AuthMiddleware.
func newPublicTestRouter() *gin.Engine {
	r := gin.New()
	g := r.Group("/api")
	g.Use(MaxBodySize())
	g.Use(RateLimit())

	// fake list endpoint that exercises ResolveOwnerForList on the public path
	g.GET("/listBucket", func(c *gin.Context) {
		owner, ok := ResolveOwnerForList(c, c.Query("owner"))
		if !ok {
			return
		}
		c.JSON(http.StatusOK, gin.H{"ok": true, "owner": owner})
	})
	return r
}

const otherAddr = "0x1111111111111111111111111111111111111111"

func TestPublicList_NoAuthAcceptsExplicitOwner(t *testing.T) {
	r := newPublicTestRouter()
	w := httptest.NewRecorder()
	// No Authorization header, explicit owner (someone else's) — must succeed.
	req := httptest.NewRequest("GET", "/api/listBucket?owner="+otherAddr, nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("want 200 on public list, got %d body=%s", w.Code, w.Body.String())
	}
	if !strings.Contains(strings.ToLower(w.Body.String()), strings.ToLower(otherAddr)) {
		t.Errorf("expected owner addr in response, got %s", w.Body.String())
	}
}

func TestPublicList_NoAuthRequiresOwner(t *testing.T) {
	r := newPublicTestRouter()
	w := httptest.NewRecorder()
	// No auth and no owner: there's no signer to default to → 400, not 401.
	req := httptest.NewRequest("GET", "/api/listBucket", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("want 400 when owner missing on public list, got %d body=%s", w.Code, w.Body.String())
	}
}

func TestPublicList_NoAuthRejectsBadOwner(t *testing.T) {
	r := newPublicTestRouter()
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/api/listBucket?owner=not-an-address", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("want 400 on malformed owner, got %d body=%s", w.Code, w.Body.String())
	}
}

func TestInfoBypass(t *testing.T) {
	r := newTestRouter()
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/api/info", nil)
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("want 200 on /api/info (bypass), got %d body=%s", w.Code, w.Body.String())
	}
}

func TestMaxBodySize_RejectsHugePayload(t *testing.T) {
	t.Setenv("HUB_MAX_JSON_BYTES", "1024") // 1 KiB cap for this test
	r := newTestRouter()
	hdr := buildHeader(t, testSK, "upload", time.Now().Unix())

	big := bytes.Repeat([]byte("a"), 8192)
	body := `{"owner":"` + testAddr + `","id":"x","message":"` + string(big) + `"}`

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/upload", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", hdr)
	r.ServeHTTP(w, req)

	// gin's ShouldBindJSON over a truncated body returns a 599 from our handler;
	// the important thing is it doesn't succeed as 200.
	if w.Code == http.StatusOK {
		t.Fatalf("expected non-200 (body too large), got 200 body=%s", w.Body.String())
	}
}

func TestRateLimit_PerOwnerKicks(t *testing.T) {
	t.Setenv("HUB_RATE_OWNER_RPS", "1")
	t.Setenv("HUB_RATE_OWNER_BURST", "2")
	t.Setenv("HUB_RATE_IP_RPS", "1000") // make IP limiter effectively no-op
	t.Setenv("HUB_RATE_IP_BURST", "1000")
	r := newTestRouter()
	hdr := buildHeader(t, testSK, "upload", time.Now().Unix())

	hits := 0
	for i := 0; i < 6; i++ {
		w := httptest.NewRecorder()
		req := httptest.NewRequest("POST", "/api/upload",
			strings.NewReader(`{"owner":"`+testAddr+`","id":"x","message":"y"}`))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", hdr)
		r.ServeHTTP(w, req)
		if w.Code == http.StatusTooManyRequests {
			hits++
		}
	}
	if hits == 0 {
		t.Fatalf("expected at least one 429 in 6 rapid requests")
	}
}
