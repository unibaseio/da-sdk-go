package sdk

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/unibaseio/da-sdk-go/lib/types"
)

// msgAuth personal_signs msg ("{addr}" replaced by the signer).
func msgAuth(t *testing.T, msg string) types.Auth {
	t.Helper()
	sk, err := crypto.GenerateKey()
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
	return types.Auth{Type: "personal_sign", Addr: addr, Time: time.Now().Unix(), Sign: sig, Msg: msg}
}

func siweLines(domain string, lines ...string) string {
	return domain + " wants you to sign in with your Ethereum account:\n{addr}\n\nSign in.\n\n" +
		strings.Join(lines, "\n") + "\nIssued At: " + time.Now().UTC().Format(time.RFC3339)
}

func TestSIWEPolicy(t *testing.T) {
	ours := SIWEPolicy{Domains: []string{"app.example.com", "localhost:3000"}, ChainIDs: []int64{84532}}
	for _, c := range []struct {
		name string
		msg  string
		p    SIWEPolicy
		ok   bool
	}{
		{"good", siweLines("app.example.com", "URI: https://app.example.com/login", "Version: 1", "Chain ID: 84532", "Nonce: abcdefgh"), ours, true},
		{"port domain", siweLines("localhost:3000", "URI: http://localhost:3000", "Chain ID: 84532"), ours, true},
		{"other chain", siweLines("app.example.com", "URI: https://app.example.com", "Chain ID: 1"), ours, false},
		{"no chain", siweLines("app.example.com", "URI: https://app.example.com"), ours, false},
		{"two chains", siweLines("app.example.com", "URI: https://app.example.com", "Chain ID: 84532", "Chain ID: 1"), ours, false},
		{"bad chain", siweLines("app.example.com", "URI: https://app.example.com", "Chain ID: 0x14a34"), ours, false},
		{"URI elsewhere", siweLines("app.example.com", "URI: https://evil.example", "Chain ID: 84532"), ours, false},
		{"URI missing", siweLines("app.example.com", "Chain ID: 84532"), ours, false},
		{"two URIs", siweLines("app.example.com", "URI: https://app.example.com", "URI: https://evil.example", "Chain ID: 84532"), ours, false},
		{"URI not http", siweLines("app.example.com", "URI: ftp://app.example.com", "Chain ID: 84532"), ours, false},
		{"foreign domain", siweLines("evil.example", "URI: https://app.example.com", "Chain ID: 84532"), ours, false},
		// no allowlist: only the chain is checked
		{"no allowlist", siweLines("evil.example", "URI: https://evil.example", "Chain ID: 84532"), SIWEPolicy{ChainIDs: []int64{84532}}, true},
		// services nobody signs in to: refused without an allowlist
		{"required allowlist", siweLines("app.example.com", "URI: https://app.example.com", "Chain ID: 84532"), SIWEPolicy{RequireDomains: true}, false},
		{"required allowlist set", siweLines("app.example.com", "URI: https://app.example.com", "Chain ID: 1"), SIWEPolicy{RequireDomains: true, Domains: []string{"app.example.com"}}, true},
		// readable non-SIWE text: no domain line, no chain line
		{"readable", "Sign in to Unibase Hub\n\nAccount:\n{addr}\n\nIssued At: " + time.Now().UTC().Format(time.RFC3339), SIWEPolicy{ChainIDs: []int64{84532}}, true},
		{"readable other chain", "Sign in\n\n{addr}\n\nChain ID: 1\nIssued At: " + time.Now().UTC().Format(time.RFC3339), SIWEPolicy{ChainIDs: []int64{84532}}, false},
		{"readable vs allowlist", "Sign in\n\n{addr}\n\nIssued At: " + time.Now().UTC().Format(time.RFC3339), ours, false},
	} {
		err := VerifyAuthFreshPolicy(msgAuth(t, c.msg), DefaultAuthDrift, c.p)
		if (err == nil) != c.ok {
			t.Errorf("%s: err=%v, want ok=%v", c.name, err, c.ok)
		}
	}
}

// Envelopes without a message (the binary hash||time formats nodes and the
// hub use) are not subject to the sign-in policy.
func TestSIWEPolicySkipsBinaryAuth(t *testing.T) {
	sk, _ := crypto.GenerateKey()
	addr := crypto.PubkeyToAddress(sk.PublicKey)
	au := BuildAuth(addr.Hex(), fmt.Sprintf("%x", crypto.FromECDSA(sk)), []byte("upload"))
	if err := VerifyAuthFreshPolicy(au, DefaultAuthDrift, SIWEPolicy{RequireDomains: true, ChainIDs: []int64{1}}); err != nil {
		t.Fatal(err)
	}
}
