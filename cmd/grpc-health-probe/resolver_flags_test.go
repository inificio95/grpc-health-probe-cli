package main

import (
	"testing"

	"github.com/spf13/cobra"
)

func TestAddResolverFlags(t *testing.T) {
	cmd := &cobra.Command{Use: "test"}
	addResolverFlags(cmd)

	flags := []string{"resolve", "resolve-prefer-ipv6", "resolve-dns-server"}
	for _, name := range flags {
		if cmd.Flags().Lookup(name) == nil {
			t.Errorf("expected flag %q to be registered", name)
		}
	}
}

func TestAddResolverFlags_NilCmd(t *testing.T) {
	// Should not panic.
	addResolverFlags(nil)
}

func TestParseResolverConfig_Defaults(t *testing.T) {
	cmd := &cobra.Command{Use: "test"}
	addResolverFlags(cmd)

	cfg := parseResolverConfig(cmd)
	if cfg == nil {
		t.Fatal("expected non-nil config")
	}
	if cfg.Enabled {
		t.Error("expected Enabled=false by default")
	}
	if cfg.PreferIPv6 {
		t.Error("expected PreferIPv6=false by default")
	}
	if cfg.CustomResolver != "" {
		t.Errorf("expected empty CustomResolver, got %q", cfg.CustomResolver)
	}
}

func TestParseResolverConfig_Enabled(t *testing.T) {
	cmd := &cobra.Command{Use: "test"}
	addResolverFlags(cmd)
	_ = cmd.Flags().Set("resolve", "true")
	_ = cmd.Flags().Set("resolve-prefer-ipv6", "true")
	_ = cmd.Flags().Set("resolve-dns-server", "1.1.1.1:53")

	cfg := parseResolverConfig(cmd)
	if !cfg.Enabled {
		t.Error("expected Enabled=true")
	}
	if !cfg.PreferIPv6 {
		t.Error("expected PreferIPv6=true")
	}
	if cfg.CustomResolver != "1.1.1.1:53" {
		t.Errorf("expected CustomResolver=1.1.1.1:53, got %q", cfg.CustomResolver)
	}
}

func TestParseResolverConfig_NilCmd(t *testing.T) {
	cfg := parseResolverConfig(nil)
	if cfg == nil {
		t.Fatal("expected non-nil config for nil cmd")
	}
	if cfg.Enabled {
		t.Error("expected Enabled=false for nil cmd")
	}
}
