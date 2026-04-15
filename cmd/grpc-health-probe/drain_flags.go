package main

import (
	"time"

	"github.com/spf13/cobra"

	"github.com/example/grpc-health-probe-cli/internal/probe"
)

// addDrainFlags registers drain-related flags on cmd.
func addDrainFlags(cmd *cobra.Command) {
	if cmd == nil {
		return
	}
	cmd.Flags().Bool("drain", false, "wait for in-flight requests to complete before exiting")
	cmd.Flags().Duration("drain-timeout", 5*time.Second, "maximum time to wait for in-flight requests during drain")
}

// parseDrainConfig builds a DrainConfig from the flags registered by addDrainFlags.
func parseDrainConfig(cmd *cobra.Command) *probe.DrainConfig {
	if cmd == nil {
		return probe.DefaultDrainConfig()
	}

	cfg := probe.DefaultDrainConfig()

	if enabled, err := cmd.Flags().GetBool("drain"); err == nil {
		cfg.Enabled = enabled
	}

	if timeout, err := cmd.Flags().GetDuration("drain-timeout"); err == nil {
		cfg.DrainTimeout = timeout
	}

	return cfg
}
