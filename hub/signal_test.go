package hub

import (
	"context"
	"sync"
	"testing"
)

// The signal handler and the daemon both call Shutdown on SIGINT/SIGTERM. It
// must run once (a second close of checkpointStop used to panic) and stop the
// drain loop.
func TestShutdownRunsOnce(t *testing.T) {
	s := &Server{
		shutdownChan:   make(chan struct{}),
		checkpointStop: make(chan struct{}),
	}
	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := s.Shutdown(context.Background()); err != nil {
				t.Error(err)
			}
		}()
	}
	wg.Wait()
	select {
	case <-s.shutdownChan:
	default:
		t.Fatal("drain loop not signalled to stop")
	}
}
