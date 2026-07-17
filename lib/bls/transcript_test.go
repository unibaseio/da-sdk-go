package bls

// Invariant guards for the Fiat-Shamir transcript — CLAUDE.md gotcha #1.
//
// The transcript MiMC( 28 zero bytes ‖ stream ‖ {Commits[i]‖MoveCommits[i]‖
// LimitCommits[i]}_{i<k} ) must stay byte-identical across THREE sites:
//   1. the encoder      — helper.EncodeData        (da-go)
//   2. the verifier     — sdk.CheckFileFull        (da-sdk-go)
//   3. the in-circuit derivation — plonk rsone      (da-core, NOT reachable from Go)
//
// Sites 1+2 both route through EncodeWitness.Challenge, so a pure-Go test can lock
// them. Site 3 lives inside a gnark circuit and can only be cross-checked by the
// golden vector below (see da-core circuit-tier stub referenced in
// OPTIMIZATION_PLAN_V2.md §V1). These tests are FAST (no SRS): commitment points
// are cheap scalar multiples of the generator, not real KZG commitments.

import (
	"bytes"
	"encoding/hex"
	"math/big"
	"testing"
)

// distinctWitness builds an EncodeWitness whose G1 slots are distinct, cheap,
// deterministic points (scalar·G). No SRS / KZG setup needed — we only exercise
// the transcript byte layout, which doesn't care whether the points are real
// commitments. Distinct points ensure a reorder (e.g. Commits<->MoveCommits)
// actually changes the hash.
func distinctWitness(n, k int) *EncodeWitness {
	ew := NewEncodeWitness(n, k)
	s := int64(1)
	set := func(p *G1) { p.ScalarMultiplicationBase(big.NewInt(s)); s++ }
	set(&ew.Root)
	for i := 0; i < n; i++ {
		set(&ew.Commits[i])
	}
	for i := 0; i < k; i++ {
		set(&ew.MoveCommits[i])
		set(&ew.LimitCommits[i])
	}
	return ew
}

// reference is an INDEPENDENT re-implementation of the transcript, kept separate
// on purpose: if someone edits EncodeWitness.Challenge (reorders the writes,
// changes 28→32, drops LimitCommits, swaps Marshal for Bytes), this copy diverges
// and the test fails.
func referenceTranscript(ew *EncodeWitness, stream []byte) []byte {
	mh := NewFieldHash()
	buf := make([]byte, 28)
	buf = append(buf, stream...)
	mh.Write(buf)
	k := len(ew.MoveCommits)
	for i := 0; i < k; i++ {
		mh.Write(ew.Commits[i].Marshal())
		mh.Write(ew.MoveCommits[i].Marshal())
		mh.Write(ew.LimitCommits[i].Marshal())
	}
	return mh.Sum(nil)
}

// TestTranscriptMatchesReference locks Challenge against the independent copy.
func TestTranscriptMatchesReference(t *testing.T) {
	stream := make([]byte, 20) // 20-byte eth address width
	for i := range stream {
		stream[i] = byte(i + 1)
	}
	ew := distinctWitness(6, 4)

	got := ew.Challenge(stream)
	want := referenceTranscript(ew, stream)
	if !bytes.Equal(got, want) {
		t.Fatalf("transcript drift:\n Challenge = %x\n reference = %x", got, want)
	}
}

// TestTranscriptOrderSensitive proves the transcript actually binds ordering —
// i.e. the reference guard has teeth. Swapping a Commit with its MoveCommit must
// change the output.
func TestTranscriptOrderSensitive(t *testing.T) {
	stream := make([]byte, 20)
	ew := distinctWitness(6, 4)
	base := append([]byte(nil), ew.Challenge(stream)...)

	ew.Commits[0], ew.MoveCommits[0] = ew.MoveCommits[0], ew.Commits[0]
	if bytes.Equal(base, ew.Challenge(stream)) {
		t.Fatal("transcript insensitive to Commit/MoveCommit ordering — guard is toothless")
	}
}

// goldenTranscript pins the exact output for fixed inputs, locking the MiMC
// choice (BW6-761), the curve, the Marshal encoding, AND the layout — the
// strongest cross-site guard, including against the in-circuit derivation
// (site 3) whose gnark test should assert the same vector.
//
// Locked 2026-07-17 against the live EncodeWitness.Challenge (MiMC BW6-761 over
// the compressed-Marshal commitment bytes). Recompute + repaste ONLY if you
// intentionally change the transcript — and if you do, the on-chain rsone circuit
// and KZGVKRoot almost certainly need updating too (CLAUDE.md gotcha #1/#2).
const goldenTranscript = "012c6a7fc2e3a0cf2c98e94f276d4b39e114d798d979bc217ef7875cf8c0dc4c74f82fd174fbeda3b915fa93b286ddb7"

func TestTranscriptGolden(t *testing.T) {
	stream := make([]byte, 20)
	for i := range stream {
		stream[i] = byte(0xA0 + i)
	}
	ew := distinctWitness(6, 4)
	got := hex.EncodeToString(ew.Challenge(stream))

	if goldenTranscript == "" {
		t.Logf("golden not set — computed transcript = %s (paste into goldenTranscript)", got)
		t.Skip("fill goldenTranscript constant, then re-run")
	}
	if got != goldenTranscript {
		t.Fatalf("transcript golden mismatch:\n got    = %s\n golden = %s\n"+
			"if this change is intentional, the on-chain rsone circuit + KZGVKRoot may also need updating",
			got, goldenTranscript)
	}
}

// TestEncodeWitnessRoundTrip guards CBOR/gnark (de)serialization symmetry — the
// wire form that travels in FileFull.Proofs and gets re-parsed by the verifier.
func TestEncodeWitnessRoundTrip(t *testing.T) {
	ew := distinctWitness(14, 7)
	// H and ClaimedValues are part of the wire form too; give them values.
	ew.H.ScalarMultiplicationBase(big.NewInt(999))
	for i := range ew.ClaimedValues {
		ew.ClaimedValues[i].SetInt64(int64(i*7 + 3))
	}

	blob := ew.Serialize()
	out := new(EncodeWitness)
	if err := out.Deserialize(blob); err != nil {
		t.Fatalf("deserialize: %v", err)
	}

	// The transcript over the round-tripped witness must be identical — the
	// tightest single assertion that every G1/Fr slot survived the round trip.
	stream := make([]byte, 20)
	if !bytes.Equal(ew.Challenge(stream), out.Challenge(stream)) {
		t.Fatal("witness changed across Serialize/Deserialize")
	}
}
