package main

import (
	"testing"
	"time"

	"github.com/spf13/cobra"
)

func newBudgetTestCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "test"}
	addTimeoutBudgetFlags(cmd)
	return cmd
}

func TestAddTimeoutBudgetFlags(t *testing.T) {
	cmd := newBudgetTestCmd()
	for _, name := range []string{"budget-enabled", "budget-total", "budget-reserve"} {
		if cmd.Flags().Lookup(name) == nil {
			t.Errorf("expected flag %q to be registered", name)
		}
	}
}

func TestAddTimeoutBudgetFlags_NilCmd(t *testing.T) {
	// Should not panic
	addTimeoutBudgetFlags(nil)
}

func TestParseTimeoutBudgetConfig_Defaults(t *testing.T) {
	cmd := newBudgetTestCmd()
	cfg, err := parseTimeoutBudgetConfig(cmd)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.Enabled {
		t.Error("expected Enabled false by default")
	}
	if cfg.TotalBudget != 30*time.Second {
		t.Errorf("expected TotalBudget 30s, got %v", cfg.TotalBudget)
	}
	if cfg.ReserveBuffer != 500*time.Millisecond {
		t.Errorf("expected ReserveBuffer 500ms, got %v", cfg.ReserveBuffer)
	}
}

func TestParseTimeoutBudgetConfig_Enabled(t *testing.T) {
	cmd := newBudgetTestCmd()
	_ = cmd.Flags().Set("budget-enabled", "true")
	_ = cmd.Flags().Set("budget-total", "20s")
	_ = cmd.Flags().Set("budget-reserve", "1s")

	cfg, err := parseTimeoutBudgetConfig(cmd)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !cfg.Enabled {
		t.Error("expected Enabled true")
	}
	if cfg.TotalBudget != 20*time.Second {
		t.Errorf("expected TotalBudget 20s, got %v", cfg.TotalBudget)
	}
	if cfg.ReserveBuffer != 1*time.Second {
		t.Errorf("expected ReserveBuffer 1s, got %v", cfg.ReserveBuffer)
	}
}

func TestParseTimeoutBudgetConfig_NilCmd(t *testing.T) {
	cfg, err := parseTimeoutBudgetConfig(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg == nil {
		t.Fatal("expected non-nil config")
	}
}
