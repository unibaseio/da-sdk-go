package hub

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"

	"github.com/unibaseio/da-sdk-go/lib/types"
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

// The tests below run against the real /v1 router (registV1 on an in-memory
// sqlite Server, see newV1TestServer): the middleware order, body caps and
// owner checks are the ones the hub binary serves.

const testSK = "0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f20"
const testAddr = "0x6370eF2f4Db3611D657b90667De398a2Cc2a370C"
const otherAddr = "0x1111111111111111111111111111111111111111"

// signedWrite is a signed /v1 write that needs no LogFS or chain: create the
// bucket name.
func signedWrite(hdr, bucket string) *http.Request {
	req := httptest.NewRequest("PUT", "/v1/buckets/"+bucket, strings.NewReader(`{"kind":"memory"}`))
	req.Header.Set("Content-Type", "application/json")
	if hdr != "" {
		req.Header.Set("Authorization", hdr)
	}
	return req
}

// sealReq is a signed POST /v1/seal carrying only the owner field: seal checks
// owner==signer before anything else, so the owner check is all that runs.
func sealReq(t *testing.T, hdr, owner string) *http.Request {
	t.Helper()
	var body bytes.Buffer
	mw := multipart.NewWriter(&body)
	if err := mw.WriteField("owner", owner); err != nil {
		t.Fatal(err)
	}
	mw.Close()
	req := httptest.NewRequest("POST", "/v1/seal", &body)
	req.Header.Set("Content-Type", mw.FormDataContentType())
	req.Header.Set("Authorization", hdr)
	return req
}

func serve(s *Server, req *http.Request) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	s.Router.ServeHTTP(w, req)
	return w
}

func TestAuthMiddleware_RejectsNoHeader(t *testing.T) {
	s := newV1TestServer(t)
	w := serve(s, signedWrite("", "b"))
	if w.Code != http.StatusUnauthorized {
		t.Fatalf("want 401, got %d body=%s", w.Code, w.Body.String())
	}
	if !strings.Contains(w.Body.String(), "missing Authorization") {
		t.Errorf("want 'missing Authorization' in body, got %s", w.Body.String())
	}
}

func TestAuthMiddleware_AcceptsValidSignature(t *testing.T) {
	s := newV1TestServer(t)
	w := serve(s, signedWrite(buildHeader(t, testSK, "hub", time.Now().Unix()), "b"))
	if w.Code != http.StatusCreated {
		t.Fatalf("want 201, got %d body=%s", w.Code, w.Body.String())
	}
	// the bucket belongs to the signer (lowercase canonical form)
	if !strings.Contains(w.Body.String(), `"owner":"`+strings.ToLower(testAddr)+`"`) {
		t.Errorf("want signer as owner, got %s", w.Body.String())
	}
}

func TestAuthMiddleware_RejectsStaleTimestamp(t *testing.T) {
	s := newV1TestServer(t)
	// 1 hour ago — outside the default 10 min window
	w := serve(s, signedWrite(buildHeader(t, testSK, "hub", time.Now().Add(-time.Hour).Unix()), "b"))
	if w.Code != http.StatusUnauthorized {
		t.Fatalf("want 401, got %d body=%s", w.Code, w.Body.String())
	}
	if !strings.Contains(w.Body.String(), "out of window") {
		t.Errorf("want 'out of window', got %s", w.Body.String())
	}
}

func TestOwnerMatch_RejectsNonHexOwner(t *testing.T) {
	s := newV1TestServer(t)
	w := serve(s, sealReq(t, buildHeader(t, testSK, "hub", time.Now().Unix()), "noah-2026"))
	if w.Code != http.StatusUnauthorized {
		t.Fatalf("want 401, got %d body=%s", w.Code, w.Body.String())
	}
	if !strings.Contains(w.Body.String(), "Ethereum address") {
		t.Errorf("want 'Ethereum address' err, got %s", w.Body.String())
	}
}

