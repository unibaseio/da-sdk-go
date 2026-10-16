package contract

import (
	"bytes"
	"errors"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	etypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/unibaseio/da-sdk-go/contract/v2/go/eproof"
)

// Audit 2026-10-13 F2: Piece/EProof emit AddPieceV2/AddReplicaV2/SubmitV2,
// which carry every argument. Decoding them reads nothing from the RPC, so a
// piece, replica or epoch proof submitted through a contract (whose tx calldata
// is not the method's arguments) is seen like a direct one.

func eproofABI(t *testing.T) abi.ABI {
	t.Helper()
	a, err := abi.JSON(strings.NewReader(eproof.EProofABI))
	if err != nil {
		t.Fatal(err)
	}
	return a
}

func v2Log(t *testing.T, cabi abi.ABI, name string, emitter, indexed common.Address, args ...interface{}) etypes.Log {
	t.Helper()
	ev := cabi.Events[name]
	data, err := ev.Inputs.NonIndexed().Pack(args...)
	if err != nil {
		t.Fatal(err)
	}
	return etypes.Log{
		Address: emitter,
		Topics:  []common.Hash{ev.ID, common.BytesToHash(indexed.Bytes())},
		Data:    data,
		// a tx no RPC knows: the decoder must not need it
		TxHash: common.HexToHash("0xf2"),
	}
}

// failAllRPC makes every call the decoders could make fail.
func failAllRPC(f *fakeRPC) {
	for _, m := range []string{"eth_getTransactionByHash", "eth_call", "eth_getTransactionReceipt"} {
		f.rpcErr[m] = jsonRPCError{Code: -32000, Message: "unavailable"}
	}
}

func TestHandleAddReplicaV2FromEventAlone(t *testing.T) {
	f := newFakeRPC(t)
	c := managerOn(t, f)
	failAllRPC(f)
	store := common.HexToAddress("0x5555")
	name := bytes.Repeat([]byte{7}, 96)
	proof := []byte("inner proof bytes")
	elog := v2Log(t, pieceABI(t), "AddReplicaV2", c.PieceAddr, store, uint64(41), uint64(9), name, uint64(3), uint8(2), proof)

	rc, err := c.HandleAddReplica(elog, pieceABI(t))
	if err != nil {
		t.Fatalf("HandleAddReplica(AddReplicaV2): %v", err)
	}
	if rc.StoredOn != store || rc.Serial != 41 || rc.Witness.Index != 9 || rc.Piece != 3 || rc.Index != 2 || !bytes.Equal(rc.Witness.Proof, proof) {
		t.Fatalf("decoded %+v", rc)
	}
	if rc.Name == "" {
		t.Fatal("no replica name")
	}
	if n := f.count("eth_getTransactionByHash"); n != 0 {
		t.Fatalf("fetched the transaction %d times", n)
	}
}

func TestHandleAddPieceV2FromEventAlone(t *testing.T) {
	f := newFakeRPC(t)
	c := managerOn(t, f)
	failAllRPC(f)
	owner := common.HexToAddress("0x0a11")
	streamer := common.HexToAddress("0x5717")
	name := bytes.Repeat([]byte{9}, 96)
	elog := v2Log(t, pieceABI(t), "AddPieceV2", c.PieceAddr, owner,
		uint64(5), uint64(77), name, big.NewInt(1e11), uint64(4096), uint64(900), uint8(6), uint8(4), streamer)

	pc, err := c.HandleAddPiece(elog, pieceABI(t))
	if err != nil {
		t.Fatalf("HandleAddPiece(AddPieceV2): %v", err)
	}
	if pc.Owner != owner || pc.Serial != 5 || pc.Start != 77 || pc.Price.Cmp(big.NewInt(1e11)) != 0 ||
		pc.Size != 4096 || pc.Expire != 900 || pc.Policy.N != 6 || pc.Policy.K != 4 || pc.Streamer != streamer || pc.Name == "" {
		t.Fatalf("decoded %+v", pc)
	}
}

func TestHandleSubmitV2FromEventAlone(t *testing.T) {
	f := newFakeRPC(t)
	c := managerOn(t, f)
	failAllRPC(f)
	store := common.HexToAddress("0x5555")
	sum := bytes.Repeat([]byte{1}, 96)
	pf := bytes.Repeat([]byte{2}, 144)
	elog := v2Log(t, eproofABI(t), "SubmitV2", c.EProofAddr, store, uint64(12), sum, pf)

	ei, err := c.HandleSubmitEProof(elog, eproofABI(t))
	if err != nil {
		t.Fatalf("HandleSubmitEProof(SubmitV2): %v", err)
	}
	if ei.Store != store || ei.Epoch != 12 || !bytes.Equal(ei.Hash, crypto.Keccak256(sum, pf)) || len(ei.Sum) == 0 {
		t.Fatalf("decoded %+v", ei)
	}
}

// The V2 fields are taken as the log states them, so only the real contract's
// logs count.
func TestV2EventsRefuseForeignLogs(t *testing.T) {
	f := newFakeRPC(t)
	c := managerOn(t, f)
	other := common.HexToAddress("0xbad")
	a := common.HexToAddress("0x5555")
	name := bytes.Repeat([]byte{7}, 96)

	logs := []struct {
		name string
		do   func() error
	}{
		{"AddReplicaV2", func() error {
			_, err := c.HandleAddReplica(v2Log(t, pieceABI(t), "AddReplicaV2", other, a, uint64(1), uint64(0), name, uint64(1), uint8(0), []byte{1}), pieceABI(t))
			return err
		}},
		{"AddPieceV2", func() error {
			_, err := c.HandleAddPiece(v2Log(t, pieceABI(t), "AddPieceV2", other, a, uint64(1), uint64(1), name, big.NewInt(1), uint64(1), uint64(2), uint8(6), uint8(4), a), pieceABI(t))
			return err
		}},
		{"SubmitV2", func() error {
			_, err := c.HandleSubmitEProof(v2Log(t, eproofABI(t), "SubmitV2", other, a, uint64(1), name, []byte{1}), eproofABI(t))
			return err
		}},
	}
	for _, l := range logs {
		if err := l.do(); !errors.Is(err, ErrForeignLog) {
			t.Errorf("%s from another contract: err = %v, want ErrForeignLog", l.name, err)
		}
	}
}

func TestEventIDs(t *testing.T) {
	cabi := pieceABI(t)
	ids := EventIDs(cabi, "AddReplica", "AddReplicaV2", "NoSuchEvent")
	if len(ids) != 2 || ids[0] != cabi.Events["AddReplica"].ID || ids[1] != cabi.Events["AddReplicaV2"].ID {
		t.Fatalf("EventIDs = %v", ids)
	}
}
