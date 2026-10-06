package contract

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"net/url"
	"testing"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/rpc"
)

func TestClassifyRPCErr(t *testing.T) {
	wrap := func(e error) error { return fmt.Errorf("fail to get tx 0xab: %w", e) }
	for _, c := range []struct {
		name string
		err  error
		want ErrClass
	}{
		{"not a direct call", fmt.Errorf("%w: routed", ErrNotDirectCall), ErrPermanent},
		{"foreign log", fmt.Errorf("%w: x", ErrForeignLog), ErrPermanent},
		{"handler logic", errors.New("invalid log data length"), ErrPermanent},
		{"panic in handler", errors.New("panic: index out of range"), ErrPermanent},
		{"revert", codeErr{3, "execution reverted"}, ErrAnswered},
		{"header not found", codeErr{-32000, "header not found"}, ErrAnswered},
		{"rate limit code", codeErr{-32005, "request limit reached"}, ErrOutage},
		{"rate limit text", codeErr{-32000, "Too Many Requests"}, ErrOutage},
		{"tx not found (wrapped)", wrap(ethereum.NotFound), ErrAnswered},
		{"http 503 (wrapped)", wrap(rpc.HTTPError{StatusCode: 503}), ErrOutage},
		{"http 429", rpc.HTTPError{StatusCode: 429}, ErrOutage},
		{"http 401", rpc.HTTPError{StatusCode: 401}, ErrOutage},
		{"http 400", rpc.HTTPError{StatusCode: 400}, ErrAnswered},
		{"dial refused", &url.Error{Op: "Post", URL: "http://x", Err: &net.OpError{Op: "dial", Err: errors.New("connection refused")}}, ErrOutage},
		{"deadline", wrap(context.DeadlineExceeded), ErrOutage},
		{"eof", wrap(io.EOF), ErrOutage},
	} {
		if got := ClassifyRPCErr(c.err); got != c.want {
			t.Errorf("%s: got %s, want %s", c.name, got, c.want)
		}
	}
}
