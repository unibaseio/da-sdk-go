package lighthash

import (
	"bytes"
	"encoding/hex"
	"math/big"
	"testing"
	"time"

	"github.com/unibaseio/da-sdk-go/lib/merkle"
)

func TestHashSpeed(t *testing.T) {
	h := New()
	nt := time.Now()

	fieldSize := h.BlockSize()
	t.Log(fieldSize, h.Size())
	var value Fr
	bigval := new(big.Int)
	value.SetRandom()
	value.BigInt(bigval)
	bval := bigval.Bytes()
	bval = append(make([]byte, fieldSize-len(bval)), bval...)
	for i := 0; i < 1024*1024; i++ {
		h.Reset()
		h.Write(bval)
		h.Write(bval)
		h.Sum(nil)
	}

	t.Fatal("cost: ", time.Since(nt))
}

func TestMerkleHashSpeed(t *testing.T) {
	h := New()
	nt := time.Now()

	mt := merkle.New(h)

	fieldSize := h.BlockSize()
	t.Log(fieldSize, h.Size())
	var value Fr
	bigval := new(big.Int)
	value.SetRandom()
	value.BigInt(bigval)
	for i := 0; i < 10*1024*1024; i++ {
		mt.Push(bigval.Bytes())
	}

	mt.Root()

	t.Fatal("cost: ", time.Since(nt))
}

func TestHashTwo(t *testing.T) {
	h := New()

	bigval := new(big.Int)
	bigval2 := new(big.Int)

	var val Fr
	val.SetRandom()
	val.BigInt(bigval)
	val.SetRandom()
	val.BigInt(bigval2)

	h.Write(bigval.Bytes())
	h.Write(bigval2.Bytes())
	h1 := h.Sum(nil)
	t.Log("h1: ", hex.EncodeToString(h1))

	h2, _ := HashHint([]*big.Int{new(big.Int).Set(bigval), new(big.Int).Set(bigval2)})
	t.Log("h2: ", hex.EncodeToString(h2.Bytes()))
	if !bytes.Equal(h1, h2.Bytes()) {
		t.Fatal("unequal")
	}
	t.Fatal()
}
