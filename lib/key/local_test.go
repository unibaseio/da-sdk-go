package key_test

import (
	"os"
	"testing"

	"github.com/unibaseio/da-sdk-go/lib/key"
	"github.com/unibaseio/da-sdk-go/sdk"
)

// A throwaway signature still verifies, and nothing is written to disk.
func TestBuildAuthLocal(t *testing.T) {
	before, _ := os.ReadDir("/tmp/.dimo")
	au := key.BuildAuthLocal([]byte("download"))
	if err := sdk.VerifyAuth(au); err != nil {
		t.Fatalf("local auth does not verify: %v", err)
	}
	if au2 := key.BuildAuthLocal([]byte("download")); au2.Addr == au.Addr {
		t.Fatal("throwaway key reused")
	}
	after, _ := os.ReadDir("/tmp/.dimo")
	if len(after) != len(before) {
		t.Fatalf("keystore files written: %d -> %d", len(before), len(after))
	}
}
