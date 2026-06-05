package bls

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"math/big"

	"github.com/unibaseio/da-sdk-go/lib/types"
	bls12377 "github.com/consensys/gnark-crypto/ecc/bls12-377"
	"github.com/consensys/gnark-crypto/ecc/bls12-377/kzg"
)

var _ types.IProof = (*MoveProof)(nil)

type MoveProof struct {
	// ClaimedValue purported value
	ClaimedValue Fr

	// H quotient polynomial (f - f(r)))/(x-r)
	H [3]G1
}

func NewMoveProof(buf []byte) (*MoveProof, error) {
	pf := MoveProof{}
	if len(buf) != FrSize+3*G1Size {
		return nil, fmt.Errorf("wrong size: %d, need %d", len(buf), FrSize+2*G1Size)
	}

	pf.ClaimedValue.SetBytes(buf[:FrSize])
	for i := 0; i < 3; i++ {
		_, err := pf.H[i].SetBytes(buf[FrSize+i*G1Size : FrSize+(i+1)*G1Size])
		if err != nil {
			return nil, err
		}
	}
	return &pf, nil
}

func (pf *MoveProof) Type() int {
	return 2
}

func (pf *MoveProof) Add(ip types.IProof) error {
	pr, ok := ip.(*MoveProof)
	if !ok {
		return fmt.Errorf("wrong move proof")
	}
	pf.ClaimedValue.Add(&pf.ClaimedValue, &pr.ClaimedValue)
	for i := 0; i < 3; i++ {
		pf.H[i].Add(&pf.H[i], &pr.H[i])
	}
	return nil
}

func (pf *MoveProof) Serialize() []byte {
	buf := make([]byte, FrSize+3*G1Size)

	copy(buf[:FrSize], pf.ClaimedValue.Marshal())
	for i := 0; i < len(pf.H); i++ {
		copy(buf[FrSize+i*G1Size:FrSize+(i+1)*G1Size], pf.H[i].Marshal())
	}

	return buf
}

func (pk *PublicKey) GenProof2(ic types.IChallenge, slen int, d []byte) (*MoveProof, error) {
	chal, ok := ic.(*Challenge)
	if !ok {
		return nil, fmt.Errorf("invalid chal")
	}

	if len(chal.Random) != 48 {
		return nil, fmt.Errorf("invalid chal random length")
	}

	if len(d) == 0 {
		return nil, fmt.Errorf("invalid data size: zero")
	}

	shards := Split(slen, d)
	if len(shards) > MaxShard {
		return nil, fmt.Errorf("invalid data shards %d: too large", len(shards))
	}

	if len(shards) < MinShard {
		var fr Fr
		shards = append(shards, fr)
	}

	var fr_r Fr
	fr_r.SetBytes(chal.Random[:32])

	srs := kzg.ProvingKey{
		G1: pk.Pk.G1,
	}
	op, err := kzg.Open(shards, fr_r, srs)
	if err != nil {
		return nil, err
	}

	k := binary.BigEndian.Uint64(chal.Random[32:40])
	if int(k)+len(shards) > MaxShard {
		return nil, fmt.Errorf("data size too large")
	}
	// (x^k*f(x)-r^k*f(r))/(x-r)
	nshards := make([]Fr, k)
	nshards = append(nshards, shards...)
	op2, err := kzg.Open(nshards, fr_r, srs)
	if err != nil {
		return nil, err
	}

	k = binary.BigEndian.Uint64(chal.Random[40:48])
	if int(k)+len(shards) > MaxShard {
		return nil, fmt.Errorf("data size too large")
	}
	// (x^k*f(x)-r^k*f(r))/(x-r)
	nshards = make([]Fr, k)
	nshards = append(nshards, shards...)
	op3, err := kzg.Open(nshards, fr_r, srs)
	if err != nil {
		return nil, err
	}

	return &MoveProof{
		H:            [3]G1{op.H, op2.H, op3.H},
		ClaimedValue: op.ClaimedValue,
	}, nil
}

func (vk *VerifyKey) VerifyProof2(ir types.IChallenge, ip types.IProof) error {
	pf, ok := ip.(*MoveProof)
	if !ok {
		return fmt.Errorf("wrong proof4")
	}

	chal, ok := ir.(*Challenge)
	if !ok {
		return fmt.Errorf("wrong chal")
	}

	if len(chal.Random) != 48 {
		return fmt.Errorf("invalid chal random")
	}

	if len(chal.Digests) != 3 {
		return fmt.Errorf("proof4 wrong chal commit")
	}

	var fr_vrandom1, fr_vrandom2 Fr
	h := sha256.New()
	h.Write(pf.H[1].Marshal())
	fr_vrandom1.SetBytes(h.Sum(nil))

	h.Reset()
	h.Write(pf.H[2].Marshal())
	fr_vrandom2.SetBytes(h.Sum(nil))

	var fr_r Fr
	fr_r.SetBytes(chal.Random[:32])
	brk := new(big.Int)

	// com0 + v1*com1+v2*com2
	var comsum, hsum, temp G1
	comsum.Set(&chal.Digests[0])
	fr_vrandom1.BigInt(brk)
	temp.ScalarMultiplication(&chal.Digests[1], brk)
	comsum.Add(&comsum, &temp)
	fr_vrandom2.BigInt(brk)
	temp.ScalarMultiplication(&chal.Digests[2], brk)
	comsum.Add(&comsum, &temp)

	// H0 + v1*H1 + v2*H2
	hsum.Set(&pf.H[0])
	fr_vrandom1.BigInt(brk)
	temp.ScalarMultiplication(&pf.H[1], brk)
	hsum.Add(&hsum, &temp)
	fr_vrandom2.BigInt(brk)
	temp.ScalarMultiplication(&pf.H[2], brk)
	hsum.Add(&hsum, &temp)

	// C+r*H
	fr_r.BigInt(brk)
	temp.ScalarMultiplication(&hsum, brk)
	comsum.Add(&comsum, &temp)

	// y(1+v1*r^k1+v2*r^k2)
	var cv, fr_rk Fr
	cv.SetOne()
	brk.SetUint64(binary.BigEndian.Uint64(chal.Random[32:40]))
	fr_rk.Exp(fr_r, brk)
	fr_rk.Mul(&fr_rk, &fr_vrandom1)
	cv.Add(&cv, &fr_rk)
	brk.SetUint64(binary.BigEndian.Uint64(chal.Random[40:48]))
	fr_rk.Exp(fr_r, brk)
	fr_rk.Mul(&fr_rk, &fr_vrandom2)
	cv.Add(&cv, &fr_rk)
	cv.Mul(&cv, &pf.ClaimedValue)

	// [f(r)]G₁
	var left, right G1
	cv.BigInt(brk)
	left.ScalarMultiplication(&vk.G1, brk)

	// [f(r) - f(s)]G₁ - r*H
	left.Sub(&left, &comsum)

	right.Set(&hsum)

	// e([f(r) - f(s)]G₁-r*H, G₂).e(H, [s]G₂) ==? 1
	check, err := bls12377.PairingCheck(
		[]G1{left, right},
		[]G2{vk.G2[0], vk.G2[1]},
	)
	if err != nil {
		return err
	}
	if !check {
		return fmt.Errorf("can't verify move opening proof")
	}

	return nil
}
