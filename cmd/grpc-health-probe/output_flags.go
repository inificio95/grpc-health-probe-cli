package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/your-org/grpc-health-probe-cli/internal/probe"
)

// addOutputFlags registers output-related CLI flags onto the given command.
// Returns an error if cmd is nil.
func addOutputFlags(cmd *cobra.Command) error {
	if cmd == nil {
		return fmt.Errorf("cmd must not be nil")
	}

	cmd.Flags().StringP(
		"format",
		"f",
		"text",
		`Output format for health check results. Supported values: "text", "json"`,
	)

	cmd.Flags().BoolP(
		"no-color",
		"",
		false,
		"Disable ANSI color codes in text output",
	)

	return nil
}

// parseOutputConfig builds a probe.OutputConfig from the flags registered by
// addOutputFlags. Falls back to probe.DefaultOutputConfig() when cmd is nil.
func parseOutputConfig(cmd *cobra.Command) *probe.OutputConfig {
	defaults := probe.DefaultOutputConfig()

	if cmd == nil {
		return defaults
	}

	cfg := &probe.OutputConfig{
		Writer: defaults.Writer,
	}

	if f := cmd.Flags().Lookup("format"); f != nil {
		cfg.Format = f.Value.String()
	} else {
		cfg.Format = defaults.Format
	}

	if f := cmd.Flags().Lookup("no-color"); f != nil {
		cfg.NoColor = f.Value.String() == "true"
	}

	return cfg
}
