package contract

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	com "github.com/unibaseio/da-sdk-go/contract/common"
)

// Windows in time and liveness parameters (da-contract fix-blocktime). Base
// Sepolia went from 2s to 0.2s blocks on 2026-10-21, so windows counted in
// blocks alone shrank 10x; the contracts now also take a minimum in seconds.
// Contracts from before that change do not have these getters: a reverted
// call reads as 0 (blocks only, no grace), which is how they behave.

func orZero(v *big.Int, err error) (uint64, error) {
	if err != nil {
		if isRevert(err) {
			return 0, nil
		}
		return 0, err
	}
	return v.Uint64(), nil
}

// GetRSMinProveSeconds reads RSProof.minProveSeconds: an RS challenge's answer
// window lasts at least this long, besides GetRSMinTime blocks.
func (c *ContractManage) GetRSMinProveSeconds() (uint64, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), readCallTimeout())
	defer cancel()
	rsp, err := c.NewRSProof(ctx)
	if err != nil {
		return 0, err
	}
	return orZero(rsp.MinProveSeconds(&bind.CallOpts{From: com.Base, Context: ctx}))
}

// GetEProofMinProveSeconds reads EProof.minProveSeconds (see GetRSMinProveSeconds).
func (c *ContractManage) GetEProofMinProveSeconds() (uint64, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), readCallTimeout())
	defer cancel()
	ep, err := c.NewEProof(ctx)
	if err != nil {
		return 0, err
	}
	return orZero(ep.MinProveSeconds(&bind.CallOpts{From: com.Base, Context: ctx}))
}

// GetSubmitGrace reads EProof.submitGrace: the epochs an epoch proof may come
// late. The proof of epoch e is due in e+1 and accepted until e+1+grace; a
// missing one can be challenged from e+2+grace.
func (c *ContractManage) GetSubmitGrace() (uint64, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), readCallTimeout())
	defer cancel()
	ep, err := c.NewEProof(ctx)
	if err != nil {
		return 0, err
	}
	g, err := ep.SubmitGrace(&bind.CallOpts{From: com.Base, Context: ctx})
	if err != nil {
		if isRevert(err) {
			return 0, nil
		}
		return 0, err
	}
	return g, nil
}

// challengeStake is what a challenger posts to spender (RSProof or EProof):
// its basePenalty. Approving a constant broke every challenge once governance
// raised the penalty (storage audit 2026-09-30 F5).
func (c *ContractManage) challengeStake(ctx context.Context, spender common.Address) (*big.Int, error) {
	opts := &bind.CallOpts{From: com.Base, Context: ctx}
	switch spender {
	case c.RSProofAddr:
		rsp, err := c.NewRSProof(ctx)
		if err != nil {
			return nil, err
		}
		return rsp.BasePenalty(opts)
	case c.EProofAddr:
		ep, err := c.NewEProof(ctx)
		if err != nil {
			return nil, err
		}
		return ep.BasePenalty(opts)
	}
	return com.DefaultPenalty, nil
}

// Window is a game's answer window as the contract enforces it: the move made
// at block base (chalTime / proveTime) may be answered while the block is
// below base+Blocks, or while the time is below MoveAt+Secs (Secs 0: blocks
// only).
type Window struct {
	MoveAt uint64 // unix time of the last move
	Blocks uint64
	Secs   uint64
}

// OpenAt reports whether the move made at block base is still open at block
// and unix time now (an emergency pause aside).
func (w Window) OpenAt(base, block, now uint64) bool {
	return block < base+w.Blocks || (w.Secs != 0 && now < w.MoveAt+w.Secs)
}

// GetRSClock reads RSProof.getClock: the window of the challenge on slot pri
// of piece pi. Against a contract without it: minProveTime blocks only.
func (c *ContractManage) GetRSClock(pi uint64, pri uint8) (Window, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), readCallTimeout())
	defer cancel()
	rsp, err := c.NewRSProof(ctx)
	if err != nil {
		return Window{}, err
	}
	w, err := rsp.GetClock(&bind.CallOpts{From: com.Base, Context: ctx}, pi, pri)
	if err != nil {
		if !isRevert(err) {
			return Window{}, err
		}
		b, err := c.GetRSMinTime()
		return Window{Blocks: b}, err
	}
	return Window{MoveAt: w.MoveAt, Blocks: w.Blocks, Secs: w.Secs}, nil
}

// GetEpochClock reads EProof.getClock: the window of store a's game on epoch ep.
func (c *ContractManage) GetEpochClock(a common.Address, ep uint64) (Window, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), readCallTimeout())
	defer cancel()
	epi, err := c.NewEProof(ctx)
	if err != nil {
		return Window{}, err
	}
	w, err := epi.GetClock(&bind.CallOpts{From: com.Base, Context: ctx}, a, ep)
	if err != nil {
		if !isRevert(err) {
			return Window{}, err
		}
		b, err := c.GetEProofMinTime()
		return Window{Blocks: b}, err
	}
	return Window{MoveAt: w.MoveAt, Blocks: w.Blocks, Secs: w.Secs}, nil
}

// LateFill reports whether Piece accepts a replica in an empty slot after the
// claim window, until the piece expires (da-contract D1). Earlier contracts
// refuse it ("too late") once start+DelaySubmit epochs have passed.
func (c *ContractManage) LateFill() (bool, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), readCallTimeout())
	defer cancel()
	pi, err := c.NewPiece(ctx)
	if err != nil {
		return false, err
	}
	if _, err = pi.ReserveOf(&bind.CallOpts{From: com.Base, Context: ctx}, 0); err != nil {
		if isRevert(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
