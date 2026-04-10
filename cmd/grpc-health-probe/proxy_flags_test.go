package main

import (
	"testing"

	"github.com/spf13/cobra"
)

func TestAddProxyFlags(t *testing.T) {
	cmd := &cobra.Command{Use: "test"}
	addProxyFlags(cmd)

	if f := cmd.Flags().Lookup("proxy"); f == nil {
		t.Error("expected --proxy flag to be registered")
	}
	if f := cmd.Flags().Lookup("proxy-url"); f == nil {
		t.Error("expected --proxy-url flag to be registered")
	}
}

func TestAddProxyFlags_NilCmd(t *testing.T) {
	// Should not panic
	addProxyFlags(nil)
}

func TestParseProxyConfig_Defaults(t *testing.T) {
	cmd := &cobra.Command{Use: "test"}
	addProxyFlags(cmd)

	cfg := parseProxyConfig(cmd)
	if cfg == nil {
		t.Fatal("expected non-nil config")
	}
	if cfg.Enabled {
		t.Error("expected proxy disabled by default")
	}
	if cfg.ProxyURL != "" {
		t.Errorf("expected empty proxy URL by default, got %q", cfg.ProxyURL)
	}
}

func TestParseProxyConfig_Enabled(t *testing.T) {
	cmd := &cobra.Command{Use: "test"}
	addProxyFlags(cmd)
	_ = cmd.Flags().Set("proxy", "true")
	_ = cmd.Flags().Set("proxy-url", "http://proxy.example.com:8080")

	cfg := parseProxyConfig(cmd)
	if !cfg.Enabled {
		t.Error("expected proxy to be enabled")
	}
	if cfg.ProxyURL != "http://proxy.example.com:8080" {
		t.Errorf("unexpected proxy URL: %q", cfg.ProxyURL)
	}
}

func TestParseProxyConfig_NilCmd(t *testing.T) {
	cfg := parseProxyConfig(nil)
	if cfg == nil {
		t.Fatal("expected non-nil config for nil cmd")
	}
	if cfg.Enabled {
		t.Error("expected proxy disabled for nil cmd")
	}
}
