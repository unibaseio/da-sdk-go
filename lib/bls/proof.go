package bls

import (
	"fmt"
	"math/big"

	"github.com/unibaseio/da-sdk-go/lib/types"

	bls "github.com/consensys/gnark-crypto/ecc/bls12-377"
	"github.com/consensys/gnark-crypto/ecc/bls12-377/kzg"
)

type EpochProof struct {
	Sum          G1
	ClaimedValue Fr
	H            G1
}

var _ types.IProof = (*Proof)(nil)

type Proof struct {
	// ClaimedValue purported value
	ClaimedValue Fr
	// H quotient polynomial (f - f(r)))/(x-r)
	H G1
}

func NewProof(buf []byte) (*Proof, error) {
	pf := Proof{}
	if len(buf) != FrSize+G1Size {
		return nil, fmt.Errorf("wrong size: %d, need %d", len(buf), FrSize+G1Size)
	}

	pf.ClaimedValue.SetBytes(buf[:FrSize])

	_, err := pf.H.SetBytes(buf[FrSize : FrSize+G1Size])
	if err != nil {
		return nil, err
	}
	return &pf, nil
}

func (pf *Proof) Type() int {
	return 1
}

func (pf *Proof) Add(ip types.IProof) error {
	pr, ok := ip.(*Proof)
	if !ok {
		return fmt.Errorf("wrong proof1")
	}
	pf.ClaimedValue.Add(&pf.ClaimedValue, &pr.ClaimedValue)
	pf.H.Add(&pf.H, &pr.H)
	return nil
}

func (pf *Proof) Serialize() []byte {
	buf := make([]byte, FrSize+G1Size)
	copy(buf[:FrSize], pf.ClaimedValue.Marshal())
	hbyte := pf.H.Bytes()
	copy(buf[FrSize:FrSize+G1Size], hbyte[:])

	return buf
}

func (pk *PublicKey) GenProof1(ic types.IChallenge, typ int, d []byte) (*Proof, error) {
	chal, ok := ic.(*Challenge)
	if !ok {
		return nil, fmt.Errorf("invalid chal")
	}
	if len(d) == 0 {
		return nil, fmt.Errorf("invalid data size: zero")
	}

	shards := Split(typ, d)
	if len(shards) > MaxShard {
		return nil, fmt.Errorf("invalid data shards %d: too large", len(shards))
	}

	if len(shards) < MinShard {
		var fr Fr
		shards = append(shards, fr)
	}

	var fr_r Fr
	fr_r.SetBytes(chal.Random)

	srs := kzg.ProvingKey{
		G1: pk.Pk.G1,
	}
	op, err := kzg.Open(shards, fr_r, srs)
	if err != nil {
		return nil, err
	}

	return &Proof{
		H:            op.H,
		ClaimedValue: op.ClaimedValue,
	}, nil
}

func (vk *VerifyKey) VerifyProof1(ir types.IChallenge, ip types.IProof) error {
	pf, ok := ip.(*Proof)
	if !ok {
		return fmt.Errorf("wrong proof")
	}

	chal, ok := ir.(*Challenge)
	if !ok {
		return fmt.Errorf("wrong chal")
	}

	var fr_r Fr
	fr_r.SetBytes(chal.Random)

	icp := ir.Commitment()
	com, ok := icp.(*Commitment)
	if !ok {
		return fmt.Errorf("wrong proof1 commit")
	}

	// [f(r)]G₁
	var left G1
	bval := new(big.Int)
	pf.ClaimedValue.BigInt(bval)
	left.ScalarMultiplication(&vk.G1, bval)

	// [f(r) - f(s)]G₁
	left.Sub(&left, &com.Value)

	var comG1 G1
	// [f(r) - f(s)]G₁ - r*H
	fr_r.BigInt(bval)
	comG1.ScalarMultiplication(&pf.H, bval)
	left.Sub(&left, &comG1)

	// e([f(r) - f(s)]G₁-r*H, G₂).e(H, [s]G₂) ==? 1
	check, err := bls.PairingCheck(
		[]G1{left, pf.H},
		[]G2{vk.G2[0], vk.G2[1]},
	)
	if err != nil {
		return err
	}
	if !check {
		return fmt.Errorf("can't verify opening proof")
	}
	return nil
}
