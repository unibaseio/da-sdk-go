package contract

import (
	"math/big"
	"testing"

	"github.com/unibaseio/da-sdk-go/lib/types"
)

// TestAddPieceCost pins the shared cost formula so v1 (AddPiece) and the v2
// /api/seal cost stay identical and match what the contract locks. Expected is
// an independent hand computation:
//
//	_size1 = 1 + (1048576-1)/(31*4)      = 1 + 8456      = 8457
//	_size  = 1 + (8457-1)/(32*1024)      = 1 + 0         = 1
//	val    = (1301-100)*1 * 1e11 + 1e12  = 1.211e14
//	val    = val * 6 (N)                 = 726600000000000
func TestAddPieceCost(t *testing.T) {
	pc := types.PieceCore{
		Policy: types.Policy{N: 6, K: 4},
		Size:   1048576, // 1 MiB
		Start:  100,
		Expire: 1301, // 1201 epochs
		Price:  big.NewInt(1e11),
	}
	want := big.NewInt(726600000000000)
	got := AddPieceCost(pc)
	if got.Cmp(want) != 0 {
		t.Fatalf("AddPieceCost = %s, want %s", got, want)
	}
}
