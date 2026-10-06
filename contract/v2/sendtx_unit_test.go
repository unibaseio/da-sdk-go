package contract

import (
	"context"
	"errors"
	"go/ast"
	"go/parser"
	"go/token"
	"net"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/rpc"
)

type codeErr struct {
	code int
	msg  string
}

func (e codeErr) Error() string  { return e.msg }
func (e codeErr) ErrorCode() int { return e.code }

func TestClassifySendErr(t *testing.T) {
	for _, c := range []struct {
		name string
		err  error
		want sendOutcome
	}{
		{"node: insufficient funds", codeErr{-32000, "insufficient funds for gas * price + value"}, sendRejected},
		{"node: underpriced", codeErr{-32000, "transaction underpriced"}, sendRejected},
		{"node: rate limited", codeErr{-32005, "limit exceeded"}, sendRejected},
		{"node: already known", codeErr{-32000, "already known"}, sendPooled},
		{"node: nonce too low", codeErr{-32000, "nonce too low"}, sendNonceUsed},
		{"node: replacement underpriced", codeErr{-32000, "replacement transaction underpriced"}, sendNonceUsed},
		{"proxy: internal error", codeErr{-32603, "internal error"}, sendUnknown},
		{"proxy: upstream timeout", codeErr{-32000, "upstream request timeout"}, sendUnknown},
		{"forward: EOF", codeErr{-32000, `failed to forward tx to sequencer: Post "https://seq": EOF`}, sendUnknown},
		{"forward: reset", codeErr{-32000, "failed to forward tx to sequencer: read tcp: connection reset by peer"}, sendUnknown},
		{"forward: 502", codeErr{-32000, "502 Bad Gateway: upstream connect error"}, sendUnknown},
		{"forward: deadline", codeErr{-32000, "context deadline exceeded"}, sendUnknown},
		{"node: funds with digits", codeErr{-32000, "insufficient funds for gas * price + value: have 5021 want 5032"}, sendRejected},
		{"http 408", rpc.HTTPError{StatusCode: 408, Status: "408 Request Timeout"}, sendUnknown},
		{"http 429", rpc.HTTPError{StatusCode: 429, Status: "429 Too Many Requests"}, sendRejected},
		{"http 502", rpc.HTTPError{StatusCode: 502, Status: "502 Bad Gateway"}, sendUnknown},
		{"connection reset", &net.OpError{Op: "read", Err: errors.New("connection reset by peer")}, sendUnknown},
		{"deadline", context.DeadlineExceeded, sendUnknown},
	} {
		if got := classifySendErr(c.err); got != c.want {
			t.Errorf("%s: got %d, want %d", c.name, got, c.want)
		}
	}
}

// Every nonce reservation must go through sendTx, which hands back nonces of
// transactions that never left the process (N5). The reserving function is
// unexported and may be called from sendTx only; the old exported MakeAuth is
// gone so no caller can bypass it.
func TestOnlySendTxReservesNonces(t *testing.T) {
	files, err := filepath.Glob("*.go")
	if err != nil {
		t.Fatal(err)
	}
	fset := token.NewFileSet()
	for _, fn := range files {
		if strings.HasSuffix(fn, "_test.go") {
			continue
		}
		src, err := os.ReadFile(fn)
		if err != nil {
			t.Fatal(err)
		}
		f, err := parser.ParseFile(fset, fn, src, 0)
		if err != nil {
			t.Fatal(err)
		}
		for _, d := range f.Decls {
			fd, ok := d.(*ast.FuncDecl)
			if !ok {
				continue
			}
			if fd.Name.Name == "MakeAuth" {
				t.Errorf("%s: exported MakeAuth is back; reserve nonces only in sendTx", fn)
			}
			ast.Inspect(fd, func(n ast.Node) bool {
				sel, ok := n.(*ast.SelectorExpr)
				if !ok {
					return true
				}
				switch sel.Sel.Name {
				case "reserveAuth", "nextNonce":
					if fd.Name.Name != "sendTx" && fd.Name.Name != "reserveAuth" {
						t.Errorf("%s: %s calls %s; only sendTx may reserve a nonce", fset.Position(sel.Pos()), fd.Name.Name, sel.Sel.Name)
					}
				case "MakeAuth", "MakeAuthBySk":
					if fd.Name.Name != "reserveAuth" {
						t.Errorf("%s: %s builds its own auth with %s", fset.Position(sel.Pos()), fd.Name.Name, sel.Sel.Name)
					}
				}
				return true
			})
		}
	}
}
