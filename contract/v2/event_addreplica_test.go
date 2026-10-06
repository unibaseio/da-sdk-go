package contract

import (
	"bytes"
	"errors"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	etypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

// addReplicaLog builds an AddReplica log and the direct addReplica call that
// emitted it, signed so an RPC can serve it, and registers the tx with f.
func addReplicaLog(t *testing.T, f *fakeRPC, pieceAddr common.Address, store common.Address, ri, ordinal uint64, name []byte, pi uint64, pri uint8, proof []byte) etypes.Log {
	t.Helper()
	cabi := pieceABI(t)
	m := cabi.Methods["addReplica"]
	args, err := m.Inputs.Pack(name, pi, pri, proof)
	if err != nil {
		t.Fatal(err)
	}
	sk, _ := crypto.GenerateKey()
	tx, err := etypes.SignNewTx(sk, etypes.LatestSignerForChainID(big.NewInt(31337)), &etypes.LegacyTx{
		Nonce: 1, To: &pieceAddr, Gas: 500000, GasPrice: big.NewInt(1), Value: big.NewInt(0),
		Data: append(append([]byte{}, m.ID...), args...),
	})
	if err != nil {
		t.Fatal(err)
	}
	f.known[tx.Hash()] = tx

	ev := cabi.Events["AddReplica"]
	data, err := ev.Inputs.NonIndexed().Pack(ri, ordinal)
	if err != nil {
		t.Fatal(err)
	}
	return etypes.Log{
		Address:     pieceAddr,
		Topics:      []common.Hash{ev.ID, common.BytesToHash(store.Bytes())},
		Data:        data,
		BlockNumber: 12,
		TxHash:      tx.Hash(),
	}
}

// N4: the AddReplica fields come from the emitting direct call; decoding them
// must not depend on reading Piece state at the latest block. With every
// eth_call failing (lagging or erroring RPC), the event still decodes — before
// the fix this returned an error and the sync loops skipped the event for good.
func TestHandleAddReplicaWithoutStateReads(t *testing.T) {
	f := newFakeRPC(t)
	c := managerOn(t, f)
	f.rpcErr["eth_call"] = jsonRPCError{Code: -32000, Message: "header not found"}

	store := common.HexToAddress("0x5555")
	name := bytes.Repeat([]byte{7}, 96)
	proof := []byte("inner proof bytes")
	elog := addReplicaLog(t, f, c.PieceAddr, store, 41, 9, name, 3, 2, proof)

	rc, err := c.HandleAddReplica(elog, pieceABI(t))
	if err != nil {
		t.Fatalf("HandleAddReplica with failing eth_call: %v", err)
	}
	if rc.Serial != 41 || rc.Witness.Index != 9 || rc.Piece != 3 || rc.Index != 2 || rc.StoredOn != store || !bytes.Equal(rc.Witness.Proof, proof) {
		t.Fatalf("decoded %+v", rc)
	}
	if n := f.count("eth_call"); n != 0 {
		t.Fatalf("%d state reads while decoding AddReplica, want 0", n)
	}
}

// The decoder relies on the log coming from Piece; a log from any other
// address is refused (permanently: retrying cannot change the emitter).
func TestHandleAddReplicaRefusesForeignLog(t *testing.T) {
	f := newFakeRPC(t)
	c := managerOn(t, f)
	other := common.HexToAddress("0x9999")
	elog := addReplicaLog(t, f, other, common.HexToAddress("0x5555"), 41, 9, []byte{1}, 3, 2, []byte{2})

	_, err := c.HandleAddReplica(elog, pieceABI(t))
	if !errors.Is(err, ErrForeignLog) {
		t.Fatalf("got %v, want ErrForeignLog", err)
	}
	if ClassifyRPCErr(err) != ErrPermanent {
		t.Fatal("a foreign log must classify as permanent")
	}
}
