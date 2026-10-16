package contract

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/rpc"
	com "github.com/unibaseio/da-sdk-go/contract/common"
	"github.com/unibaseio/da-sdk-go/contract/v2/go/eproof"
	"github.com/unibaseio/da-sdk-go/contract/v2/go/rsproof"
	"github.com/unibaseio/da-sdk-go/lib/bls"
)

// ErrProofRejected means the chain refuses an answer: the on-chain move made
// with it reverts. An RPC failure says nothing about the answer and is
// returned as is.
var ErrProofRejected = errors.New("the chain refuses this answer")

// Simulating an answer runs the very move a store is about to send — prove,
// proveKZG, proveCom, proveOne — as an eth_call from the store's address. The
// contract then computes every public input the verifier gets from chain state
// (piece and replica names, the stored inner-proof root, the epoch seed, the
// game's commitments). TestProveRS/TestProveEpoch checked only the VK root and
// took the other public inputs from the prover, so a prover could return a
// valid proof of another statement that passed the check and reverted on
// chain (audit 2026-10-13 H1).

type rawCaller interface {
	Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error
}

func (c *ContractManage) simulate(ctx context.Context, raw rawCaller, method string, args ...interface{}) error {
	var out []interface{}
	err := raw.Call(&bind.CallOpts{From: c.From(), Context: ctx}, &out, method, args...)
	if err == nil {
		return nil
	}
	if isRevert(err) {
		return fmt.Errorf("%w: %s: %v", ErrProofRejected, method, err)
	}
	return err
}

// isRevert tells a revert (the call ran and failed) from a failure to run it.
func isRevert(err error) bool {
	var de rpc.DataError
	if errors.As(err, &de) && de.ErrorData() != nil {
		return true
	}
	return strings.Contains(err.Error(), "execution reverted")
}

// SimulateProveRS runs ProveRS(_pn, _rn, _pri, _pf) as an eth_call.
func (c *ContractManage) SimulateProveRS(_pn, _rn string, _pri uint8, _pf []byte) error {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Minute)
	defer cancel()
	rsp, err := c.NewRSProof(ctx)
	if err != nil {
		return err
	}
	pname, err := com.G1StringInSolidity(_pn)
	if err != nil {
		return err
	}
	rname, err := com.G1StringInSolidity(_rn)
	if err != nil {
		return err
	}
	return c.simulate(ctx, &rsproof.RSProofRaw{Contract: rsp}, "prove", pname, rname, _pri, CurrentProofVersion, _pf)
}

// SimulateProveKZG runs ProveKZG(_ep, _wroot, _pf) as an eth_call.
func (c *ContractManage) SimulateProveKZG(_ep uint64, _wroot []byte, _pf []byte) error {
	if len(_wroot) != 32 {
		return fmt.Errorf("invalid witness root length")
	}
	ctx, cancel := context.WithTimeout(context.TODO(), time.Minute)
	defer cancel()
	pi, err := c.NewEProof(ctx)
	if err != nil {
		return err
	}
	var wt [32]byte
	copy(wt[:], _wroot)
	return c.simulate(ctx, &eproof.EProofRaw{Contract: pi}, "proveKZG", _ep, wt, CurrentProofVersion, _pf)
}

// SimulateProveSum runs ProveSum(_ep, coms, _pf) as an eth_call (proveCom moves
// no tokens, so ProveSum's approve is not needed for it).
func (c *ContractManage) SimulateProveSum(_ep uint64, coms []bls.G1, _pf []byte) error {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Minute)
	defer cancel()
	pi, err := c.NewEProof(ctx)
	if err != nil {
		return err
	}
	cs := make([][]byte, len(coms))
	for i := range coms {
		cs[i] = com.G1InSolidity(coms[i])
	}
	return c.simulate(ctx, &eproof.EProofRaw{Contract: pi}, "proveCom", _ep, cs, CurrentProofVersion, _pf)
}

// SimulateProveOne runs ProveOne(_ep, _com, _pf) as an eth_call.
func (c *ContractManage) SimulateProveOne(_ep uint64, _com bls.G1, _pf []byte) error {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Minute)
	defer cancel()
	pi, err := c.NewEProof(ctx)
	if err != nil {
		return err
	}
	return c.simulate(ctx, &eproof.EProofRaw{Contract: pi}, "proveOne", _ep, com.G1InSolidity(_com), CurrentProofVersion, _pf)
}