func TestOwnerMatch_RejectsOtherAddress(t *testing.T) {
	s := newV1TestServer(t)
	w := serve(s, sealReq(t, buildHeader(t, testSK, "hub", time.Now().Unix()), otherAddr))
	if w.Code != http.StatusUnauthorized {
		t.Fatalf("want 401, got %d body=%s", w.Code, w.Body.String())
	}
	if !strings.Contains(w.Body.String(), "does not match signer") {
		t.Errorf("want 'does not match signer', got %s", w.Body.String())
	}
}

// The signer's own address passes the owner check whatever its case, and the
// request reaches seal's own validation.
func TestOwnerMatch_AcceptsSigner(t *testing.T) {
	s := newV1TestServer(t)
	for _, owner := range []string{testAddr, strings.ToLower(testAddr)} {
		w := serve(s, sealReq(t, buildHeader(t, testSK, "hub", time.Now().Unix()), owner))
		if w.Code != http.StatusBadRequest || !strings.Contains(w.Body.String(), "rsn/rsk") {
			t.Fatalf("owner %s: want 400 from seal validation, got %d body=%s", owner, w.Code, w.Body.String())
		}
	}
}

// Enumeration without an owner is scoped to the signer, and naming another
// owner is refused.
func TestListDefaultsToSigner(t *testing.T) {
	s := newV1TestServer(t)
	s.gdb.Create(&types.Bucket{Name: "mine", Owner: strings.ToLower(testAddr), Kind: "memory"})
	s.gdb.Create(&types.Bucket{Name: "theirs", Owner: strings.ToLower(otherAddr), Kind: "memory"})
	hdr := buildHeader(t, testSK, "hub", time.Now().Unix())

	req := httptest.NewRequest("GET", "/v1/buckets", nil)
	req.Header.Set("Authorization", hdr)
	w := serve(s, req)
	if w.Code != http.StatusOK {
		t.Fatalf("want 200, got %d body=%s", w.Code, w.Body.String())
	}
	if !strings.Contains(w.Body.String(), `"mine"`) || strings.Contains(w.Body.String(), `"theirs"`) {
		t.Fatalf("listing not scoped to the signer: %s", w.Body.String())
	}

	req = httptest.NewRequest("GET", "/v1/buckets?owner="+otherAddr, nil)
	req.Header.Set("Authorization", hdr)
	if w := serve(s, req); w.Code != http.StatusUnauthorized || !strings.Contains(w.Body.String(), "does not match signer") {
		t.Fatalf("listing another owner: want 401, got %d body=%s", w.Code, w.Body.String())
	}
}

// Public point reads (ResolveOwnerForList without a signer): owner is
// optional, any valid owner may be named, case doesn't matter, junk is a 400.
func newPublicReadServer(t *testing.T) *Server {
	s := newV1TestServer(t)
	s.gdb.Create(&types.Bucket{Name: "pub", Owner: strings.ToLower(testAddr), Kind: "memory"})
	s.gdb.Create(&types.Needle{Owner: strings.ToLower(testAddr), Bucket: "pub", Name: "k", File: 1, Size: 3})
	return s
}

func getStatus(s *Server, path string) (int, string) {
	w := serve(s, httptest.NewRequest("GET", path, nil))
	return w.Code, w.Body.String()
}

func TestPublicRead_NoAuthAcceptsExplicitOwner(t *testing.T) {
	s := newPublicReadServer(t)
	if code, body := getStatus(s, "/v1/buckets/pub/objects/k?owner="+strings.ToLower(testAddr)); code != http.StatusOK {
		t.Fatalf("want 200 on public read, got %d body=%s", code, body)
	}
	// another owner is a valid scope too: not found, but not refused
	if code, body := getStatus(s, "/v1/buckets/pub/objects/k?owner="+otherAddr); code != http.StatusNotFound {
		t.Fatalf("want 404 for another owner's scope, got %d body=%s", code, body)
	}
}

