package main

import (
	"github.com/example/grpc-health-probe-cli/internal/probe"
	"github.com/spf13/cobra"
)

// addNamespaceFlags registers namespace-related flags on the given command.
func addNamespaceFlags(cmd *cobra.Command) {
	if cmd == nil {
		return
	}
	cmd.Flags().Bool("namespace-enabled", false, "Enable namespace qualification for probe identifiers")
	cmd.Flags().String("namespace", "", "Namespace prefix to qualify probe names")
	cmd.Flags().String("namespace-separator", ".", "Separator between namespace and probe name")
}

// parseNamespaceConfig builds a NamespaceConfig from the command's parsed flags.
func parseNamespaceConfig(cmd *cobra.Command) *probe.NamespaceConfig {
	if cmd == nil {
		return probe.DefaultNamespaceConfig()
	}

	cfg := probe.DefaultNamespaceConfig()

	if v, err := cmd.Flags().GetBool("namespace-enabled"); err == nil {
		cfg.Enabled = v
	}
	if v, err := cmd.Flags().GetString("namespace"); err == nil && v != "" {
		cfg.Namespace = v
	}
	if v, err := cmd.Flags().GetString("namespace-separator"); err == nil && v != "" {
		cfg.Separator = v
	}

	return cfg
}
