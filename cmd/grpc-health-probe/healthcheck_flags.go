package main

import (
	"time"

	"github.com/spf13/cobra"

	"github.com/example/grpc-health-probe-cli/internal/probe"
)

// addHealthCheckFlags registers health-check related flags on cmd.
func addHealthCheckFlags(cmd *cobra.Command) {
	if cmd == nil {
		return
	}
	cmd.Flags().Bool("hc-enabled", true, "Enable health-check specific request logic")
	cmd.Flags().Bool("hc-include-service", true, "Include the service name in the health check request")
	cmd.Flags().Duration("hc-interval", 5*time.Second, "Minimum interval between health checks in watch mode")
	cmd.Flags().Int("hc-max-failures", 0, "Max consecutive failures before hard failure (0 = unlimited)")
}

// parseHealthCheckConfig builds a HealthCheckConfig from the flags registered
// by addHealthCheckFlags. It returns the default config when cmd is nil.
func parseHealthCheckConfig(cmd *cobra.Command) *probe.HealthCheckConfig {
	cfg := probe.DefaultHealthCheckConfig()
	if cmd == nil {
		return cfg
	}

	if v, err := cmd.Flags().GetBool("hc-enabled"); err == nil {
		cfg.Enabled = v
	}
	if v, err := cmd.Flags().GetBool("hc-include-service"); err == nil {
		cfg.IncludeServiceName = v
	}
	if v, err := cmd.Flags().GetDuration("hc-interval"); err == nil {
		cfg.CheckInterval = v
	}
	if v, err := cmd.Flags().GetInt("hc-max-failures"); err == nil {
		cfg.MaxConsecutiveFailures = v
	}
	return cfg
}
