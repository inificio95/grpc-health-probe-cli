package main

import (
	"testing"
	"time"

	"github.com/spf13/cobra"
)

func TestAddDeadlineFlags(t *testing.T) {
	cmd := &cobra.Command{Use: "test"}
	addDeadlineFlags(cmd)

	if !cmd.Flags().HasFlags() {
		t.Fatal("expected flags to be registered")
	}
	if cmd.Flags().Lookup("deadline") == nil {
		t.Error("expected --deadline flag")
	}
	if cmd.Flags().Lookup("deadline-duration") == nil {
		t.Error("expected --deadline-duration flag")
	}
}

func TestAddDeadlineFlags_NilCmd(t *testing.T) {
	// Should not panic
	addDeadlineFlags(nil)
}

func TestParseDeadlineConfig_Defaults(t *testing.T) {
	cmd := &cobra.Command{Use: "test"}
	addDeadlineFlags(cmd)

	cfg, err := parseDeadlineConfig(cmd)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.Enabled {
		t.Error("expected deadline disabled by default")
	}
	if cfg.Duration != 10*time.Second {
		t.Errorf("expected 10s default, got %v", cfg.Duration)
	}
}

func TestParseDeadlineConfig_Enabled(t *testing.T) {
	cmd := &cobra.Command{Use: "test"}
	addDeadlineFlags(cmd)
	_ = cmd.Flags().Set("deadline", "true")
	_ = cmd.Flags().Set("deadline-duration", "5s")

	cfg, err := parseDeadlineConfig(cmd)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !cfg.Enabled {
		t.Error("expected deadline enabled")
	}
	if cfg.Duration != 5*time.Second {
		t.Errorf("expected 5s, got %v", cfg.Duration)
	}
}

func TestParseDeadlineConfig_NilCmd(t *testing.T) {
	_, err := parseDeadlineConfig(nil)
	if err == nil {
		t.Error("expected error for nil command")
	}
}
