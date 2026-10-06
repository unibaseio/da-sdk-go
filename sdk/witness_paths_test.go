package sdk

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/unibaseio/da-sdk-go/lib/types"
)

// S3: the hub's seal and drain hand a stream's upload answer to CheckFileFull,
// the SDK upload to CheckFileFullShape. A witness declaring 2^32-1 commitments
// must be refused by the shape walk, before the gnark decoder would allocate
// the declared length.
func TestStreamWitnessShapeCheckedFirst(t *testing.T) {
	data := []byte(strings.Repeat("x", 1000))
	fp := filepath.Join(t.TempDir(), "f")
	if err := os.WriteFile(fp, data, 0o600); err != nil {
		t.Fatal(err)
	}
	sum := sha256.Sum256(data)
	evil := []byte{0xDA, 0, 1}               // versioned frame
	evil = append(evil, make([]byte, 96)...) // root (raw G1)
	evil = binary.BigEndian.AppendUint32(evil, 0xFFFFFFFF)
	pol := types.Policy{N: 6, K: 4}
	ff := types.FileFull{
		FileReceipt: types.FileReceipt{
			FileCore: types.FileCore{Policy: pol, Hash: hex.EncodeToString(sum[:]), Size: int64(len(data))},
			Pieces:   []string{strings.Repeat("ab", 48)},
		},
		Proofs:     [][]byte{evil},
		PieceSizes: []int64{int64(len(data))},
	}

	var before, after runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&before)
	_, err := CheckFileFull(ff, common.HexToAddress("0x1"), fp)
	runtime.ReadMemStats(&after)
	if err == nil || !strings.Contains(err.Error(), "witness") {
		t.Fatalf("CheckFileFull: %v", err)
	}
	if grew := after.TotalAlloc - before.TotalAlloc; grew > 16<<20 {
		t.Fatalf("rejecting the witness allocated %d bytes", grew)
	}
	if err := CheckFileFullShape(ff, pol); err == nil {
		t.Fatal("CheckFileFullShape accepted the witness")
	}
}
