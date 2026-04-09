package main

import (
	"github.com/spf13/cobra"
	"github.com/your-org/grpc-health-probe-cli/internal/probe"
)

// addServiceFlags registers service-related CLI flags onto the given command.
func addServiceFlags(cmd *cobra.Command, cfg *probe.ServiceConfig) {
	if cfg == nil {
		return
	}

	cmd.Flags().StringVar(
		&cfg.Service,
		"service",
		cfg.Service,
		`gRPC health check service name. Leave empty to check server-level health.
Example: --service=mypackage.MyService`,
	)
}

// parseServiceConfig builds a ServiceConfig from the parsed CLI flags.
func parseServiceConfig(cmd *cobra.Command) (*probe.ServiceConfig, error) {
	cfg := probe.DefaultServiceConfig()

	if cmd.Flags().Changed("service") {
		val, err := cmd.Flags().GetString("service")
		if err != nil {
			return nil, err
		}
		cfg.Service = val
	}

	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	return cfg, nil
}
