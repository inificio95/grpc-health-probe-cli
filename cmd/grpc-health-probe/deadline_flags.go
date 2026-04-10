package main

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/user/grpc-health-probe-cli/internal/probe"
)

// addDeadlineFlags registers deadline-related CLI flags onto cmd.
func addDeadlineFlags(cmd *cobra.Command) {
	if cmd == nil {
		return
	}
	cmd.Flags().Bool("deadline", false, "Enable a global deadline for the entire probe")
	cmd.Flags().Duration("deadline-duration", 30*time.Second, "Maximum total time allowed for the probe (e.g. 30s, 1m)")
}

// parseDeadlineConfig builds a DeadlineConfig from the flags registered on cmd.
func parseDeadlineConfig(cmd *cobra.Command) (*probe.DeadlineConfig, error) {
	if cmd == nil {
		return probe.DefaultDeadlineConfig(), nil
	}

	enabled, err := cmd.Flags().GetBool("deadline")
	if err != nil {
		return nil, fmt.Errorf("reading --deadline flag: %w", err)
	}

	duration, err := cmd.Flags().GetDuration("deadline-duration")
	if err != nil {
		return nil, fmt.Errorf("reading --deadline-duration flag: %w", err)
	}

	cfg := &probe.DeadlineConfig{
		Enabled:  enabled,
		Duration: duration,
	}

	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("invalid deadline configuration: %w", err)
	}

	return cfg, nil
}
