package contract

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"time"

	com "github.com/unibaseio/da-sdk-go/contract/common"
	"github.com/unibaseio/da-sdk-go/contract/v2/go/token"
	"github.com/unibaseio/da-sdk-go/lib/bls"
	"github.com/unibaseio/da-sdk-go/lib/types"
	"github.com/unibaseio/da-sdk-go/lib/utils"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// CurrentProofVersion is the proof-format version this client generates (V6-B2).
// It is passed to the on-chain prove functions, which dispatch to the verifier
// registered for that version — so a future circuit upgrade bumps this + registers
// a new verifier, while proofs from old clients keep verifying against version 1.
const CurrentProofVersion uint8 = 1

// waitForAllowance polls until the RPC observes owner's token allowance for
// spender to be at least min. Guards against RPC read-after-write lag on
// load-balanced endpoints (e.g. Alchemy on base-sepolia): an Approve tx can be
// mined (CheckTx confirms it) yet the very next estimateGas / eth_call land on
// a replica that hasn't caught up, read a stale allowance of 0, and revert the
// in-contract safeTransferFrom with "execution reverted". CheckTx confirms the
// tx is mined; this confirms the read side caught up. Mirrors waitForCode in
// contract/deploy/deploy_v2.go, which fixed the same race for a fresh proxy.
func (c *ContractManage) waitForAllowance(ctx context.Context, ti *token.Token, owner, spender common.Address, min *big.Int) error {
	for i := 0; i < 30; i++ {
		cur, err := ti.Allowance(&bind.CallOpts{Context: ctx, From: owner}, owner, spender)
		if err == nil && cur.Cmp(min) >= 0 {
			return nil
		}
		select {
		case <-ctx.Done():
			return fmt.Errorf("wait for allowance %s -> %s aborted: %w", owner.Hex(), spender.Hex(), ctx.Err())
		case <-time.After(2 * time.Second):
		}
	}
	return fmt.Errorf("timed out waiting for allowance %s -> %s >= %s", owner.Hex(), spender.Hex(), min.String())
}

// Attest records this validator's on-chain liveness for the given epoch in the
// reward pool (FixB+A2). No-op if the pool address isn't configured.
func (c *ContractManage) Attest(epoch uint64) error {
	if c.ValidatorRewardAddr == (common.Address{}) {
		return nil
	}
	ctx, cancel := context.WithTimeout(context.TODO(), 1*time.Minute)
	defer cancel()
	vi, err := c.NewValidatorReward(ctx)
	if err != nil {
		return err
	}
	au, err := c.MakeAuth()
	if err != nil {
		return err
	}
	tx, err := vi.Attest(au, epoch)
	if err != nil {
		return err
	}
	return c.CheckTx(tx.Hash())
}

// Claim pulls this validator's allocated reward from the pool. No-op if the
// pool address isn't configured.
func (c *ContractManage) Claim() error {
	if c.ValidatorRewardAddr == (common.Address{}) {
		return nil
	}
	ctx, cancel := context.WithTimeout(context.TODO(), 1*time.Minute)
	defer cancel()
	vi, err := c.NewValidatorReward(ctx)
	if err != nil {
		return err
	}
	au, err := c.MakeAuth()
	if err != nil {
		return err
	}
	tx, err := vi.Claim(au)
	if err != nil {
		return err
	}
	return c.CheckTx(tx.Hash())
}

// AddPieceCost is the exact token amount AddPiece locks (IncreaseAllowance) for
// a piece, in wei. Single source of truth shared by AddPiece (v1) and the hub
// /api/seal v2 path (which must return a cost matching what the contract locks,
// or the client's IncreaseAllowance underfunds and addPiece reverts). pc.Start/
// Expire/Price must already be resolved (non-zero).
func AddPieceCost(pc types.PieceCore) *big.Int {
	_size := 1 + (pc.Size-1)/(31*int64(pc.Policy.K))
	_size = 1 + (_size-1)/(32*1024)

	val := big.NewInt(int64(pc.Expire-pc.Start) * _size)
	val.Mul(val, pc.Price)
	val.Add(val, big.NewInt(int64(com.DefaultStreamPrice)))
	val.Mul(val, big.NewInt(int64(pc.Policy.N)))
	return val
}

