package types

// V6 (FORMAT_VERSIONING_DESIGN.md) — guards the CBOR wire-version behavior:
// new writes are stamped, round-trips preserve fields, legacy (no Version) blobs
// decode as v1, and a future version fails closed instead of being misparsed.

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/fxamacker/cbor/v2"
)

func TestWireVersionStamped(t *testing.T) {
	rc := &ReplicaCore{Name: "abc", Index: 3}
	b, err := rc.Serialize()
	if err != nil {
		t.Fatal(err)
	}
	if rc.Version != CurrentWireVersion {
		t.Fatalf("Serialize did not stamp version: got %d want %d", rc.Version, CurrentWireVersion)
	}
	out := &ReplicaCore{}
	if err := out.Deserialize(b); err != nil {
		t.Fatal(err)
	}
	if out.Version != CurrentWireVersion || out.Name != "abc" || out.Index != 3 {
		t.Fatalf("round-trip mismatch: %+v", out)
	}
}

// Embedding: FileFull → FileReceipt → FileCore.Version must resolve and stamp.
func TestWireVersionEmbeddedStamped(t *testing.T) {
	ff := &FileFull{}
	ff.Name = "file1"
	ff.Owner = common.HexToAddress("0x01")
	if _, err := ff.Serialize(); err != nil {
		t.Fatal(err)
	}
	if ff.Version != CurrentWireVersion {
		t.Fatalf("embedded Version not stamped: %d", ff.Version)
	}
}

// A record written before versioning existed has no Version key → decodes to 0,
// which we treat as v1 (no error).
func TestWireVersionLegacyDecodesAsV1(t *testing.T) {
	legacy, err := cbor.Marshal(map[string]interface{}{"Name": "legacy", "Index": uint64(9)})
	if err != nil {
		t.Fatal(err)
	}
	rc := &ReplicaCore{}
	if err := rc.Deserialize(legacy); err != nil {
		t.Fatalf("legacy blob rejected: %v", err)
	}
	if rc.Version != 0 || rc.Name != "legacy" || rc.Index != 9 {
		t.Fatalf("legacy decode wrong: %+v", rc)
	}
}

// A record written by a newer node must fail closed, not be misparsed.
func TestWireVersionFutureRejected(t *testing.T) {
	rc := &ReplicaCore{Name: "future", Version: CurrentWireVersion + 5}
	b, err := rc.Serialize() // stampVersion leaves an already-set version alone
	if err != nil {
		t.Fatal(err)
	}
	if err := (&ReplicaCore{}).Deserialize(b); err == nil {
		t.Fatal("expected rejection of a future wire version")
	}
}
