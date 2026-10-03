package contract

import (
	"errors"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	etypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/unibaseio/da-sdk-go/contract/v2/go/piece"
)

func pieceABI(t *testing.T) abi.ABI {
	t.Helper()
	a, err := abi.JSON(strings.NewReader(piece.PieceABI))
	if err != nil {
		t.Fatal(err)
	}
	return a
}

func txTo(to *common.Address, data []byte) *etypes.Transaction {
	return etypes.NewTx(&etypes.LegacyTx{To: to, Data: data, Gas: 1, GasPrice: big.NewInt(1), Value: big.NewInt(0)})
}

// Events must be decoded only from a direct call to the emitting contract with
// the method's own selector. Short calldata used to panic every syncing node
// (inputData[4:]); calls routed through a contract carry other arguments.
func TestDirectCallInputs(t *testing.T) {
	cabi := pieceABI(t)
	m := cabi.Methods["addReplica"]
	pieceAddr := common.HexToAddress("0x1000")
	wrapper := common.HexToAddress("0x2000")

	good, err := m.Inputs.Pack([]byte("replica"), uint64(7), uint8(3), []byte{1, 2, 3})
	if err != nil {
		t.Fatal(err)
	}
	call := append(append([]byte{}, m.ID...), good...)

	in, err := directCallInputs(txTo(&pieceAddr, call), pieceAddr, m)
	if err != nil || len(in) != 4 || in[1].(uint64) != 7 {
		t.Fatalf("direct call: %v %v", in, err)
	}

	for _, c := range []struct {
		name string
		tx   *etypes.Transaction
	}{
		{"empty calldata (used to panic)", txTo(&pieceAddr, nil)},
		{"three bytes", txTo(&pieceAddr, []byte{1, 2, 3})},
		{"other method's selector", txTo(&pieceAddr, append(append([]byte{}, cabi.Methods["addPiece"].ID...), good...))},
		{"routed through a wrapper contract", txTo(&wrapper, call)},
		{"contract creation", txTo(nil, call)},
	} {
		if _, err := directCallInputs(c.tx, pieceAddr, m); !errors.Is(err, ErrNotDirectCall) {
			t.Errorf("%s: got %v, want ErrNotDirectCall", c.name, err)
		}
	}

	// decodeAddPieceFields on short calldata errors instead of panicking
	if _, err := decodeAddPieceFields(cabi, []byte{1}); !errors.Is(err, ErrNotDirectCall) {
		t.Errorf("short addPiece calldata: got %v", err)
	}
}
