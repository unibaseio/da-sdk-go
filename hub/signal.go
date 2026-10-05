package hub

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// signalShutdownTimeout bounds the graceful shutdown a signal triggers.
const signalShutdownTimeout = 10 * time.Second

// SetupSignalHandler shuts the server down gracefully on a termination signal
// and exits.
//
// It used to run a separate "emergency" path (close DB + repo, os.Exit(1)) in
// parallel with the daemon's own SIGINT/SIGTERM handler calling Shutdown: both
// fired on the same signal, so the repo could be closed under the LogFS
// instances Shutdown was still flushing, or the process exited mid-flush. Both
// now go through Shutdown, which runs exactly once (shutdownOnce) — a second
// caller waits for the first to finish — so the order (HTTP → stats → LogFS →
// DB → repo) always holds.
func (s *Server) SetupSignalHandler() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan,
		syscall.SIGINT,  // Ctrl+C
		syscall.SIGTERM, // Termination request
		syscall.SIGQUIT, // Quit from keyboard
		syscall.SIGHUP,  // Hangup detected
	)

	go func() {
		sig := <-sigChan
		logger.Warnf("received signal: %v, shutting down...", sig)

		ctx, cancel := context.WithTimeout(context.Background(), signalShutdownTimeout)
		defer cancel()
		if err := s.Shutdown(ctx); err != nil {
			logger.Errorf("shutdown: %v", err)
			os.Exit(1)
		}
		os.Exit(0)
	}()
}
