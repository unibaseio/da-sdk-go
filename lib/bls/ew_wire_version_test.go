package bls

// V6 §B1 (EncodeWitness half) — versioned wire framing with legacy coexistence.
// distinctWitness is defined in transcript_test.go (same package).

import (
	"bytes"
	"math/big"
	"testing"
)

func TestEWFrameVersioned(t *testing.T) {
	b := distinctWitness(6, 4).Serialize()
	if len(b) < 3 || b[0] != ewMagic {
		t.Fatalf("serialized blob not framed: len=%d first=%#x", len(b), b[0])
	}
}

func TestEWVersionedRoundTrip(t *testing.T) {
	ew := distinctWitness(6, 4)
	ew.H.ScalarMultiplicationBase(big.NewInt(77))
	for i := range ew.ClaimedValues {
		ew.ClaimedValues[i].SetInt64(int64(i + 1))
	}
	out := new(EncodeWitness)
	if err := out.Deserialize(ew.Serialize()); err != nil {
		t.Fatal(err)
	}
	stream := make([]byte, 20)
	if !bytes.Equal(ew.Challenge(stream), out.Challenge(stream)) {
		t.Fatal("round-trip changed the witness")
	}
}

// A pre-versioning blob is exactly the framed payload minus the 3-byte header;
// its first byte has bit7==0, so Deserialize must take the legacy path.
func TestEWLegacyBlobDecodes(t *testing.T) {
	ew := distinctWitness(6, 4)
	legacy := ew.Serialize()[3:]
	if legacy[0]&0x80 != 0 {
		t.Fatalf("legacy first byte must have bit7==0, got %#x", legacy[0])
	}
	out := new(EncodeWitness)
	if err := out.Deserialize(legacy); err != nil {
		t.Fatalf("legacy blob rejected: %v", err)
	}
	stream := make([]byte, 20)
	if !bytes.Equal(ew.Challenge(stream), out.Challenge(stream)) {
		t.Fatal("legacy decode mismatch")
	}
}

func TestEWFutureVersionRejected(t *testing.T) {
	b := distinctWitness(6, 4).Serialize()
	b[2] = 99 // bump the low version byte in the frame
	if err := (new(EncodeWitness)).Deserialize(b); err == nil {
		t.Fatal("expected rejection of a future EncodeWitness wire version")
	}
}
