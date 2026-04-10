package main

import (
	"github.com/spf13/cobra"

	"github.com/your-org/grpc-health-probe-cli/internal/probe"
)

// addResolverFlags registers DNS resolver flags onto cmd.
func addResolverFlags(cmd *cobra.Command) {
	if cmd == nil {
		return
	}
	f := cmd.Flags()
	f.Bool("resolve", false, "Pre-resolve the target hostname via DNS before dialing")
	f.Bool("resolve-prefer-ipv6", false, "Prefer IPv6 addresses when resolving the target hostname")
	f.String("resolve-dns-server", "", "Custom DNS server to use for resolution (host:port, e.g. 8.8.8.8:53)")
}

// parseResolverConfig builds a ResolverConfig from the flags registered by
// addResolverFlags. It returns the default config when cmd is nil.
func parseResolverConfig(cmd *cobra.Command) *probe.ResolverConfig {
	cfg := probe.DefaultResolverConfig()
	if cmd == nil {
		return cfg
	}

	if v, err := cmd.Flags().GetBool("resolve"); err == nil {
		cfg.Enabled = v
	}
	if v, err := cmd.Flags().GetBool("resolve-prefer-ipv6"); err == nil {
		cfg.PreferIPv6 = v
	}
	if v, err := cmd.Flags().GetString("resolve-dns-server"); err == nil {
		cfg.CustomResolver = v
	}

	return cfg
}
