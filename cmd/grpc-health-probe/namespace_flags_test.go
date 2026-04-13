package main

import (
	"testing"

	"github.com/spf13/cobra"
)

func newNamespaceTestCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "test"}
	addNamespaceFlags(cmd)
	return cmd
}

func TestAddNamespaceFlags(t *testing.T) {
	cmd := newNamespaceTestCmd()
	if cmd.Flags().Lookup("namespace") == nil {
		t.Error("expected --namespace flag to be registered")
	}
	if cmd.Flags().Lookup("namespace-enabled") == nil {
		t.Error("expected --namespace-enabled flag to be registered")
	}
}

func TestAddNamespaceFlags_NilCmd(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic for nil cmd")
		}
	}()
	addNamespaceFlags(nil)
}

func TestParseNamespaceConfig_Defaults(t *testing.T) {
	cmd := newNamespaceTestCmd()
	cfg := parseNamespaceConfig(cmd)
	if cfg == nil {
		t.Fatal("expected non-nil config")
	}
	if cfg.Enabled {
		t.Error("expected Enabled=false by default")
	}
	if cfg.Namespace != "" {
		t.Errorf("expected empty Namespace, got %q", cfg.Namespace)
	}
}

func TestParseNamespaceConfig_Enabled(t *testing.T) {
	cmd := newNamespaceTestCmd()
	_ = cmd.Flags().Set("namespace-enabled", "true")
	_ = cmd.Flags().Set("namespace", "production")
	cfg := parseNamespaceConfig(cmd)
	if !cfg.Enabled {
		t.Error("expected Enabled=true")
	}
	if cfg.Namespace != "production" {
		t.Errorf("expected Namespace='production', got %q", cfg.Namespace)
	}
}

func TestParseNamespaceConfig_NilCmd(t *testing.T) {
	cfg := parseNamespaceConfig(nil)
	if cfg == nil {
		t.Fatal("expected non-nil default config for nil cmd")
	}
	if cfg.Enabled {
		t.Error("expected Enabled=false for nil cmd")
	}
}
