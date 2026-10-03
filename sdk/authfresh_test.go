package sdk

import (
	"crypto/sha256"
	"encoding/binary"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/unibaseio/da-sdk-go/lib/types"
)

// signedAt builds a legacy auth (signature over hash || be64(ts)).
func signedAt(t *testing.T, ts int64) types.Auth {
	t.Helper()
	sk, err := crypto.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}
	hash := []byte("requestPiece")
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(ts))
	sum := sha256.Sum256(append(append([]byte{}, hash...), b...))
	sig, err := crypto.Sign(sum[:], sk)
	if err != nil {
		t.Fatal(err)
	}
	return types.Auth{Addr: crypto.PubkeyToAddress(sk.PublicKey), Hash: hash, Time: ts, Sign: sig}
}

func TestVerifyAuthFresh(t *testing.T) {
	now := time.Now().Unix()

	if err := VerifyAuthFresh(signedAt(t, now), DefaultAuthDrift); err != nil {
		t.Fatalf("fresh header rejected: %v", err)
	}
	if err := VerifyAuthFresh(signedAt(t, now-DefaultAuthDrift+30), DefaultAuthDrift); err != nil {
		t.Fatalf("header inside the window rejected: %v", err)
	}

	// a captured header replayed an hour later
	err := VerifyAuthFresh(signedAt(t, now-3600), DefaultAuthDrift)
	if err == nil || !strings.Contains(err.Error(), "out of window") {
		t.Fatalf("stale header: got %v, want out of window", err)
	}
	// a header dated in the future
	if err := VerifyAuthFresh(signedAt(t, now+3600), DefaultAuthDrift); err == nil {
		t.Fatal("future header accepted")
	}

	// moving the envelope's time into the window breaks the signature
	au := signedAt(t, now-3600)
	au.Time = now
	if err := VerifyAuthFresh(au, DefaultAuthDrift); err == nil || !strings.Contains(err.Error(), "verify auth") {
		t.Fatalf("re-dated header: got %v, want a signature failure", err)
	}
}
