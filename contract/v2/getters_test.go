package contract

import (
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/unibaseio/da-sdk-go/contract/v2/go/eproof"
	"github.com/unibaseio/da-sdk-go/contract/v2/go/node"
	"github.com/unibaseio/da-sdk-go/contract/v2/go/piece"
)

func selector(t *testing.T, abiJSON, method string) string {
	t.Helper()
	a, err := abi.JSON(strings.NewReader(abiJSON))
	if err != nil {
		t.Fatal(err)
	}
	return common.Bytes2Hex(a.Methods[method].ID)
}

func word(v uint64) []byte { return common.LeftPadBytes(new(big.Int).SetUint64(v).Bytes(), 32) }

// The getters the responder (ND3/ND5) and the replication scan (ND2) rely on
// call the deployed contracts' own view functions.
func TestChainGetters(t *testing.T) {
	f := newFakeRPC(t)
	c := managerOn(t, f)
	c.NodeAddr = common.HexToAddress("0x1004")

	f.callResult[selector(t, node.NodeABI, "emergencyLastPauseBlock")] = word(777)
	f.callResult[selector(t, node.NodeABI, "emergencyPaused")] = word(1)
	f.callResult[selector(t, eproof.EProofABI, "challengeWindow")] = word(7)
	f.callResult[selector(t, piece.PieceABI, "delay")] = word(9)

	ep, err := c.GetEmergencyPause()
	if err != nil || !ep.Paused || ep.LastBlock != 777 {
		t.Fatalf("GetEmergencyPause = %+v, %v", ep, err)
	}
	if w, err := c.GetEProofChallengeWindow(); err != nil || w != 7 {
		t.Fatalf("GetEProofChallengeWindow = %d, %v", w, err)
	}
	if d, err := c.GetPieceDelay(); err != nil || d != 9 {
		t.Fatalf("GetPieceDelay = %d, %v", d, err)
	}
}
