package main

import (
	"time"

	"github.com/spf13/cobra"

	"github.com/example/grpc-health-probe/internal/probe"
)

// addSemaphoreFlags registers semaphore-related flags on cmd.
func addSemaphoreFlags(cmd *cobra.Command) {
	if cmd == nil {
		return
	}
	cmd.Flags().Bool("semaphore", false, "Enable semaphore-based concurrency limiting")
	cmd.Flags().Int("semaphore-max-tickets", 10, "Maximum concurrent in-flight health checks")
	cmd.Flags().Duration("semaphore-acquire-timeout", 5*time.Second, "Maximum time to wait for a semaphore ticket")
}

// parseSemaphoreConfig builds a SemaphoreConfig from the flags registered by
// addSemaphoreFlags. It falls back to DefaultSemaphoreConfig values when the
// command is nil.
func parseSemaphoreConfig(cmd *cobra.Command) *probe.SemaphoreConfig {
	cfg := probe.DefaultSemaphoreConfig()
	if cmd == nil {
		return cfg
	}

	if v, err := cmd.Flags().GetBool("semaphore"); err == nil {
		cfg.Enabled = v
	}
	if v, err := cmd.Flags().GetInt("semaphore-max-tickets"); err == nil {
		cfg.MaxTickets = v
	}
	if v, err := cmd.Flags().GetDuration("semaphore-acquire-timeout"); err == nil {
		cfg.AcquireTimeout = v
	}

	return cfg
}
