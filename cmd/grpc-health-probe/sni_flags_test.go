package main

import (
	"testing"

	"github.com/spf13/cobra"
)

func newSNITestCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "test"}
	addSNIFlags(cmd)
	return cmd
}

func TestAddSNIFlags(t *testing.T) {
	cmd := newSNITestCmd()
	if cmd.Flags().Lookup("sni") == nil {
		t.Error("expected --sni flag to be registered")
	}
	if cmd.Flags().Lookup("sni-server-name") == nil {
		t.Error("expected --sni-server-name flag to be registered")
	}
}

func TestAddSNIFlags_NilCmd(t *testing.T) {
	// should not panic
	addSNIFlags(nil)
}

func TestParseSNIConfig_Defaults(t *testing.T) {
	cmd := newSNITestCmd()
	cfg := parseSNIConfig(cmd)
	if cfg == nil {
		t.Fatal("expected non-nil config")
	}
	if cfg.Enabled {
		t.Error("expected Enabled to be false by default")
	}
	if cfg.ServerName != "" {
		t.Errorf("expected empty ServerName by default, got %q", cfg.ServerName)
	}
}

func TestParseSNIConfig_Enabled(t *testing.T) {
	cmd := newSNITestCmd()
	_ = cmd.Flags().Set("sni", "true")
	_ = cmd.Flags().Set("sni-server-name", "override.example.com")

	cfg := parseSNIConfig(cmd)
	if !cfg.Enabled {
		t.Error("expected Enabled to be true")
	}
	if cfg.ServerName != "override.example.com" {
		t.Errorf("expected ServerName %q, got %q", "override.example.com", cfg.ServerName)
	}
}

func TestParseSNIConfig_NilCmd(t *testing.T) {
	cfg := parseSNIConfig(nil)
	if cfg == nil {
		t.Fatal("expected non-nil config for nil cmd")
	}
	if cfg.Enabled {
		t.Error("expected Enabled to be false for nil cmd")
	}
}
