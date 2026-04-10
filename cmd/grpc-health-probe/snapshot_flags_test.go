package main

import (
	"testing"

	"github.com/spf13/cobra"
)

func newSnapshotTestCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "test"}
	addSnapshotFlags(cmd)
	return cmd
}

func TestAddSnapshotFlags(t *testing.T) {
	cmd := newSnapshotTestCmd()
	if cmd.Flags().Lookup("snapshot") == nil {
		t.Error("expected --snapshot flag")
	}
	if cmd.Flags().Lookup("snapshot-path") == nil {
		t.Error("expected --snapshot-path flag")
	}
	if cmd.Flags().Lookup("snapshot-format") == nil {
		t.Error("expected --snapshot-format flag")
	}
}

func TestAddSnapshotFlags_NilCmd(t *testing.T) {
	// Should not panic
	addSnapshotFlags(nil)
}

func TestParseSnapshotConfig_Defaults(t *testing.T) {
	cmd := newSnapshotTestCmd()
	cfg, err := parseSnapshotConfig(cmd)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.Enabled {
		t.Error("expected Enabled=false by default")
	}
	if cfg.Format != "json" {
		t.Errorf("expected Format=json, got %q", cfg.Format)
	}
}

func TestParseSnapshotConfig_Enabled(t *testing.T) {
	cmd := newSnapshotTestCmd()
	_ = cmd.Flags().Set("snapshot", "true")
	_ = cmd.Flags().Set("snapshot-path", "/tmp/result.json")
	_ = cmd.Flags().Set("snapshot-format", "json")

	cfg, err := parseSnapshotConfig(cmd)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !cfg.Enabled {
		t.Error("expected Enabled=true")
	}
	if cfg.FilePath != "/tmp/result.json" {
		t.Errorf("expected FilePath=/tmp/result.json, got %q", cfg.FilePath)
	}
}

func TestParseSnapshotConfig_NilCmd(t *testing.T) {
	cfg, err := parseSnapshotConfig(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg == nil {
		t.Fatal("expected non-nil config")
	}
}

func TestParseSnapshotConfig_InvalidFormat(t *testing.T) {
	cmd := newSnapshotTestCmd()
	_ = cmd.Flags().Set("snapshot", "true")
	_ = cmd.Flags().Set("snapshot-path", "/tmp/result.out")
	_ = cmd.Flags().Set("snapshot-format", "csv")

	_, err := parseSnapshotConfig(cmd)
	if err == nil {
		t.Error("expected error for invalid snapshot format")
	}
}
