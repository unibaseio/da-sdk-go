package bls

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"io"
	"math/big"

	"github.com/unibaseio/da-sdk-go/lib/types"

	"github.com/consensys/gnark-crypto/ecc/bls12-377/kzg"
)

// in case too large file
func (pk *PublicKey) GenCommitments(slen int, r io.Reader) ([]types.ICommitment, error) {
	res := make([]types.ICommitment, 0, 1)
	data := make([]byte, MaxSize)

	breakFlag := false
	for !breakFlag {
		n, err := io.ReadAtLeast(r, data, MaxSize)
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			breakFlag = true
		} else if err != nil {
			return nil, err
		}
		if n == 0 {
			continue
		}

		c, err := pk.GenCommitment(slen, data[:n], 0)
		if err != nil {
			return nil, err
		}
		res = append(res, c)
	}

	return res, nil
}

func (pk *PublicKey) GenCommitment(slen int, d []byte, offset int) (types.ICommitment, error) {
	if len(d) == 32 && slen > 32 {
		return pk.genCommitmentInTree(slen, d)
	}

	if len(d) == 0 {
		return nil, fmt.Errorf("zero size")
	}

	shards := Split(slen, d)
	if len(shards) > MaxShard {
		return nil, fmt.Errorf("invalid data shards %d: too large", len(shards))
	}

	srs := kzg.ProvingKey{
		G1: pk.Pk.G1[offset:],
	}
	temp, err := kzg.Commit(shards, srs)
	if err != nil {
		return nil, err
	}

	return NewCommitment(temp.Marshal())
}

// for commit tree
func (pk *PublicKey) genCommitmentInTree(slen int, v []byte) (types.ICommitment, error) {
	if len(v) != 32 {
		return nil, fmt.Errorf("zero size")
	}

	if slen < 2 {
		return nil, fmt.Errorf("small slen")
	}

	shards := make([]Fr, slen)
	var fr_v, fr_e, fr_c Fr
	nv := sha256.Sum256(v[:20])
	fr_v.SetBytes(nv[:])
	fr_big := new(big.Int)
	fr_v.BigInt(fr_big)
	fr_e.Exp(fr_v, fr_big)
	fr_c.SetBytes(v)

	for i := 0; i < slen; i++ {
		fr_e.Mul(&fr_e, &fr_v)
		shards[i].Add(&fr_e, &fr_c)
	}

	temp, err := kzg.Commit(shards, pk.Pk)
	if err != nil {
		return nil, err
	}

	return NewCommitment(temp.Marshal())
}

func (pk *PublicKey) GenProofs(ic types.IChallenge, slen int, r io.Reader) ([]types.IProof, error) {
	res := make([]types.IProof, 0, 1)
	data := make([]byte, MaxSize)
	breakFlag := false
	for !breakFlag {
		n, err := io.ReadAtLeast(r, data, MaxSize)
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			breakFlag = true
		} else if err != nil {
			return nil, err
		}

		if n == 0 {
			continue
		}

		p, err := pk.GenProof(ic, slen, data[:n])
		if err != nil {
			return nil, err
		}
		res = append(res, p)
	}
	return res, nil
}

func (pk *PublicKey) GenProof(ic types.IChallenge, slen int, d []byte) (types.IProof, error) {
	switch ic.Type() {
	case 0, 1:
		return pk.GenProof1(ic, slen, d)
	case 2:
		return pk.GenProof2(ic, slen, d)
	default:
		return nil, fmt.Errorf("unsupported proof type: %d", ic.Type())
	}
}

var _ types.IVerifyKey = (*VerifyKey)(nil)

type VerifyKey struct {
	*kzg.VerifyingKey
}

func NewKZGVerifyKey(data []byte) (*VerifyKey, error) {
	vk := &VerifyKey{
		VerifyingKey: new(kzg.VerifyingKey),
	}
	err := vk.Deserialize(data)
	return vk, err
}

func (vk *VerifyKey) VerifyProof(ir types.IChallenge, ip types.IProof) error {
	switch ip.Type() {
	case 1:
		return vk.VerifyProof1(ir, ip)
	case 2:
		return vk.VerifyProof2(ir, ip)
	default:
		return fmt.Errorf("unsupported proof type: %d", ip.Type())
	}
}

func (vk *VerifyKey) Serialize() []byte {
	var w bytes.Buffer
	_, err := vk.VerifyingKey.WriteTo(&w)
	if err != nil {
		panic(err)
	}
	return w.Bytes()
}

func (vk *VerifyKey) Deserialize(buf []byte) error {
	if vk.VerifyingKey == nil {
		vk.VerifyingKey = new(kzg.VerifyingKey)
	}
	_, err := vk.VerifyingKey.ReadFrom(bytes.NewReader(buf))
	return err
}

var _ types.IPublicKey = (*PublicKey)(nil)

type PublicKey struct {
	*kzg.SRS
}

func NewKZGPublicKey(data []byte) (*PublicKey, error) {
	pk := &PublicKey{
		SRS: new(kzg.SRS),
	}
	err := pk.Deserialize(data)
	return pk, err
}

func (pk *PublicKey) VerifyKey() types.IVerifyKey {
	return &VerifyKey{
		VerifyingKey: &pk.Vk,
	}
}

func (pk *PublicKey) Serialize() []byte {
	var w bytes.Buffer
	_, err := pk.SRS.Pk.WriteRawTo(&w)
	if err != nil {
		panic(err)
	}
	_, err = pk.SRS.Vk.WriteRawTo(&w)
	if err != nil {
		panic(err)
	}
	return w.Bytes()
}

func (pk *PublicKey) Deserialize(buf []byte) error {
	r := bytes.NewReader(buf)
	_, err := pk.SRS.Pk.UnsafeReadFrom(r)
	if err != nil {
		return err
	}
	_, err = pk.SRS.Vk.ReadFrom(r)
	return err
}

func GenKZGKey(num uint64, seed *big.Int) *PublicKey {
	kzgSRS, err := kzg.NewSRS(num, seed)
	if err != nil {
		panic(err)
	}

	pk := &PublicKey{
		SRS: kzgSRS,
	}
	return pk
}
