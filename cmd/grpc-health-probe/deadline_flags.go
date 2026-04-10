package main

import (
	"errors"
	"time"

	"github.com/spf13/cobra"

	"github.com/user/grpc-health-probe-cli/internal/probe"
)

// addDeadlineFlags registers deadline-related flags onto cmd.
func addDeadlineFlags(cmd *cobra.Command) {
	if cmd == nil {
		return
	}
	cmd.Flags().Bool("deadline", false, "Enable an overall deadline for the probe")
	cmd.Flags().Duration("deadline-duration", 10*time.Second, "Overall deadline duration (requires --deadline)")
}

// parseDeadlineConfig builds a DeadlineConfig from the flags set on cmd.
func parseDeadlineConfig(cmd *cobra.Command) (*probe.DeadlineConfig, error) {
	if cmd == nil {
		return nil, errors.New("parseDeadlineConfig: nil command")
	}

	cfg := probe.DefaultDeadlineConfig()

	if v, err := cmd.Flags().GetBool("deadline"); err == nil {
		cfg.Enabled = v
	}
	if v, err := cmd.Flags().GetDuration("deadline-duration"); err == nil {
		cfg.Duration = v
	}

	if err := cfg.Validate(); err != nil {
		return nil, err
	}
	return cfg, nil
}
