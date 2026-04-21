package main

import (
	"testing"
	"time"

	"github.com/spf13/cobra"
)

func newBurstTestCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "test"}
	addBurstFlags(cmd)
	return cmd
}

func TestAddBurstFlags(t *testing.T) {
	cmd := newBurstTestCmd()
	for _, name := range []string{"burst-enabled", "burst-max", "burst-interval", "burst-cooldown"} {
		if cmd.Flags().Lookup(name) == nil {
			t.Errorf("expected flag %q to be registered", name)
		}
	}
}

func TestAddBurstFlags_NilCmd(t *testing.T) {
	// should not panic
	addBurstFlags(nil)
}

func TestParseBurstConfig_Defaults(t *testing.T) {
	cmd := newBurstTestCmd()
	_ = cmd.Flags().Parse([]string{})
	cfg := parseBurstConfig(cmd)
	if cfg == nil {
		t.Fatal("expected non-nil config")
	}
	if cfg.Enabled {
		t.Error("expected Enabled=false by default")
	}
	if cfg.MaxBurst != 5 {
		t.Errorf("expected MaxBurst=5, got %d", cfg.MaxBurst)
	}
	if cfg.BurstInterval != 500*time.Millisecond {
		t.Errorf("unexpected BurstInterval: %s", cfg.BurstInterval)
	}
	if cfg.Cooldown != 2*time.Second {
		t.Errorf("unexpected Cooldown: %s", cfg.Cooldown)
	}
}

func TestParseBurstConfig_Enabled(t *testing.T) {
	cmd := newBurstTestCmd()
	_ = cmd.Flags().Parse([]string{
		"--burst-enabled",
		"--burst-max=10",
		"--burst-interval=1s",
		"--burst-cooldown=5s",
	})
	cfg := parseBurstConfig(cmd)
	if !cfg.Enabled {
		t.Error("expected Enabled=true")
	}
	if cfg.MaxBurst != 10 {
		t.Errorf("expected MaxBurst=10, got %d", cfg.MaxBurst)
	}
	if cfg.BurstInterval != time.Second {
		t.Errorf("unexpected BurstInterval: %s", cfg.BurstInterval)
	}
	if cfg.Cooldown != 5*time.Second {
		t.Errorf("unexpected Cooldown: %s", cfg.Cooldown)
	}
}

func TestParseBurstConfig_NilCmd(t *testing.T) {
	cfg := parseBurstConfig(nil)
	if cfg == nil {
		t.Fatal("expected non-nil config from nil cmd")
	}
	if cfg.Enabled {
		t.Error("expected Enabled=false for nil cmd")
	}
}
