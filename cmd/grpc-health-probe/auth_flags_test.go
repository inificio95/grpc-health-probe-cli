package main

import (
	"testing"

	"github.com/spf13/cobra"
)

func TestAddAuthFlags(t *testing.T) {
	cmd := &cobra.Command{Use: "test"}
	addAuthFlags(cmd)

	flags := []string{"auth-type", "auth-token", "auth-username", "auth-password"}
	for _, name := range flags {
		if cmd.Flags().Lookup(name) == nil {
			t.Errorf("expected flag %q to be registered", name)
		}
	}
}

func TestAddAuthFlags_NilCmd(t *testing.T) {
	// Should not panic
	addAuthFlags(nil)
}

func TestParseAuthConfig_Defaults(t *testing.T) {
	cmd := &cobra.Command{Use: "test"}
	addAuthFlags(cmd)

	cfg, err := parseAuthConfig(cmd)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg == nil {
		t.Fatal("expected non-nil config")
	}
	if cfg.Type != "none" {
		t.Errorf("expected auth type %q, got %q", "none", cfg.Type)
	}
}

func TestParseAuthConfig_Bearer(t *testing.T) {
	cmd := &cobra.Command{Use: "test"}
	addAuthFlags(cmd)
	_ = cmd.Flags().Set("auth-type", "bearer")
	_ = cmd.Flags().Set("auth-token", "my-secret-token")

	cfg, err := parseAuthConfig(cmd)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.Type != "bearer" {
		t.Errorf("expected type %q, got %q", "bearer", cfg.Type)
	}
	if cfg.Token != "my-secret-token" {
		t.Errorf("expected token %q, got %q", "my-secret-token", cfg.Token)
	}
}

func TestParseAuthConfig_Basic(t *testing.T) {
	cmd := &cobra.Command{Use: "test"}
	addAuthFlags(cmd)
	_ = cmd.Flags().Set("auth-type", "basic")
	_ = cmd.Flags().Set("auth-username", "admin")
	_ = cmd.Flags().Set("auth-password", "secret")

	cfg, err := parseAuthConfig(cmd)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.Username != "admin" {
		t.Errorf("expected username %q, got %q", "admin", cfg.Username)
	}
	if cfg.Password != "secret" {
		t.Errorf("expected password %q, got %q", "secret", cfg.Password)
	}
}

func TestParseAuthConfig_NilCmd(t *testing.T) {
	cfg, err := parseAuthConfig(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg == nil {
		t.Fatal("expected non-nil default config")
	}
}

func TestParseAuthConfig_BearerMissingToken(t *testing.T) {
	cmd := &cobra.Command{Use: "test"}
	addAuthFlags(cmd)
	_ = cmd.Flags().Set("auth-type", "bearer")
	// token intentionally omitted

	_, err := parseAuthConfig(cmd)
	if err == nil {
		t.Fatal("expected validation error for bearer without token")
	}
}
