package main

import (
	"testing"

	"github.com/spf13/cobra"
)

func TestAddCompressionFlags(t *testing.T) {
	cmd := &cobra.Command{}
	addCompressionFlags(cmd)

	f := cmd.Flags().Lookup("compression")
	if f == nil {
		t.Fatal("expected --compression flag to be registered")
	}
	if f.DefValue != "none" {
		t.Errorf("expected default compression 'none', got %q", f.DefValue)
	}
}

func TestAddCompressionFlags_NilCmd(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic for nil cmd")
		}
	}()
	addCompressionFlags(nil)
}

func TestParseCompressionConfig_Defaults(t *testing.T) {
	cmd := &cobra.Command{}
	addCompressionFlags(cmd)
	_ = cmd.ParseFlags([]string{})

	cfg, err := parseCompressionConfig(cmd)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.Type != "none" {
		t.Errorf("expected type 'none', got %q", cfg.Type)
	}
	if cfg.Enabled {
		t.Error("expected Enabled to be false by default")
	}
}

func TestParseCompressionConfig_Gzip(t *testing.T) {
	cmd := &cobra.Command{}
	addCompressionFlags(cmd)
	_ = cmd.ParseFlags([]string{"--compression", "gzip"})

	cfg, err := parseCompressionConfig(cmd)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.Type != "gzip" {
		t.Errorf("expected type 'gzip', got %q", cfg.Type)
	}
	if !cfg.Enabled {
		t.Error("expected Enabled to be true for gzip")
	}
}

func TestParseCompressionConfig_Invalid(t *testing.T) {
	cmd := &cobra.Command{}
	addCompressionFlags(cmd)
	_ = cmd.ParseFlags([]string{"--compression", "zstd"})

	_, err := parseCompressionConfig(cmd)
	if err == nil {
		t.Error("expected error for unsupported compression type")
	}
}
