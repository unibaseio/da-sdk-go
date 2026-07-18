package hub

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// TestDownloadSingleflightDedup verifies the server's dlSF group collapses
// concurrent fallbacks for the same key into a single execution (the guarantee
// download() relies on to avoid N expensive DA reconstructs for one hot object).
func TestDownloadSingleflightDedup(t *testing.T) {
	s := &Server{}
	var calls int32
	const n = 32
	start := make(chan struct{})
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			<-start
			_, _, _ = s.dlSF.Do("owner/name", func() (interface{}, error) {
				atomic.AddInt32(&calls, 1)
				time.Sleep(30 * time.Millisecond) // hold the flight so others join
				return []byte("data"), nil
			})
		}()
	}
	close(start) // release all at once
	wg.Wait()

	if got := atomic.LoadInt32(&calls); got != 1 {
		t.Fatalf("expected the fallback to run once for concurrent same-key requests, ran %d times", got)
	}

	// a later, non-overlapping call runs again (the flight already completed)
	_, _, _ = s.dlSF.Do("owner/name", func() (interface{}, error) {
		atomic.AddInt32(&calls, 1)
		return nil, nil
	})
	if got := atomic.LoadInt32(&calls); got != 2 {
		t.Fatalf("expected a fresh call after the first flight finished, got %d", got)
	}
}
