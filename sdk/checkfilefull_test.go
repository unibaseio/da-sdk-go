package sdk

// CheckFileFull without the 3 GB SRS: the tests use their own toy trusted
// setup — a τ they know — so every commitment and the accumulated opening H
// are computed exactly (C = [f(τ)]G₁, H = [(F(τ)-F(r))/(τ-r)]G₁) and checked
// against the matching verifying key. That lets every rejection be tested
// against an otherwise valid answer.

import (
	"crypto/sha256"
	"encoding/hex"
	"math/big"
	"os"
	"path/filepath"
	"strings"
	"testing"

	bls12377 "github.com/consensys/gnark-crypto/ecc/bls12-377"
	"github.com/consensys/gnark-crypto/ecc/bls12-377/kzg"
	"github.com/ethereum/go-ethereum/common"
	"github.com/unibaseio/da-sdk-go/lib/bls"
	"github.com/unibaseio/da-sdk-go/lib/bls/erasure"
	"github.com/unibaseio/da-sdk-go/lib/types"
)

var testStream = common.HexToAddress("0x0102030405060708090a0b0c0d0e0f1011121314")

// testTau is the toy setup's secret; useToyKey points CheckFileFull at the
// matching verifying key for the duration of a test.
var testTau = func() bls.Fr { var x bls.Fr; x.SetUint64(0x5eed1234abcd); return x }()

func useToyKey(t *testing.T) {
	t.Helper()
	_, _, g1, g2 := bls12377.Generators()
	vk := &bls.VerifyKey{VerifyingKey: new(kzg.VerifyingKey)}
	vk.G1 = g1
	vk.G2[0] = g2
	tb := new(big.Int)
	testTau.BigInt(tb)
	vk.G2[1].ScalarMultiplication(&g2, tb)
	bls.PrecomputeLines(vk.VerifyingKey)
	old := encodingVerifyKey
	encodingVerifyKey = func() (*bls.VerifyKey, error) { return vk, nil }
	t.Cleanup(func() { encodingVerifyKey = old })
}

func gmul(p *bls.G1, s bls.Fr) {
	var b big.Int
	s.BigInt(&b)
	p.ScalarMultiplicationBase(&b)
}

func pow(x bls.Fr, e int64) bls.Fr {
	var r bls.Fr
	r.Exp(x, big.NewInt(e))
	return r
}

