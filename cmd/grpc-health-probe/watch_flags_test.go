package main

import (
	"testing"
	"time"

	"github.com/spf13/cobra"
)

func TestAddWatchFlags(t *testing.T) {
	cmd := &cobra.Command{Use: "test"}
	addWatchFlags(cmd)

	flags := cmd.Flags()
	if flags.Lookup("watch") == nil {
		t.Error("expected --watch flag to be registered")
	}
	if flags.Lookup("watch-interval") == nil {
		t.Error("expected --watch-interval flag to be registered")
	}
}

func TestAddWatchFlags_NilCmd(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic when cmd is nil")
		}
	}()
	addWatchFlags(nil)
}

func TestParseWatchConfig_Defaults(t *testing.T) {
	cmd := &cobra.Command{Use: "test"}
	addWatchFlags(cmd)

	cfg, err := parseWatchConfig(cmd)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.Enabled {
		t.Error("expected watch to be disabled by default")
	}
	if cfg.Interval != 10*time.Second {
		t.Errorf("expected default interval 10s, got %v", cfg.Interval)
	}
}

func TestParseWatchConfig_Enabled(t *testing.T) {
	cmd := &cobra.Command{Use: "test"}
	addWatchFlags(cmd)
	_ = cmd.Flags().Set("watch", "true")
	_ = cmd.Flags().Set("watch-interval", "5s")

	cfg, err := parseWatchConfig(cmd)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !cfg.Enabled {
		t.Error("expected watch to be enabled")
	}
	if cfg.Interval != 5*time.Second {
		t.Errorf("expected interval 5s, got %v", cfg.Interval)
	}
}

func TestParseWatchConfig_NilCmd(t *testing.T) {
	_, err := parseWatchConfig(nil)
	if err == nil {
		t.Error("expected error when cmd is nil")
	}
}

func TestParseWatchConfig_InvalidInterval(t *testing.T) {
	cmd := &cobra.Command{Use: "test"}
	addWatchFlags(cmd)
	_ = cmd.Flags().Set("watch", "true")
	_ = cmd.Flags().Set("watch-interval", "0s")

	cfg, err := parseWatchConfig(cmd)
	if err != nil {
		t.Fatalf("unexpected error parsing flags: %v", err)
	}
	if err := cfg.Validate(); err == nil {
		t.Error("expected validation error for zero interval with watch enabled")
	}
}
