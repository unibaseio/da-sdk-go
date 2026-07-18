package erasure

// RS-code homomorphism on the G1 COMMITMENT path.
//
// CLAUDE.md: rs.Check works on BOTH Fr (32B data) and G1 (48/96B commitments)
// via the same linear combination — this homomorphism is exactly why parity
// *commitments* verify as a valid RS code of the data commitments WITHOUT the
// data (CheckWitness → rs.Check). The existing code_test.go covers only the Fr
// path; this file covers the G1 path that the verifier actually relies on.
//
// Slow tier: builds a small KZG key. Skipped under `go test -short`.

import (
	"math/big"
	"testing"

	"github.com/unibaseio/da-sdk-go/lib/bls"
	"github.com/unibaseio/da-sdk-go/lib/utils"
)

// buildRSCommitments mirrors helper.EncodeData's parity construction: commit the
// k data shards, RS-encode them column-wise to n-k parity shards, commit those.
// Returns marshaled data-commit + parity-commit bytes as rs.Check consumes them.
func buildRSCommitments(t *testing.T, n, k, slen int) (dv, pv [][]byte, need []int) {
	t.Helper()
	pk := bls.GenKZGKey(256, big.NewInt(20260717)) // degree ≫ slen; fast

	data := make([][]byte, n)
	for j := 0; j < k; j++ {
		data[j] = utils.RandomBytes(slen * bls.PadSize)
	}
	for j := k; j < n; j++ {
		data[j] = make([]byte, 0, slen*bls.PadSize)
	}

	rs, err := NewRS(n, k)
	if err != nil {
		t.Fatal(err)
	}
	need = make([]int, n-k)
	for i := range need {
		need[i] = k + i
	}

	// column-wise RS encode (PadSize at a time), same as EncodeData
	col := make([][]byte, k)
	for c := 0; c < slen; c++ {
		for j := 0; j < k; j++ {
			col[j] = data[j][c*bls.PadSize : (c+1)*bls.PadSize]
		}
		par, err := rs.Encode(col, need)
		if err != nil {
			t.Fatal(err)
		}
		for j := 0; j < n-k; j++ {
			data[k+j] = append(data[k+j], par[j]...)
		}
	}

	commit := func(d []byte) []byte {
		c, err := pk.GenCommitment(bls.PadSize, d, 0)
		if err != nil {
			t.Fatal(err)
		}
		return c.Raw()
	}
	dv = make([][]byte, k)
	for j := 0; j < k; j++ {
		dv[j] = commit(data[j])
	}
	pv = make([][]byte, n-k)
	for j := 0; j < n-k; j++ {
		pv[j] = commit(data[k+j])
	}
	return dv, pv, need
}

func TestRSCheckG1Homomorphism(t *testing.T) {
	if testing.Short() {
		t.Skip("needs KZG key; slow tier")
	}
	rs, err := NewRS(6, 4)
	if err != nil {
		t.Fatal(err)
	}
	dv, pv, need := buildRSCommitments(t, 6, 4, 2)

	if err := rs.Check(dv, pv, need); err != nil {
		t.Fatalf("valid parity commitments rejected: %v", err)
	}
}

func TestRSCheckG1RejectsCorruptParity(t *testing.T) {
	if testing.Short() {
		t.Skip("needs KZG key; slow tier")
	}
	rs, err := NewRS(6, 4)
	if err != nil {
		t.Fatal(err)
	}
	dv, pv, need := buildRSCommitments(t, 6, 4, 2)

	// flip a parity commitment to a different valid point → must fail the code check
	var bad bls.G1
	bad.ScalarMultiplicationBase(big.NewInt(7))
	pv[0] = bad.Marshal()

	if err := rs.Check(dv, pv, need); err == nil {
		t.Fatal("corrupt parity commitment accepted — G1 homomorphism check is not binding")
	}
}
