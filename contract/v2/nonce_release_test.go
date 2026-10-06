package contract

import (
	"math/big"
	"testing"
)

// A nonce reserved for a tx that was never broadcast must not leave a gap:
// every later tx would queue behind it until the stuck heuristic resyncs.
func TestReleaseNonce(t *testing.T) {
	// the most recent reservation is handed back and reused
	c := &ContractManage{nonceReady: true, localNonce: 10}
	c.releaseNonce(9)
	if c.localNonce != 9 || !c.nonceReady {
		t.Fatalf("release of the last nonce: localNonce=%d ready=%v, want 9 true", c.localNonce, c.nonceReady)
	}

	// a later nonce is already out: the released one is kept for reuse and
	// tracking stays (no chain re-read that would count up through the later ones)
	c = &ContractManage{nonceReady: true, localNonce: 12}
	c.releaseNonce(9)
	if !c.nonceReady || c.localNonce != 12 || len(c.freeNonces) != 1 || c.freeNonces[0] != 9 {
		t.Fatalf("release behind a later nonce: localNonce=%d ready=%v free=%v, want 12 true [9]", c.localNonce, c.nonceReady, c.freeNonces)
	}
	c.releaseNonce(9) // twice: still one entry
	if len(c.freeNonces) != 1 {
		t.Fatalf("double release: free=%v", c.freeNonces)
	}

	// releasing the top nonce also folds free nonces just below it back
	c = &ContractManage{nonceReady: true, localNonce: 12, freeNonces: []uint64{10, 7}}
	c.releaseNonce(11)
	if c.localNonce != 10 || len(c.freeNonces) != 1 || c.freeNonces[0] != 7 {
		t.Fatalf("fold: localNonce=%d free=%v, want 10 [7]", c.localNonce, c.freeNonces)
	}

	// not tracking yet: nothing to release
	c = &ContractManage{nonceReady: false, localNonce: 0}
	c.releaseNonce(0)
	if c.nonceReady || c.localNonce != 0 {
		t.Fatalf("release while untracked changed state: localNonce=%d ready=%v", c.localNonce, c.nonceReady)
	}
}

// N5 audit: a nonce released while later reservations are in flight must be
// reused once, never re-derived from the chain. A reserves 0, B reserves 1 and
// is still building its tx, A fails before broadcast: the old release re-read
// the chain (pending nonce still 0) and handed out 0 and then 1 again — B's.
func TestReleasedNonceNotDuplicated(t *testing.T) {
	f := newFakeRPC(t)
	c := managerOn(t, f)
	a, err := c.nextNonce()
	if err != nil {
		t.Fatal(err)
	}
	b, _ := c.nextNonce()
	c.releaseNonce(a) // A never broadcast; B still in flight
	cn, _ := c.nextNonce()
	d, _ := c.nextNonce()
	if a != 0 || b != 1 || cn != 0 || d != 2 {
		t.Fatalf("got A=%d B=%d C=%d D=%d, want 0 1 0 2 (D must not reuse B's nonce)", a, b, cn, d)
	}
	// the chain moves past a free nonce (signed elsewhere): it is dropped
	c.releaseNonce(d) // top: localNonce back to 2
	c.releaseNonce(cn)
	f.mu.Lock()
	f.pendingNonce = 5
	f.mu.Unlock()
	if e, _ := c.nextNonce(); e != 5 {
		t.Fatalf("after the chain moved to 5: got %d, want 5", e)
	}
}

// Transfer and TransferToken share the nonce tracker: with a nonce already
// reserved by another goroutine, they must not reuse it (they used to read
// PendingNonceAt on their own, which still says 0).
func TestTransfersUseNonceTracker(t *testing.T) {
	f := newFakeRPC(t)
	c := managerOn(t, f)
	if n, err := c.nextNonce(); err != nil || n != 0 {
		t.Fatalf("reserve: %d %v", n, err)
	}
	if err := c.Transfer(someStore, big.NewInt(1)); err != nil {
		t.Fatalf("Transfer: %v", err)
	}
	if err := c.TransferToken(someStore, big.NewInt(1)); err != nil {
		t.Fatalf("TransferToken: %v", err)
	}
	if len(f.sent) != 2 || f.sent[0].Nonce() != 1 || f.sent[1].Nonce() != 2 {
		var got []uint64
		for _, tx := range f.sent {
			got = append(got, tx.Nonce())
		}
		t.Fatalf("sent nonces %v, want [1 2]", got)
	}
}
