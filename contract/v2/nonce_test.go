package contract

import (
	"context"
	"math/big"
	"sync"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const anvilRPC = "http://127.0.0.1:8545"

// reachableAnvil skips the test unless a local Anvil is answering.
func reachableAnvil(t *testing.T) {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	cl, err := ethclient.DialContext(ctx, anvilRPC)
	if err != nil {
		t.Skip("no local anvil at " + anvilRPC)
	}
	defer cl.Close()
	if _, err := cl.ChainID(ctx); err != nil {
		t.Skip("no local anvil at " + anvilRPC)
	}
}

// TestNextNonceConcurrent is the regression guard for the nonce-collision fix:
// many goroutines sharing one key (as a store node does — epoch proofs, replica
// adds, challenge answers) must each get a DISTINCT, CONTIGUOUS nonce. Before
// the fix every goroutine read PendingNonceAt independently and collided.
func TestNextNonceConcurrent(t *testing.T) {
	reachableAnvil(t)

	sk, err := crypto.GenerateKey() // fresh account → chain pending nonce 0
	if err != nil {
		t.Fatal(err)
	}
	c := &ContractManage{RPC: anvilRPC, ChainID: big.NewInt(31337), sk: sk}

	const N = 64
	var wg sync.WaitGroup
	nonces := make([]uint64, N)
	errs := make([]error, N)
	for i := 0; i < N; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			nonces[i], errs[i] = c.nextNonce()
		}(i)
	}
	wg.Wait()

	seen := make(map[uint64]bool, N)
	for i := 0; i < N; i++ {
		if errs[i] != nil {
			t.Fatalf("nextNonce[%d]: %v", i, errs[i])
		}
		if seen[nonces[i]] {
			t.Fatalf("duplicate nonce %d handed out twice", nonces[i])
		}
		seen[nonces[i]] = true
	}
	// fresh account: the 64 nonces must be exactly {0..63} — no dup, no gap.
	for n := uint64(0); n < N; n++ {
		if !seen[n] {
			t.Fatalf("missing nonce %d — allocation left a gap", n)
		}
	}
}
