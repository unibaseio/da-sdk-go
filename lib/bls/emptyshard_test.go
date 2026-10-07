package bls

import (
	"bytes"
	"testing"
)

func TestNonEmptyDataShards(t *testing.T) {
	for _, c := range []struct{ elems, slen, k, want int }{
		{1, 1, 4, 1}, {2, 1, 4, 2}, {4, 1, 4, 4}, {5, 2, 4, 3}, {8, 2, 4, 4}, {9, 3, 4, 3}, {1, 33, 4, 1}, {100, 1, 4, 4},
	} {
		if got := NonEmptyDataShards(c.elems, c.slen, c.k); got != c.want {
			t.Errorf("NonEmptyDataShards(%d, %d, %d) = %d, want %d", c.elems, c.slen, c.k, got, c.want)
		}
	}
}

// The filler is one padded element, depends on the shard index and on the
// non-empty commitments, and is deterministic.
func TestEmptyShardElement(t *testing.T) {
	var a, b G1
	a.X.SetUint64(1)
	b.X.SetUint64(2)
	e2 := EmptyShardElement([]G1{a}, 2)
	if len(e2) != PadSize || e2[0] != PadByte {
		t.Fatalf("not one padded element: %x", e2)
	}
	if !bytes.Equal(e2, EmptyShardElement([]G1{a}, 2)) {
		t.Fatal("not deterministic")
	}
	if bytes.Equal(e2, EmptyShardElement([]G1{a}, 3)) {
		t.Fatal("same element for two shard indexes")
	}
	if bytes.Equal(e2, EmptyShardElement([]G1{b}, 2)) {
		t.Fatal("same element for different data")
	}
	if bytes.Equal(e2, LegacyEmptyShardElement()) {
		t.Fatal("equals the legacy constant")
	}
}