func (c *ContractManage) UpdateEpoch() (uint64, error) {
	ctx, cancle := context.WithTimeout(context.TODO(), 1*time.Minute)
	defer cancle()
	ei, err := c.NewEpoch(ctx)
	if err != nil {
		return 0, err
	}

	au, err := c.MakeAuth()
	if err != nil {
		return 0, err
	}

	tx, err := ei.Check(au)
	if err != nil {
		return 0, err
	}
	err = c.CheckTx(tx.Hash())
	if err != nil {
		return 0, err
	}

	return ei.Current(&bind.CallOpts{From: au.From})
}

func (c *ContractManage) RegisterNode(_typ uint8, val *big.Int) error {
	ctx, cancle := context.WithTimeout(context.TODO(), 3*time.Minute)
	defer cancle()
	ni, err := c.NewNode(ctx)
	if err != nil {
		return err
	}

	au, err := c.MakeAuth()
	if err != nil {
		return err
	}

	ti, err := c.NewToken(ctx)
	if err != nil {
		return err
	}

	if val == nil {
		isActive, _, err := ni.Check(&bind.CallOpts{From: au.From}, au.From, _typ)
		if err == nil && isActive {
			com.Logger.Debugf("%s already pledge enough money in type %d", au.From, _typ)
			return nil
		}

		pval, err := ni.MinStakeOf(&bind.CallOpts{From: au.From}, _typ)
		if err != nil {
			return err
		}
		pinfo, err := ni.NodeInfoOf(&bind.CallOpts{From: au.From}, au.From)
		if err != nil {
			return err
		}

		pinfo.StakedAmount.Sub(pinfo.StakedAmount, pinfo.LockedAmount)

		if pval.Cmp(pinfo.StakedAmount) > 0 {
			pval.Sub(pval, pinfo.StakedAmount)
			val = pval
		} else {
			com.Logger.Debug("no need more pledge")
			return nil
		}
	}

	if val.Cmp(big.NewInt(0)) < 0 {
		return fmt.Errorf("negative value")
	}

	com.Logger.Debug("register node: ", au.From, val)
	tx, err := ti.Approve(au, c.NodeAddr, val)
	if err != nil {
		return err
	}
	err = c.CheckTx(tx.Hash())
	if err != nil {
		return err
	}
	// Wait for the RPC to observe the approve before Stake reads the allowance
	// on a possibly-stale replica. See waitForAllowance.
	if err = c.waitForAllowance(ctx, ti, au.From, c.NodeAddr, val); err != nil {
		return err
	}
	au, err = c.MakeAuth() // fresh nonce: the prior au's nonce is already spent
	if err != nil {
		return err
	}
	tx, err = ni.Stake(au, _typ, val)
	if err != nil {
		return err
	}
	err = c.CheckTx(tx.Hash())
	if err != nil {
		return err
	}
	_, _, err = ni.Check(&bind.CallOpts{From: au.From}, au.From, _typ)
	return err
}

// AddPiece registers a piece on-chain (IncreaseAllowance + addPiece) under the
// default 3-minute budget. See AddPieceCtx for a caller-supplied deadline.
func (c *ContractManage) AddPiece(pc types.PieceCore) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	return c.addPieceImpl(ctx, pc, nil)
}

// AddPieceCtx is AddPiece with a caller-supplied context. ctx bounds the whole
// flow — the contract-binding calls AND the two CheckTx receipt waits (which
// previously escaped AddPiece's own timeout). A synchronous HTTP handler passes
// a short ctx so a stuck tx can't hold the request open; on ctx expiry the blob
// is already staged, so the (idempotent) seal can be retried to confirm.
func (c *ContractManage) AddPieceCtx(ctx context.Context, pc types.PieceCore) (string, error) {
	return c.addPieceImpl(ctx, pc, nil)
}

// AddPieceFor submits a piece sponsored by this account (the payer / relayer)
// but attributed on-chain to `owner` (the end user) via the contract's
// addPieceFor. This account must hold RELAYER_ROLE on the Piece contract. Used
// by the hub's sponsored /api/seal path so piece.owner reflects the real user.
func (c *ContractManage) AddPieceFor(pc types.PieceCore, owner common.Address) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	return c.addPieceImpl(ctx, pc, &owner)
}

