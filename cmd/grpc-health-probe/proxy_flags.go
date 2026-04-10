package main

import (
	"github.com/spf13/cobra"

	"github.com/yourusername/grpc-health-probe-cli/internal/probe"
)

// addProxyFlags registers proxy-related flags on the given command.
func addProxyFlags(cmd *cobra.Command) {
	if cmd == nil {
		return
	}
	cmd.Flags().Bool("proxy", false, "Enable HTTP/HTTPS proxy for gRPC connections")
	cmd.Flags().String("proxy-url", "", "Proxy URL (e.g. http://proxy.example.com:8080)")
}

// parseProxyConfig builds a ProxyConfig from the command's flags.
func parseProxyConfig(cmd *cobra.Command) *probe.ProxyConfig {
	if cmd == nil {
		return probe.DefaultProxyConfig()
	}

	cfg := probe.DefaultProxyConfig()

	if enabled, err := cmd.Flags().GetBool("proxy"); err == nil {
		cfg.Enabled = enabled
	}
	if proxyURL, err := cmd.Flags().GetString("proxy-url"); err == nil {
		cfg.ProxyURL = proxyURL
	}

	return cfg
}
