package contract

import (
	"bytes"
	"encoding/hex"
	"errors"
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

// ErrNotDirectCall means the transaction that emitted an event did not call the
// contract's method directly (it went through another contract), so its
// calldata is not that method's arguments and cannot be decoded as such.
var ErrNotDirectCall = errors.New("event's transaction is not a direct call to the method")

// ErrForeignLog means a log handed to a decoder was not emitted by the
// contract that decoder is for.
var ErrForeignLog = errors.New("log not emitted by the expected contract")

// directCallInputs decodes the arguments of method m from tx's calldata, but
// only when tx calls the emitting contract directly with m's selector. Any
// other shape (a call routed through a contract, or crafted calldata) is
// refused: decoding it as m's arguments would read attacker-chosen bytes, and
// slicing short calldata used to crash every syncing node.
func directCallInputs(tx *etypes.Transaction, emitter common.Address, m abi.Method) ([]interface{}, error) {
	if tx.To() == nil || *tx.To() != emitter {
		return nil, fmt.Errorf("%w: tx %s targets %v, event emitted by %s", ErrNotDirectCall, tx.Hash().Hex(), tx.To(), emitter.Hex())
	}
	data := tx.Data()
	if len(data) < 4 || !bytes.Equal(data[:4], m.ID) {
		return nil, fmt.Errorf("%w: tx %s is not %s", ErrNotDirectCall, tx.Hash().Hex(), m.Name)
	}
	return m.Inputs.UnpackValues(data[4:])
}

// decodeAddPieceFields recovers the piece fields
// [pn, price, size, expire, rsn, rsk, streamer] from a tx's calldata. The
// AddPiece event is emitted by BOTH addPiece (owner == caller) and addPieceFor
// (owner == attributed user, caller == relayer/hub); addPieceFor has an extra
// leading `owner` arg that shifts every field by one. Dispatch on the 4-byte
// selector so attributed pieces aren't silently mis-decoded with addPiece's
// layout. The owner itself comes from the indexed event topic, not here.
func decodeAddPieceFields(cabi abi.ABI, inputData []byte) ([]interface{}, error) {
	if len(inputData) < 4 {
		return nil, fmt.Errorf("%w: calldata shorter than a selector", ErrNotDirectCall)
	}
	selector := inputData[:4]

	if m, ok := cabi.Methods["addPieceFor"]; ok && bytes.Equal(selector, m.ID) {
		in, err := m.Inputs.UnpackValues(inputData[4:])
		if err != nil {
			return nil, err
		}
		if len(in) != 8 {
			return nil, fmt.Errorf("invalid addPieceFor input length")
		}
		return in[1:], nil // drop the leading owner arg
	}
	if m, ok := cabi.Methods["addPiece"]; ok && bytes.Equal(selector, m.ID) {
		in, err := m.Inputs.UnpackValues(inputData[4:])
		if err != nil {
			return nil, err
		}
		if len(in) != 7 {
			return nil, fmt.Errorf("invalid addPiece input length")
		}
		return in, nil
	}
	return nil, fmt.Errorf("unexpected method selector %x for AddPiece event", selector)
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

	if tx.To() == nil || *tx.To() != elog.Address {
		return pc, fmt.Errorf("%w: AddPiece tx %s targets %v, event emitted by %s", ErrNotDirectCall, elog.TxHash.Hex(), tx.To(), elog.Address.Hex())
	}
	fields, err := decodeAddPieceFields(cabi, tx.Data())
	if err != nil {
		return pc, err
	}

	g1, err := com.SolidityToG1(fields[0].([]byte))
	if err == nil {
		pc.Name = com.G1ToString(g1)
	} else {
		pc.Name = hex.EncodeToString(fields[0].([]byte))
	}
	pc.Price = fields[1].(*big.Int)
	pc.Size = int64(fields[2].(uint64))
	pc.Expire = fields[3].(uint64)
	pc.Policy.N = fields[4].(uint8)
	pc.Policy.K = fields[5].(uint8)
	pc.Streamer = fields[6].(common.Address)

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

	if elog.Address != c.PieceAddr {
		return rc, fmt.Errorf("%w: AddReplica log from %s, Piece is %s", ErrForeignLog, elog.Address.Hex(), c.PieceAddr.Hex())
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

	inputs, err := directCallInputs(tx, elog.Address, method)
	if err != nil {
		return rc, err
	}

	if len(inputs) != 4 {
		return rc, fmt.Errorf("invalid input length")
	}
	// No re-read of these fields from Piece state (N4). The log comes from
	// Piece (checked above), whose only AddReplica emit is at the end of
	// addReplica, and directCallInputs has proved the tx is a direct call to
	// Piece.addReplica — so a successful tx runs exactly one addReplica, with
	// exactly these arguments, and it stored rmap[name]=ri, prmap[pi][pri]=ri,
	// storedOn=caller and root=keccak(proof) from them; Piece never rewrites
	// those entries. A re-read adds nothing, and reading at the latest block
	// made a lagging or failing RPC drop the event for good.

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

	inputs, err := directCallInputs(tx, elog.Address, method)
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

	// A proveCom answer is the 8 child commitments it stored in EVerify
	// (commit[store][epoch][i]); read them from there rather than from the
	// tx's calldata, which a call routed through a contract does not carry.
	// Prove events from proveKZG (round 0) and proveOne (the leaf) have none.
	ci, err := c.GetEpochChalDetail(ei.Store, ei.Epoch)
	if err != nil {
		return ei, err
	}
	if ei.Round == 0 || ei.Round > ci.Round {
		return ei, nil
	}
	if ci.RoundAt != ei.Round {
		// the game has moved past this event; the current commits are not its own
		return ei, fmt.Errorf("stale Prove event: round %d, game now at %d", ei.Round, ci.RoundAt)
	}
	coms, err := c.GetEpochCommits(ei.Store, ei.Epoch)
	if err != nil {
		return ei, err
	}
	ei.Coms = coms
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
