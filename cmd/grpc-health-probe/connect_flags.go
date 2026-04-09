package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/user/grpc-health-probe-cli/internal/probe"
)

// addConnectFlags registers connection-related flags onto cmd.
func addConnectFlags(cmd *cobra.Command) {
	if cmd == nil {
		return
	}
	cmd.Flags().String("user-agent", "grpc-health-probe-cli/1.0", "User-Agent header sent with each request")
	cmd.Flags().Bool("connect-block", false, "Block until the connection is established before sending the health check")
	cmd.Flags().Bool("connect-fail-fast", false, "Fail immediately on non-temporary dial errors")
}

// parseConnectConfig reads connection flags from cmd and returns a ConnectConfig.
func parseConnectConfig(cmd *cobra.Command) (*probe.ConnectConfig, error) {
	if cmd == nil {
		return probe.DefaultConnectConfig(), nil
	}

	cfg := probe.DefaultConnectConfig()

	if ua, err := cmd.Flags().GetString("user-agent"); err == nil {
		cfg.UserAgent = ua
	}

	if block, err := cmd.Flags().GetBool("connect-block"); err == nil {
		cfg.Block = block
	}

	if failFast, err := cmd.Flags().GetBool("connect-fail-fast"); err == nil {
		cfg.FailOnNonTempDialError = failFast
	}

	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("invalid connect config: %w", err)
	}

	return cfg, nil
}
