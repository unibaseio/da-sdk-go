package hub

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
)

// msgHeader personal_signs msg (with "{addr}" replaced by the signer) and
// returns the Authorization envelope.
func msgHeader(t *testing.T, msg string) string {
	t.Helper()
	sk, err := crypto.HexToECDSA(testSK)
	if err != nil {
		t.Fatal(err)
	}
	addr := crypto.PubkeyToAddress(sk.PublicKey)
	msg = strings.ReplaceAll(msg, "{addr}", addr.Hex())
	digest := crypto.Keccak256([]byte(fmt.Sprintf("\x19Ethereum Signed Message:\n%d", len(msg)) + msg))
	sig, err := crypto.Sign(digest, sk)
	if err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(map[string]any{
		"Type": "personal_sign", "Addr": addr.Hex(), "Time": time.Now().Unix(),
		"Sign": fmt.Sprintf("0x%x", sig), "Msg": msg,
	})
	return string(b)
}

func siweText(domain, uri, chain string) string {
	return domain + " wants you to sign in with your Ethereum account:\n{addr}\n\nSign in to Unibase Hub.\n\n" +
		"URI: " + uri + "\nVersion: 1\nChain ID: " + chain + "\nNonce: n0nce1234567\n" +
		"Issued At: " + time.Now().UTC().Format(time.RFC3339)
}

// S6/H5: a sign-in message for another chain is refused.
func TestAuthMiddleware_SIWE_ChainID(t *testing.T) {
	t.Setenv("HUB_SIWE_CHAIN_IDS", "84532")
	s := newV1TestServer(t)
	if w := serve(s, signedWrite(msgHeader(t, siweText("hub.unibase.io", "https://hub.unibase.io", "97")), "c1")); w.Code != http.StatusUnauthorized {
		t.Fatalf("sign-in for chain 97 on an 84532 hub: %d %s", w.Code, w.Body.String())
	}
	if w := serve(s, signedWrite(msgHeader(t, siweText("hub.unibase.io", "https://hub.unibase.io", "84532")), "c2")); w.Code != http.StatusCreated {
		t.Fatalf("sign-in for this chain: %d %s", w.Code, w.Body.String())
	}
}

// S6/H5: with HUB_SIWE_DOMAINS set, the URI must be on an allowed domain too.
func TestAuthMiddleware_SIWE_URI(t *testing.T) {
	t.Setenv("HUB_SIWE_DOMAINS", "hub.unibase.io")
	s := newV1TestServer(t)
	if w := serve(s, signedWrite(msgHeader(t, siweText("hub.unibase.io", "https://evil.example/login", "84532")), "u1")); w.Code != http.StatusUnauthorized {
		t.Fatalf("URI on another site: %d %s", w.Code, w.Body.String())
	}
	if w := serve(s, signedWrite(msgHeader(t, siweText("hub.unibase.io", "https://hub.unibase.io/login", "84532")), "u2")); w.Code != http.StatusCreated {
		t.Fatalf("URI on our domain: %d %s", w.Code, w.Body.String())
	}
}

// The readable (non-SIWE) sign-in text the membase extension and product-web
// sign has no Chain ID or URI line: still accepted when no domain allowlist is
// set, as before.
func TestAuthMiddleware_ReadableMessageCompat(t *testing.T) {
	t.Setenv("HUB_SIWE_CHAIN_IDS", "84532")
	s := newV1TestServer(t)
	msg := "Sign in to Unibase Hub\n\nThis is an off-chain signature proving you control this wallet.\n\n" +
		"Account:\n{addr}\n\nIssued At: " + time.Now().UTC().Format(time.RFC3339) + "\nNonce: 0123456789abcdef"
	if w := serve(s, signedWrite(msgHeader(t, msg), "r1")); w.Code != http.StatusCreated {
		t.Fatalf("readable sign-in message: %d %s", w.Code, w.Body.String())
	}
}

// The hub's configured chain is the default when HUB_SIWE_CHAIN_IDS is unset.
func TestSIWEChainIDsDefault(t *testing.T) {
	defer hubChainID.Store(hubChainID.Load())
	hubChainID.Store(chainIDOf("base-sepolia"))
	if ids := siweChainIDs(); len(ids) != 1 || ids[0] != 84532 {
		t.Fatalf("default chain ids %v", ids)
	}
	t.Setenv("HUB_SIWE_CHAIN_IDS", "8453, bad ,1")
	if ids := siweChainIDs(); len(ids) != 2 || ids[0] != 8453 || ids[1] != 1 {
		t.Fatalf("configured chain ids %v", ids)
	}
	if chainIDOf("nope") != 0 {
		t.Fatal("unknown chain type")
	}
}
