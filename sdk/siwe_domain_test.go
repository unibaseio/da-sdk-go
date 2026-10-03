package sdk

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/unibaseio/da-sdk-go/lib/types"
)

// siweAuth builds a personal_signed EIP-4361 auth for domain with extra lines.
func siweAuth(t *testing.T, domain string, issued time.Time, extra string) types.Auth {
	t.Helper()
	sk, err := crypto.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}
	addr := crypto.PubkeyToAddress(sk.PublicKey)
	msg := domain + " wants you to sign in with your Ethereum account:\n" + addr.Hex() + "\n\n" +
		"Sign in.\n\nURI: https://" + domain + "\nVersion: 1\nChain ID: 84532\nNonce: abcdef12\n" +
		"Issued At: " + issued.UTC().Format(time.RFC3339) + extra
	digest := crypto.Keccak256([]byte(fmt.Sprintf("\x19Ethereum Signed Message:\n%d", len(msg)) + msg))
	sig, err := crypto.Sign(digest, sk)
	if err != nil {
		t.Fatal(err)
	}
	return types.Auth{Type: "personal_sign", Addr: addr, Time: issued.Unix(), Sign: sig, Msg: msg}
}

func TestSIWEDomainAndValidity(t *testing.T) {
	now := time.Now()
	ours := []string{"app.membase.io", "hub.unibase.io"}

	if err := VerifyAuthFreshDomains(siweAuth(t, "app.membase.io", now, ""), DefaultAuthDrift, ours); err != nil {
		t.Fatalf("our own domain rejected: %v", err)
	}
	// a sign-in another site collected from the user
	err := VerifyAuthFreshDomains(siweAuth(t, "evil.example", now, ""), DefaultAuthDrift, ours)
	if err == nil || !strings.Contains(err.Error(), "not for this service") {
		t.Fatalf("foreign domain: got %v", err)
	}
	// no allowlist configured: domain not checked
	if err := VerifyAuthFreshDomains(siweAuth(t, "evil.example", now, ""), DefaultAuthDrift, nil); err != nil {
		t.Fatalf("no allowlist: %v", err)
	}

	expired := "\nExpiration Time: " + now.Add(-time.Minute).UTC().Format(time.RFC3339)
	if err := VerifyAuthFresh(siweAuth(t, "app.membase.io", now, expired), DefaultAuthDrift); err == nil {
		t.Fatal("expired message accepted")
	}
	notYet := "\nNot Before: " + now.Add(time.Hour).UTC().Format(time.RFC3339)
	if err := VerifyAuthFresh(siweAuth(t, "app.membase.io", now, notYet), DefaultAuthDrift); err == nil {
		t.Fatal("not-yet-valid message accepted")
	}
	valid := "\nExpiration Time: " + now.Add(time.Hour).UTC().Format(time.RFC3339)
	if err := VerifyAuthFresh(siweAuth(t, "app.membase.io", now, valid), DefaultAuthDrift); err != nil {
		t.Fatalf("message inside its validity window rejected: %v", err)
	}
}
