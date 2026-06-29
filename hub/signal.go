package hub

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// SetupSignalHandler sets up signal handling for graceful shutdown
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
		logger.Warnf("received signal: %v, initiating emergency shutdown...", sig)

		// Create a context with timeout for emergency shutdown
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// Perform emergency data persistence
		s.emergencyShutdown(ctx)

		// Exit the process
		os.Exit(1)
	}()
}

// emergencyShutdown performs critical data persistence operations quickly
func (s *Server) emergencyShutdown(ctx context.Context) {
	logger.Warn("performing emergency data persistence...")

	// Force database checkpoint
	if s.gdb != nil {
		logger.Info("emergency: forcing database checkpoint...")
		if err := s.gdb.Exec("PRAGMA wal_checkpoint(FULL);").Error; err != nil {
			logger.Errorf("emergency checkpoint failed: %v", err)
		}

		// Get underlying SQL database and close it
		if sqlDB, err := s.gdb.DB(); err == nil {
			sqlDB.Close()
		}
	}

	// Close repository
	if s.rp != nil {
		logger.Info("emergency: closing repository...")
		s.rp.Close()
	}

	logger.Warn("emergency shutdown completed")
}
