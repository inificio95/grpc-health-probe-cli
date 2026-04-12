package main

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/user/grpc-health-probe-cli/internal/probe"
)

// addExitCodeFlags registers exit-code-related flags onto cmd.
func addExitCodeFlags(cmd *cobra.Command) {
	if cmd == nil {
		return
	}
	cmd.Flags().Bool("exit-code-enabled", true, "Map probe results to specific exit codes")
	cmd.Flags().Int("exit-code-not-serving", 1, "Exit code when service is not serving")
	cmd.Flags().Int("exit-code-unknown", 2, "Exit code when service status is unknown")
	cmd.Flags().Int("exit-code-timeout", 4, "Exit code when probe times out")
}

// parseExitCodeConfig builds an ExitCodeConfig from cobra flags.
func parseExitCodeConfig(cmd *cobra.Command) (*probe.ExitCodeConfig, error) {
	if cmd == nil {
		return probe.DefaultExitCodeConfig(), nil
	}
	cfg := probe.DefaultExitCodeConfig()

	if f := cmd.Flags().Lookup("exit-code-enabled"); f != nil {
		v, err := cmd.Flags().GetBool("exit-code-enabled")
		if err != nil {
			return nil, fmt.Errorf("exit-code-enabled: %w", err)
		}
		cfg.Enabled = v
	}
	if f := cmd.Flags().Lookup("exit-code-not-serving"); f != nil {
		v, err := cmd.Flags().GetInt("exit-code-not-serving")
		if err != nil {
			return nil, fmt.Errorf("exit-code-not-serving: %w", err)
		}
		cfg.NotServingCode = v
	}
	if f := cmd.Flags().Lookup("exit-code-unknown"); f != nil {
		v, err := cmd.Flags().GetInt("exit-code-unknown")
		if err != nil {
			return nil, fmt.Errorf("exit-code-unknown: %w", err)
		}
		cfg.UnknownCode = v
	}
	if f := cmd.Flags().Lookup("exit-code-timeout"); f != nil {
		v, err := cmd.Flags().GetInt("exit-code-timeout")
		if err != nil {
			return nil, fmt.Errorf("exit-code-timeout: %w", err)
		}
		cfg.TimeoutCode = v
	}
	return cfg, nil
}
