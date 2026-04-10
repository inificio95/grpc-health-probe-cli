package main

import (
	"testing"
	"time"

	"github.com/spf13/cobra"
)

func newCBTestCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "test"}
	addCircuitBreakerFlags(cmd)
	return cmd
}

func TestAddCircuitBreakerFlags(t *testing.T) {
	cmd := newCBTestCmd()
	for _, name := range []string{
		"circuit-breaker",
		"circuit-breaker-max-failures",
		"circuit-breaker-open-duration",
		"circuit-breaker-half-open-requests",
	} {
		if cmd.Flags().Lookup(name) == nil {
			t.Errorf("expected flag %q to be registered", name)
		}
	}
}

func TestAddCircuitBreakerFlags_NilCmd(t *testing.T) {
	// Should not panic
	addCircuitBreakerFlags(nil)
}

func TestParseCircuitBreakerConfig_Defaults(t *testing.T) {
	cmd := newCBTestCmd()
	_ = cmd.ParseFlags([]string{})
	cfg := parseCircuitBreakerConfig(cmd)
	if cfg.Enabled {
		t.Error("expected Enabled=false by default")
	}
	if cfg.MaxFailures != 5 {
		t.Errorf("expected MaxFailures=5, got %d", cfg.MaxFailures)
	}
	if cfg.OpenDuration != 30*time.Second {
		t.Errorf("expected OpenDuration=30s, got %s", cfg.OpenDuration)
	}
	if cfg.HalfOpenRequests != 1 {
		t.Errorf("expected HalfOpenRequests=1, got %d", cfg.HalfOpenRequests)
	}
}

func TestParseCircuitBreakerConfig_Enabled(t *testing.T) {
	cmd := newCBTestCmd()
	_ = cmd.ParseFlags([]string{
		"--circuit-breaker",
		"--circuit-breaker-max-failures=3",
		"--circuit-breaker-open-duration=10s",
		"--circuit-breaker-half-open-requests=2",
	})
	cfg := parseCircuitBreakerConfig(cmd)
	if !cfg.Enabled {
		t.Error("expected Enabled=true")
	}
	if cfg.MaxFailures != 3 {
		t.Errorf("expected MaxFailures=3, got %d", cfg.MaxFailures)
	}
	if cfg.OpenDuration != 10*time.Second {
		t.Errorf("expected OpenDuration=10s, got %s", cfg.OpenDuration)
	}
	if cfg.HalfOpenRequests != 2 {
		t.Errorf("expected HalfOpenRequests=2, got %d", cfg.HalfOpenRequests)
	}
}

func TestParseCircuitBreakerConfig_NilCmd(t *testing.T) {
	cfg := parseCircuitBreakerConfig(nil)
	if cfg == nil {
		t.Fatal("expected non-nil config for nil cmd")
	}
	if cfg.Enabled {
		t.Error("expected Enabled=false for nil cmd")
	}
}
