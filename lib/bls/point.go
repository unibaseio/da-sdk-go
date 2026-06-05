package bls

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"math/big"

	"github.com/unibaseio/da-sdk-go/lib/types"

	bls "github.com/consensys/gnark-crypto/ecc/bls12-377"
	"github.com/consensys/gnark-crypto/ecc/bls12-377/kzg"
	gkzg "github.com/consensys/gnark-crypto/kzg"
)

func (ppk *PointPublicKey) GenCommitments(slen int, r io.Reader) ([]types.ICommitment, error) {
	return nil, fmt.Errorf("unsupported method")
}

func (ppk *PointPublicKey) GenCommitment(slen int, d []byte) (types.ICommitment, error) {
	if len(d) == 0 {
		return nil, fmt.Errorf("zero size")
	}

	shards := Split(slen, d)
	if len(shards) > MaxShard {
		return nil, fmt.Errorf("invalid data shards %d: too large", len(shards))
	}

	srs := kzg.ProvingKey{
		G1: ppk.G1L,
	}
	temp, err := kzg.Commit(shards, srs)
	if err != nil {
		return nil, err
	}

	return &Commitment{
		Value: temp,
	}, nil
}

func (ppk *PointPublicKey) GenProofs(ic types.IChallenge, slen int, r io.Reader) ([]types.IProof, error) {
	return nil, fmt.Errorf("unsupported method")
}

func (ppk *PointPublicKey) GenProof(ic types.IChallenge, slen int, d []byte) (types.IProof, error) {
	chal, ok := ic.(*Challenge)
	if !ok {
		return nil, fmt.Errorf("invalid chal")
	}

	if len(d) == 0 {
		return nil, fmt.Errorf("zero size")
	}

	shards := Split(slen, d)
	if len(shards) > MaxShard {
		return nil, fmt.Errorf("invalid data shards %d: too large", len(shards))
	}

	rnd := int(binary.BigEndian.Uint64(chal.Random))
	if rnd >= len(shards) {
		return nil, fmt.Errorf("random too large")
	}

	if len(shards) < MinShard {
		var fr Fr
		shards = append(shards, fr)
	}

	res := &Proof{
		ClaimedValue: shards[rnd],
	}

	if rnd > 0 {
		srs := kzg.ProvingKey{
			G1: ppk.G1L[ppk.N-rnd:],
		}
		sum, err := kzg.Commit(shards[:rnd], srs)
		if err != nil {
			return nil, err
		}
		res.H.Add(&res.H, &sum)
	}

	if rnd < ppk.N-1 && rnd < len(shards)-1 {
		srs := kzg.ProvingKey{
			G1: ppk.G1R[:ppk.N-rnd-1],
		}

		sum, err := kzg.Commit(shards[rnd+1:], srs)
		if err != nil {
			return nil, err
		}
		res.H.Add(&res.H, &sum)
	}
	return res, nil
}

func (pvk *PointVerifyKey) VerifyProof(ic types.IChallenge, ip types.IProof) error {
	chal, ok := ic.(*Challenge)
	if !ok {
		return fmt.Errorf("wrong chal")
	}

	pr, ok := ip.(*Proof)
	if !ok {
		return fmt.Errorf("wrong proof")
	}

	rnd := int(binary.BigEndian.Uint64(chal.Random))
	if rnd > pvk.N {
		return fmt.Errorf("invalid challenge random")
	}

	// should handle zero?
	if pr.ClaimedValue.IsZero() {
		var right G1
		right.Neg(&pr.H)
		gt, err := bls.Pair([]G1{chal.Sum, right}, []G2{pvk.G2[pvk.N-rnd], pvk.G2[0]})
		if err != nil {
			return err
		}

		var onegt GT
		onegt.SetOne()
		if !gt.Equal(&onegt) {
			return fmt.Errorf("can't verify point proof at sero value")
		}
	} else {
		bval := new(big.Int)
		var val Fr
		val.Inverse(&pr.ClaimedValue)
		val.BigInt(bval)

		var left, right G1
		left.ScalarMultiplication(&chal.Sum, bval)
		right.ScalarMultiplication(&pr.H, bval)
		right.Neg(&right)

		gt, err := bls.Pair([]G1{left, right}, []G2{pvk.G2[pvk.N-rnd], pvk.G2[0]})
		if err != nil {
			return err
		}
		if !gt.Equal(&pvk.GT) {
			return fmt.Errorf("can't verify point proof")
		}
	}

	return nil
}

