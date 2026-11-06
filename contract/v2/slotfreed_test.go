package contract

import (
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	etypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/unibaseio/da-sdk-go/contract/v2/go/piece"
)

func TestHandleSlotFreed(t *testing.T) {
	cabi, err := abi.JSON(strings.NewReader(piece.PieceABI))
	if err != nil {
		t.Fatal(err)
	}
	c := &ContractManage{}
	c.PieceAddr = common.HexToAddress("0xb1ece")
	ev := cabi.Events["SlotFreed"]
	store := common.HexToAddress("0x5701")
	data, err := ev.Inputs.NonIndexed().Pack(uint8(3), uint64(42), store, big.NewInt(1234))
	if err != nil {
		t.Fatal(err)
	}
	log := etypes.Log{Address: c.PieceAddr, Topics: []common.Hash{ev.ID, common.BigToHash(big.NewInt(17))}, Data: data}
	sf, err := c.HandleSlotFreed(log, cabi)
	if err != nil {
		t.Fatal(err)
	}
	if sf.Piece != 17 || sf.Index != 3 || sf.Replica != 42 || sf.Store != store || sf.Refund.Int64() != 1234 {
		t.Fatalf("decoded %+v", sf)
	}
	log.Address = common.HexToAddress("0xbad")
	if _, err := c.HandleSlotFreed(log, cabi); err == nil {
		t.Fatal("a SlotFreed log from another contract was accepted")
	}
}
