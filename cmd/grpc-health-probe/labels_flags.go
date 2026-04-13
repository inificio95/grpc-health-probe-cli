package main

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/user/grpc-health-probe-cli/internal/probe"
)

// addLabelsFlags registers label-related flags onto cmd.
func addLabelsFlags(cmd *cobra.Command) {
	if cmd == nil {
		return
	}
	cmd.Flags().StringArray("label", []string{}, "Attach a key=value label to the probe result (repeatable)")
}

// parseLabelsConfig builds a LabelsConfig from cmd flags.
func parseLabelsConfig(cmd *cobra.Command) (*probe.LabelsConfig, error) {
	if cmd == nil {
		return probe.DefaultLabelsConfig(), nil
	}

	cfg := probe.DefaultLabelsConfig()

	raw, err := cmd.Flags().GetStringArray("label")
	if err != nil || len(raw) == 0 {
		return cfg, nil
	}

	cfg.Enabled = true
	for _, entry := range raw {
		parts := strings.SplitN(entry, "=", 2)
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid label format %q: expected key=value", entry)
		}
		cfg.Labels[parts[0]] = parts[1]
	}

	return cfg, nil
}
