package main

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/your-org/grpc-health-probe-cli/internal/probe"
)

func newTestCmd() *cobra.Command {
	return &cobra.Command{
		Use: "test",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
}

func TestAddServiceFlags(t *testing.T) {
	cmd := newTestCmd()
	cfg := probe.DefaultServiceConfig()
	addServiceFlags(cmd, cfg)

	if !cmd.Flags().Lookup("service").Changed {
		// flag registered but not changed — that is expected
	}

	f := cmd.Flags().Lookup("service")
	if f == nil {
		t.Fatal("expected --service flag to be registered")
	}
	if f.DefValue != "" {
		t.Errorf("expected default value to be empty string, got %q", f.DefValue)
	}
}

func TestAddServiceFlags_NilConfig(t *testing.T) {
	cmd := newTestCmd()
	// Should not panic
	addServiceFlags(cmd, nil)

	if cmd.Flags().Lookup("service") != nil {
		t.Error("expected no --service flag when config is nil")
	}
}

func TestParseServiceConfig_Defaults(t *testing.T) {
	cmd := newTestCmd()
	cfg := probe.DefaultServiceConfig()
	addServiceFlags(cmd, cfg)

	parsed, err := parseServiceConfig(cmd)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if parsed.Service != "" {
		t.Errorf("expected empty service name, got %q", parsed.Service)
	}
	if !parsed.IsServerLevel() {
		t.Error("expected server-level check with empty service name")
	}
}

func TestParseServiceConfig_Named(t *testing.T) {
	cmd := newTestCmd()
	cfg := probe.DefaultServiceConfig()
	addServiceFlags(cmd, cfg)

	if err := cmd.Flags().Set("service", "mypackage.MyService"); err != nil {
		t.Fatalf("failed to set flag: %v", err)
	}

	parsed, err := parseServiceConfig(cmd)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if parsed.Service != "mypackage.MyService" {
		t.Errorf("expected service name %q, got %q", "mypackage.MyService", parsed.Service)
	}
	if parsed.IsServerLevel() {
		t.Error("expected named service check, not server-level")
	}
}
