package contract

import "testing"

// A nonce reserved for a tx that was never broadcast must not leave a gap:
// every later tx would queue behind it until the stuck heuristic resyncs.
func TestReleaseNonce(t *testing.T) {
	// the most recent reservation is handed back and reused
	c := &ContractManage{nonceReady: true, localNonce: 10}
	c.releaseNonce(9)
	if c.localNonce != 9 || !c.nonceReady {
		t.Fatalf("release of the last nonce: localNonce=%d ready=%v, want 9 true", c.localNonce, c.nonceReady)
	}

	// a later nonce is already out: fall back to re-reading the chain
	c = &ContractManage{nonceReady: true, localNonce: 12}
	c.releaseNonce(9)
	if c.nonceReady || c.localNonce != 12 {
		t.Fatalf("release behind a later nonce: localNonce=%d ready=%v, want 12 false", c.localNonce, c.nonceReady)
	}

	// not tracking yet: nothing to release
	c = &ContractManage{nonceReady: false, localNonce: 0}
	c.releaseNonce(0)
	if c.nonceReady || c.localNonce != 0 {
		t.Fatalf("release while untracked changed state: localNonce=%d ready=%v", c.localNonce, c.nonceReady)
	}
}
