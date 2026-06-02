package common

import (
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/unibaseio/da-sdk-go/lib/bls"
)

func G1ToString(g bls.G1) string {
	gbyte := g.Bytes()
	return hex.EncodeToString(gbyte[:])
}

func G1InSolidity(g bls.G1) []byte {
	val := new(big.Int)
	g.X.BigInt(val)
	res := ToBytes(6, val)

	g.Y.BigInt(val)
	resy := ToBytes(6, val)
	res = append(res, resy...)
	return res
}

func G1StringInSolidity(s string) ([]byte, error) {
	b, err := hex.DecodeString(s)
	if err != nil {
		return nil, err
	}
	var g bls.G1
	_, err = g.SetBytes(b)
	if err != nil {
		return nil, err
	}

	return G1InSolidity(g), nil
}

func SolidityToG1(buf []byte) (bls.G1, error) {
	var res bls.G1
	if len(buf) != 96 {
		return res, fmt.Errorf("short g1")
	}

	val := ToValue(buf[:48])
	res.X.SetBigInt(val)

	val = ToValue(buf[48:96])
	res.Y.SetBigInt(val)

	return res, nil
}

func FrInSolidity(g bls.Fr) []byte {
	val := new(big.Int)
	g.BigInt(val)
	res := ToBytes(6, val)
	return res
}

func SolidityToFr(buf []byte) (bls.Fr, error) {
	var res bls.Fr
	if len(buf) != 48 {
		return res, fmt.Errorf("short fr")
	}

	val := ToValue(buf[:48])
	res.SetBigInt(val)

	return res, nil
}

func ToBytes(fc int, val *big.Int) []byte {
	base := new(big.Int).Lsh(big.NewInt(1), 64)
	tmpb := new(big.Int).Set(val)

	res := make([]byte, 0, 8*fc)
	for i := 0; i < fc; i++ {
		tmp := new(big.Int).Mod(tmpb, base)
		tmpb.Rsh(tmpb, 64)
		res = append(res, make([]byte, 8-len(tmp.Bytes()))...)
		res = append(res, tmp.Bytes()...)
	}
	return res
}

func ToValue(buf []byte) *big.Int {
	base := big.NewInt(1)
	tmp := new(big.Int)
	res := new(big.Int)

	for i := 0; i < len(buf)/8; i++ {
		tmp.SetBytes(buf[i*8 : (i+1)*8])
		tmp.Mul(tmp, base)
		res.Add(res, tmp)
		base.Lsh(base, 64)
	}
	return res
}