func TestPublicRead_CanonicalizesOwner(t *testing.T) {
	// Ethereum addresses are case-insensitive (EIP-55 case is just a UI
	// checksum): the EIP-55 form finds rows stored under the lowercase form.
	s := newPublicReadServer(t)
	if code, body := getStatus(s, "/v1/buckets/pub/objects/k?owner="+testAddr); code != http.StatusOK {
		t.Fatalf("want 200 for the EIP-55 owner form, got %d body=%s", code, body)
	}
}

func TestPublicRead_NoAuthNoOwner(t *testing.T) {
	s := newPublicReadServer(t)
	if code, body := getStatus(s, "/v1/buckets/pub/objects/k"); code != http.StatusOK {
		t.Fatalf("want 200 (unscoped) when owner is omitted, got %d body=%s", code, body)
	}
}

func TestPublicRead_NoAuthRejectsBadOwner(t *testing.T) {
	s := newPublicReadServer(t)
	if code, body := getStatus(s, "/v1/buckets/pub/objects/k?owner=not-an-address"); code != http.StatusBadRequest {
		t.Fatalf("want 400 on malformed owner, got %d body=%s", code, body)
	}
}

func TestInfoIsPublic(t *testing.T) {
	s := newV1TestServer(t)
	if code, body := getStatus(s, "/v1/info"); code != http.StatusOK {
		t.Fatalf("want 200 on /v1/info without auth, got %d body=%s", code, body)
	}
}

func TestWriteBodyCap(t *testing.T) {
	t.Setenv("HUB_MAX_MULTIPART_BYTES", "1024") // read when the write group is mounted
	s := newV1TestServer(t)
	hdr := buildHeader(t, testSK, "hub", time.Now().Unix())
	if w := serve(s, signedWrite(hdr, "b")); w.Code != http.StatusCreated {
		t.Fatalf("create bucket: got %d body=%s", w.Code, w.Body.String())
	}

	req := httptest.NewRequest("PUT", "/v1/buckets/b/objects/k", bytes.NewReader(bytes.Repeat([]byte("a"), 8192)))
	req.Header.Set("Authorization", hdr)
	w := serve(s, req)
	if w.Code >= 200 && w.Code < 300 {
		t.Fatalf("expected the oversized body to be refused, got %d body=%s", w.Code, w.Body.String())
	}
	if !strings.Contains(w.Body.String(), "too large") {
		t.Errorf("want 'too large', got %s", w.Body.String())
	}
}

func TestRateLimit_PerIPKicks(t *testing.T) {
	t.Setenv("HUB_RATE_IP_RPS", "1")
	t.Setenv("HUB_RATE_IP_BURST", "3")
	s := newV1TestServer(t)

	hits429 := 0
	for i := 0; i < 10; i++ {
		req := httptest.NewRequest("GET", "/v1/info", nil)
		req.RemoteAddr = "203.0.113.7:1234" // same IP each time → shared bucket
		if serve(s, req).Code == http.StatusTooManyRequests {
			hits429++
		}
	}
	if hits429 == 0 {
		t.Fatalf("expected at least one 429 from a single IP doing 10 rapid requests (burst=3)")
	}
}

// The per-owner tier keys on the recovered signer, so one wallet can't escape
// it by spreading writes over many IPs.
func TestRateLimit_PerOwnerKicks(t *testing.T) {
	t.Setenv("HUB_RATE_OWNER_RPS", "0.001")
	t.Setenv("HUB_RATE_OWNER_BURST", "1")
	s := newV1TestServer(t)
	hdr := buildHeader(t, testSK, "hub", time.Now().Unix())

	codes := map[int]int{}
	for i := 0; i < 3; i++ {
		req := signedWrite(hdr, fmt.Sprintf("b%d", i))
		req.RemoteAddr = fmt.Sprintf("203.0.113.%d:1", 10+i)
		codes[serve(s, req).Code]++
	}
	if codes[http.StatusCreated] != 1 || codes[http.StatusTooManyRequests] != 2 {
		t.Fatalf("want 1x201 then 429s, got %v", codes)
	}
}
