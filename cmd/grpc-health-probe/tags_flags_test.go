package main

import (
	"testing"

	"github.com/spf13/cobra"
)

func newTagsTestCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "test"}
	addTagsFlags(cmd)
	return cmd
}

func TestAddTagsFlags(t *testing.T) {
	cmd := newTagsTestCmd()
	if cmd.Flags().Lookup("tag") == nil {
		t.Error("expected --tag flag to be registered")
	}
}

func TestAddTagsFlags_NilCmd(t *testing.T) {
	// Should not panic.
	addTagsFlags(nil)
}

func TestParseTagsConfig_Defaults(t *testing.T) {
	cmd := newTagsTestCmd()
	cfg := parseTagsConfig(cmd)
	if cfg == nil {
		t.Fatal("expected non-nil config")
	}
	if cfg.Enabled {
		t.Error("expected Enabled false when no tags provided")
	}
}

func TestParseTagsConfig_SingleTag(t *testing.T) {
	cmd := newTagsTestCmd()
	_ = cmd.Flags().Set("tag", "env=prod")
	cfg := parseTagsConfig(cmd)
	if !cfg.Enabled {
		t.Error("expected Enabled true when tag provided")
	}
	if v, ok := cfg.Tags["env"]; !ok || v != "prod" {
		t.Errorf("expected tag env=prod, got %v", cfg.Tags)
	}
}

func TestParseTagsConfig_MultipleTags(t *testing.T) {
	cmd := newTagsTestCmd()
	_ = cmd.Flags().Set("tag", "env=staging")
	_ = cmd.Flags().Set("tag", "region=eu-west-1")
	cfg := parseTagsConfig(cmd)
	if len(cfg.Tags) != 2 {
		t.Errorf("expected 2 tags, got %d", len(cfg.Tags))
	}
}

func TestParseTagsConfig_NilCmd(t *testing.T) {
	cfg := parseTagsConfig(nil)
	if cfg == nil {
		t.Fatal("expected non-nil config for nil cmd")
	}
	if cfg.Enabled {
		t.Error("expected Enabled false for nil cmd")
	}
}

func TestParseTagsConfig_MalformedEntry(t *testing.T) {
	cmd := newTagsTestCmd()
	// Malformed entry without '=' should be skipped.
	_ = cmd.Flags().Set("tag", "badentry")
	cfg := parseTagsConfig(cmd)
	if cfg.Enabled {
		t.Error("expected Enabled false when only malformed tags provided")
	}
}