type MultiProof struct {
	H            G1
	ClaimedValue []Fr
}

func (ppk *PointPublicKey) GenParamProof(randoms []int) (*MultiProof, error) {
	shards := make([]Fr, ppk.N)
	for i := 0; i < ppk.N; i++ {
		shards[i].SetInt64(int64(i + 1))
	}
	return ppk.genMultiProof(randoms, shards)
}

func (ppk *PointPublicKey) GenMultiProof(randoms []int, slen int, d []byte) (*MultiProof, error) {
	if len(d) == 0 {
		return nil, fmt.Errorf("invalid data size: zero")
	}

	shards := Split(slen, d)
	if len(shards) > ppk.N {
		return nil, fmt.Errorf("data size too large")
	}
	return ppk.genMultiProof(randoms, shards)
}

func (ppk *PointPublicKey) genMultiProof(randoms []int, shards []Fr) (*MultiProof, error) {
	before := make([]Fr, ppk.N)
	after := make([]Fr, ppk.N-1)

	res := &MultiProof{
		ClaimedValue: make([]Fr, len(randoms)),
	}

	for i, rnd := range randoms {
		if rnd >= len(shards) {
			return nil, fmt.Errorf("random too large")
		}
		res.ClaimedValue[i].Set(&shards[rnd])

		for j := 0; j < rnd; j++ {
			before[ppk.N-rnd+j].Add(&before[ppk.N-rnd+j], &shards[j])
		}
		for j := 0; j < ppk.N-rnd-1 && j < len(shards)-rnd-1; j++ {
			after[j].Add(&after[j], &shards[rnd+1+j])
		}
	}

	srs := kzg.ProvingKey{
		G1: ppk.G1L,
	}
	sum, err := kzg.Commit(before, srs)
	if err != nil {
		return nil, err
	}
	res.H.Add(&res.H, &sum)

	srs = kzg.ProvingKey{
		G1: ppk.G1R,
	}
	sum1, err := kzg.Commit(after, srs)
	if err != nil {
		return nil, err
	}
	res.H.Add(&res.H, &sum1)

	return res, nil
}

func (pvk *PointVerifyKey) VerifyMultiProof(commit G1, randoms []int, pf *MultiProof) error {

	var gsum G2
	var vsum Fr
	for i, rnd := range randoms {
		if rnd > pvk.N {
			return fmt.Errorf("invalid challenge random")
		}
		gsum.Add(&gsum, &pvk.G2[pvk.N-rnd])
		vsum.Add(&vsum, &pf.ClaimedValue[i])

	}

	bval := new(big.Int)
	var val Fr
	val.Inverse(&vsum)
	val.BigInt(bval)

	var left, right G1
	left.ScalarMultiplication(&commit, bval)
	right.ScalarMultiplication(&pf.H, bval)
	right.Neg(&right)

	gt, err := bls.Pair([]G1{left, right}, []G2{gsum, pvk.G2[0]})
	if err != nil {
		return err
	}
	if !gt.Equal(&pvk.GT) {
		return fmt.Errorf("can't verify multi point proof")
	}

	return nil
}

type PointCircuitKey struct {
	G1 G1 // sum of g1s
	G2 G2 // g2 one
	GT GT
}

