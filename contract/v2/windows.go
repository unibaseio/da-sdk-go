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
