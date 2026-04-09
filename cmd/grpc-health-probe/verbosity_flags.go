package main

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/user/grpc-health-probe-cli/internal/probe"
)

// addVerbosityFlags registers verbosity-related flags on the given command.
func addVerbosityFlags(cmd *cobra.Command) {
	if cmd == nil {
		return
	}
	cmd.Flags().BoolP("quiet", "q", false, "suppress all output except errors")
	cmd.Flags().BoolP("verbose", "v", false, "enable verbose diagnostic output")
}

// parseVerbosityConfig reads verbosity flags from the command and returns a VerbosityConfig.
func parseVerbosityConfig(cmd *cobra.Command) (*probe.VerbosityConfig, error) {
	if cmd == nil {
		return probe.DefaultVerbosityConfig(), nil
	}

	cfg := probe.DefaultVerbosityConfig()

	quiet, err := cmd.Flags().GetBool("quiet")
	if err != nil {
		return nil, fmt.Errorf("failed to parse --quiet flag: %w", err)
	}

	verbose, err := cmd.Flags().GetBool("verbose")
	if err != nil {
		return nil, fmt.Errorf("failed to parse --verbose flag: %w", err)
	}

	if quiet && verbose {
		return nil, fmt.Errorf("--quiet and --verbose are mutually exclusive")
	}

	switch {
	case quiet:
		cfg.Level = probe.VerbosityQuiet
	case verbose:
		cfg.Level = probe.VerbosityVerbose
	default:
		cfg.Level = probe.VerbosityNormal
	}

	return cfg, nil
}
