package main

import (
	"testing"

	"github.com/spf13/cobra"
)

func TestAddTLSFlags(t *testing.T) {
	cmd := &cobra.Command{Use: "test"}
	addTLSFlags(cmd)

	expected := []string{
		"tls",
		"tls-insecure-skip-verify",
		"tls-ca-cert",
		"tls-client-cert",
		"tls-client-key",
		"tls-server-name",
	}
	for _, name := range expected {
		if cmd.Flags().Lookup(name) == nil {
			t.Errorf("expected flag %q to be registered", name)
		}
	}
}

func TestAddTLSFlags_NilCmd(t *testing.T) {
	// Should not panic
	addTLSFlags(nil)
}

func TestParseTLSConfig_Defaults(t *testing.T) {
	cmd := &cobra.Command{Use: "test"}
	addTLSFlags(cmd)

	cfg, err := parseTLSConfig(cmd)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.Enabled {
		t.Error("expected TLS to be disabled by default")
	}
	if cfg.InsecureSkipVerify {
		t.Error("expected InsecureSkipVerify to be false by default")
	}
	if cfg.CACertFile != "" || cfg.ClientCertFile != "" || cfg.ClientKeyFile != "" || cfg.ServerName != "" {
		t.Error("expected all string fields to be empty by default")
	}
}

func TestParseTLSConfig_WithValues(t *testing.T) {
	cmd := &cobra.Command{Use: "test"}
	addTLSFlags(cmd)
	_ = cmd.Flags().Set("tls", "true")
	_ = cmd.Flags().Set("tls-insecure-skip-verify", "true")
	_ = cmd.Flags().Set("tls-ca-cert", "/path/to/ca.crt")
	_ = cmd.Flags().Set("tls-client-cert", "/path/to/client.crt")
	_ = cmd.Flags().Set("tls-client-key", "/path/to/client.key")
	_ = cmd.Flags().Set("tls-server-name", "example.com")

	cfg, err := parseTLSConfig(cmd)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !cfg.Enabled {
		t.Error("expected TLS to be enabled")
	}
	if !cfg.InsecureSkipVerify {
		t.Error("expected InsecureSkipVerify to be true")
	}
	if cfg.CACertFile != "/path/to/ca.crt" {
		t.Errorf("unexpected CACertFile: %q", cfg.CACertFile)
	}
	if cfg.ClientCertFile != "/path/to/client.crt" {
		t.Errorf("unexpected ClientCertFile: %q", cfg.ClientCertFile)
	}
	if cfg.ClientKeyFile != "/path/to/client.key" {
		t.Errorf("unexpected ClientKeyFile: %q", cfg.ClientKeyFile)
	}
	if cfg.ServerName != "example.com" {
		t.Errorf("unexpected ServerName: %q", cfg.ServerName)
	}
}

func TestParseTLSConfig_NilCmd(t *testing.T) {
	cfg, err := parseTLSConfig(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg == nil {
		t.Fatal("expected non-nil default TLS config")
	}
}
