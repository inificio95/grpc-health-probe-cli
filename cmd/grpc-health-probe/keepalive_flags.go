package main

import (
	"time"

	"github.com/spf13/cobra"
	"github.com/your-org/grpc-health-probe-cli/internal/probe"
)

// addKeepaliveFlags registers keepalive-related flags on the given command.
func addKeepaliveFlags(cmd *cobra.Command) {
	if cmd == nil {
		return
	}
	cmd.Flags().Bool("keepalive", false, "Enable gRPC client keepalive")
	cmd.Flags().Duration("keepalive-time", 30*time.Second, "Interval between keepalive pings")
	cmd.Flags().Duration("keepalive-timeout", 10*time.Second, "Timeout waiting for keepalive ping ack")
	cmd.Flags().Bool("keepalive-permit-without-stream", false, "Send keepalive pings even without active RPCs")
}

// parseKeepaliveConfig reads keepalive flags from the command and returns a KeepaliveConfig.
func parseKeepaliveConfig(cmd *cobra.Command) *probe.KeepaliveConfig {
	cfg := probe.DefaultKeepaliveConfig()
	if cmd == nil {
		return cfg
	}

	if v, err := cmd.Flags().GetBool("keepalive"); err == nil {
		cfg.Enabled = v
	}
	if v, err := cmd.Flags().GetDuration("keepalive-time"); err == nil {
		cfg.Time = v
	}
	if v, err := cmd.Flags().GetDuration("keepalive-timeout"); err == nil {
		cfg.Timeout = v
	}
	if v, err := cmd.Flags().GetBool("keepalive-permit-without-stream"); err == nil {
		cfg.PermitWithoutStream = v
	}
	return cfg
}
