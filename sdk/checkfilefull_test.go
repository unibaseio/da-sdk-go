package sdk

// CheckFileFull without the SRS: the client check never opens H, so a witness
// built from arbitrary data commitments (with RS-consistent parity commitments
// and claimed values evaluated at the real Fiat-Shamir point) passes it. That
// lets every rejection be tested against an otherwise valid answer.

import (
	"crypto/sha256"
	"encoding/hex"
	"math/big"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/unibaseio/da-sdk-go/lib/bls"
	"github.com/unibaseio/da-sdk-go/lib/bls/erasure"
	"github.com/unibaseio/da-sdk-go/lib/types"
)

var testStream = common.HexToAddress("0x0102030405060708090a0b0c0d0e0f1011121314")

// fakeEncode answers like a stream would for data held in one piece.
func fakeEncode(t *testing.T, data []byte, pol types.Policy, stream common.Address) types.FileFull {
	t.Helper()
	n, k := int(pol.N), int(pol.K)
	ew := bls.NewEncodeWitness(n, k)

	rs, err := erasure.NewRS(n, k)
	if err != nil {
		t.Fatal(err)
	}
	scal := make([]bls.Fr, k)
	for j := 0; j < k; j++ {
		scal[j].SetUint64(uint64(1000 + j))
	}
	at := make([]int, n-k)
	for i := range at {
		at[i] = k + i
	}
	par, err := rs.EncodeFr(scal, at)
	if err != nil {
		t.Fatal(err)
	}
	gmul := func(p *bls.G1, s *bls.Fr) {
		var b big.Int
		s.BigInt(&b)
		p.ScalarMultiplicationBase(&b)
	}
	for j := 0; j < k; j++ {
		gmul(&ew.Commits[j], &scal[j])
		ew.MoveCommits[j].ScalarMultiplicationBase(big.NewInt(int64(2000 + j)))
		ew.LimitCommits[j].ScalarMultiplicationBase(big.NewInt(int64(3000 + j)))
		ew.Root.Add(&ew.Root, &ew.MoveCommits[j])
	}
	for i := range par {
		gmul(&ew.Commits[k+i], &par[i])
	}

	var rnd bls.Fr
	rnd.SetBytes(ew.Challenge(stream.Bytes()))
	size := int64(len(data))
	slen := (1 + (size-1)/(31*int64(k))) * 31
	rest := data
	for j := 0; j < k; j++ {
		m := slen
		if int64(len(rest)) < m {
			m = int64(len(rest))
		}
		ew.ClaimedValues[j] = bls.Eval(bls.Split(31, rest[:m]), rnd)
		rest = rest[m:]
	}

	root := ew.Root.Bytes()
	sum := sha256.Sum256(data)
	return types.FileFull{
		FileReceipt: types.FileReceipt{
			FileCore: types.FileCore{Policy: pol, Hash: hex.EncodeToString(sum[:]), Size: size},
			Pieces:   []string{hex.EncodeToString(root[:])},
		},
		Proofs:     [][]byte{ew.Serialize()},
		PieceSizes: []int64{size},
	}
}

func writeTemp(t *testing.T, data []byte) string {
	t.Helper()
	fp := filepath.Join(t.TempDir(), "f.bin")
	if err := os.WriteFile(fp, data, 0o600); err != nil {
		t.Fatal(err)
	}
	return fp
}

func testData() []byte {
	b := make([]byte, 4000)
	for i := range b {
		b[i] = byte(i*7 + 3)
	}
	return b
}

func TestCheckFileFullAcceptsHonestAnswer(t *testing.T) {
	pol := types.Policy{N: 6, K: 4}
	data := testData()
	ff := fakeEncode(t, data, pol, testStream)
	pcs, err := CheckFileFullPolicy(ff, testStream, writeTemp(t, data), pol)
	if err != nil {
		t.Fatalf("honest answer rejected: %v", err)
	}
	if len(pcs) != 1 || pcs[0].Name != ff.Pieces[0] || pcs[0].Size != int64(len(data)) || pcs[0].Policy != pol {
		t.Fatalf("pieces: %+v", pcs)
	}
	if _, err := CheckFileFull(ff, testStream, writeTemp(t, data)); err != nil {
		t.Fatalf("CheckFileFull (any policy): %v", err)
	}
}

