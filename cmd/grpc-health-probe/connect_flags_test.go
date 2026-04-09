package main

import (
	"testing"

	"github.com/spf13/cobra"
)

func TestAddConnectFlags(t *testing.T) {
	cmd := &cobra.Command{Use: "test"}
	addConnectFlags(cmd)

	flags := []string{"user-agent", "connect-block", "connect-fail-fast"}
	for _, name := range flags {
		if cmd.Flags().Lookup(name) == nil {
			t.Errorf("expected flag %q to be registered", name)
		}
	}
}

func TestAddConnectFlags_NilCmd(t *testing.T) {
	// Should not panic
	addConnectFlags(nil)
}

func TestParseConnectConfig_Defaults(t *testing.T) {
	cmd := &cobra.Command{Use: "test"}
	addConnectFlags(cmd)

	cfg, err := parseConnectConfig(cmd)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.UserAgent == "" {
		t.Error("expected non-empty default UserAgent")
	}
	if cfg.Block {
		t.Error("expected Block to default to false")
	}
	if cfg.FailOnNonTempDialError {
		t.Error("expected FailOnNonTempDialError to default to false")
	}
}

func TestParseConnectConfig_WithValues(t *testing.T) {
	cmd := &cobra.Command{Use: "test"}
	addConnectFlags(cmd)
	_ = cmd.Flags().Set("user-agent", "my-custom-agent/2.0")
	_ = cmd.Flags().Set("connect-block", "true")
	_ = cmd.Flags().Set("connect-fail-fast", "true")

	cfg, err := parseConnectConfig(cmd)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.UserAgent != "my-custom-agent/2.0" {
		t.Errorf("expected UserAgent %q, got %q", "my-custom-agent/2.0", cfg.UserAgent)
	}
	if !cfg.Block {
		t.Error("expected Block to be true")
	}
	if !cfg.FailOnNonTempDialError {
		t.Error("expected FailOnNonTempDialError to be true")
	}
}

func TestParseConnectConfig_NilCmd(t *testing.T) {
	cfg, err := parseConnectConfig(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg == nil {
		t.Fatal("expected non-nil config for nil cmd")
	}
}
