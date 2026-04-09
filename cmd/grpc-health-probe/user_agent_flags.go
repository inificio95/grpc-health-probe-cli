package main

import (
	"github.com/example/grpc-health-probe-cli/internal/probe"
	"github.com/spf13/cobra"
)

// addUserAgentFlags registers user-agent related flags on cmd.
func addUserAgentFlags(cmd *cobra.Command) {
	if cmd == nil {
		return
	}
	defaults := probe.DefaultUserAgentConfig()
	cmd.Flags().String("user-agent-name", defaults.AppName, "Application name included in the gRPC user-agent header")
	cmd.Flags().String("user-agent-version", defaults.AppVersion, "Application version included in the gRPC user-agent header")
	cmd.Flags().Bool("no-user-agent", defaults.Disabled, "Disable the custom gRPC user-agent header")
}

// parseUserAgentConfig builds a UserAgentConfig from cobra flags.
func parseUserAgentConfig(cmd *cobra.Command) *probe.UserAgentConfig {
	if cmd == nil {
		return probe.DefaultUserAgentConfig()
	}
	cfg := probe.DefaultUserAgentConfig()
	if name, err := cmd.Flags().GetString("user-agent-name"); err == nil {
		cfg.AppName = name
	}
	if version, err := cmd.Flags().GetString("user-agent-version"); err == nil {
		cfg.AppVersion = version
	}
	if disabled, err := cmd.Flags().GetBool("no-user-agent"); err == nil {
		cfg.Disabled = disabled
	}
	return cfg
}