func TestCheckFileFullRejects(t *testing.T) {
	pol := types.Policy{N: 6, K: 4}
	data := testData()

	cases := []struct {
		name   string
		mutate func(ff *types.FileFull) // on the honest answer
		file   []byte                   // nil: data
		want   types.Policy             // zero: pol
		errHas string
	}{
		{name: "unsupported policy", mutate: func(ff *types.FileFull) { ff.Policy = types.Policy{N: 5, K: 4} }, errHas: "policy"},
		{name: "policy not the requested one", want: types.Policy{N: 14, K: 7}, errHas: "requested"},
		{name: "blank hash", mutate: func(ff *types.FileFull) { ff.Hash = "" }, errHas: "sha256"},
		{name: "wrong hash", mutate: func(ff *types.FileFull) { ff.Hash = strings.Repeat("ab", 32) }, errHas: "hash"},
		{name: "proof missing", mutate: func(ff *types.FileFull) { ff.Proofs = nil }, errHas: "proofs"},
		{name: "size missing", mutate: func(ff *types.FileFull) { ff.PieceSizes = nil }, errHas: "sizes"},
		{name: "no pieces for a non-empty file", mutate: func(ff *types.FileFull) {
			ff.Pieces, ff.Proofs, ff.PieceSizes = nil, nil, nil
		}, errHas: "cover"},
		{name: "pieces cover a prefix", mutate: func(ff *types.FileFull) { ff.PieceSizes[0] = 1000 }, errHas: "cover"},
		{name: "piece size over the piece maximum", mutate: func(ff *types.FileFull) {
			ff.PieceSizes[0] = MaxPieceSize(pol) + 1
			ff.Size = ff.PieceSizes[0]
		}, errHas: "size"},
		{name: "receipt size is a prefix of the file", mutate: func(ff *types.FileFull) {
			ff.PieceSizes[0], ff.Size = 1000, 1000
		}, errHas: "file is"},
		{name: "witness for another policy", mutate: func(ff *types.FileFull) {
			ff.Proofs[0] = bls.NewEncodeWitness(14, 7).Serialize()
		}, errHas: "witness"},
		{name: "piece named for another root", mutate: func(ff *types.FileFull) {
			ff.Pieces[0] = strings.Repeat("00", 48)
		}, errHas: "root"},
		{name: "file differs from what was encoded", file: append([]byte{0xff}, data[1:]...), errHas: "unequal"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ff := fakeEncode(t, data, pol, testStream)
			if c.mutate != nil {
				c.mutate(&ff)
			}
			file := c.file
			if file == nil {
				file = data
			}
			want := c.want
			if want == (types.Policy{}) {
				want = pol
			}
			_, err := CheckFileFullPolicy(ff, testStream, writeTemp(t, file), want)
			if err == nil {
				t.Fatal("accepted")
			}
			if !strings.Contains(err.Error(), c.errHas) {
				t.Fatalf("err %q, want it to mention %q", err, c.errHas)
			}
		})
	}
}

func TestCheckFileFullRejectsDirectory(t *testing.T) {
	pol := types.Policy{N: 6, K: 4}
	ff := fakeEncode(t, testData(), pol, testStream)
	if _, err := CheckFileFull(ff, testStream, t.TempDir()); err == nil {
		t.Fatal("accepted a directory")
	}
}

func TestCheckFileFullEmptyFile(t *testing.T) {
	sum := sha256.Sum256(nil)
	ff := types.FileFull{FileReceipt: types.FileReceipt{FileCore: types.FileCore{
		Policy: types.Policy{N: 6, K: 4}, Hash: hex.EncodeToString(sum[:]),
	}}}
	pcs, err := CheckFileFull(ff, testStream, writeTemp(t, nil))
	if err != nil || len(pcs) != 0 {
		t.Fatalf("empty file: %v %v", pcs, err)
	}
}
