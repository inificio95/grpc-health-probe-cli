package main

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/example/grpc-health-probe-cli/internal/probe"
)

// addQuorumFlags registers quorum-related flags onto cmd.
func addQuorumFlags(cmd *cobra.Command) {
	if cmd == nil {
		return
	}
	cmd.Flags().Bool("quorum", false, "Enable quorum mode: require min-success out of total checks")
	cmd.Flags().Int("quorum-min-success", 1, "Minimum successful checks required to satisfy quorum")
	cmd.Flags().Int("quorum-total", 1, "Total number of checks to perform for quorum evaluation")
}

// parseQuorumConfig builds a QuorumConfig from the flags registered on cmd.
// Returns an error if cmd is nil or if the resulting config fails validation.
func parseQuorumConfig(cmd *cobra.Command) (*probe.QuorumConfig, error) {
	if cmd == nil {
		return nil, fmt.Errorf("command must not be nil")
	}

	enabled, err := cmd.Flags().GetBool("quorum")
	if err != nil {
		return nil, fmt.Errorf("reading --quorum: %w", err)
	}

	minSuccess, err := cmd.Flags().GetInt("quorum-min-success")
	if err != nil {
		return nil, fmt.Errorf("reading --quorum-min-success: %w", err)
	}

	total, err := cmd.Flags().GetInt("quorum-total")
	if err != nil {
		return nil, fmt.Errorf("reading --quorum-total: %w", err)
	}

	cfg := &probe.QuorumConfig{
		Enabled:    enabled,
		MinSuccess: minSuccess,
		Total:      total,
	}

	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("invalid quorum config: %w", err)
	}

	return cfg, nil
}
