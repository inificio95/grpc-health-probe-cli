package main

import (
	"testing"
	"time"

	"github.com/spf13/cobra"
)

func TestAddKeepaliveFlags(t *testing.T) {
	cmd := &cobra.Command{Use: "test"}
	addKeepaliveFlags(cmd)

	if cmd.Flags().Lookup("keepalive-time") == nil {
		t.Error("expected --keepalive-time flag to be registered")
	}
	if cmd.Flags().Lookup("keepalive-timeout") == nil {
		t.Error("expected --keepalive-timeout flag to be registered")
	}
	if cmd.Flags().Lookup("keepalive-permit-without-stream") == nil {
		t.Error("expected --keepalive-permit-without-stream flag to be registered")
	}
}

func TestAddKeepaliveFlags_NilCmd(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic for nil cmd")
		}
	}()
	addKeepaliveFlags(nil)
}

func TestParseKeepaliveConfig_Defaults(t *testing.T) {
	cmd := &cobra.Command{Use: "test"}
	addKeepaliveFlags(cmd)

	cfg, err := parseKeepaliveConfig(cmd)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.Enabled {
		t.Error("expected keepalive to be disabled by default")
	}
}

func TestParseKeepaliveConfig_Enabled(t *testing.T) {
	cmd := &cobra.Command{Use: "test"}
	addKeepaliveFlags(cmd)

	_ = cmd.Flags().Set("keepalive-time", "30s")
	_ = cmd.Flags().Set("keepalive-timeout", "10s")
	_ = cmd.Flags().Set("keepalive-permit-without-stream", "true")

	cfg, err := parseKeepaliveConfig(cmd)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !cfg.Enabled {
		t.Error("expected keepalive to be enabled when time is set")
	}
	if cfg.Time != 30*time.Second {
		t.Errorf("expected time=30s, got %s", cfg.Time)
	}
	if cfg.Timeout != 10*time.Second {
		t.Errorf("expected timeout=10s, got %s", cfg.Timeout)
	}
	if !cfg.PermitWithoutStream {
		t.Error("expected permit_without_stream=true")
	}
}

func TestParseKeepaliveConfig_NilCmd(t *testing.T) {
	_, err := parseKeepaliveConfig(nil)
	if err == nil {
		t.Error("expected error for nil cmd")
	}
}