// addPieceImpl is the shared body: owner==nil -> addPiece (self), owner!=nil ->
// addPieceFor (sponsored, attributed to *owner). ctx bounds the binding calls +
// both CheckTx receipt waits. Payment is always pulled from this account.
func (c *ContractManage) addPieceImpl(ctx context.Context, pc types.PieceCore, owner *common.Address) (string, error) {
	com.Logger.Debug("add piece: ", pc)

	ce, err := c.GetEpoch()
	if err != nil {
		return "", err
	}

	if pc.Expire == 0 {
		pc.Start = ce
		pc.Expire = ce + uint64(com.DefaultStoreEpoch)
	}
	if pc.Price == nil {
		pc.Price = big.NewInt(int64(com.DefaultReplicaPrice))
	}

	val := AddPieceCost(pc)
	com.Logger.Debug("submitpiece val: ", utils.FormatEth(val))

	au, err := c.MakeAuth()
	if err != nil {
		return "", err
	}
	au.Context = ctx // a cancelled ctx aborts the send's EstimateGas too

	ti, err := c.NewToken(ctx)
	if err != nil {
		return "", err
	}

	gtoken := c.BalanceOf(au.From)
	com.Logger.Debug("submitpiece0: ", gtoken)
	tx, err := ti.Approve(au, c.PieceAddr, val)
	if err != nil {
		return "", err
	}
	err = c.CheckTxCtx(ctx, tx.Hash())
	if err != nil {
		return "", err
	}
	// Wait for the RPC's read side to observe the approve before addPiece's
	// estimateGas reads the allowance (else a stale replica reverts the
	// in-contract safeTransferFrom). See waitForAllowance.
	if err = c.waitForAllowance(ctx, ti, au.From, c.PieceAddr, val); err != nil {
		return "", err
	}

	fi, err := c.NewPiece(ctx)
	if err != nil {
		return "", err
	}

	pb, err := com.G1StringInSolidity(pc.Name)
	if err != nil {
		return "", err
	}

	com.Logger.Debug("add piece: ", pc)
	com.Logger.Debug("submitpiece1: ", c.BalanceOf(au.From))
	au, err = c.MakeAuth() // fresh nonce: the prior au's nonce is already spent
	if err != nil {
		return "", err
	}
	au.Context = ctx
	if owner != nil {
		tx, err = fi.AddPieceFor(au, *owner, pb, pc.Price, uint64(pc.Size), pc.Expire, pc.Policy.N, pc.Policy.K, pc.Streamer)
	} else {
		tx, err = fi.AddPiece(au, pb, pc.Price, uint64(pc.Size), pc.Expire, pc.Policy.N, pc.Policy.K, pc.Streamer)
	}
	if err != nil {
		return "", err
	}
	err = c.CheckTxCtx(ctx, tx.Hash())
	if err != nil {
		return "", err
	}
	com.Logger.Debug("submitpiece2: ", c.BalanceOf(au.From))
	com.Logger.Debug("submitpiece cost: ", utils.FormatEth(gtoken.Sub(gtoken, c.BalanceOf(au.From))))

	return tx.Hash().String(), nil
}

func (c *ContractManage) AddReplica(rc types.ReplicaCore, pf []byte) error {
	com.Logger.Debug("add replica: ", rc)
	rb, err := com.G1StringInSolidity(rc.Name)
	if err != nil {
		return err
	}
	ctx, cancle := context.WithTimeout(context.TODO(), 3*time.Minute)
	defer cancle()
	fi, err := c.NewPiece(ctx)
	if err != nil {
		return err
	}

	au, err := c.MakeAuth()
	if err != nil {
		return err
	}

	pbyte, err := com.G1StringInSolidity(rc.Piece)
	if err != nil {
		return err
	}
	_pi, err := fi.GetPIndex(&bind.CallOpts{From: au.From}, pbyte)
	if err != nil {
		return err
	}

	if _pi == 0 {
		return fmt.Errorf("%s is not on chain", rc.Piece)
	}

	_ri, err := fi.GetRIndex(&bind.CallOpts{From: au.From}, rb)
	if err != nil {
		return err
	}

	if _ri > 0 {
		return fmt.Errorf("%s is already on chain", rc.Name)
	}

	gtoken := c.BalanceOf(au.From)
	com.Logger.Debug("add replica: ", _pi, rc)
	com.Logger.Debug("submitreplica0: ", c.BalanceOf(au.From))
	tx, err := fi.AddReplica(au, rb, _pi, rc.Index, pf)
	if err != nil {
		return err
	}
	err = c.CheckTx(tx.Hash())
	if err != nil {
		return err
	}
	com.Logger.Debug("submitreplica1: ", c.BalanceOf(au.From))
	com.Logger.Debug("submitreplica cost: ", utils.FormatEth(gtoken.Sub(gtoken, c.BalanceOf(au.From))))

	return nil
}

