package log

import (
	"bytes"
	"context"
	stdlog "log"
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
