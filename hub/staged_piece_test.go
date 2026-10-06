package hub

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/unibaseio/da-sdk-go/lib/types"
)

// A staged piece is registered on the hub's own terms: a stream's receipt
// cannot set the price or expiry (the bond the hub locks), and a receipt for
// another piece, stream, policy or size is refused (audit N10).
func TestStagedPieceCore(t *testing.T) {
	stream := common.HexToAddress("0x5DA4EAB14B739CACD400dE5D79549c38a23fac8B")
	pol := types.Policy{N: 6, K: 4}
	honest := types.PieceReceipt{PieceCore: types.PieceCore{
		Policy: pol, Name: "ab12", Size: 1000, Streamer: stream,
		Price: big.NewInt(1e18), Expire: 1 << 40, Owner: common.HexToAddress("0x1"),
	}}
	pc, err := stagedPieceCore(honest, stream, "AB12", pol, 1000)
	if err != nil {
		t.Fatal(err)
	}
	if pc.Price != nil || pc.Expire != 0 || pc.Start != 0 {
		t.Fatalf("receipt's terms leaked into the registration: price %v expire %d start %d", pc.Price, pc.Expire, pc.Start)
	}
	if pc.Size != 1000 || pc.Streamer != stream || pc.Policy != pol || pc.Name != "ab12" || pc.Owner != (common.Address{}) {
		t.Fatalf("unexpected core %+v", pc)
	}

	bad := map[string]func(*types.PieceReceipt) (common.Address, string, int64){
		"other piece": func(r *types.PieceReceipt) (common.Address, string, int64) { return stream, "cd34", 1000 },
		"other stream": func(r *types.PieceReceipt) (common.Address, string, int64) {
			return common.HexToAddress("0x2"), "ab12", 1000
		},
		"other policy": func(r *types.PieceReceipt) (common.Address, string, int64) {
			r.Policy = types.Policy{N: 14, K: 7}
			return stream, "ab12", 1000
		},
		"other size":   func(r *types.PieceReceipt) (common.Address, string, int64) { return stream, "ab12", 999 },
		"empty volume": func(r *types.PieceReceipt) (common.Address, string, int64) { r.Size = 0; return stream, "ab12", 0 },
	}
	for name, mod := range bad {
		r := honest
		st, pn, size := mod(&r)
		if _, err := stagedPieceCore(r, st, pn, pol, size); err == nil {
			t.Errorf("%s: accepted", name)
		}
	}
}
