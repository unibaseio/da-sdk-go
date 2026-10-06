package bls

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"sync"

	bls12377 "github.com/consensys/gnark-crypto/ecc/bls12-377"
	"github.com/consensys/gnark-crypto/ecc/bls12-377/kzg"
)

// The KZG verifying key of the shared ceremony SRS (BLS12-377): [1]G1, [1]G2
// and [τ]G2, compressed. A client checks an encoding's opening with these
// three points instead of the 3 GB SRS. da-go's TestEmbeddedKZGVK checks them
// against the chain's KZGVKRoot and the cached SRS.
const (
	srsVKG1  = "a08848defe740a67c8fc6225bf87ff5485951e2caa9d41bb188282c8bd37cb5cd5481512ffcd394eeab9b16eb21be9ef"
	srsVKG20 = "a0ea6040e700403170dc5a51b1b140d5532777ee6651cecbe7223ece0799c9de5cf89984bff76fe6b26bfefa6ea16afe018480be71c785fec89630a2a3841d01c565f071203e50317ea501f557db6b9b71889f52bb53540274e3e48f7c005196"
	srsVKG21 = "80552e9e3d32dbfa51619c66799dd184d1029f0ebbcabb1c2679e499749b7eb254948157606554ed03c1a6c31a30d53f00fa7c3f9a88fd62d0a23c96f8158d1e2995fb13fc0a93f96ee7e7b83bbeb136276581f9f64926891656e1a358d6d1bf"
)

var (
	srsVKOnce sync.Once
	srsVK     *VerifyKey
	srsVKErr  error
)

// SRSVerifyKey returns the embedded verifying key of the shared SRS.
func SRSVerifyKey() (*VerifyKey, error) {
	srsVKOnce.Do(func() {
		vk := &VerifyKey{VerifyingKey: new(kzg.VerifyingKey)}
		for _, p := range []struct {
			hex string
			set func([]byte) error
		}{
			{srsVKG1, func(b []byte) error { _, err := vk.G1.SetBytes(b); return err }},
			{srsVKG20, func(b []byte) error { _, err := vk.G2[0].SetBytes(b); return err }},
			{srsVKG21, func(b []byte) error { _, err := vk.G2[1].SetBytes(b); return err }},
		} {
			b, err := hex.DecodeString(p.hex)
			if err != nil {
				srsVKErr = err
				return
			}
			if err := p.set(b); err != nil {
				srsVKErr = err
				return
			}
		}
		PrecomputeLines(vk.VerifyingKey)
		srsVK = vk
	})
	return srsVK, srsVKErr
}

// Accumulation formats of an encoding's opening H (EncodeWitness.Format).
const (
	// FormatV1: shard i's C_i, M_i, L_i are weighted 1, v_i, v_i² with
	// v_i = MiMC(y_i). The data commitments share weight 1, so only their
	// sum is bound (2026-10-05 audit, N1). Kept to verify existing encodings.
	FormatV1 uint8 = 1
	// FormatV2: one coefficient γ derived after every commitment and claimed
	// value is fixed; shard i's C_i, M_i, L_i are weighted γ^(3i), γ^(3i+1),
	// γ^(3i+2), so every commitment is bound on its own.
	FormatV2 uint8 = 2
)

// BatchCoefficient is FormatV2's γ = MiMC(point ‖ y_0 ‖ … ‖ y_{k-1}): point
// is the raw Fiat-Shamir output (ew.Challenge, which already binds every
// commitment) and each y_i the shard's claimed value, every input one
// 48-byte field element. The RSOne v2 circuit derives the same value.
func (ew *EncodeWitness) BatchCoefficient(point []byte) Fr {
	h := NewFieldHash()
	h.Write(pad48(point))
	for i := range ew.ClaimedValues {
		b := ew.ClaimedValues[i].Marshal()
		h.Write(pad48(b))
	}
	var g Fr
	g.SetBytes(h.Sum(nil))
	return g
}

func pad48(b []byte) []byte {
	if len(b) >= 48 {
		return b
	}
	out := make([]byte, 48-len(b), 48)
	return append(out, b...)
}

// AccCoefficient is v = MiMC(y), the per-shard coefficient of the encoding's
// accumulated opening (format 1): the shard's move and limit commitments are
// weighted v and v². The encoder (da-core acc) and the RSOne circuits use
// the same value.
func AccCoefficient(y Fr) Fr {
	h := NewFieldHash()
	b := y.Marshal()
	buf := make([]byte, 48-len(b))
	buf = append(buf, b...)
	h.Write(buf)
	var v Fr
	v.SetBytes(h.Sum(nil))
	return v
}

// VerifyOpening checks the encoder's accumulated KZG opening H under the
// witness's format. slen is the shard length in field elements. Every point is
// checked to lie in the prime-order subgroup first: the witness comes from the
// stream and is decoded without subgroup checks.
func (ew *EncodeWitness) VerifyOpening(vk *VerifyKey, stream []byte, slen int) error {
	k := len(ew.MoveCommits)
	if k == 0 || len(ew.LimitCommits) != k || len(ew.ClaimedValues) != k || len(ew.Commits) < k {
		return fmt.Errorf("witness shape does not match its K=%d", k)
	}
	if slen <= 0 || slen*k > MaxShard {
		return fmt.Errorf("shard length %d out of range", slen)
	}
	pts := append([]G1{ew.Root, ew.H}, ew.Commits...)
	pts = append(pts, ew.MoveCommits...)
	pts = append(pts, ew.LimitCommits...)
	for i := range pts {
		if !pts[i].IsInSubGroup() {
			return fmt.Errorf("witness point %d is not in the G1 subgroup", i)
		}
	}
	switch ew.format() {
	case FormatV1:
		return ew.verifyOpeningV1(vk, stream, slen)
	case FormatV2:
		return ew.verifyOpeningV2(vk, stream, slen)
	}
	return fmt.Errorf("unknown witness format %d", ew.Format)
}

