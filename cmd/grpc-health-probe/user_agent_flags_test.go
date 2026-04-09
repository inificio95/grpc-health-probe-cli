package main

import (
	"testing"

	"github.com/spf13/cobra"
)

func TestAddUserAgentFlags(t *testing.T) {
	cmd := &cobra.Command{Use: "test"}
	addUserAgentFlags(cmd)

	for _, flag := range []string{"user-agent-name", "user-agent-version", "no-user-agent"} {
		if cmd.Flags().Lookup(flag) == nil {
			t.Errorf("expected flag %q to be registered", flag)
		}
	}
}

func TestAddUserAgentFlags_NilCmd(t *testing.T) {
	// Should not panic
	addUserAgentFlags(nil)
}

func TestParseUserAgentConfig_Defaults(t *testing.T) {
	cmd := &cobra.Command{Use: "test"}
	addUserAgentFlags(cmd)
	_ = cmd.ParseFlags([]string{})

	cfg := parseUserAgentConfig(cmd)
	if cfg.AppName == "" {
		t.Error("expected non-empty default AppName")
	}
	if cfg.AppVersion == "" {
		t.Error("expected non-empty default AppVersion")
	}
	if cfg.Disabled {
		t.Error("expected Disabled to be false by default")
	}
}

func TestParseUserAgentConfig_CustomValues(t *testing.T) {
	cmd := &cobra.Command{Use: "test"}
	addUserAgentFlags(cmd)
	_ = cmd.ParseFlags([]string{"--user-agent-name=myapp", "--user-agent-version=2.0.0"})

	cfg := parseUserAgentConfig(cmd)
	if cfg.AppName != "myapp" {
		t.Errorf("expected AppName %q, got %q", "myapp", cfg.AppName)
	}
	if cfg.AppVersion != "2.0.0" {
		t.Errorf("expected AppVersion %q, got %q", "2.0.0", cfg.AppVersion)
	}
}

func TestParseUserAgentConfig_Disabled(t *testing.T) {
	cmd := &cobra.Command{Use: "test"}
	addUserAgentFlags(cmd)
	_ = cmd.ParseFlags([]string{"--no-user-agent"})

	cfg := parseUserAgentConfig(cmd)
	if !cfg.Disabled {
		t.Error("expected Disabled to be true")
	}
}

func TestParseUserAgentConfig_NilCmd(t *testing.T) {
	cfg := parseUserAgentConfig(nil)
	if cfg == nil {
		t.Fatal("expected non-nil config from nil cmd")
	}
	if cfg.AppName == "" {
		t.Error("expected default AppName for nil cmd")
	}
}
