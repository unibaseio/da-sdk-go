package sdk

// Fast negative guards for CheckWitness — the structural gate the verifier runs
// before the (expensive) RS-homomorphism check. These lock the early rejections:
// shard-count mismatches and the Σ MoveCommits == Root invariant. A fully-valid
// witness (which also needs a real RS-coded parity set) is exercised by the slow
// end-to-end test in da-go/lib/helper (respecting the da-go→da-sdk-go dep order).

import (
	"math/big"
	"testing"

	"github.com/unibaseio/da-sdk-go/lib/bls"
)

// witnessWithRootConsistent builds an (n,k) witness whose Root == Σ MoveCommits
// but whose parity commits are NOT a valid RS code — so it passes every check
// EXCEPT the final rs.Check. Enough to drive the count/Root branches.
func witnessWithRootConsistent(n, k int) *bls.EncodeWitness {
	ew := bls.NewEncodeWitness(n, k)
	s := int64(1)
	set := func(p *bls.G1) { p.ScalarMultiplicationBase(big.NewInt(s)); s++ }
	for i := 0; i < n; i++ {
		set(&ew.Commits[i])
	}
	for i := 0; i < k; i++ {
		set(&ew.MoveCommits[i])
		set(&ew.LimitCommits[i])
		ew.Root.Add(&ew.Root, &ew.MoveCommits[i]) // Root = Σ MoveCommits
	}
	return ew
}

func TestCheckWitnessRejectsWrongCommitCount(t *testing.T) {
	ew := witnessWithRootConsistent(6, 4)
	ew.Commits = ew.Commits[:5] // n=6 declared, 5 supplied
	if err := CheckWitness(6, 4, ew); err == nil {
		t.Fatal("expected rejection on short Commits")
	}
}

func TestCheckWitnessRejectsWrongMoveCount(t *testing.T) {
	ew := witnessWithRootConsistent(6, 4)
	ew.MoveCommits = ew.MoveCommits[:3] // k=4 declared, 3 supplied
	if err := CheckWitness(6, 4, ew); err == nil {
		t.Fatal("expected rejection on short MoveCommits")
	}
}

func TestCheckWitnessRejectsBadRoot(t *testing.T) {
	ew := witnessWithRootConsistent(6, 4)
	ew.Root.ScalarMultiplicationBase(big.NewInt(424242)) // no longer Σ MoveCommits
	if err := CheckWitness(6, 4, ew); err == nil {
		t.Fatal("expected rejection: Root != Σ MoveCommits")
	}
}