func (pck *PointCircuitKey) Serialize() []byte {
	var w bytes.Buffer
	enc := bls.NewEncoder(&w, bls.RawEncoding())
	toEncode := []interface{}{
		pck.G1,
		pck.G2,
		pck.GT,
	}
	for _, v := range toEncode {
		if err := enc.Encode(v); err != nil {
			panic(err)
		}
	}

	return w.Bytes()
}

func (pck *PointCircuitKey) Deserialize(res []byte) error {
	dec := bls.NewDecoder(bytes.NewReader(res), bls.NoSubgroupChecks())
	toDecode := []interface{}{
		&pck.G1,
		&pck.G2,
		&pck.GT,
	}
	for _, v := range toDecode {
		if err := dec.Decode(v); err != nil {
			return err
		}
	}
	return nil
}

var _ types.IVerifyKey = (*PointVerifyKey)(nil)

type PointVerifyKey struct {
	N  int
	G2 []G2 // N+1
	GT GT
}

func NewPointVerifyKey(data []byte) (*PointVerifyKey, error) {
	pvk := new(PointVerifyKey)
	err := pvk.Deserialize(data)
	return pvk, err
}

func (pvk *PointVerifyKey) Serialize() []byte {
	var w bytes.Buffer
	enc := bls.NewEncoder(&w, bls.RawEncoding())
	toEncode := []interface{}{
		pvk.G2,
		pvk.GT,
	}
	for _, v := range toEncode {
		if err := enc.Encode(v); err != nil {
			panic(err)
		}
	}

	return w.Bytes()
}

func (pvk *PointVerifyKey) Deserialize(res []byte) error {
	dec := bls.NewDecoder(bytes.NewReader(res), bls.NoSubgroupChecks())
	toDecode := []interface{}{
		&pvk.G2,
		&pvk.GT,
	}
	for _, v := range toDecode {
		if err := dec.Decode(v); err != nil {
			return err
		}
	}
	pvk.N = len(pvk.G2) - 1
	return nil
}

var _ gkzg.SRS = (*PointPublicKey)(nil)

type PointPublicKey struct {
	N      int
	G1L    []G1 //N
	G1R    []G1 //N-1
	G2     []G2 //N+1
	GT     GT
	Commit G1
}

func NewPointPublicKey(data []byte) (*PointPublicKey, error) {
	ppk := new(PointPublicKey)
	err := ppk.Deserialize(data)
	ppk.CalCommit()
	return ppk, err
}

func (ppk *PointPublicKey) WriteDump(w io.Writer, maxPkPoints ...int) error {
	return fmt.Errorf("method unsupported")
}

func (ppk *PointPublicKey) ReadDump(r io.Reader, maxPkPoints ...int) error {
	return fmt.Errorf("method unsupported")
}

func (ppk *PointPublicKey) ToKZG() *PublicKey {
	return &PublicKey{
		SRS: &kzg.SRS{
			Pk: kzg.ProvingKey{
				G1: ppk.G1L,
			},
			Vk: kzg.VerifyingKey{
				G1: ppk.G1L[0],
				G2: [2]G2{ppk.G2[0], ppk.G2[1]},
			},
		},
	}
}

func (ppk *PointPublicKey) ToCirciutKey() *PointCircuitKey {
	return &PointCircuitKey{
		G1: ppk.Commit,
		G2: ppk.G2[0],
		GT: ppk.GT,
	}
}

func (ppk *PointPublicKey) ToVerifyKey() *PointVerifyKey {
	return &PointVerifyKey{
		N:  ppk.N,
		G2: ppk.G2,
		GT: ppk.GT,
	}
}

func (ppk *PointPublicKey) VerifyKey() types.IVerifyKey {
	return ppk.ToVerifyKey()
}

func (ppk *PointPublicKey) WriteTo(w io.Writer) (int64, error) {
	enc := bls.NewEncoder(w)
	toEncode := []interface{}{
		ppk.G1L,
		ppk.G1R,
		ppk.G2,
		ppk.GT,
	}
	for _, v := range toEncode {
		if err := enc.Encode(v); err != nil {
			panic(err)
		}
	}

	return enc.BytesWritten(), nil
}

