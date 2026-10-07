package bls

import (
	"crypto/sha256"
	"encoding/binary"
)

// Empty data shards.
//
// A piece's data is laid out over its K data shards, slen padded elements
// each; a small piece leaves the last shards wholly empty. Such a shard must
// not be the zero polynomial (an infinity commitment no RSOne circuit can
// prove), and its commitment — which is its replica name — must differ from
// every other piece's and from the piece's other empty shards: the contract
// keeps replica names unique, so equal names can be registered only once,
// network-wide (2026-10-07: with one constant filler, a 50-byte piece got
// 5/6 replicas and every later small piece collided with it).
//
// An empty shard therefore holds one element derived from the commitments of
// the piece's non-empty data shards (bound to the data by the encoder's
// opening, see EncodeWitness.VerifyOpening) and its own index; the rest of
// the shard stays zero. Encoder (da-go EncodeData) and client
// (sdk.CheckFileFull) both call EmptyShardElement. Unpad drops the element's
// prefix byte and downloads truncate to the piece size, so the filler never
// reaches the data.

const emptyShardDomain = "unibase-da/empty-shard/v1"

// NonEmptyDataShards is how many of a piece's k data shards hold data, for a
// piece of elems padded elements laid out slen elements per shard. Shard i is
// empty iff i >= NonEmptyDataShards; shard 0 always holds data.
func NonEmptyDataShards(elems, slen, k int) int {
	if slen <= 0 || elems <= 0 {
		return 1
	}
	n := (elems + slen - 1) / slen
	if n < 1 {
		n = 1
	}
	if n > k {
		n = k
	}
	return n
}

// EmptyShardElement is the padded element (PadSize bytes) at the start of
// empty data shard i: Pad(sha256(domain ‖ Marshal(C_0..C_{e-1}) ‖ be32(i))[:31]),
// where C_0..C_{e-1} are the commitments of the piece's non-empty data shards.
func EmptyShardElement(nonEmpty []G1, i int) []byte {
	h := sha256.New()
	h.Write([]byte(emptyShardDomain))
	for j := range nonEmpty {
		h.Write(nonEmpty[j].Marshal())
	}
	var b [4]byte
	binary.BigEndian.PutUint32(b[:], uint32(i))
	h.Write(b[:])
	sum := h.Sum(nil)
	return Pad(sum[:UnPadSize])
}

// LegacyEmptyShardElement is what the 0.1.34 encoder put in every empty data
// shard (Pad of nothing). Clients still accept it so receipts from a stream
// not yet upgraded verify during a rollout; it concerns filler only, never
// the piece's data.
func LegacyEmptyShardElement() []byte {
	return Pad(nil)
}
