package main

import (
	"time"

	"github.com/spf13/cobra"

	"github.com/user/grpc-health-probe-cli/internal/probe"
)

// addBurstFlags registers burst-related CLI flags onto cmd.
func addBurstFlags(cmd *cobra.Command) {
	if cmd == nil {
		return
	}
	cmd.Flags().Bool("burst-enabled", false, "Enable burst detection and suppression")
	cmd.Flags().Int("burst-max", 5, "Maximum number of probes allowed within the burst interval")
	cmd.Flags().Duration("burst-interval", 500*time.Millisecond, "Time window used to count burst attempts")
	cmd.Flags().Duration("burst-cooldown", 2*time.Second, "Cooldown period after a burst is detected")
}

// parseBurstConfig reads burst flags from cmd and returns a BurstConfig.
func parseBurstConfig(cmd *cobra.Command) *probe.BurstConfig {
	if cmd == nil {
		return probe.DefaultBurstConfig()
	}
	cfg := probe.DefaultBurstConfig()
	if v, err := cmd.Flags().GetBool("burst-enabled"); err == nil {
		cfg.Enabled = v
	}
	if v, err := cmd.Flags().GetInt("burst-max"); err == nil {
		cfg.MaxBurst = v
	}
	if v, err := cmd.Flags().GetDuration("burst-interval"); err == nil {
		cfg.BurstInterval = v
	}
	if v, err := cmd.Flags().GetDuration("burst-cooldown"); err == nil {
		cfg.Cooldown = v
	}
	return cfg
}