func (c *ContractManage) UpdateStore(store common.Address) error {
	com.Logger.Debug("update store: ", store)

	ctx, cancle := context.WithTimeout(context.TODO(), 1*time.Minute)
	defer cancle()
	fi, err := c.NewPiece(ctx)
	if err != nil {
		return err
	}

	au, err := c.MakeAuth()
	if err != nil {
		return err
	}

	tx, err := fi.CheckStore(au, store)
	if err != nil {
		return err
	}

	err = c.CheckTx(tx.Hash())
	if err != nil {
		return err
	}

	return nil
}

func (c *ContractManage) ChallengeRS(_pn, _rn string, _pri uint8) error {
	ctx, cancle := context.WithTimeout(context.TODO(), 3*time.Minute)
	defer cancle()

	au, err := c.MakeAuth()
	if err != nil {
		return err
	}

	ti, err := c.NewToken(ctx)
	if err != nil {
		return err
	}

	tx, err := ti.Approve(au, c.RSProofAddr, com.DefaultPenalty)
	if err != nil {
		return err
	}
	err = c.CheckTx(tx.Hash())
	if err != nil {
		return err
	}
	// Wait for the RPC to observe the penalty approve before Challenge reads the
	// allowance on a possibly-stale replica. See waitForAllowance.
	if err = c.waitForAllowance(ctx, ti, au.From, c.RSProofAddr, com.DefaultPenalty); err != nil {
		return err
	}

	rsp, err := c.NewRSProof(ctx)
	if err != nil {
		return err
	}
	pname, err := com.G1StringInSolidity(_pn)
	if err != nil {
		pname, err = hex.DecodeString(_pn)
		if err != nil {
			return err
		}
	}

	rname, err := com.G1StringInSolidity(_rn)
	if err != nil {
		rname, err = hex.DecodeString(_rn)
		if err != nil {
			return err
		}
	}

	com.Logger.Debug("challenge rs proof: ", _rn, _pn, _pri)
	au, err = c.MakeAuth() // fresh nonce: the prior au's nonce is already spent
	if err != nil {
		return err
	}
	tx, err = rsp.Challenge(au, pname, rname, _pri)
	if err != nil {
		return err
	}

	err = c.CheckTx(tx.Hash())
	if err != nil {
		return err
	}

	return nil
}

func (c *ContractManage) ProveRS(_pn, _rn string, _pri uint8, _pf []byte) error {
	ctx, cancle := context.WithTimeout(context.TODO(), 3*time.Minute)
	defer cancle()
	rsp, err := c.NewRSProof(ctx)
	if err != nil {
		return err
	}

	au, err := c.MakeAuth()
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

	tx, err := rsp.Prove(au, pname, rname, _pri, CurrentProofVersion, _pf)
	if err != nil {
		return err
	}

	err = c.CheckTx(tx.Hash())
	if err != nil {
		return err
	}

	return nil
}

func (c *ContractManage) CheckRSChallenge(_pn, _rn string, _pri uint8) error {
	ctx, cancle := context.WithTimeout(context.TODO(), 1*time.Minute)
	defer cancle()
	rsp, err := c.NewRSProof(ctx)
	if err != nil {
		return err
	}

	au, err := c.MakeAuth()
	if err != nil {
		return err
	}

	pname, err := com.G1StringInSolidity(_pn)
	if err != nil {
		pname, err = hex.DecodeString(_pn)
		if err != nil {
			return err
		}
	}

	rname, err := com.G1StringInSolidity(_rn)
	if err != nil {
		rname, err = hex.DecodeString(_rn)
		if err != nil {
			return err
		}
	}

	piece, err := c.NewPiece(ctx)
	if err != nil {
		return err
	}

	_pi, err := piece.GetPIndex(&bind.CallOpts{From: au.From}, pname)
	if err != nil {
		return err
	}

	_ri, err := piece.GetRIndex(&bind.CallOpts{From: au.From}, rname)
	if err != nil {
		return err
	}

	tx, err := rsp.Check(au, _pi, _ri, _pri)
	if err != nil {
		return err
	}

	err = c.CheckTx(tx.Hash())
	if err != nil {
		return err
	}

	return nil
}

