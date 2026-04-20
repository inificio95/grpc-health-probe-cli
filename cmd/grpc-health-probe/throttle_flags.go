package main

import (
	"time"

	"github.com/spf13/cobra"

	"github.com/user/grpc-health-probe-cli/internal/probe"
)

// addThrottleFlags registers throttle-related CLI flags onto cmd.
func addThrottleFlags(cmd *cobra.Command) {
	if cmd == nil {
		return
	}
	cmd.Flags().Bool("throttle", false, "enable request throttling")
	cmd.Flags().Duration("throttle-min-interval", 500*time.Millisecond, "minimum interval between probe calls")
	cmd.Flags().Int("throttle-burst", 1, "number of burst requests allowed before throttling")
}

// parseThrottleConfig builds a ThrottleConfig from parsed CLI flags.
func parseThrottleConfig(cmd *cobra.Command) *probe.ThrottleConfig {
	if cmd == nil {
		return probe.DefaultThrottleConfig()
	}
	cfg := probe.DefaultThrottleConfig()

	if v, err := cmd.Flags().GetBool("throttle"); err == nil {
		cfg.Enabled = v
	}
	if v, err := cmd.Flags().GetDuration("throttle-min-interval"); err == nil {
		cfg.MinInterval = v
	}
	if v, err := cmd.Flags().GetInt("throttle-burst"); err == nil {
		cfg.Burst = v
	}
	return cfg
}
