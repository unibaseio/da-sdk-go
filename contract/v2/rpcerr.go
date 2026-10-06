package contract

import (
	"context"
	"errors"
	"io"
	"net"
	"net/url"
	"strings"
	"syscall"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/rpc"
)

// ErrClass says how an event handler's error should be retried (N4).
type ErrClass int

const (
	// ErrPermanent: retrying cannot help (the log or its tx cannot be decoded,
	// the tx went through another contract, a handler's own logic refused it).
	ErrPermanent ErrClass = iota
	// ErrAnswered: an RPC node answered with an error (a revert, "not found",
	// a 4xx). Usually it is lag — the node is behind the block that emitted the
	// event — and goes away, but it can also be deterministic, so retrying it
	// must be bounded.
	ErrAnswered
	// ErrOutage: no usable answer at all (connection failure, timeout, HTTP
	// 5xx/429/401/403, rate limiting). Nothing can be concluded about the
	// event; retry until the endpoint answers.
	ErrOutage
)

func (c ErrClass) String() string {
	switch c {
	case ErrOutage:
		return "outage"
	case ErrAnswered:
		return "answered"
	default:
		return "permanent"
	}
}

// ClassifyRPCErr classifies an error from decoding/handling a chain event.
// Pure.
func ClassifyRPCErr(err error) ErrClass {
	if err == nil {
		return ErrPermanent
	}
	if errors.Is(err, ErrNotDirectCall) || errors.Is(err, ErrForeignLog) {
		return ErrPermanent
	}
	if errors.Is(err, context.DeadlineExceeded) || errors.Is(err, context.Canceled) ||
		errors.Is(err, io.EOF) || errors.Is(err, io.ErrUnexpectedEOF) ||
		errors.Is(err, syscall.ECONNREFUSED) || errors.Is(err, syscall.ECONNRESET) {
		return ErrOutage
	}
	var he rpc.HTTPError
	if errors.As(err, &he) {
		switch {
		case he.StatusCode >= 500, he.StatusCode == 429, he.StatusCode == 401, he.StatusCode == 403, he.StatusCode == 408:
			return ErrOutage
		default:
			return ErrAnswered
		}
	}
	var re rpc.Error
	if errors.As(err, &re) {
		msg := strings.ToLower(re.Error())
		if re.ErrorCode() == -32005 || strings.Contains(msg, "rate limit") || strings.Contains(msg, "limit exceeded") ||
			strings.Contains(msg, "too many requests") || strings.Contains(msg, "timeout") || strings.Contains(msg, "timed out") {
			return ErrOutage
		}
		return ErrAnswered
	}
	if errors.Is(err, ethereum.NotFound) {
		return ErrAnswered // e.g. a replica behind the block, or a pruned tx index
	}
	var ne net.Error
	if errors.As(err, &ne) {
		return ErrOutage
	}
	var ue *url.Error
	if errors.As(err, &ue) {
		return ErrOutage
	}
	return ErrPermanent
}