// verifyOpeningV1 checks FormatV1: at the Fiat-Shamir point r,
//
//	Σ_i (C_i + v_i·M_i + v_i²·L_i)  opens to  Σ_i y_i·(1 + v_i·r^(i·slen) + v_i²·r^(Max-slen))
//
// with v_i = AccCoefficient(y_i), slen the shard length in field elements.
// Together with the claimed values y_i (which the caller checks against the
// real data) this binds every move and limit commitment, hence the piece
// root, to the data. In this format the data commitments C_i share the
// coefficient 1, so only their sum is bound (2026-10-05 audit, N1).
func (ew *EncodeWitness) verifyOpeningV1(vk *VerifyKey, stream []byte, slen int) error {
	k := len(ew.MoveCommits)
	var r Fr
	r.SetBytes(ew.Challenge(stream))
	rLen := new(big.Int)
	var rk1, rStep, rk2 Fr
	rStep.Exp(r, rLen.SetInt64(int64(slen)))
	rk2.Exp(r, rLen.SetInt64(int64(MaxShard-slen)))
	rk1.SetOne()

	var sum G1
	var value Fr
	for i := 0; i < k; i++ {
		y := ew.ClaimedValues[i]
		v := AccCoefficient(y)
		vb := new(big.Int)
		v.BigInt(vb)

		// C_i + v·(M_i + v·L_i)
		var t G1
		t.ScalarMultiplication(&ew.LimitCommits[i], vb)
		t.Add(&t, &ew.MoveCommits[i])
		t.ScalarMultiplication(&t, vb)
		t.Add(&t, &ew.Commits[i])
		sum.Add(&sum, &t)

		// y_i·(1 + v·r^k1 + v²·r^k2)
		var f, tmp Fr
		f.Mul(&v, &v).Mul(&f, &rk2)
		tmp.Mul(&v, &rk1)
		f.Add(&f, &tmp)
		tmp.SetOne()
		f.Add(&f, &tmp).Mul(&f, &y)
		value.Add(&value, &f)

		rk1.Mul(&rk1, &rStep)
	}

	return checkOpening(vk, &sum, &value, &r, &ew.H)
}

// verifyOpeningV2 checks FormatV2: with γ = BatchCoefficient(r) and
// a_i = γ^(3i), at r
//
//	Σ_i a_i·(C_i + γ·M_i + γ²·L_i)  opens to  Σ_i a_i·y_i·(1 + γ·r^(i·slen) + γ²·r^(Max-slen))
//
// Each commitment has its own weight, so each is bound to its claimed value.
func (ew *EncodeWitness) verifyOpeningV2(vk *VerifyKey, stream []byte, slen int) error {
	k := len(ew.MoveCommits)
	point := ew.Challenge(stream)
	var r Fr
	r.SetBytes(point)
	g := ew.BatchCoefficient(point)
	var g2, g3 Fr
	g2.Mul(&g, &g)
	g3.Mul(&g2, &g)

	e := new(big.Int)
	var rStep, rk1, rk2 Fr
	rStep.Exp(r, e.SetInt64(int64(slen)))
	rk2.Exp(r, e.SetInt64(int64(MaxShard-slen)))
	rk1.SetOne()

	var sum G1
	var value Fr
	var a Fr
	a.SetOne()
	gb, g2b := new(big.Int), new(big.Int)
	g.BigInt(gb)
	g2.BigInt(g2b)
	for i := 0; i < k; i++ {
		// a·(C + γ·M + γ²·L)
		var t, m, l G1
		m.ScalarMultiplication(&ew.MoveCommits[i], gb)
		l.ScalarMultiplication(&ew.LimitCommits[i], g2b)
		t.Add(&ew.Commits[i], &m)
		t.Add(&t, &l)
		ab := new(big.Int)
		a.BigInt(ab)
		t.ScalarMultiplication(&t, ab)
		sum.Add(&sum, &t)

		// a·y·(1 + γ·r^k1 + γ²·r^k2)
		var f, tmp Fr
		f.Mul(&g2, &rk2)
		tmp.Mul(&g, &rk1)
		f.Add(&f, &tmp)
		tmp.SetOne()
		f.Add(&f, &tmp).Mul(&f, &ew.ClaimedValues[i]).Mul(&f, &a)
		value.Add(&value, &f)

		rk1.Mul(&rk1, &rStep)
		a.Mul(&a, &g3)
	}
	return checkOpening(vk, &sum, &value, &r, &ew.H)
}

// checkOpening is the KZG check e([value]G₁ - com - r·H, G₂)·e(H, [τ]G₂) == 1.
func checkOpening(vk *VerifyKey, com *G1, value, r *Fr, h *G1) error {
	var left, rh G1
	b := new(big.Int)
	value.BigInt(b)
	left.ScalarMultiplication(&vk.G1, b)
	left.Sub(&left, com)
	r.BigInt(b)
	rh.ScalarMultiplication(h, b)
	left.Sub(&left, &rh)
	ok, err := bls12377.PairingCheck([]G1{left, *h}, []G2{vk.G2[0], vk.G2[1]})
	if err != nil {
		return err
	}
	if !ok {
		return fmt.Errorf("encoding opening does not verify")
	}
	return nil
}
