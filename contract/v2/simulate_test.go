package contract

import (
	"errors"
	"testing"

	"github.com/unibaseio/da-sdk-go/lib/bls"
)

// Audit 2026-10-13 H1: an answer is checked by running the on-chain move as an
// eth_call. A revert means the chain refuses the answer (the store then asks
// another prover); an RPC failure says nothing about the answer.
func TestSimulateTellsRevertFromRPCFailure(t *testing.T) {
	f := newFakeRPC(t)
	c := managerOn(t, f)
	var g bls.G1
	moves := map[string]func() error{
		"proveKZG": func() error { return c.SimulateProveKZG(3, make([]byte, 32), []byte{1}) },
		"proveCom": func() error { return c.SimulateProveSum(3, make([]bls.G1, 8), []byte{1}) },
		"proveOne": func() error { return c.SimulateProveOne(3, g, []byte{1}) },
	}

	for name, move := range moves {
		if err := move(); err != nil {
			t.Fatalf("%s: a call that succeeds: %v", name, err)
		}
	}

	f.rpcErr["eth_call"] = jsonRPCError{Code: 3, Message: "execution reverted: inv pf", Data: "0x08c379a0"}
	for name, move := range moves {
		if err := move(); !errors.Is(err, ErrProofRejected) {
			t.Errorf("%s: revert gave %v, want ErrProofRejected", name, err)
		}
	}

	f.rpcErr["eth_call"] = jsonRPCError{Code: -32000, Message: "header not found"}
	for name, move := range moves {
		if err := move(); err == nil || errors.Is(err, ErrProofRejected) {
			t.Errorf("%s: RPC failure gave %v, want a plain error", name, err)
		}
	}
}
