package main

import (
	"errors"
	"time"

	"github.com/spf13/cobra"

	"github.com/user/grpc-health-probe-cli/internal/probe"
)

// addWatchFlags registers watch/polling flags onto cmd.
func addWatchFlags(cmd *cobra.Command) {
	if cmd == nil {
		return
	}
	cmd.Flags().Bool("watch", false, "Enable polling watch mode")
	cmd.Flags().Duration("watch-interval", 5*time.Second, "Interval between health checks in watch mode")
	cmd.Flags().Int("watch-max-checks", 0, "Maximum number of checks to perform (0 = unlimited)")
}

// parseWatchConfig builds a WatchConfig from the flags registered on cmd.
func parseWatchConfig(cmd *cobra.Command) (*probe.WatchConfig, error) {
	if cmd == nil {
		return nil, errors.New("command must not be nil")
	}

	cfg := probe.DefaultWatchConfig()

	enabled, err := cmd.Flags().GetBool("watch")
	if err != nil {
		return nil, err
	}
	if enabled {
		cfg.Mode = probe.WatchModePolling
	}

	interval, err := cmd.Flags().GetDuration("watch-interval")
	if err != nil {
		return nil, err
	}
	cfg.Interval = interval

	maxChecks, err := cmd.Flags().GetInt("watch-max-checks")
	if err != nil {
		return nil, err
	}
	cfg.MaxChecks = maxChecks

	if err := cfg.Validate(); err != nil {
		return nil, err
	}
	return cfg, nil
}
