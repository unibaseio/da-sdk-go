package log

import (
	"bytes"
	"context"
	stdlog "log"
	"net/http"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const testKey = "2QQUkFakeKeyFakeKey0123"

func TestRedactURL(t *testing.T) {
	for in, want := range map[string]string{
		"https://base-sepolia.g.alchemy.com/v2/" + testKey: "https://base-sepolia.g.alchemy.com/" + redacted,
		"https://mainnet.infura.io/v3/abc?x=1":             "https://mainnet.infura.io/" + redacted,
		"https://user:pw@rpc.example.com":                  "https://rpc.example.com/" + redacted,
		"https://base-sepolia-rpc.publicnode.com":          "https://base-sepolia-rpc.publicnode.com",
		"http://127.0.0.1:8545/":                           "http://127.0.0.1:8545",
		"not a url":                                        redacted,
	} {
		if got := RedactURL(in); got != want {
			t.Errorf("RedactURL(%q) = %q, want %q", in, got, want)
		}
	}
}

// The key must vanish from zap entries and stdlib log lines alike, including
// the transport error net/http builds around the full URL.
func TestScrubbedOutput(t *testing.T) {
	ep := "http://127.0.0.1:1/v2/" + testKey // nothing listens: dialing fails
	RegisterURL(ep)
	RegisterURL("http://127.0.0.1:1/rpc") // too short to be a secret: no effect

	cl, err := ethclient.DialContext(context.Background(), ep)
	if err != nil {
		t.Fatal(err)
	}
	_, rpcErr := cl.ChainID(context.Background())
	if rpcErr == nil || !strings.Contains(rpcErr.Error(), testKey) {
		t.Fatalf("expected a transport error quoting the URL, got %v", rpcErr)
	}

	var zbuf bytes.Buffer
	core := zapcore.NewCore(getEncoder(), scrubWriteSyncer(zapcore.AddSync(&zbuf)), zapcore.DebugLevel)
	zap.New(core).Sugar().Errorw("call failed: ", "err", rpcErr, "ep", ep)
	zap.New(core).Sugar().Info("connected to chain: ", ep)

	var sbuf bytes.Buffer
	sl := stdlog.New(ScrubWriter(&sbuf), "", 0)
	sl.Println("sync err:", rpcErr)

	for name, out := range map[string]string{"zap": zbuf.String(), "stdlib": sbuf.String()} {
		if strings.Contains(out, testKey) {
			t.Errorf("%s output leaks the key: %s", name, out)
		}
		if !strings.Contains(out, "/"+redacted) {
			t.Errorf("%s output not redacted: %s", name, out)
		}
		if !strings.Contains(out, "127.0.0.1:1") {
			t.Errorf("%s output lost the host: %s", name, out)
		}
	}
	if got := Scrub("GET /rpc ok"); got != "GET /rpc ok" {
		t.Errorf("short path was scrubbed: %q", got)
	}
}

// S11: a key in the URL's userinfo is scrubbed — the username too, which
// net/http does not mask in its errors (it masks only a password).
func TestScrubUserinfo(t *testing.T) {
	const userKey = "u5erKeyFake0123456789abc"
	const passKey = "pa55KeyFake0123456789xyz"
	userURL := "http://" + userKey + "@127.0.0.1:1/"
	RegisterURL(userURL)
	RegisterURL("https://someone:" + passKey + "@rpc.example.net/")

	_, err := http.Get(userURL) // nothing listens: the error quotes the URL
	if err == nil || !strings.Contains(err.Error(), userKey) {
		t.Fatalf("expected a transport error quoting the username, got %v", err)
	}
	for _, line := range []string{err.Error(), "dial https://someone:" + passKey + "@rpc.example.net/ failed", "key " + passKey} {
		if out := Scrub(line); strings.Contains(out, userKey) || strings.Contains(out, passKey) {
			t.Errorf("userinfo key survived: %q", out)
		}
	}
}

// S11: a key carried as a subdomain label is scrubbed; ordinary labels are not.
func TestScrubHostLabelKey(t *testing.T) {
	const hostKey = "h0stKeyFake0123456789def"
	RegisterURL("https://" + hostKey + ".rpc.example.io/")
	RegisterURL("https://base-sepolia-rpc.publicnode.com")
	out := Scrub(`Post "https://` + hostKey + `.rpc.example.io/": EOF`)
	if strings.Contains(out, hostKey) {
		t.Errorf("host-label key survived: %q", out)
	}
	if !strings.Contains(out, ".rpc.example.io") {
		t.Errorf("the rest of the host was lost: %q", out)
	}
	if got := Scrub("using base-sepolia-rpc.publicnode.com"); got != "using base-sepolia-rpc.publicnode.com" {
		t.Errorf("an ordinary host label was scrubbed: %q", got)
	}
}

// S11: registering an endpoint must not scrub its ordinary path segments
// everywhere: only token-like segments are secrets.
func TestScrubKeepsOrdinaryPaths(t *testing.T) {
	const pathKey = "p4thKeyFake0123456789ghi"
	RegisterURL("https://rpc.example.org/v1/ethereum/stats")
	RegisterURL("https://rpc.example.org/v1/ethereum/" + pathKey)
	if got := Scrub("GET /v1/ethereum/stats 200"); got != "GET /v1/ethereum/stats 200" {
		t.Errorf("a non-secret path was scrubbed: %q", got)
	}
	if got := Scrub("GET /v1/ethereum/" + pathKey); strings.Contains(got, pathKey) || !strings.Contains(got, "/v1/ethereum/") {
		t.Errorf("path key: %q", got)
	}
}

func TestTokenLike(t *testing.T) {
	for _, c := range []struct {
		s    string
		min  int
		want bool
	}{
		{"2QQUkFakeKeyFakeKey0123", minPathTokenLen, true},
		{"0123456789abcdef0123456789abcdef", minPathTokenLen, true}, // hex key
		{"123e4567-e89b-12d3-a456-426614174000", minPathTokenLen, true},
		{"ethereum", minPathTokenLen, false},
		{"stats", minPathTokenLen, false},
		{"v1", minPathTokenLen, false},
		{"base-sepolia-rpc", minHostTokenLen, false},
		{"eth-mainnet-2024-a", minHostTokenLen, false},         // under 20
		{"abcdefghijklmnopqrstuvwxyz", minHostTokenLen, false}, // no digit
		{"12345678901234567890123", minHostTokenLen, false},    // no letter
		{"key.with.dots.0123456789", minPathTokenLen, false},
	} {
		if got := tokenLike(c.s, c.min); got != c.want {
			t.Errorf("tokenLike(%q, %d) = %v, want %v", c.s, c.min, got, c.want)
		}
	}
}
