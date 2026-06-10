package contract

import (
	"context"
	"encoding/binary"
	"fmt"
	"math/big"
	"time"

	com "github.com/unibaseio/da-sdk-go/contract/common"
	"github.com/unibaseio/da-sdk-go/contract/v2/go/eproof"
	"github.com/unibaseio/da-sdk-go/contract/v2/go/everify"
	"github.com/unibaseio/da-sdk-go/contract/v2/go/node"
	"github.com/unibaseio/da-sdk-go/contract/v2/go/piece"
	"github.com/unibaseio/da-sdk-go/contract/v2/go/rsproof"
	"github.com/unibaseio/da-sdk-go/lib/types"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func (c *ContractManage) getOrder(count, pcnt uint64) (uint8, error) {
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancle()

	ri, err := c.NewEVerify(ctx)
	if err != nil {
		return 0, err
	}
	return ri.GetOrder(&bind.CallOpts{From: com.Base}, count, pcnt)
}

func GetOrder(count, pcnt uint64) (uint8, uint64) {
	if count == 0 {
		return 0, 0
	}
	total := uint64(8)
	order := uint8(1)
	for total < count {
		total *= 8
		order += 1
	}
	if order > 4 {
		total /= 8
		order -= 1
	}

	for order > 6 {
		total /= 8
		order -= 1
	}

	for total < count-pcnt {
		total *= 8
		order += 1
	}

	return order, total
}

func (c *ContractManage) choose(addr common.Address, seed [32]byte, count, pcnt uint64, i uint64) (uint64, error) {
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancle()
	ri, err := c.NewEVerify(ctx)
	if err != nil {
		return 0, err
	}
	return ri.Choose(&bind.CallOpts{From: com.Base}, addr, seed, count, pcnt, i)
}

func Choose(addr common.Address, seed [32]byte, count, pcnt uint64, index uint64) uint64 {
	if pcnt+index < count {
		return pcnt + index
	}
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, index)
	if pcnt > 0 {
		buf = crypto.Keccak256(seed[:], addr.Bytes(), buf)
		res := new(big.Int).SetBytes(buf)
		res.Mod(res, big.NewInt(int64(pcnt)))
		return res.Uint64()
	}
	return pcnt
}

func (c *ContractManage) GetBlockNumber() (uint64, error) {
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancle()
	client, err := c.Client(ctx)
	if err != nil {
		return 0, err
	}
	return client.BlockNumber(ctx)
}

func (c *ContractManage) GetEpochBlocks() (uint64, error) {
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancle()

	ei, err := c.NewEpoch(ctx)
	if err != nil {
		return 0, err
	}

	return ei.Slots(&bind.CallOpts{From: com.Base})
}

func (c *ContractManage) GetEpoch() (uint64, error) {
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancle()

	ei, err := c.NewEpoch(ctx)
	if err != nil {
		return 0, err
	}

	return ei.Current(&bind.CallOpts{From: com.Base})
}

func (c *ContractManage) GetEpochInfo(_epoch uint64) (*big.Int, [32]byte, error) {
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancle()

	ei, err := c.NewEpoch(ctx)
	if err != nil {
		return nil, [32]byte{}, err
	}

	return ei.GetEpoch(&bind.CallOpts{From: com.Base}, _epoch)
}

func (c *ContractManage) CheckNode(addr common.Address, _typ uint8) error {
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancle()

	ni, err := c.NewNode(ctx)
	if err != nil {
		return err
	}

	_, _, err = ni.Check(&bind.CallOpts{From: addr}, addr, _typ)
	return err
}

func (c *ContractManage) GetPieceSerial(_pn string) (uint64, error) {
	pnb, err := com.G1StringInSolidity(_pn)
	if err != nil {
		return 0, err
	}
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancle()
	gi, err := c.NewPiece(ctx)
	if err != nil {
		return 0, err
	}
	return gi.GetPIndex(&bind.CallOpts{From: com.Base}, pnb)
}

func (c *ContractManage) GetReplicaSerial(_pn string) (uint64, error) {
	pnb, err := com.G1StringInSolidity(_pn)
	if err != nil {
		return 0, err
	}
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancle()
	gi, err := c.NewPiece(ctx)
	if err != nil {
		return 0, err
	}
	return gi.GetRIndex(&bind.CallOpts{From: com.Base}, pnb)
}

func (c *ContractManage) GetPiece(_pi uint64) (piece.IPiecePieceInfo, error) {
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancle()
	fi, err := c.NewPiece(ctx)
	if err != nil {
		return piece.IPiecePieceInfo{}, err
	}
	pb, err := fi.GetPiece(&bind.CallOpts{From: com.Base}, _pi)
	if err != nil {
		return piece.IPiecePieceInfo{}, err
	}
	return pb, nil
}

func (c *ContractManage) GetReplica(_ri uint64) (piece.IPieceReplicaInfo, error) {
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancle()
	fi, err := c.NewPiece(ctx)
	if err != nil {
		return piece.IPieceReplicaInfo{}, err
	}
	pb, err := fi.GetReplica(&bind.CallOpts{From: com.Base}, _ri)
	if err != nil {
		return piece.IPieceReplicaInfo{}, err
	}

	return pb, nil
}

