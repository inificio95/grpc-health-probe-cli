package main

import (
	"testing"

	"github.com/spf13/cobra"
)

func TestAddOutputFlags(t *testing.T) {
	cmd := &cobra.Command{Use: "test"}
	if err := addOutputFlags(cmd); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if cmd.Flags().Lookup("format") == nil {
		t.Error("expected --format flag to be registered")
	}
	if cmd.Flags().Lookup("no-color") == nil {
		t.Error("expected --no-color flag to be registered")
	}
}

func TestAddOutputFlags_NilCmd(t *testing.T) {
	if err := addOutputFlags(nil); err == nil {
		t.Error("expected error for nil cmd, got nil")
	}
}

func TestParseOutputConfig_Defaults(t *testing.T) {
	cmd := &cobra.Command{Use: "test"}
	if err := addOutputFlags(cmd); err != nil {
		t.Fatalf("setup error: %v", err)
	}

	cfg := parseOutputConfig(cmd)
	if cfg == nil {
		t.Fatal("expected non-nil config")
	}
	if cfg.Format != "text" {
		t.Errorf("expected default format \"text\", got %q", cfg.Format)
	}
	if cfg.NoColor {
		t.Error("expected NoColor to be false by default")
	}
}

func TestParseOutputConfig_JSONFormat(t *testing.T) {
	cmd := &cobra.Command{Use: "test"}
	if err := addOutputFlags(cmd); err != nil {
		t.Fatalf("setup error: %v", err)
	}
	if err := cmd.Flags().Set("format", "json"); err != nil {
		t.Fatalf("failed to set flag: %v", err)
	}

	cfg := parseOutputConfig(cmd)
	if cfg.Format != "json" {
		t.Errorf("expected format \"json\", got %q", cfg.Format)
	}
}

func TestParseOutputConfig_NoColor(t *testing.T) {
	cmd := &cobra.Command{Use: "test"}
	if err := addOutputFlags(cmd); err != nil {
		t.Fatalf("setup error: %v", err)
	}
	if err := cmd.Flags().Set("no-color", "true"); err != nil {
		t.Fatalf("failed to set flag: %v", err)
	}

	cfg := parseOutputConfig(cmd)
	if !cfg.NoColor {
		t.Error("expected NoColor to be true")
	}
}

func TestParseOutputConfig_NilCmd(t *testing.T) {
	cfg := parseOutputConfig(nil)
	if cfg == nil {
		t.Fatal("expected non-nil config even for nil cmd")
	}
	if cfg.Format != "text" {
		t.Errorf("expected default format \"text\", got %q", cfg.Format)
	}
}