func (c *ContractManage) SubmitProof(_ep uint64, _pf bls.EpochProof) error {
	ctx, cancle := context.WithTimeout(context.TODO(), 1*time.Minute)
	defer cancle()
	pi, err := c.NewEProof(ctx)
	if err != nil {
		return err
	}

	au, err := c.MakeAuth()
	if err != nil {
		return err
	}

	_sum := com.G1InSolidity(_pf.Sum)
	_pfb := com.G1InSolidity(_pf.H)
	_frb := com.FrInSolidity(_pf.ClaimedValue)
	_pfb = append(_pfb, _frb...)

	gtoken := c.BalanceOf(au.From)
	com.Logger.Debug("submit epoch proof: ", au.From, _ep)
	com.Logger.Debug("submitproof0: ", c.BalanceOf(au.From))
	tx, err := pi.Submit(au, _ep, _sum, _pfb)
	if err != nil {
		return err
	}

	err = c.CheckTx(tx.Hash())
	if err != nil {
		return err
	}
	com.Logger.Debug("submitproof1: ", c.BalanceOf(au.From))
	com.Logger.Debug("submitproof cost: ", utils.FormatEth(gtoken.Sub(gtoken, c.BalanceOf(au.From))))
	return nil
}

func (c *ContractManage) ChallengeKZG(addr common.Address, _ep uint64) error {
	ctx, cancle := context.WithTimeout(context.TODO(), 3*time.Minute)
	defer cancle()

	au, err := c.MakeAuth()
	if err != nil {
		return err
	}

	ti, err := c.NewToken(ctx)
	if err != nil {
		return err
	}

	tx, err := ti.Approve(au, c.EProofAddr, com.DefaultPenalty)
	if err != nil {
		return err
	}
	err = c.CheckTx(tx.Hash())
	if err != nil {
		return err
	}
	// Wait for the RPC to observe the penalty approve before ChalKZG reads the
	// allowance on a possibly-stale replica. See waitForAllowance.
	if err = c.waitForAllowance(ctx, ti, au.From, c.EProofAddr, com.DefaultPenalty); err != nil {
		return err
	}

	pi, err := c.NewEProof(ctx)
	if err != nil {
		return err
	}

	com.Logger.Debug("challenge eproof: ", addr, _ep)
	au, err = c.MakeAuth() // fresh nonce: the prior au's nonce is already spent
	if err != nil {
		return err
	}
	tx, err = pi.ChalKZG(au, addr, _ep)
	if err != nil {
		return err
	}

	err = c.CheckTx(tx.Hash())
	if err != nil {
		return err
	}

	return nil
}

func (c *ContractManage) ProveKZG(_ep uint64, _wroot []byte, _pf []byte) error {
	if len(_wroot) != 32 {
		return fmt.Errorf("invalid witness root length")
	}

	ctx, cancle := context.WithTimeout(context.TODO(), 3*time.Minute)
	defer cancle()
	pi, err := c.NewEProof(ctx)
	if err != nil {
		return err
	}

	au, err := c.MakeAuth()
	if err != nil {
		return err
	}

	var _wt [32]byte
	copy(_wt[:], _wroot)

	tx, err := pi.ProveKZG(au, _ep, _wt, CurrentProofVersion, _pf)
	if err != nil {
		return err
	}

	err = c.CheckTx(tx.Hash())
	if err != nil {
		return err
	}

	return nil
}

