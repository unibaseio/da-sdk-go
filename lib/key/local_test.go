package key_test

import (
	"testing"

	"github.com/unibaseio/da-sdk-go/lib/key"
	"github.com/unibaseio/da-sdk-go/sdk"
)

// A throwaway signature still verifies, and every call uses a fresh key.
func TestBuildAuthLocal(t *testing.T) {
	au := key.BuildAuthLocal([]byte("download"))
	if err := sdk.VerifyAuth(au); err != nil {
		t.Fatalf("local auth does not verify: %v", err)
	}
	if au2 := key.BuildAuthLocal([]byte("download")); au2.Addr == au.Addr {
		t.Fatal("throwaway key reused")
	}
}
