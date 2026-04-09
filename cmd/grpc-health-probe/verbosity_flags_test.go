package main

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/your-org/grpc-health-probe-cli/internal/probe"
)

func TestAddVerbosityFlags(t *testing.T) {
	cmd := &cobra.Command{}
	addVerbosityFlags(cmd)

	f := cmd.Flags().Lookup("verbosity")
	if f == nil {
		t.Fatal("expected --verbosity flag to be registered")
	}
	if f.DefValue != "info" {
		t.Errorf("expected default verbosity 'info', got %q", f.DefValue)
	}

	q := cmd.Flags().Lookup("quiet")
	if q == nil {
		t.Fatal("expected --quiet flag to be registered")
	}
}

func TestAddVerbosityFlags_NilCmd(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic for nil cmd")
		}
	}()
	addVerbosityFlags(nil)
}

func TestParseVerbosityConfig_Defaults(t *testing.T) {
	cmd := &cobra.Command{}
	addVerbosityFlags(cmd)
	_ = cmd.ParseFlags([]string{})

	cfg, err := parseVerbosityConfig(cmd)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.Level != "info" {
		t.Errorf("expected level 'info', got %q", cfg.Level)
	}
	if cfg.Quiet {
		t.Error("expected quiet to be false by default")
	}
}

func TestParseVerbosityConfig_Debug(t *testing.T) {
	cmd := &cobra.Command{}
	addVerbosityFlags(cmd)
	_ = cmd.ParseFlags([]string{"--verbosity", "debug"})

	cfg, err := parseVerbosityConfig(cmd)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.Level != "debug" {
		t.Errorf("expected level 'debug', got %q", cfg.Level)
	}
}

func TestParseVerbosityConfig_Quiet(t *testing.T) {
	cmd := &cobra.Command{}
	addVerbosityFlags(cmd)
	_ = cmd.ParseFlags([]string{"--quiet"})

	cfg, err := parseVerbosityConfig(cmd)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !cfg.Quiet {
		t.Error("expected quiet to be true")
	}
}

func TestParseVerbosityConfig_Invalid(t *testing.T) {
	cmd := &cobra.Command{}
	addVerbosityFlags(cmd)
	_ = cmd.ParseFlags([]string{"--verbosity", "trace"})

	_, err := parseVerbosityConfig(cmd)
	if err == nil {
		t.Error("expected error for invalid verbosity level")
	}
}

func TestParseVerbosityConfig_IsQuiet(t *testing.T) {
	cfg := &probe.VerbosityConfig{Quiet: true, Level: "info"}
	if !cfg.IsQuiet() {
		t.Error("expected IsQuiet() to return true")
	}
}