func (c *ContractManage) ChallengeSum(addr common.Address, _ep uint64, _qIndex uint8, sum string) error {
	ctx, cancle := context.WithTimeout(context.TODO(), 3*time.Minute)
	defer cancle()

	au, err := c.MakeAuth()
	if err != nil {
		return err
	}
	ti, err := c.NewToken(ctx)
	if err != nil {
		return err
	}

	if len(sum) > 0 {
		tx, err := ti.Approve(au, c.EProofAddr, com.DefaultPenalty)
		if err != nil {
			return err
		}
		err = c.CheckTx(tx.Hash())
		if err != nil {
			return err
		}
		// Wait for the RPC to observe the penalty approve before Challenge reads
		// the allowance on a possibly-stale replica. See waitForAllowance.
		if err = c.waitForAllowance(ctx, ti, au.From, c.EProofAddr, com.DefaultPenalty); err != nil {
			return err
		}
	}

	pi, err := c.NewEProof(ctx)
	if err != nil {
		return err
	}

	if len(sum) > 0 {
		_sum, err := com.G1StringInSolidity(sum)
		if err != nil {
			_sum, err = hex.DecodeString(sum)
			if err != nil {
				return err
			}
		}
		com.Logger.Debug("challenge eproof sum0: ", addr, _ep)
		au, err = c.MakeAuth() // fresh nonce: the prior au's nonce is already spent
		if err != nil {
			return err
		}
		tx, err := pi.Challenge(au, addr, _ep, _sum)
		if err != nil {
			return err
		}
		err = c.CheckTx(tx.Hash())
		if err != nil {
			return err
		}
	} else {
		com.Logger.Debug("challenge eproof sum: ", addr, _ep, _qIndex)
		tx, err := pi.ChalCom(au, addr, _ep, _qIndex)
		if err != nil {
			return err
		}

		err = c.CheckTx(tx.Hash())
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *ContractManage) ProveSum(_ep uint64, coms []bls.G1, _pf []byte) error {
	ctx, cancle := context.WithTimeout(context.TODO(), 3*time.Minute)
	defer cancle()

	au, err := c.MakeAuth()
	if err != nil {
		return err
	}

	ti, err := c.NewToken(ctx)
	if err != nil {
		return err
	}

	tx, err := ti.Approve(au, c.EProofAddr, com.DefaultPenalty)
	if err != nil {
		return err
	}
	err = c.CheckTx(tx.Hash())
	if err != nil {
		return err
	}
	// Wait for the RPC to observe the penalty approve before ProveCom reads the
	// allowance on a possibly-stale replica. See waitForAllowance.
	if err = c.waitForAllowance(ctx, ti, au.From, c.EProofAddr, com.DefaultPenalty); err != nil {
		return err
	}

	pi, err := c.NewEProof(ctx)
	if err != nil {
		return err
	}

	_coms := make([][]byte, len(coms))
	for i := 0; i < len(coms); i++ {
		_coms[i] = com.G1InSolidity(coms[i])
	}

	com.Logger.Debug("prove eproof sum: ", au.From, _ep)
	au, err = c.MakeAuth() // fresh nonce: the prior au's nonce is already spent
	if err != nil {
		return err
	}
	tx, err = pi.ProveCom(au, _ep, _coms, CurrentProofVersion, _pf)
	if err != nil {
		return err
	}

	err = c.CheckTx(tx.Hash())
	if err != nil {
		return err
	}

	return nil
}

func (c *ContractManage) ChallengeOne(addr common.Address, _ep uint64, _qIndex uint8) error {
	ctx, cancle := context.WithTimeout(context.TODO(), 1*time.Minute)
	defer cancle()

	au, err := c.MakeAuth()
	if err != nil {
		return err
	}

	pi, err := c.NewEProof(ctx)
	if err != nil {
		return err
	}

	com.Logger.Debug("challenge eproof one: ", addr, _ep, _qIndex)
	tx, err := pi.ChalOne(au, addr, _ep, _qIndex)
	if err != nil {
		return err
	}

	err = c.CheckTx(tx.Hash())
	if err != nil {
		return err
	}

	return nil
}

func (c *ContractManage) ProveOne(_ep uint64, _com bls.G1, _pf []byte) error {
	ctx, cancle := context.WithTimeout(context.TODO(), 3*time.Minute)
	defer cancle()

	au, err := c.MakeAuth()
	if err != nil {
		return err
	}

	pi, err := c.NewEProof(ctx)
	if err != nil {
		return err
	}

	_commit := com.G1InSolidity(_com)
	com.Logger.Debug("prove eproof one: ", au.From, _ep)
	tx, err := pi.ProveOne(au, _ep, _commit, CurrentProofVersion, _pf)
	if err != nil {
		return err
	}

	err = c.CheckTx(tx.Hash())
	if err != nil {
		return err
	}

	return nil
}

func (c *ContractManage) CheckEpochChallenge(addr common.Address, _ep uint64) error {
	ctx, cancle := context.WithTimeout(context.TODO(), 1*time.Minute)
	defer cancle()

	au, err := c.MakeAuth()
	if err != nil {
		return err
	}

	ep, err := c.NewEProof(ctx)
	if err != nil {
		return err
	}

	com.Logger.Debug("check eproof: ", addr, _ep)
	tx, err := ep.Check(au, addr, _ep)
	if err != nil {
		return err
	}

	err = c.CheckTx(tx.Hash())
	if err != nil {
		return err
	}

	return nil
}

func (c *ContractManage) TestProveRS(rsn, rsk uint8, pub []*big.Int, _pf []byte) error {
	if len(pub) != 3 {
		return fmt.Errorf("invalid public length")
	}
	ctx, cancle := context.WithTimeout(context.TODO(), 1*time.Minute)
	defer cancle()
	rsp, err := c.NewRSProof(ctx)
	if err != nil {
		return err
	}

	vt, err := rsp.GetVKRoot(&bind.CallOpts{From: com.Base}, rsn, rsk)
	if err != nil {
		return err
	}

	if vt.Cmp(pub[0]) != 0 {
		return fmt.Errorf("unequal vkroot")
	}

	rsv, err := c.NewRSOne(ctx)
	if err != nil {
		return err
	}

	ok, err := rsv.Verify(&bind.CallOpts{From: com.Base}, _pf, pub)
	if err != nil {
		return err
	}
	if !ok {
		return fmt.Errorf("invalid rs proof")
	}

	return nil
}

func (c *ContractManage) TestProveEpoch(_key string, pub []*big.Int, _pf []byte) error {
	if len(pub) != 2 {
		return fmt.Errorf("invalid public length")
	}
	ctx, cancle := context.WithTimeout(context.TODO(), 1*time.Minute)
	defer cancle()
	ev, err := c.NewEVerify(ctx)
	if err != nil {
		return err
	}

	vt := new(big.Int)
	switch _key {
	case "kzg":
		vt, err = ev.INKZGVK(&bind.CallOpts{From: com.Base})
	case "mul":
		vt, err = ev.INMULVK(&bind.CallOpts{From: com.Base})
	case "add":
		vt, err = ev.INADDVK(&bind.CallOpts{From: com.Base})
	default:
		return fmt.Errorf("unsupported inner circuit: %s", _key)
	}
	if err != nil {
		return err
	}

	if vt.Cmp(pub[0]) != 0 {
		return fmt.Errorf("unequal vkroot")
	}

	switch _key {
	case "kzg":
		pv, err := c.NewKZGPlonk(ctx)
		if err != nil {
			return err
		}
		ok, err := pv.Verify(&bind.CallOpts{From: com.Base}, _pf, pub)
		if err != nil {
			return err
		}
		if !ok {
			return fmt.Errorf("invalid rs proof kzg")
		}
	case "add":
		pv, err := c.NewAddPlonk(ctx)
		if err != nil {
			return err
		}
		ok, err := pv.Verify(&bind.CallOpts{From: com.Base}, _pf, pub)
		if err != nil {
			return err
		}
		if !ok {
			return fmt.Errorf("invalid rs proof add")
		}
	case "mul":
		pv, err := c.NewMulPlonk(ctx)
		if err != nil {
			return err
		}
		ok, err := pv.Verify(&bind.CallOpts{From: com.Base}, _pf, pub)
		if err != nil {
			return err
		}
		if !ok {
			return fmt.Errorf("invalid rs proof mul")
		}
	default:
		return fmt.Errorf("unsupported key")
	}

	return nil
}

func (c *ContractManage) Settle(_money *big.Int) error {
	ctx, cancle := context.WithTimeout(context.TODO(), 1*time.Minute)
	defer cancle()
	fi, err := c.NewPiece(ctx)
	if err != nil {
		return err
	}
	au, err := c.MakeAuth()
	if err != nil {
		return err
	}

	tx, err := fi.Settle(au, _money)
	if err != nil {
		return err
	}

	err = c.CheckTx(tx.Hash())
	if err != nil {
		return err
	}

	return nil
}

func (c *ContractManage) WithdrawRevenue(_money *big.Int) error {
	ctx, cancle := context.WithTimeout(context.TODO(), 1*time.Minute)
	defer cancle()
	fi, err := c.NewPiece(ctx)
	if err != nil {
		return err
	}
	au, err := c.MakeAuth()
	if err != nil {
		return err
	}

	tx, err := fi.Withdraw(au, _money)
	if err != nil {
		return err
	}

	err = c.CheckTx(tx.Hash())
	if err != nil {
		return err
	}

	return nil
}