func (c *ContractManage) GetMinPledge(_typ uint8) (*big.Int, error) {
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancle()
	ni, err := c.NewNode(ctx)
	if err != nil {
		return nil, err
	}
	return ni.MinStakeOf(&bind.CallOpts{From: com.Base}, _typ)
}

func (c *ContractManage) GetPledgeInfo(addr common.Address) (node.INodeNodeInfo, error) {
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancle()
	ni, err := c.NewNode(ctx)
	if err != nil {
		return node.INodeNodeInfo{}, err
	}
	return ni.NodeInfoOf(&bind.CallOpts{From: addr}, addr)
}

func (c *ContractManage) GetStore(addr common.Address) (piece.IPieceStoreInfo, error) {
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancle()
	fi, err := c.NewPiece(ctx)
	if err != nil {
		return piece.IPieceStoreInfo{}, err
	}
	fss, err := fi.GetStore(&bind.CallOpts{From: addr}, addr)
	if err != nil {
		return piece.IPieceStoreInfo{}, err
	}
	return fss, nil
}

func (c *ContractManage) GetStoreStat(addr common.Address, _epoch uint64) (piece.IPieceStoreStat, error) {
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancle()
	fi, err := c.NewPiece(ctx)
	if err != nil {
		return piece.IPieceStoreStat{}, err
	}
	fss, err := fi.GetSStat(&bind.CallOpts{From: addr}, addr, _epoch)
	if err != nil {
		return piece.IPieceStoreStat{}, err
	}
	return fss, nil
}

func (c *ContractManage) GetStoreReplica(_a common.Address, _ri uint64) (uint64, error) {
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancle()
	fi, err := c.NewPiece(ctx)
	if err != nil {
		return 0, err
	}
	return fi.GetSRAt(&bind.CallOpts{From: com.Base}, _a, _ri)
}

func (c *ContractManage) GetRSChalInfo(_pi uint64, _pri uint8) (rsproof.IRSProofProofInfo, error) {
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancle()
	rsp, err := c.NewRSProof(ctx)
	if err != nil {
		return rsproof.IRSProofProofInfo{}, err
	}

	return rsp.GetProof(&bind.CallOpts{From: com.Base}, _pi, _pri)
}

func (c *ContractManage) GetRSMinTime() (uint64, error) {
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancle()
	rsp, err := c.NewRSProof(ctx)
	if err != nil {
		return 0, err
	}

	t, err := rsp.MinProveTime(&bind.CallOpts{From: com.Base})
	if err != nil {
		return 0, err
	}

	return t.Uint64(), nil
}

func (c *ContractManage) GetEpochChalInfo(_a common.Address, _ep uint64) (eproof.IEProofProofInfo, error) {
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancle()
	ep, err := c.NewEProof(ctx)
	if err != nil {
		return eproof.IEProofProofInfo{}, err
	}

	return ep.GetEProof(&bind.CallOpts{From: com.Base}, _a, _ep)
}

func (c *ContractManage) GetEpochChalDetail(_a common.Address, _ep uint64) (everify.IEVerifyCInfo, error) {
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancle()
	ep, err := c.NewEVerify(ctx)
	if err != nil {
		return everify.IEVerifyCInfo{}, err
	}

	return ep.GetCInfo(&bind.CallOpts{From: com.Base}, _a, _ep)
}

func (c *ContractManage) GetEProofMinTime() (uint64, error) {
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancle()
	rsp, err := c.NewEProof(ctx)
	if err != nil {
		return 0, err
	}

	t, err := rsp.MinProveTime(&bind.CallOpts{From: com.Base})
	if err != nil {
		return 0, err
	}

	return t.Uint64(), nil
}

func (c *ContractManage) GetRevenue(addr common.Address, typ string) (*big.Int, error) {
	res := big.NewInt(0)
	ctx, cancle := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancle()

	switch typ {
	case types.StoreType:
		gi, err := c.NewPiece(ctx)
		if err != nil {
			return res, err
		}
		si, err := gi.GetStore(&bind.CallOpts{From: addr}, addr)
		if err != nil {
			return res, err
		}

		res.Set(si.Revenue)
		return res, nil
	case types.StreamType:
		gi, err := c.NewPiece(ctx)
		if err != nil {
			return res, err
		}
		val, err := gi.GetRevenue(&bind.CallOpts{From: addr}, addr)
		if err != nil {
			return res, err
		}

		res.Set(val)
		return res, nil
	default:
		return res, fmt.Errorf("unsupported")
	}
}

func (c *ContractManage) CheckBalance(addr common.Address) error {
	val := c.BalanceOf(addr)
	if val.Cmp(big.NewInt(0)) == 0 {
		time.Sleep(30 * time.Second)
		val = c.BalanceOf(addr)
		if val.Cmp(big.NewInt(0)) == 0 {
			return fmt.Errorf("not has gas token")
		}
	}

	val = c.BalanceOfToken(addr)
	retry := 10
	for retry > 0 {
		if val.Cmp(big.NewInt(0)) > 0 {
			return nil
		}
		time.Sleep(30 * time.Second)
		val = c.BalanceOfToken(addr)
		retry--
	}
	return fmt.Errorf("not has erc20 token")
}
