package contract

import (
	"context"
	"math/big"
	"reflect"
	"testing"
	"time"
)

func TestSplitEndpoints(t *testing.T) {
	cases := []struct {
		in   string
		want []string
	}{
		{"http://a", []string{"http://a"}},
		{"http://a,http://b", []string{"http://a", "http://b"}},
		{" http://a , http://b ,", []string{"http://a", "http://b"}},
	}
	for _, c := range cases {
		if got := splitEndpoints(c.in); !reflect.DeepEqual(got, c.want) {
			t.Errorf("splitEndpoints(%q) = %v, want %v", c.in, got, c.want)
		}
	}
}

// TestFilterRPCFailover proves the same rotation for the SEPARATE filter-log
// path: FilterBlockNumber on a dead endpoint fails and rotates, then succeeds
// on the live one. Event sync depends on this — a store that can't read its
// challenge would be wrongly slashed.
func TestFilterRPCFailover(t *testing.T) {
	reachableAnvil(t)

	const dead = "http://127.0.0.1:1"
	cm := &ContractManage{
		filterRPCs:      []string{dead, anvilRPC},
		filterIdx:       0,
		RPCForFilterLog: dead,
		ChainID:         big.NewInt(31337),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	_, err := cm.FilterBlockNumber(ctx) // dead endpoint -> error + internal rotate
	cancel()
	if err == nil {
		t.Fatal("expected dead filter endpoint to fail")
	}
	if cm.activeFilterRPC() != anvilRPC {
		t.Fatalf("after failover active filter=%s, want %s", cm.activeFilterRPC(), anvilRPC)
	}

	ctx, cancel = context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if _, err := cm.FilterBlockNumber(ctx); err != nil {
		t.Fatalf("filter call after failover should succeed: %v", err)
	}
}

// activeFilterRPC test accessor mirrors activeRPC for assertions.
func (c *ContractManage) activeFilterRPC() string {
	c.filterMu.Lock()
	defer c.filterMu.Unlock()
	return c.RPCForFilterLog
}

// TestRotateRPCFailover proves the rotation mechanism: a dead active endpoint
// fails the call, rotateRPC advances to the next endpoint and drops the cached
// client, and subsequent calls succeed on the live endpoint.
func TestRotateRPCFailover(t *testing.T) {
	reachableAnvil(t)

	const dead = "http://127.0.0.1:1" // nothing listening → connection refused
	cm := &ContractManage{
		rpcs:    []string{dead, anvilRPC},
		rpcIdx:  0,
		RPC:     dead,
		ChainID: big.NewInt(31337),
	}

	// active endpoint is dead: dial is lazy (no error) but a real call fails.
	cl, err := cm.Client(context.Background())
	if err != nil {
		t.Fatalf("dial dead endpoint: %v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	_, err = cl.ChainID(ctx)
	cancel()
	if err == nil {
		t.Fatal("expected the dead endpoint to fail a call")
	}

	// failover to the next endpoint.
	cm.rotateRPC()
	if cm.activeRPC() != anvilRPC {
		t.Fatalf("after rotate active=%s, want %s", cm.activeRPC(), anvilRPC)
	}

	// the shared client was dropped; Client re-dials the live endpoint.
	cl, err = cm.Client(context.Background())
	if err != nil {
		t.Fatalf("dial after failover: %v", err)
	}
	ctx, cancel = context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if _, err := cl.ChainID(ctx); err != nil {
		t.Fatalf("call after failover should succeed: %v", err)
	}
}
