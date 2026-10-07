package hub

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/unibaseio/da-sdk-go/lib/types"

	"github.com/unibaseio/da-sdk-go/lib/kv"
	"github.com/unibaseio/da-sdk-go/lib/logfs"
)

// S9: the reclaim gate reads the same drain offset drainInstance advances.
func TestDrainNextGate(t *testing.T) {
	ds, err := kv.NewBadgerStore(filepath.Join(t.TempDir(), "kv"), &kv.DefaultOptions)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { ds.Close() })
	first := logfs.GetIndex("0xhub", "0xowner")
	if got := drainNext(ds, "0xhub", "0xowner"); got != first {
		t.Fatalf("nothing drained: next %d, want the first volume %d", got, first)
	}
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, first+3) // as drainInstance records a commit
	if err := ds.Put(drainOffsetKey("0xowner"), buf); err != nil {
		t.Fatal(err)
	}
	if got := drainNext(ds, "0xhub", "0xowner"); got != first+3 {
		t.Fatalf("next %d, want %d", got, first+3)
	}
	if drainNext(ds, "0xhub", "0xother") != logfs.GetIndex("0xhub", "0xother") {
		t.Fatal("offsets are per owner")
	}
}

// The drain offset never passes an uncommitted volume: when volume i fails
// and i+1 would succeed, the offset stays at i and i+1 is not tried this
// pass. (The old loop moved on after a volume whose pieces were not all
// registered, and i+1's commit wrote i+2, so the gate freed i's file.)
func TestDrainVolumesStopsAtUncommitted(t *testing.T) {
	ds, err := kv.NewBadgerStore(filepath.Join(t.TempDir(), "kv"), &kv.DefaultOptions)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { ds.Close() })
	first := logfs.GetIndex("0xhub", "0xowner")
	key := drainOffsetKey("0xowner")
	ok := map[uint64]bool{first: true, first + 1: false, first + 2: true}
	var tried []uint64
	drainVolumes(ds, key, first, first+3, func(i uint64) bool {
		tried = append(tried, i)
		return ok[i]
	})
	if got := drainNext(ds, "0xhub", "0xowner"); got != first+1 {
		t.Fatalf("offset %d after volume %d failed, want %d", got, first+1, first+1)
	}
	if len(tried) != 2 {
		t.Fatalf("tried %v: a volume after the uncommitted one was attempted", tried)
	}
	// next pass: the failed volume commits now, the rest follow
	ok[first+1] = true
	drainVolumes(ds, key, drainNext(ds, "0xhub", "0xowner"), first+3, func(i uint64) bool { return ok[i] })
	if got := drainNext(ds, "0xhub", "0xowner"); got != first+3 {
		t.Fatalf("offset %d after all committed, want %d", got, first+3)
	}
}

// A volume whose bytes this hub already recorded (gateway: "already has file
// with hash H") is committed rather than retried forever — but only when H is
// both the stream receipt's hash and the local volume's.
func TestSameContentRecorded(t *testing.T) {
	fp := filepath.Join(t.TempDir(), "1.vol")
	body := []byte("identical volume bytes")
	if err := os.WriteFile(fp, body, 0o600); err != nil {
		t.Fatal(err)
	}
	sum := sha256.Sum256(body)
	h := hex.EncodeToString(sum[:])
	other := strings.Repeat("ab", 32)
	res := types.FileFull{FileReceipt: types.FileReceipt{FileCore: types.FileCore{Hash: h}, Pieces: []string{"p"}}}
	msg := func(hash string) string {
		return `response: 500 Internal Server Error, msg: {"Type":"file","Message":"already has file with hash ` + hash + `"}`
	}
	if !sameContentRecorded(msg(h), res, fp) {
		t.Fatal("same bytes recorded: want commit")
	}
	if sameContentRecorded(msg(other), res, fp) {
		t.Fatal("the gateway names another hash: must not commit")
	}
	bad := res
	bad.Hash = other
	if sameContentRecorded(msg(other), bad, fp) {
		t.Fatal("receipt hash is not the local volume's: must not commit")
	}
	if sameContentRecorded("response: 500 Internal Server Error, msg: db down", res, fp) {
		t.Fatal("another error: must not commit")
	}
	none := res
	none.Pieces = nil
	if sameContentRecorded(msg(h), none, fp) {
		t.Fatal("receipt without pieces: must not commit")
	}
}
