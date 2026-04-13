package main

import (
	"github.com/spf13/cobra"

	"github.com/user/grpc-health-probe-cli/internal/probe"
)

// addSNIFlags registers SNI-related flags on the given command.
func addSNIFlags(cmd *cobra.Command) {
	if cmd == nil {
		return
	}
	cmd.Flags().Bool("sni", false, "Enable TLS SNI (Server Name Indication) override")
	cmd.Flags().String("sni-server-name", "", "Override the server name used in TLS SNI handshake")
}

// parseSNIConfig reads SNI flags from the command and returns a populated SNIConfig.
func parseSNIConfig(cmd *cobra.Command) *probe.SNIConfig {
	if cmd == nil {
		return probe.DefaultSNIConfig()
	}

	cfg := probe.DefaultSNIConfig()

	if enabled, err := cmd.Flags().GetBool("sni"); err == nil {
		cfg.Enabled = enabled
	}
	if serverName, err := cmd.Flags().GetString("sni-server-name"); err == nil {
		cfg.ServerName = serverName
	}

	return cfg
}
