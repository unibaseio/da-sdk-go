package contract

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"time"

	com "github.com/unibaseio/da-sdk-go/contract/common"
	"github.com/unibaseio/da-sdk-go/lib/bls"
	"github.com/unibaseio/da-sdk-go/lib/types"
	"github.com/unibaseio/da-sdk-go/lib/utils"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func (c *ContractManage) Set(_typ string, ca common.Address) error {
	ctx, cancle := context.WithTimeout(context.TODO(), 1*time.Minute)
	defer cancle()
	bi, err := c.NewBank(ctx)
	if err != nil {
		return err
	}

	au, err := c.MakeAuth()
	if err != nil {
		return err
	}

	tx, err := bi.Set(au, _typ, ca)
	if err != nil {
		return err
	}

	err = com.CheckTx(c.RPC, tx.Hash())
	if err != nil {
		return err
	}
	com.Logger.Debug(_typ, " set success to: ", ca.String())
	return nil
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
	err = com.CheckTx(c.RPC, tx.Hash())
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
		_, err = ni.Check(&bind.CallOpts{From: au.From}, au.From, _typ)
		if err == nil {
			com.Logger.Debugf("%s already pledge enough money in type %d", au.From, _typ)
			return nil
		}

		pval, err := ni.GetMinPledge(&bind.CallOpts{From: au.From}, _typ)
		if err != nil {
			return err
		}
		pinfo, err := ni.GetPledge(&bind.CallOpts{From: au.From}, au.From, _typ)
		if err != nil {
			return err
		}

		pinfo.Value.Sub(pinfo.Value, pinfo.Lock)

		if pval.Cmp(pinfo.Value) > 0 {
			pval.Sub(pval, pinfo.Value)
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
	tx, err := ti.IncreaseAllowance(au, c.BankAddr, val)
	if err != nil {
		return err
	}
	err = com.CheckTx(c.RPC, tx.Hash())
	if err != nil {
		return err
	}
	tx, err = ni.Pledge(au, _typ, val)
	if err != nil {
		return err
	}
	err = com.CheckTx(c.RPC, tx.Hash())
	if err != nil {
		return err
	}
	_, err = ni.Check(&bind.CallOpts{From: au.From}, au.From, _typ)
	return err
}

func (c *ContractManage) AddPiece(pc types.PieceCore) (string, error) {
	com.Logger.Debug("add piece: ", pc)
	ctx, cancle := context.WithTimeout(context.TODO(), 3*time.Minute)
	defer cancle()

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

	_size := 1 + (pc.Size-1)/(31*int64(pc.Policy.K))
	_size = 1 + (_size-1)/(32*1024)

	val := big.NewInt(int64(pc.Expire-pc.Start) * _size)
	val.Mul(val, pc.Price)
	val.Add(val, big.NewInt(int64(com.DefaultStreamPrice)))
	val.Mul(val, big.NewInt(int64(pc.Policy.N)))

	au, err := c.MakeAuth()
	if err != nil {
		return "", err
	}

	ti, err := c.NewToken(ctx)
	if err != nil {
		return "", err
	}

	gtoken := c.BalanceOf(au.From)
	fmt.Println("submitpiece0: ", gtoken)
	tx, err := ti.IncreaseAllowance(au, c.BankAddr, val)
	if err != nil {
		return "", err
	}
	err = com.CheckTx(c.RPC, tx.Hash())
	if err != nil {
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
	fmt.Println("submitpiece1: ", c.BalanceOf(au.From))
	tx, err = fi.AddPiece(au, pb, pc.Price, uint64(pc.Size), pc.Expire, pc.Policy.N, pc.Policy.K, pc.Streamer)
	if err != nil {
		return "", err
	}
	err = com.CheckTx(c.RPC, tx.Hash())
	if err != nil {
		return "", err
	}
	fmt.Println("submitpiece2: ", c.BalanceOf(au.From))
	fmt.Println("submitpiece cost: ", utils.FormatEth(gtoken.Sub(gtoken, c.BalanceOf(au.From))))

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
	fmt.Println("submitreplica0: ", c.BalanceOf(au.From))
	tx, err := fi.AddReplica(au, rb, _pi, rc.Index, pf)
	if err != nil {
		return err
	}
	err = com.CheckTx(c.RPC, tx.Hash())
	if err != nil {
		return err
	}
	fmt.Println("submitreplica1: ", c.BalanceOf(au.From))
	fmt.Println("submitreplica cost: ", utils.FormatEth(gtoken.Sub(gtoken, c.BalanceOf(au.From))))

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

	err = com.CheckTx(c.RPC, tx.Hash())
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

	tx, err := ti.IncreaseAllowance(au, c.BankAddr, com.DefaultPenalty)
	if err != nil {
		return err
	}
	err = com.CheckTx(c.RPC, tx.Hash())
	if err != nil {
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
	tx, err = rsp.Challenge(au, pname, rname, _pri)
	if err != nil {
		return err
	}

	err = com.CheckTx(c.RPC, tx.Hash())
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

	tx, err := rsp.Prove(au, pname, rname, _pri, _pf)
	if err != nil {
		return err
	}

	err = com.CheckTx(c.RPC, tx.Hash())
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

	err = com.CheckTx(c.RPC, tx.Hash())
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
	fmt.Println("submitproof0: ", c.BalanceOf(au.From))
	tx, err := pi.Submit(au, _ep, _sum, _pfb)
	if err != nil {
		return err
	}

	err = com.CheckTx(c.RPC, tx.Hash())
	if err != nil {
		return err
	}
	fmt.Println("submitproof1: ", c.BalanceOf(au.From))
	fmt.Println("submitproof cost: ", utils.FormatEth(gtoken.Sub(gtoken, c.BalanceOf(au.From))))
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

	tx, err := ti.IncreaseAllowance(au, c.BankAddr, com.DefaultPenalty)
	if err != nil {
		return err
	}
	err = com.CheckTx(c.RPC, tx.Hash())
	if err != nil {
		return err
	}

	pi, err := c.NewEProof(ctx)
	if err != nil {
		return err
	}

	com.Logger.Debug("challenge eproof: ", addr, _ep)
	tx, err = pi.ChalKZG(au, addr, _ep)
	if err != nil {
		return err
	}

	err = com.CheckTx(c.RPC, tx.Hash())
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

	tx, err := pi.ProveKZG(au, _ep, _wt, _pf)
	if err != nil {
		return err
	}

	err = com.CheckTx(c.RPC, tx.Hash())
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
		tx, err := ti.IncreaseAllowance(au, c.BankAddr, com.DefaultPenalty)
		if err != nil {
			return err
		}
		err = com.CheckTx(c.RPC, tx.Hash())
		if err != nil {
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
		tx, err := pi.Challenge(au, addr, _ep, _sum)
		if err != nil {
			return err
		}
		err = com.CheckTx(c.RPC, tx.Hash())
		if err != nil {
			return err
		}
	} else {
		com.Logger.Debug("challenge eproof sum: ", addr, _ep, _qIndex)
		tx, err := pi.ChalCom(au, addr, _ep, _qIndex)
		if err != nil {
			return err
		}

		err = com.CheckTx(c.RPC, tx.Hash())
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

	tx, err := ti.IncreaseAllowance(au, c.BankAddr, com.DefaultPenalty)
	if err != nil {
		return err
	}
	err = com.CheckTx(c.RPC, tx.Hash())
	if err != nil {
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
	tx, err = pi.ProveCom(au, _ep, _coms, _pf)
	if err != nil {
		return err
	}

	err = com.CheckTx(c.RPC, tx.Hash())
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

	err = com.CheckTx(c.RPC, tx.Hash())
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
	tx, err := pi.ProveOne(au, _ep, _commit, _pf)
	if err != nil {
		return err
	}

	err = com.CheckTx(c.RPC, tx.Hash())
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

	err = com.CheckTx(c.RPC, tx.Hash())
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

	err = com.CheckTx(c.RPC, tx.Hash())
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

	err = com.CheckTx(c.RPC, tx.Hash())
	if err != nil {
		return err
	}

	return nil
}

func (c *ContractManage) SettleReward(addr common.Address, _ep uint64) error {
	ctx, cancle := context.WithTimeout(context.TODO(), 1*time.Minute)
	defer cancle()
	fi, err := c.NewReward(ctx)
	if err != nil {
		return err
	}
	au, err := c.MakeAuth()
	if err != nil {
		return err
	}

	tx, err := fi.Settle(au, addr, _ep)
	if err != nil {
		return err
	}

	err = com.CheckTx(c.RPC, tx.Hash())
	if err != nil {
		return err
	}

	return nil
}

func (c *ContractManage) WithdrawReward(_money *big.Int) error {
	ctx, cancle := context.WithTimeout(context.TODO(), 1*time.Minute)
	defer cancle()
	fi, err := c.NewReward(ctx)
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

	err = com.CheckTx(c.RPC, tx.Hash())
	if err != nil {
		return err
	}

	return nil
}

func (c *ContractManage) AddModel(mc types.ModelMeta) error {
	com.Logger.Debug("add model: ", mc)
	ctx, cancle := context.WithTimeout(context.TODO(), 1*time.Minute)
	defer cancle()
	mi, err := c.NewModel(ctx)
	if err != nil {
		return err
	}
	au, err := c.MakeAuth()
	if err != nil {
		return err
	}

	_rt, err := com.G1StringInSolidity(mc.Hash)
	if err != nil {
		return err
	}

	com.Logger.Debug("add model: ", mc.Name)
	tx, err := mi.Add(au, mc.Name, _rt)
	if err != nil {
		return err
	}

	err = com.CheckTx(c.RPC, tx.Hash())
	if err != nil {
		return err
	}

	return nil
}

func (c *ContractManage) AddGPU(_gn string) error {
	ctx, cancle := context.WithTimeout(context.TODO(), 1*time.Minute)
	defer cancle()
	gi, err := c.NewGPU(ctx)
	if err != nil {
		return err
	}
	au, err := c.MakeAuth()
	if err != nil {
		return err
	}

	tx, err := gi.Add(au, _gn)
	if err != nil {
		return err
	}

	err = com.CheckTx(c.RPC, tx.Hash())
	if err != nil {
		return err
	}

	return nil
}

func (c *ContractManage) AddSpace(msm types.SpaceMeta) error {
	ctx, cancle := context.WithTimeout(context.TODO(), 3*time.Minute)
	defer cancle()
	si, err := c.NewSpace(ctx)
	if err != nil {
		return err
	}

	gi, err := c.NewGPU(ctx)
	if err != nil {
		return err
	}

	mi, err := c.NewModel(ctx)
	if err != nil {
		return err
	}

	au, err := c.MakeAuth()
	if err != nil {
		return err
	}

	_mi, err := mi.GetIndex(&bind.CallOpts{From: au.From}, msm.Model)
	if err != nil {
		return err
	}

	_gi, err := gi.GetIndex(&bind.CallOpts{From: au.From}, msm.GPU)
	if err != nil {
		return err
	}

	ce, err := c.UpdateEpoch()
	if err != nil {
		return err
	}

	val := big.NewInt(int64(com.DefaultSpacePrice))
	val.Mul(val, big.NewInt(int64(com.DefaultSpaceEpoch)))

	ti, err := c.NewToken(ctx)
	if err != nil {
		return err
	}

	tx, err := ti.IncreaseAllowance(au, c.BankAddr, val)
	if err != nil {
		return err
	}
	err = com.CheckTx(c.RPC, tx.Hash())
	if err != nil {
		return err
	}

	tx, err = si.Add(au, msm.Name, _mi, _gi, big.NewInt(int64(com.DefaultSpacePrice)), ce+uint64(com.DefaultSpaceEpoch))
	if err != nil {
		return err
	}

	err = com.CheckTx(c.RPC, tx.Hash())
	if err != nil {
		return err
	}

	return nil
}

func (c *ContractManage) ActivateSpace(sn, root string, pfbyte []byte) error {
	ctx, cancle := context.WithTimeout(context.TODO(), 3*time.Minute)
	defer cancle()
	si, err := c.NewSpace(ctx)
	if err != nil {
		return err
	}

	au, err := c.MakeAuth()
	if err != nil {
		return err
	}

	_ai, err := si.GetIndex(&bind.CallOpts{From: au.From}, sn)
	if err != nil {
		return err
	}

	tx, err := si.Activate(au, _ai)
	if err != nil {
		return err
	}

	err = com.CheckTx(c.RPC, tx.Hash())
	if err != nil {
		return err
	}

	return nil
}

func (c *ContractManage) ShutdownSpace(_ai uint64) error {
	ctx, cancle := context.WithTimeout(context.TODO(), 3*time.Minute)
	defer cancle()
	si, err := c.NewSpace(ctx)
	if err != nil {
		return err
	}

	au, err := c.MakeAuth()
	if err != nil {
		return err
	}

	tx, err := si.Shutdown(au, _ai)
	if err != nil {
		return err
	}

	err = com.CheckTx(c.RPC, tx.Hash())
	if err != nil {
		return err
	}

	return nil
}
