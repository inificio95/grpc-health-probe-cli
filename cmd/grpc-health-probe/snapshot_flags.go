package main

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/user/grpc-health-probe-cli/internal/probe"
)

// addSnapshotFlags registers snapshot-related CLI flags on cmd.
func addSnapshotFlags(cmd *cobra.Command) {
	if cmd == nil {
		return
	}
	cmd.Flags().Bool("snapshot", false, "Enable writing probe result to a snapshot file")
	cmd.Flags().String("snapshot-path", "", "Path to the snapshot output file")
	cmd.Flags().String("snapshot-format", "json", "Snapshot output format: text or json")
}

// parseSnapshotConfig builds a SnapshotConfig from the command's parsed flags.
func parseSnapshotConfig(cmd *cobra.Command) (*probe.SnapshotConfig, error) {
	if cmd == nil {
		return probe.DefaultSnapshotConfig(), nil
	}

	cfg := probe.DefaultSnapshotConfig()

	if v, err := cmd.Flags().GetBool("snapshot"); err == nil {
		cfg.Enabled = v
	}
	if v, err := cmd.Flags().GetString("snapshot-path"); err == nil {
		cfg.FilePath = v
	}
	if v, err := cmd.Flags().GetString("snapshot-format"); err == nil {
		cfg.Format = v
	}

	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("invalid snapshot configuration: %w", err)
	}
	return cfg, nil
}
