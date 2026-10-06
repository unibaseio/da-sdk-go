package hub

import (
	"encoding/binary"
	"path/filepath"
	"testing"

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
