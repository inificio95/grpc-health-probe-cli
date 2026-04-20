package main

import (
	"testing"
	"time"

	"github.com/spf13/cobra"
)

func newThrottleTestCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "test"}
	addThrottleFlags(cmd)
	return cmd
}

func TestAddThrottleFlags(t *testing.T) {
	cmd := newThrottleTestCmd()
	if cmd.Flags().Lookup("throttle") == nil {
		t.Error("expected --throttle flag")
	}
	if cmd.Flags().Lookup("throttle-min-interval") == nil {
		t.Error("expected --throttle-min-interval flag")
	}
	if cmd.Flags().Lookup("throttle-burst") == nil {
		t.Error("expected --throttle-burst flag")
	}
}

func TestAddThrottleFlags_NilCmd(t *testing.T) {
	// should not panic
	addThrottleFlags(nil)
}

func TestParseThrottleConfig_Defaults(t *testing.T) {
	cmd := newThrottleTestCmd()
	cfg := parseThrottleConfig(cmd)
	if cfg.Enabled {
		t.Error("expected throttle disabled by default")
	}
	if cfg.MinInterval != 500*time.Millisecond {
		t.Errorf("unexpected default min_interval: %s", cfg.MinInterval)
	}
	if cfg.Burst != 1 {
		t.Errorf("unexpected default burst: %d", cfg.Burst)
	}
}

func TestParseThrottleConfig_Enabled(t *testing.T) {
	cmd := newThrottleTestCmd()
	_ = cmd.Flags().Set("throttle", "true")
	_ = cmd.Flags().Set("throttle-min-interval", "250ms")
	_ = cmd.Flags().Set("throttle-burst", "3")

	cfg := parseThrottleConfig(cmd)
	if !cfg.Enabled {
		t.Error("expected throttle enabled")
	}
	if cfg.MinInterval != 250*time.Millisecond {
		t.Errorf("unexpected min_interval: %s", cfg.MinInterval)
	}
	if cfg.Burst != 3 {
		t.Errorf("unexpected burst: %d", cfg.Burst)
	}
}

func TestParseThrottleConfig_NilCmd(t *testing.T) {
	cfg := parseThrottleConfig(nil)
	if cfg == nil {
		t.Fatal("expected non-nil config")
	}
	if cfg.Enabled {
		t.Error("expected throttle disabled for nil cmd")
	}
}
