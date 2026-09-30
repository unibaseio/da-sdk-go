package contract

import (
	"errors"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

// Node.check reports an inactive node as (false, 0) instead of reverting, so
// CheckNode must turn the flag into an error callers can recognise; before
// this, an exited or slashed store looked healthy and never re-staked.
func TestNodeStatusErr(t *testing.T) {
	a := common.HexToAddress("0x2c8057f261d18b8e5ba62f0fdb0b63e0b9b59b33")

	if err := nodeStatusErr(true, a, 1); err != nil {
		t.Fatalf("active node: got %v, want nil", err)
	}

	err := nodeStatusErr(false, a, 1)
	if !errors.Is(err, ErrNodeInactive) {
		t.Fatalf("inactive node: %v does not wrap ErrNodeInactive", err)
	}

	// an RPC failure must stay distinguishable from "not active"
	if errors.Is(errors.New("dial tcp: i/o timeout"), ErrNodeInactive) {
		t.Fatal("unrelated error matched ErrNodeInactive")
	}
}
