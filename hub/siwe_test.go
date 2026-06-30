package hub

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
)

// siweMessage is the readable EIP-4361 / SIWE text the client personal_signs.
// The hub never reconstructs it — it verifies the received bytes and only
// parses the address line + "Issued At". This mirror lets the test produce a
// valid signed envelope.
func siweMessage(addr string, issuedAt time.Time) string {
	return "hub.unibase.io wants you to sign in with your Ethereum account:\n" +
		addr + "\n\n" +
		"Sign in to Unibase Hub.\n\n" +
		"URI: https://hub.unibase.io\n" +
		"Version: 1\n" +
		"Chain ID: 97\n" +
		"Nonce: testnonce123456\n" +
		"Issued At: " + issuedAt.UTC().Format(time.RFC3339)
}

// buildSiweHeader is the SIWE counterpart of buildHeader: personal_signs the
// readable message and returns the Authorization envelope (with Msg set).
func buildSiweHeader(t *testing.T, skHex string, issuedAt time.Time) string {
	t.Helper()
	sk, err := crypto.HexToECDSA(strings.TrimPrefix(skHex, "0x"))
	if err != nil {
		t.Fatalf("HexToECDSA: %v", err)
	}
	addr := crypto.PubkeyToAddress(sk.PublicKey)
	msg := siweMessage(addr.Hex(), issuedAt)

	// EIP-191 personal_sign digest over the exact message bytes.
	pre := fmt.Sprintf("\x19Ethereum Signed Message:\n%d", len(msg))
	digest := crypto.Keccak256([]byte(pre + msg))
	sig, err := crypto.Sign(digest, sk)
	if err != nil {
		t.Fatalf("Sign: %v", err)
	}

	payload := map[string]any{
		"Type": "personal_sign",
		"Addr": addr.Hex(),
		"Time": issuedAt.Unix(),
		"Sign": "0x" + fmt.Sprintf("%x", sig),
		"Msg":  msg,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}
	return string(b)
}

func siweUploadReq(hdr string) *http.Request {
	req := httptest.NewRequest("POST", "/api/upload",
		strings.NewReader(`{"owner":"`+testAddr+`","id":"x","message":"y"}`))
	req.Header.Set("Content-Type", "application/json")
	if hdr != "" {
		req.Header.Set("Authorization", hdr)
	}
	return req
}

func TestAuthMiddleware_AcceptsSIWE(t *testing.T) {
	r := newTestRouter()
	w := httptest.NewRecorder()
	r.ServeHTTP(w, siweUploadReq(buildSiweHeader(t, testSK, time.Now().UTC())))
	if w.Code != http.StatusOK {
		t.Fatalf("want 200, got %d body=%s", w.Code, w.Body.String())
	}
}

func TestAuthMiddleware_SIWE_RejectsStale(t *testing.T) {
	r := newTestRouter()
	// default drift is 600s; 1h old is well outside the window.
	w := httptest.NewRecorder()
	r.ServeHTTP(w, siweUploadReq(buildSiweHeader(t, testSK, time.Now().UTC().Add(-time.Hour))))
	if w.Code != http.StatusUnauthorized {
		t.Fatalf("want 401, got %d body=%s", w.Code, w.Body.String())
	}
	if !strings.Contains(w.Body.String(), "out of window") {
		t.Errorf("want 'out of window', got %s", w.Body.String())
	}
}

func TestAuthMiddleware_SIWE_RejectsTamperedMessage(t *testing.T) {
	r := newTestRouter()
	hdr := buildSiweHeader(t, testSK, time.Now().UTC())

	// Tamper the signed text after signing: signature no longer recovers Addr.
	var env map[string]any
	if err := json.Unmarshal([]byte(hdr), &env); err != nil {
		t.Fatal(err)
	}
	env["Msg"] = env["Msg"].(string) + " (tampered)"
	tb, _ := json.Marshal(env)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, siweUploadReq(string(tb)))
	if w.Code != http.StatusUnauthorized {
		t.Fatalf("want 401 on tampered message, got %d body=%s", w.Code, w.Body.String())
	}
}

// TestAuthMiddleware_AcceptsLegacyPersonalSign locks backward compatibility for
// the format the CURRENTLY DEPLOYED extension sends: personal_sign over
// label||be64(ts), with NO Msg field. Must keep verifying via the legacy path.
func TestAuthMiddleware_AcceptsLegacyPersonalSign(t *testing.T) {
	r := newTestRouter()
	sk, err := crypto.HexToECDSA(testSK)
	if err != nil {
		t.Fatal(err)
	}
	addr := crypto.PubkeyToAddress(sk.PublicKey)
	label := "upload"
	ts := time.Now().Unix()

	b := make([]byte, len(label)+8)
	copy(b, label)
	binary.BigEndian.PutUint64(b[len(label):], uint64(ts))
	pre := fmt.Sprintf("\x19Ethereum Signed Message:\n%d", len(b))
	sig, err := crypto.Sign(crypto.Keccak256(append([]byte(pre), b...)), sk)
	if err != nil {
		t.Fatal(err)
	}

	payload := map[string]any{
		"Type": "personal_sign",
		"Addr": addr.Hex(),
		"Time": ts,
		"Hash": "0x" + fmt.Sprintf("%x", []byte(label)),
		"Sign": "0x" + fmt.Sprintf("%x", sig),
		// no Msg → legacy path
	}
	jb, _ := json.Marshal(payload)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, siweUploadReq(string(jb)))
	if w.Code != http.StatusOK {
		t.Fatalf("legacy personal_sign want 200, got %d body=%s", w.Code, w.Body.String())
	}
}

// TestAuthMiddleware_SIWE_RejectsAddrSwap proves the envelope Addr can't be
// swapped to a victim: VerifySIWE recovers the real signer from Msg and it must
// equal Addr, so claiming a different Addr fails.
func TestAuthMiddleware_SIWE_RejectsAddrSwap(t *testing.T) {
	r := newTestRouter()
	hdr := buildSiweHeader(t, testSK, time.Now().UTC())

	var env map[string]any
	_ = json.Unmarshal([]byte(hdr), &env)
	env["Addr"] = "0x000000000000000000000000000000000000dEaD"
	tb, _ := json.Marshal(env)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, siweUploadReq(string(tb)))
	if w.Code != http.StatusUnauthorized {
		t.Fatalf("want 401 on Addr swap, got %d body=%s", w.Code, w.Body.String())
	}
}
