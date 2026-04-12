package main

import (
	"testing"

	"github.com/spf13/cobra"
)

func newExitCodeTestCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "test"}
	addExitCodeFlags(cmd)
	return cmd
}

func TestAddExitCodeFlags(t *testing.T) {
	cmd := newExitCodeTestCmd()
	for _, flag := range []string{
		"exit-code-enabled",
		"exit-code-not-serving",
		"exit-code-unknown",
		"exit-code-timeout",
	} {
		if cmd.Flags().Lookup(flag) == nil {
			t.Errorf("expected flag %q to be registered", flag)
		}
	}
}

func TestAddExitCodeFlags_NilCmd(t *testing.T) {
	// Should not panic
	addExitCodeFlags(nil)
}

func TestParseExitCodeConfig_Defaults(t *testing.T) {
	cmd := newExitCodeTestCmd()
	cfg, err := parseExitCodeConfig(cmd)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !cfg.Enabled {
		t.Error("expected Enabled true by default")
	}
	if cfg.NotServingCode != 1 {
		t.Errorf("expected NotServingCode 1, got %d", cfg.NotServingCode)
	}
	if cfg.UnknownCode != 2 {
		t.Errorf("expected UnknownCode 2, got %d", cfg.UnknownCode)
	}
	if cfg.TimeoutCode != 4 {
		t.Errorf("expected TimeoutCode 4, got %d", cfg.TimeoutCode)
	}
}

func TestParseExitCodeConfig_CustomValues(t *testing.T) {
	cmd := newExitCodeTestCmd()
	_ = cmd.Flags().Set("exit-code-not-serving", "10")
	_ = cmd.Flags().Set("exit-code-unknown", "11")
	_ = cmd.Flags().Set("exit-code-timeout", "12")
	cfg, err := parseExitCodeConfig(cmd)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.NotServingCode != 10 {
		t.Errorf("expected NotServingCode 10, got %d", cfg.NotServingCode)
	}
	if cfg.UnknownCode != 11 {
		t.Errorf("expected UnknownCode 11, got %d", cfg.UnknownCode)
	}
	if cfg.TimeoutCode != 12 {
		t.Errorf("expected TimeoutCode 12, got %d", cfg.TimeoutCode)
	}
}

func TestParseExitCodeConfig_Disabled(t *testing.T) {
	cmd := newExitCodeTestCmd()
	_ = cmd.Flags().Set("exit-code-enabled", "false")
	cfg, err := parseExitCodeConfig(cmd)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.Enabled {
		t.Error("expected Enabled false")
	}
}

func TestParseExitCodeConfig_NilCmd(t *testing.T) {
	cfg, err := parseExitCodeConfig(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg == nil {
		t.Fatal("expected non-nil config")
	}
}
