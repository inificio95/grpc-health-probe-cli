package main

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/your-org/grpc-health-probe-cli/internal/probe"
)

// addTimeoutBudgetFlags registers timeout budget flags onto cmd.
func addTimeoutBudgetFlags(cmd *cobra.Command) {
	if cmd == nil {
		return
	}
	cmd.Flags().Bool("budget-enabled", false, "Enable a shared timeout budget across all retry attempts")
	cmd.Flags().Duration("budget-total", 30*time.Second, "Total time budget for all retry attempts combined")
	cmd.Flags().Duration("budget-reserve", 500*time.Millisecond, "Time reserved as a buffer before the budget deadline")
}

// parseTimeoutBudgetConfig builds a TimeoutBudgetConfig from parsed cobra flags.
func parseTimeoutBudgetConfig(cmd *cobra.Command) (*probe.TimeoutBudgetConfig, error) {
	if cmd == nil {
		return probe.DefaultTimeoutBudgetConfig(), nil
	}

	cfg := probe.DefaultTimeoutBudgetConfig()

	if f := cmd.Flags().Lookup("budget-enabled"); f != nil {
		v, err := cmd.Flags().GetBool("budget-enabled")
		if err != nil {
			return nil, fmt.Errorf("timeout budget: %w", err)
		}
		cfg.Enabled = v
	}

	if f := cmd.Flags().Lookup("budget-total"); f != nil {
		v, err := cmd.Flags().GetDuration("budget-total")
		if err != nil {
			return nil, fmt.Errorf("timeout budget: %w", err)
		}
		cfg.TotalBudget = v
	}

	if f := cmd.Flags().Lookup("budget-reserve"); f != nil {
		v, err := cmd.Flags().GetDuration("budget-reserve")
		if err != nil {
			return nil, fmt.Errorf("timeout budget: %w", err)
		}
		cfg.ReserveBuffer = v
	}

	return cfg, nil
}
