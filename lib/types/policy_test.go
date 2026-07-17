package types

// Boundary guard for the RS (N,K) allow-list. SupportedPolicies is the single
// source of truth (CLAUDE.md data-model: only 6/4, 14/7, 32/16, 64/32). A silent
// change here loosens what the encoder/circuit assume about shard counts.

import "testing"

func TestPolicyCheckAccepts(t *testing.T) {
	want := []Policy{{N: 6, K: 4}, {N: 14, K: 7}, {N: 32, K: 16}, {N: 64, K: 32}}
	for _, p := range want {
		if err := p.Check(); err != nil {
			t.Errorf("policy %d/%d should be supported: %v", p.N, p.K, err)
		}
	}
	// Guard against the list silently growing/shrinking.
	if len(SupportedPolicies) != len(want) {
		t.Fatalf("SupportedPolicies size changed: got %d, expected %d — "+
			"confirm encoder + on-chain circuit agree before editing the allow-list",
			len(SupportedPolicies), len(want))
	}
}

func TestPolicyCheckRejects(t *testing.T) {
	bad := []Policy{
		{N: 0, K: 0},
		{N: 6, K: 3},   // right N, wrong K
		{N: 10, K: 5},  // plausible but unsupported ratio
		{N: 4, K: 6},   // K > N
		{N: 64, K: 16}, // mismatched pairing
	}
	for _, p := range bad {
		if err := p.Check(); err == nil {
			t.Errorf("policy %d/%d must be rejected", p.N, p.K)
		}
	}
}
