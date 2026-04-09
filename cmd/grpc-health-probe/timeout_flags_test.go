package main

import (
	"testing"
	"time"

	"github.com/spf13/cobra"
	"grpc-health-probe-cli/internal/probe"
)

func TestAddTimeoutFlags(t *testing.T) {
	cmd := &cobra.Command{
		Use: "test",
	}

	config := probe.DefaultTimeoutConfig()
	addTimeoutFlags(cmd, config)

	// Verify flags are registered
	if cmd.Flags().Lookup("timeout") == nil {
		t.Error("timeout flag not registered")
	}
	if cmd.Flags().Lookup("dial-timeout") == nil {
		t.Error("dial-timeout flag not registered")
	}
	if cmd.Flags().Lookup("connect-timeout") == nil {
		t.Error("connect-timeout flag not registered")
	}
}

func TestAddTimeoutFlags_NilConfig(t *testing.T) {
	cmd := &cobra.Command{
		Use: "test",
	}

	// Should not panic with nil config
	addTimeoutFlags(cmd, nil)

	if cmd.Flags().Lookup("timeout") != nil {
		t.Error("flags should not be added with nil config")
	}
}

func TestParseTimeoutConfig(t *testing.T) {
	cmd := &cobra.Command{
		Use: "test",
	}

	config := probe.DefaultTimeoutConfig()
	addTimeoutFlags(cmd, config)

	// Set custom values
	cmd.Flags().Set("timeout", "10s")
	cmd.Flags().Set("dial-timeout", "3s")
	cmd.Flags().Set("connect-timeout", "7s")

	parsed, err := parseTimeoutConfig(cmd)
	if err != nil {
		t.Fatalf("parseTimeoutConfig failed: %v", err)
	}

	if parsed.RequestTimeout != 10*time.Second {
		t.Errorf("expected RequestTimeout 10s, got %v", parsed.RequestTimeout)
	}
	if parsed.DialTimeout != 3*time.Second {
		t.Errorf("expected DialTimeout 3s, got %v", parsed.DialTimeout)
	}
	if parsed.ConnectTimeout != 7*time.Second {
		t.Errorf("expected ConnectTimeout 7s, got %v", parsed.ConnectTimeout)
	}
}

func TestParseTimeoutConfig_Defaults(t *testing.T) {
	cmd := &cobra.Command{
		Use: "test",
	}

	config := probe.DefaultTimeoutConfig()
	addTimeoutFlags(cmd, config)

	parsed, err := parseTimeoutConfig(cmd)
	if err != nil {
		t.Fatalf("parseTimeoutConfig failed: %v", err)
	}

	defaults := probe.DefaultTimeoutConfig()
	if parsed.RequestTimeout != defaults.RequestTimeout {
		t.Errorf("expected default RequestTimeout %v, got %v", defaults.RequestTimeout, parsed.RequestTimeout)
	}
	if parsed.DialTimeout != defaults.DialTimeout {
		t.Errorf("expected default DialTimeout %v, got %v", defaults.DialTimeout, parsed.DialTimeout)
	}
	if parsed.ConnectTimeout != defaults.ConnectTimeout {
		t.Errorf("expected default ConnectTimeout %v, got %v", defaults.ConnectTimeout, parsed.ConnectTimeout)
	}
}