// fakeEncode answers like an honest stream would for data held in one
// piece, under the toy setup.
func fakeEncode(t *testing.T, data []byte, pol types.Policy, stream common.Address) types.FileFull {
	t.Helper()
	useToyKey(t)
	n, k := int(pol.N), int(pol.K)
	ew := bls.NewEncodeWitness(n, k)
	size := int64(len(data))
	slen := (1 + (size-1)/(31*int64(k))) * 31
	elems := slen / 31

	// shard polynomials and their values at τ
	shards := make([][]bls.Fr, k)
	atTau := make([]bls.Fr, k)
	rest := data
	for j := 0; j < k; j++ {
		m := slen
		if int64(len(rest)) < m {
			m = int64(len(rest))
		}
		shards[j] = bls.Split(31, rest[:m])
		if len(shards[j]) == 0 {
			shards[j] = []bls.Fr{{}}
		}
		atTau[j] = bls.Eval(shards[j], testTau)
		rest = rest[m:]
	}

	rs, err := erasure.NewRS(n, k)
	if err != nil {
		t.Fatal(err)
	}
	at := make([]int, n-k)
	for i := range at {
		at[i] = k + i
	}
	par, err := rs.EncodeFr(atTau, at)
	if err != nil {
		t.Fatal(err)
	}
	k2 := pow(testTau, int64(bls.MaxShard)-elems)
	mT := make([]bls.Fr, k)
	lT := make([]bls.Fr, k)
	for j := 0; j < k; j++ {
		k1 := pow(testTau, int64(j)*elems)
		mT[j].Mul(&atTau[j], &k1)
		lT[j].Mul(&atTau[j], &k2)
		gmul(&ew.Commits[j], atTau[j])
		gmul(&ew.MoveCommits[j], mT[j])
		gmul(&ew.LimitCommits[j], lT[j])
		ew.Root.Add(&ew.Root, &ew.MoveCommits[j])
	}
	for i := range par {
		gmul(&ew.Commits[k+i], par[i])
	}

	var r bls.Fr
	r.SetBytes(ew.Challenge(stream.Bytes()))
	for j := 0; j < k; j++ {
		ew.ClaimedValues[j] = bls.Eval(shards[j], r)
	}
	ew.H = toyOpening(ew, r, elems, atTau, mT, lT)

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

// toyOpening is H = [(F(τ)-F(r))/(τ-r)]G₁ for the format-1 accumulated
// polynomial F, given the discrete logs of C_i, M_i, L_i (known only under
// the toy setup). Forgery tests pass the logs of their forged commitments.
func toyOpening(ew *bls.EncodeWitness, r bls.Fr, elems int64, cT, mT, lT []bls.Fr) bls.G1 {
	k := len(ew.MoveCommits)
	var fTau, fR bls.Fr
	rk2 := pow(r, int64(bls.MaxShard)-elems)
	for i := 0; i < k; i++ {
		y := ew.ClaimedValues[i]
		v := bls.AccCoefficient(y)
		var v2 bls.Fr
		v2.Mul(&v, &v)
		// F(r): y·(1 + v r^k1 + v² r^k2)
		rk1 := pow(r, int64(i)*elems)
		var f, tmp bls.Fr
		f.Mul(&v2, &rk2)
		tmp.Mul(&v, &rk1)
		f.Add(&f, &tmp)
		tmp.SetOne()
		f.Add(&f, &tmp).Mul(&f, &y)
		fR.Add(&fR, &f)
		// F(τ): c + v·m + v²·l
		f.Mul(&v2, &lT[i])
		tmp.Mul(&v, &mT[i])
		f.Add(&f, &tmp).Add(&f, &cT[i])
		fTau.Add(&fTau, &f)
	}
	var q, d bls.Fr
	q.Sub(&fTau, &fR)
	d.Sub(&testTau, &r)
	d.Inverse(&d)
	q.Mul(&q, &d)
	var h bls.G1
	gmul(&h, q)
	return h
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

// reencode rewrites a witness into the receipt after a test changed it.
func reencode(ff *types.FileFull, ew *bls.EncodeWitness) {
	root := ew.Root.Bytes()
	ff.Pieces[0] = hex.EncodeToString(root[:])
	ff.Proofs[0] = ew.Serialize()
}

func witnessOf(t *testing.T, ff types.FileFull) *bls.EncodeWitness {
	t.Helper()
	ew := new(bls.EncodeWitness)
	if err := ew.Deserialize(ff.Proofs[0]); err != nil {
		t.Fatal(err)
	}
	return ew
}

// The accumulated opening is checked at all: before 2026-10-06 (audit N8)
// CheckFileFull never looked at H, so any H — and any move commitment the
// root was rebuilt from — passed.
func TestCheckFileFullChecksOpening(t *testing.T) {
	data := testData()
	pol := types.Policy{N: 6, K: 4}
	fp := writeTemp(t, data)

	// a wrong H
	ff := fakeEncode(t, data, pol, testStream)
	ew := witnessOf(t, ff)
	_, _, g1, _ := bls12377.Generators()
	ew.H.Add(&ew.H, &g1)
	reencode(&ff, ew)
	if _, err := CheckFileFull(ff, testStream, fp); err == nil || !strings.Contains(err.Error(), "opening") {
		t.Fatalf("wrong H: %v", err)
	}

	// a move commitment that is not the shard's, the root rebuilt from it
	// (so the root/RS checks pass) and the honest H kept
	ff = fakeEncode(t, data, pol, testStream)
	ew = witnessOf(t, ff)
	ew.MoveCommits[1].Add(&ew.MoveCommits[1], &g1)
	ew.Root.Add(&ew.Root, &g1)
	reencode(&ff, ew)
	if _, err := CheckFileFull(ff, testStream, fp); err == nil || !strings.Contains(err.Error(), "opening") {
		t.Fatalf("forged move commitment: %v", err)
	}
}

// Format 1's known limit (audit N1): the data commitments enter the opening
// with coefficient 1, so shifting Δ from one to another — with everything
// else honest, recomputed by an attacker who has the data but not τ — still
// verifies. This test pins that limit down; format 2 removes it.
func TestFormat1SumShiftStillVerifies(t *testing.T) {
	data := testData()
	pol := types.Policy{N: 6, K: 4}
	fp := writeTemp(t, data)
	ff := fakeEncode(t, data, pol, testStream)
	ew := witnessOf(t, ff)
	k := 4
	size := int64(len(data))
	elems := 1 + (size-1)/(31*int64(k))

	// recover the honest discrete logs the way fakeEncode built them
	slen := elems * 31
	cT := make([]bls.Fr, k)
	mT := make([]bls.Fr, k)
	lT := make([]bls.Fr, k)
	shards := make([][]bls.Fr, k)
	rest := data
	k2 := pow(testTau, int64(bls.MaxShard)-elems)
	for j := 0; j < k; j++ {
		m := slen
		if int64(len(rest)) < m {
			m = int64(len(rest))
		}
		shards[j] = bls.Split(31, rest[:m])
		rest = rest[m:]
		cT[j] = bls.Eval(shards[j], testTau)
		k1 := pow(testTau, int64(j)*elems)
		mT[j].Mul(&cT[j], &k1)
		lT[j].Mul(&cT[j], &k2)
	}

	// Δ = [d]G₁ for a d the attacker picks: C₀+Δ, C₁−Δ; parity recomputed
	var d bls.Fr
	d.SetUint64(424242)
	var dG bls.G1
	gmul(&dG, d)
	ew.Commits[0].Add(&ew.Commits[0], &dG)
	ew.Commits[1].Sub(&ew.Commits[1], &dG)
	cT[0].Add(&cT[0], &d)
	cT[1].Sub(&cT[1], &d)
	rs, err := erasure.NewRS(6, k)
	if err != nil {
		t.Fatal(err)
	}
	par, err := rs.EncodeFr(cT, []int{4, 5})
	if err != nil {
		t.Fatal(err)
	}
	for i := range par {
		gmul(&ew.Commits[k+i], par[i])
	}
	var r bls.Fr
	r.SetBytes(ew.Challenge(testStream.Bytes()))
	for j := 0; j < k; j++ {
		ew.ClaimedValues[j] = bls.Eval(shards[j], r)
	}
	// the opening of the attacker's polynomial: the C-part sums to the
	// honest one, so this is the honest opening at r, computable without τ
	honestC := make([]bls.Fr, k)
	copy(honestC, cT)
	honestC[0].Sub(&honestC[0], &d)
	honestC[1].Add(&honestC[1], &d)
	ew.H = toyOpening(ew, r, elems, honestC, mT, lT)
	reencode(&ff, ew)
	if _, err := CheckFileFull(ff, testStream, fp); err != nil {
		t.Fatalf("format 1 was expected to accept the sum shift (its known limit), got %v", err)
	}
}
