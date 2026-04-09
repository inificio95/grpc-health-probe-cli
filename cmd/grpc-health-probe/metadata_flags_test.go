package main

import (
	"testing"

	"github.com/spf13/cobra"
)

func TestAddMetadataFlags(t *testing.T) {
	cmd := newTestCmd()
	addMetadataFlags(cmd)

	f := cmd.Flags().Lookup("header")
	if f == nil {
		t.Fatal("expected --header flag to be registered")
	}
	if f.Shorthand != "H" {
		t.Errorf("expected shorthand H, got %q", f.Shorthand)
	}
}

func TestAddMetadataFlags_NilCmd(t *testing.T) {
	// Should not panic when cmd is nil.
	addMetadataFlags(nil)
}

func TestParseMetadataConfig_Defaults(t *testing.T) {
	cmd := newTestCmd()
	addMetadataFlags(cmd)

	cfg, err := parseMetadataConfig(cmd)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg == nil {
		t.Fatal("expected non-nil config")
	}
	if len(cfg.Entries) != 0 {
		t.Errorf("expected 0 entries, got %d", len(cfg.Entries))
	}
}

func TestParseMetadataConfig_SingleHeader(t *testing.T) {
	cmd := newTestCmd()
	addMetadataFlags(cmd)
	_ = cmd.Flags().Set("header", "x-request-id=abc123")

	cfg, err := parseMetadataConfig(cmd)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(cfg.Entries) != 1 {
		t.Fatalf("expected 1 entry, got %d", len(cfg.Entries))
	}
	if cfg.Entries[0].Key != "x-request-id" {
		t.Errorf("expected key x-request-id, got %q", cfg.Entries[0].Key)
	}
	if cfg.Entries[0].Value != "abc123" {
		t.Errorf("expected value abc123, got %q", cfg.Entries[0].Value)
	}
}

func TestParseMetadataConfig_MultipleHeaders(t *testing.T) {
	cmd := newTestCmd()
	addMetadataFlags(cmd)
	_ = cmd.Flags().Set("header", "authorization=Bearer tok")
	_ = cmd.Flags().Set("header", "x-trace-id=xyz")

	cfg, err := parseMetadataConfig(cmd)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(cfg.Entries) != 2 {
		t.Fatalf("expected 2 entries, got %d", len(cfg.Entries))
	}
}

func TestParseMetadataConfig_NilCmd(t *testing.T) {
	cfg, err := parseMetadataConfig(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg == nil {
		t.Fatal("expected non-nil default config")
	}
}

func TestParseMetadataConfig_MalformedHeader(t *testing.T) {
	cmd := &cobra.Command{Use: "test"}
	addMetadataFlags(cmd)
	// Malformed entry (no '=') should be silently skipped.
	_ = cmd.Flags().Set("header", "badvalue")

	cfg, err := parseMetadataConfig(cmd)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(cfg.Entries) != 0 {
		t.Errorf("expected 0 entries after skipping malformed header, got %d", len(cfg.Entries))
	}
}
