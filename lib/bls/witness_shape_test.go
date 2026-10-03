package bls

import (
	"bytes"
	"encoding/binary"
	"math/big"
	"testing"

	gbls "github.com/consensys/gnark-crypto/ecc/bls12-377"
)

func sampleWitness(n, k int) *EncodeWitness {
	ew := NewEncodeWitness(n, k)
	_, _, g1, _ := gbls.Generators()
	for i := range ew.Commits {
		ew.Commits[i].ScalarMultiplication(&g1, big.NewInt(int64(i+2)))
	}
	ew.Root = ew.Commits[0]
	return ew
}

func TestCheckEncodeWitnessShape(t *testing.T) {
	ew := sampleWitness(6, 4)

	raw := ew.Serialize()
	if err := CheckEncodeWitnessShape(raw, 6, 4); err != nil {
		t.Fatalf("raw (current) encoding: %v", err)
	}
	// compressed points inside a frame (the decoder accepts either size)
	var legacy bytes.Buffer
	legacy.Write(raw[:3])
	enc := gbls.NewEncoder(&legacy)
	for _, v := range []interface{}{&ew.Root, ew.Commits, ew.MoveCommits, ew.LimitCommits, &ew.H, ew.ClaimedValues} {
		if err := enc.Encode(v); err != nil {
			t.Fatal(err)
		}
	}
	if err := CheckEncodeWitnessShape(legacy.Bytes(), 6, 4); err != nil {
		t.Fatalf("framed compressed encoding: %v", err)
	}
	var back EncodeWitness
	if err := back.Deserialize(legacy.Bytes()); err != nil {
		t.Fatal(err)
	}

	if err := CheckEncodeWitnessShape(raw, 14, 7); err == nil {
		t.Fatal("accepted a 6/4 witness as 14/7")
	}
	if err := CheckEncodeWitnessShape(append(append([]byte{}, raw...), 0), 6, 4); err == nil {
		t.Fatal("accepted trailing bytes")
	}
	if err := CheckEncodeWitnessShape(raw[:len(raw)-1], 6, 4); err == nil {
		t.Fatal("accepted a truncated witness")
	}

	// the attack: a frame + root, then a slice declaring 2^32-1 points
	evil := append([]byte{}, raw[:3+96]...)
	evil = binary.BigEndian.AppendUint32(evil, 0xFFFFFFFF)
	if err := CheckEncodeWitnessShape(evil, 6, 4); err == nil {
		t.Fatal("accepted a 2^32-point slice")
	}
}
