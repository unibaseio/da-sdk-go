package bls

import (
	"fmt"
	"math/big"

	"github.com/unibaseio/da-sdk-go/lib/lighthash"
)

func SlothEncode(in, v []byte) error {
	cnt := 1 + (len(in)-1)/PadSize
	if len(in) != PadSize*cnt {
		return fmt.Errorf("invalid in length, should align to %d", PadSize)
	}

	var tmp, fr_v Fr
	fr_v.SetBytes(v)
	big_inv, _ := new(big.Int).SetString(SlothEnc, 16)
	for i := 0; i < cnt; i++ {
		tmp.SetBytes(in[PadSize*i : PadSize*(i+1)])
		tmp.Add(&tmp, &fr_v)
		tmp.Exp(tmp, big_inv)
		copy(in[PadSize*i:PadSize*(i+1)], tmp.Marshal())
	}
	return nil
}

func SlothDecode(in, v []byte) error {
	cnt := 1 + (len(in)-1)/PadSize
	if len(in) != PadSize*cnt {
		return fmt.Errorf("invalid in length, should align to %d", PadSize)
	}
	var tmp, fr_v Fr
	fr_v.SetBytes(v)
	big_e := big.NewInt(SlothDec)
	for i := 0; i < cnt; i++ {
		tmp.SetBytes(in[PadSize*i : PadSize*(i+1)])
		tmp.Exp(tmp, big_e)
		tmp.Sub(&tmp, &fr_v)
		copy(in[PadSize*i:PadSize*(i+1)], tmp.Marshal())
	}
	return nil
}

// 32MB: 10.5s- 12s
// 7-> 7.9s
func SlothEncodeV2(in, v []byte) error {
	cnt := 1 + (len(in)-1)/PadSize
	if len(in) != PadSize*cnt {
		return fmt.Errorf("invalid in length, should align to %d", PadSize)
	}

	h := lighthash.New()
	var tmp, tmpv, fr_v Fr
	fr_v.SetBytes(v)
	big_inv, _ := new(big.Int).SetString(SlothEnc, 16)
	for i := 0; i < cnt; i++ {
		h.Reset()
		h.Write(v)
		tmp.SetInt64(int64(i))
		h.Write(tmp.Marshal())
		tmpv.SetBytes(h.Sum(nil))
		tmp.SetBytes(in[PadSize*i : PadSize*(i+1)])
		tmp.Add(&tmp, &tmpv)
		tmp.Exp(tmp, big_inv)
		copy(in[PadSize*i:PadSize*(i+1)], tmp.Marshal())
	}
	return nil
}

// 32MB: 350ms -> 1.8s
// 160ms -> 780 ms
func SlothDecodeV2(in, v []byte) error {
	cnt := 1 + (len(in)-1)/PadSize
	if len(in) != PadSize*cnt {
		return fmt.Errorf("invalid in length, should align to %d", PadSize)
	}

	h := lighthash.New()
	var tmp, tmpv, fr_v Fr
	fr_v.SetBytes(v)
	big_e := big.NewInt(SlothDec)
	for i := 0; i < cnt; i++ {
		h.Reset()
		h.Write(v)
		tmp.SetInt64(int64(i))
		h.Write(tmp.Marshal())
		tmpv.SetBytes(h.Sum(nil))

		tmp.SetBytes(in[PadSize*i : PadSize*(i+1)])
		tmp.Exp(tmp, big_e)
		tmp.Sub(&tmp, &tmpv)
		copy(in[PadSize*i:PadSize*(i+1)], tmp.Marshal())
	}
	return nil
}

func SlothEncodeV3(in, v []byte) error {
	cnt := 1 + (len(in)-1)/PadSize
	if len(in) != PadSize*cnt {
		return fmt.Errorf("invalid in length, should align to %d", PadSize)
	}

	var tmp, tmpv, fr_v Fr
	fr_v.SetBytes(v)
	tmp.Set(&fr_v)
	for i := 0; i < cnt; i++ {
		tmpv.SetBytes(in[PadSize*i : PadSize*(i+1)])
		tmpv.Add(&tmpv, &fr_v)
		tmpv.Div(&tmpv, &tmp)
		tmp.Set(&tmpv)
		copy(in[PadSize*i:PadSize*(i+1)], tmp.Marshal())
	}
	return nil
}

func SlothDecodeV3(in, v []byte) error {
	cnt := 1 + (len(in)-1)/PadSize
	if len(in) != PadSize*cnt {
		return fmt.Errorf("invalid in length, should align to %d", PadSize)
	}

	var tmp, tmpv, fr_v Fr
	fr_v.SetBytes(v)
	tmp.Set(&fr_v)
	for i := 0; i < cnt; i++ {
		tmpv.SetBytes(in[PadSize*i : PadSize*(i+1)])
		tmp.Mul(&tmp, &tmpv)
		tmp.Sub(&tmp, &fr_v)
		copy(in[PadSize*i:PadSize*(i+1)], tmp.Marshal())
		tmp.Set(&tmpv)
	}
	return nil
}
