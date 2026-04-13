package main

import (
	"github.com/spf13/cobra"

	"github.com/your-org/grpc-health-probe-cli/internal/probe"
)

// addNamespaceFlags registers namespace-related CLI flags onto cmd.
func addNamespaceFlags(cmd *cobra.Command) {
	cmd.Flags().Bool("namespace-enabled", false, "Enable namespace prefixing for the service name")
	cmd.Flags().String("namespace", "", "Namespace prefix to prepend to the service name (e.g. 'prod')")
}

// parseNamespaceConfig reads namespace flags from cmd and returns a NamespaceConfig.
// If cmd is nil, the default config is returned.
func parseNamespaceConfig(cmd *cobra.Command) *probe.NamespaceConfig {
	defaults := probe.DefaultNamespaceConfig()
	if cmd == nil {
		return defaults
	}

	enabled, err := cmd.Flags().GetBool("namespace-enabled")
	if err != nil {
		return defaults
	}

	ns, err := cmd.Flags().GetString("namespace")
	if err != nil {
		return defaults
	}

	return &probe.NamespaceConfig{
		Enabled:   enabled,
		Namespace: ns,
	}
}
