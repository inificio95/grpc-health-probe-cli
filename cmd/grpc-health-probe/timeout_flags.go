package main

import (
	"time"

	"github.com/spf13/cobra"
	"grpc-health-probe-cli/internal/probe"
)

// addTimeoutFlags adds timeout-related flags to the provided command.
func addTimeoutFlags(cmd *cobra.Command, config *probe.TimeoutConfig) {
	if config == nil {
		return
	}

	cmd.Flags().DurationVar(
		&config.RequestTimeout,
		"timeout",
		probe.DefaultTimeoutConfig().RequestTimeout,
		"Timeout for the health check request",
	)

	cmd.Flags().DurationVar(
		&config.DialTimeout,
		"dial-timeout",
		probe.DefaultTimeoutConfig().DialTimeout,
		"Timeout for establishing connection to gRPC server",
	)

	cmd.Flags().DurationVar(
		&config.ConnectTimeout,
		"connect-timeout",
		probe.DefaultTimeoutConfig().ConnectTimeout,
		"Timeout for the entire connection process",
	)
}

// parseTimeoutConfig creates a TimeoutConfig from command flags.
func parseTimeoutConfig(cmd *cobra.Command) (*probe.TimeoutConfig, error) {
	requestTimeout, err := cmd.Flags().GetDuration("timeout")
	if err != nil {
		return nil, err
	}

	dialTimeout, err := cmd.Flags().GetDuration("dial-timeout")
	if err != nil {
		return nil, err
	}

	connectTimeout, err := cmd.Flags().GetDuration("connect-timeout")
	if err != nil {
		return nil, err
	}

	return &probe.TimeoutConfig{
		RequestTimeout:  requestTimeout,
		DialTimeout:     dialTimeout,
		ConnectTimeout:  connectTimeout,
	}, nil
}