func (ppk *PointPublicKey) WriteRawTo(w io.Writer) (int64, error) {
	// encode the SRS
	enc := bls.NewEncoder(w, bls.RawEncoding())
	toEncode := []interface{}{
		ppk.G1L,
		ppk.G1R,
		ppk.G2,
		ppk.GT,
	}
	for _, v := range toEncode {
		if err := enc.Encode(v); err != nil {
			panic(err)
		}
	}

	return enc.BytesWritten(), nil
}

func (ppk *PointPublicKey) ReadFrom(r io.Reader) (n int64, err error) {
	dec := bls.NewDecoder(r)
	toDecode := []interface{}{
		&ppk.G1L,
		&ppk.G1R,
		&ppk.G2,
		&ppk.GT,
	}
	for _, v := range toDecode {
		if err := dec.Decode(v); err != nil {
			return dec.BytesRead(), err
		}
	}
	ppk.N = len(ppk.G1L)

	return dec.BytesRead(), nil
}

func (ppk *PointPublicKey) UnsafeReadFrom(r io.Reader) (n int64, err error) {
	dec := bls.NewDecoder(r, bls.NoSubgroupChecks())
	toDecode := []interface{}{
		&ppk.G1L,
		&ppk.G1R,
		&ppk.G2,
		&ppk.GT,
	}
	for _, v := range toDecode {
		if err := dec.Decode(v); err != nil {
			return dec.BytesRead(), err
		}
	}
	ppk.N = len(ppk.G1L)

	return dec.BytesRead(), nil
}

func (ppk *PointPublicKey) Serialize() []byte {
	var w bytes.Buffer
	ppk.WriteRawTo(&w)
	return w.Bytes()
}

func (ppk *PointPublicKey) Deserialize(res []byte) error {
	_, err := ppk.UnsafeReadFrom(bytes.NewReader(res))

	return err
}

func (ppk *PointPublicKey) CalCommit() {
	shards := make([]Fr, ppk.N)
	for i := 0; i < ppk.N; i++ {
		shards[i].SetInt64(int64(i + 1))
	}

	srs := kzg.ProvingKey{
		G1: ppk.G1L,
	}
	temp, err := kzg.Commit(shards, srs)
	if err != nil {
		panic(err)
	}
	ppk.Commit = temp
}

func GenPointKey(num uint64, seed *big.Int) *PointPublicKey {
	var alpha Fr
	kzgSRS, err := kzg.NewSRS(num, seed)
	if err != nil {
		return nil
	}

	ppk := &PointPublicKey{
		N:   int(num),
		G1L: kzgSRS.Pk.G1,
		G2:  make([]G2, num+1),
		G1R: make([]G1, num-1),
	}
	ppk.G2[0] = kzgSRS.Vk.G2[0]

	alpha.SetBigInt(seed)
	alphas := make([]Fr, num)
	alphas[0].Set(&alpha)
	for i := 1; i < int(num); i++ {
		alphas[i].Mul(&alphas[i-1], &alpha)
	}
	g2s := bls.BatchScalarMultiplicationG2(&kzgSRS.Vk.G2[0], alphas)
	copy(ppk.G2[1:], g2s)

	alphas[0].Mul(&alphas[num-1], &alpha) // N+1
	for i := 1; i < int(num)-1; i++ {
		alphas[i].Mul(&alphas[i-1], &alpha)
	}
	g1rs := bls.BatchScalarMultiplicationG1(&kzgSRS.Pk.G1[0], alphas[:num-1])
	copy(ppk.G1R, g1rs)

	gt, err := bls.Pair([]G1{ppk.G1L[num-1]}, []G2{ppk.G2[1]})
	if err != nil {
		panic(err)
	}
	ppk.GT.Set(&gt)
	ppk.CalCommit()

	return ppk
}
