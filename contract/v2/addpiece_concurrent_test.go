package contract

import (
	"context"
	"math/big"
	"sync"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/unibaseio/da-sdk-go/lib/types"
)

// TestAddPieceConcurrentPairsApprove: the UB token's approve SETS the
// allowance, so concurrent AddPiece calls on one manager must not interleave —
// each approve has to be followed by its own addPiece before the next approve.
// Unserialized, the sends came out approve, approve, addPiece, addPiece: the
// second approve overwrote the first, and one paid piece failed to register.
func TestAddPieceConcurrentPairsApprove(t *testing.T) {
	f := newFakeRPC(t)
	f.knowsSent = true
	// allowance(owner, spender) reads back enough for any approval
	f.callResult["dd62ed3e"] = common.LeftPadBytes(new(big.Int).Lsh(big.NewInt(1), 200).Bytes(), 32)
	cm := managerOn(t, f)

	pc := types.PieceCore{
		Policy: types.Policy{N: 6, K: 4},
		// a real base-sepolia piece name (compressed G1 point)
		Name:   "80aa2ad90c5fdd15a048a03738d53129f022f72c3dd5413167233be4c0279a3a2fd58d8b4824a1040c106c1bc250118f",
		Size:   1 << 20,
		Start:  1,
		Expire: 1000,
		Price:  big.NewInt(1e11),
	}

	const n = 4
	var wg sync.WaitGroup
	errs := make(chan error, n)
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()
			if _, err := cm.AddPieceCtx(ctx, pc); err != nil {
				errs <- err
			}
		}()
	}
	wg.Wait()
	close(errs)
	for err := range errs {
		t.Fatalf("AddPiece: %v", err)
	}

	approve := common.Hex2Bytes("095ea7b3")
	f.mu.Lock()
	defer f.mu.Unlock()
	if len(f.sent) != 2*n {
		t.Fatalf("sent %d txs, want %d", len(f.sent), 2*n)
	}
	for i := 0; i < len(f.sent); i += 2 {
		isApprove := func(j int) bool { return string(f.sent[j].Data()[:4]) == string(approve) }
		if !isApprove(i) || isApprove(i+1) {
			t.Fatalf("tx %d/%d not an approve→addPiece pair: approves and addPieces interleaved", i, i+1)
		}
	}
}
