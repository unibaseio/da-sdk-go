package contract

import (
	"encoding/hex"
	"fmt"
	"math/big"

	com "github.com/unibaseio/da-sdk-go/contract/common"
	"github.com/unibaseio/da-sdk-go/lib/types"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	etypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

func (c *ContractManage) HandleSetEpoch(elog etypes.Log, cabi abi.ABI) (types.EpochInfo, error) {
	ei := types.EpochInfo{}

	evInfo, ok := cabi.Events["SetEpoch"]
	if !ok {
		return ei, fmt.Errorf("no event 'SetEpoch' in ABI")
	}

	ld, err := cabi.Unpack(evInfo.Name, elog.Data)
	if err != nil {
		return ei, err
	}
	if len(ld) != 1 {
		return ei, fmt.Errorf("invalid log data length")
	}
	ei.Epoch = ld[0].(uint64)
	bh, seed, err := c.GetEpochInfo(ei.Epoch)
	if err != nil {
		return ei, err
	}
	ei.Height = bh
	ei.Seed = seed
	return ei, nil
}

func (c *ContractManage) HandleAddPiece(elog etypes.Log, cabi abi.ABI) (types.PieceCore, error) {
	pc := types.PieceCore{
		TxHash: elog.TxHash.String(),
	}

	evInfo, ok := cabi.Events["AddPiece"]
	if !ok {
		return pc, fmt.Errorf("no event 'AddPiece' in ABI")
	}

	if len(elog.Topics) != 2 {
		return pc, fmt.Errorf("invalid log topic length")
	}
	pc.Owner = common.HexToAddress(elog.Topics[1].Hex())

	ld, err := cabi.Unpack(evInfo.Name, elog.Data)
	if err != nil {
		return pc, err
	}
	if len(ld) != 2 {
		return pc, fmt.Errorf("invalid log data length")
	}
	pc.Serial = ld[0].(uint64)
	pc.Start = ld[1].(uint64)

	tx, err := com.GetTransactionRetry(c.RPC, elog.TxHash)
	if err != nil {
		return pc, err
	}

	method, ok := cabi.Methods["addPiece"]
	if !ok {
		return pc, fmt.Errorf("no method 'addPiece' in ABI")
	}

	inputData := tx.Data()
	inputs, err := method.Inputs.UnpackValues(inputData[4:])
	if err != nil {
		return pc, err
	}

	if len(inputs) != 7 {
		return pc, fmt.Errorf("invalid input length")
	}

	g1, err := com.SolidityToG1(inputs[0].([]byte))
	if err == nil {
		pc.Name = com.G1ToString(g1)
	} else {
		pc.Name = hex.EncodeToString(inputs[0].([]byte))
	}
	pc.Price = inputs[1].(*big.Int)
	pc.Size = int64(inputs[2].(uint64))
	pc.Expire = inputs[3].(uint64)
	pc.Policy.N = inputs[4].(uint8)
	pc.Policy.K = inputs[5].(uint8)
	pc.Streamer = inputs[6].(common.Address)

	return pc, nil
}

func (c *ContractManage) HandleAddReplica(elog etypes.Log, cabi abi.ABI) (types.ReplicaInChain, error) {
	rc := types.ReplicaInChain{
		TxHash:  elog.TxHash.String(),
		Witness: types.ReplicaWitness{},
	}

	evInfo, ok := cabi.Events["AddReplica"]
	if !ok {
		return rc, fmt.Errorf("no event 'AddReplica' in ABI")
	}

	if len(elog.Topics) != 2 {
		return rc, fmt.Errorf("invalid log topic length")
	}
	rc.StoredOn = common.HexToAddress(elog.Topics[1].Hex())

	ld, err := cabi.Unpack(evInfo.Name, elog.Data)
	if err != nil {
		return rc, err
	}
	if len(ld) != 2 {
		return rc, fmt.Errorf("invalid log data length")
	}
	rc.Serial = ld[0].(uint64)
	rc.Witness.Index = ld[1].(uint64)

	tx, err := com.GetTransactionRetry(c.RPC, elog.TxHash)
	if err != nil {
		return rc, err
	}

	method, ok := cabi.Methods["addReplica"]
	if !ok {
		return rc, fmt.Errorf("no method 'addReplica' in ABI")
	}

	inputData := tx.Data()
	inputs, err := method.Inputs.UnpackValues(inputData[4:])
	if err != nil {
		return rc, err
	}

	if len(inputs) != 4 {
		return rc, fmt.Errorf("invalid input length")
	}

	g1, err := com.SolidityToG1(inputs[0].([]byte))
	if err == nil {
		rc.Name = com.G1ToString(g1)
	} else {
		rc.Name = hex.EncodeToString(inputs[0].([]byte))
	}
	rc.Piece = inputs[1].(uint64)
	rc.Index = inputs[2].(uint8)
	rc.Witness.Proof = inputs[3].([]byte)

	return rc, nil
}

func (c *ContractManage) HandleRSChallenge(elog etypes.Log, cabi abi.ABI) (types.RSChalInChain, error) {
	ei := types.RSChalInChain{}

	evInfo, ok := cabi.Events["Challenge"]
	if !ok {
		return ei, fmt.Errorf("no event 'Challenge' in ABI")
	}

	if len(elog.Topics) != 2 {
		return ei, fmt.Errorf("invalid log topic length")
	}
	ei.Store = common.HexToAddress(elog.Topics[1].Hex())

	ld, err := cabi.Unpack(evInfo.Name, elog.Data)
	if err != nil {
		return ei, err
	}
	if len(ld) != 1 {
		return ei, fmt.Errorf("invalid log data length")
	}
	ei.Replica = ld[0].(uint64)
	return ei, nil
}

func (c *ContractManage) HandleRSFake(elog etypes.Log, cabi abi.ABI) (types.RSChalInChain, error) {
	ei := types.RSChalInChain{}

	evInfo, ok := cabi.Events["Forge"]
	if !ok {
		return ei, fmt.Errorf("no event 'Forge' in ABI")
	}

	if len(elog.Topics) != 2 {
		return ei, fmt.Errorf("invalid log topic length")
	}
	ei.Store = common.HexToAddress(elog.Topics[1].Hex())

	ld, err := cabi.Unpack(evInfo.Name, elog.Data)
	if err != nil {
		return ei, err
	}
	if len(ld) != 1 {
		return ei, fmt.Errorf("invalid log data length")
	}
	ei.Replica = ld[0].(uint64)
	return ei, nil
}

func (c *ContractManage) HandleSubmitEProof(elog etypes.Log, cabi abi.ABI) (types.EProofInChain, error) {
	ei := types.EProofInChain{
		TxHash: elog.TxHash.String(),
	}

	evInfo, ok := cabi.Events["Submit"]
	if !ok {
		return ei, fmt.Errorf("no event 'Submit' in ABI")
	}

	if len(elog.Topics) != 2 {
		return ei, fmt.Errorf("invalid log topic length")
	}
	ei.Store = common.HexToAddress(elog.Topics[1].Hex())

	ld, err := cabi.Unpack(evInfo.Name, elog.Data)
	if err != nil {
		return ei, err
	}
	if len(ld) != 1 {
		return ei, fmt.Errorf("invalid log data length")
	}
	ei.Epoch = ld[0].(uint64)

	tx, err := com.GetTransactionRetry(c.RPC, elog.TxHash)
	if err != nil {
		return ei, err
	}

	method, ok := cabi.Methods["submit"]
	if !ok {
		return ei, fmt.Errorf("no method 'submit' in ABI")
	}

	inputData := tx.Data()
	inputs, err := method.Inputs.UnpackValues(inputData[4:])
	if err != nil {
		return ei, err
	}

	if len(inputs) != 3 {
		return ei, fmt.Errorf("invalid input length")
	}

	sum := inputs[1].([]byte)
	g1, err := com.SolidityToG1(sum)
	if err == nil {
		g1b := g1.Bytes()
		ei.Sum = g1b[:]
	} else {
		ei.Sum = sum
	}

	pf := inputs[2].(([]byte))
	if len(pf) == 144 {
		g1, err := com.SolidityToG1(pf[:96])
		if err == nil {
			g1b := g1.Bytes()
			ei.H = g1b[:]
		}

		fr, err := com.SolidityToFr(pf[96:144])
		if err == nil {
			ei.Value = fr.Marshal()
		}
	}
	ei.Hash = crypto.Keccak256(sum, pf)

	return ei, nil
}

func (c *ContractManage) HandleEPChallenge(elog etypes.Log, cabi abi.ABI) (types.EPChalInChain, error) {
	ei := types.EPChalInChain{}

	evInfo, ok := cabi.Events["Challenge"]
	if !ok {
		return ei, fmt.Errorf("no event 'Challenge' in ABI")
	}

	if len(elog.Topics) != 2 {
		return ei, fmt.Errorf("invalid log topic length")
	}
	ei.Store = common.HexToAddress(elog.Topics[1].Hex())

	ld, err := cabi.Unpack(evInfo.Name, elog.Data)
	if err != nil {
		return ei, err
	}
	if len(ld) != 3 {
		return ei, fmt.Errorf("invalid log data length")
	}
	ei.Epoch = ld[0].(uint64)
	ei.Round = ld[1].(uint8)
	ei.QIndex = ld[2].(uint8)

	return ei, nil
}

func (c *ContractManage) HandleEPProve(elog etypes.Log, cabi abi.ABI) (types.EPChalInChain, error) {
	ei := types.EPChalInChain{}

	evInfo, ok := cabi.Events["Prove"]
	if !ok {
		return ei, fmt.Errorf("no event 'Prove' in ABI")
	}

	if len(elog.Topics) != 2 {
		return ei, fmt.Errorf("invalid log topic length")
	}
	ei.Store = common.HexToAddress(elog.Topics[1].Hex())

	ld, err := cabi.Unpack(evInfo.Name, elog.Data)
	if err != nil {
		return ei, err
	}
	if len(ld) != 2 {
		return ei, fmt.Errorf("invalid log data length")
	}
	ei.Epoch = ld[0].(uint64)
	ei.Round = ld[1].(uint8)

	tx, err := com.GetTransactionRetry(c.RPC, elog.TxHash)
	if err != nil {
		return ei, err
	}

	method, ok := cabi.Methods["proveCom"]
	if !ok {
		return ei, fmt.Errorf("no method 'proveCom' in ABI")
	}

	inputData := tx.Data()
	inputs, err := method.Inputs.UnpackValues(inputData[4:])
	if err != nil {
		return ei, nil
	}

	if len(inputs) != 3 {
		return ei, nil
	}

	coms := inputs[1].([][]byte)
	ei.Coms = make([][]byte, 0, len(coms))
	for i := 0; i < len(coms); i++ {
		g1, err := com.SolidityToG1(coms[i])
		if err == nil {
			g1b := g1.Bytes()
			ei.Coms = append(ei.Coms, g1b[:])
		}
	}

	return ei, nil
}

func (c *ContractManage) HandleEPFake(elog etypes.Log, cabi abi.ABI) (types.EPChalInChain, error) {
	ei := types.EPChalInChain{}

	evInfo, ok := cabi.Events["Fake"]
	if !ok {
		return ei, fmt.Errorf("no event 'Fake' in ABI")
	}

	if len(elog.Topics) != 2 {
		return ei, fmt.Errorf("invalid log topic length")
	}
	ei.Store = common.HexToAddress(elog.Topics[1].Hex())

	ld, err := cabi.Unpack(evInfo.Name, elog.Data)
	if err != nil {
		return ei, err
	}
	if len(ld) != 1 {
		return ei, fmt.Errorf("invalid log data length")
	}
	ei.Epoch = ld[0].(uint64)
	return ei, nil
}
