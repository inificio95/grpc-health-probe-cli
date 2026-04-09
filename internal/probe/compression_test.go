package probe

import (
	"testing"

	"google.golang.org/grpc/encoding/gzip"
)

func TestDefaultCompressionConfig(t *testing.T) {
	cfg := DefaultCompressionConfig()

	if cfg.Enabled {
		t.Error("expected compression to be disabled by default")
	}

	if cfg.Type != CompressionNone {
		t.Errorf("expected type to be 'none', got %s", cfg.Type)
	}
}

func TestCompressionConfig_Validate_Disabled(t *testing.T) {
	cfg := &CompressionConfig{
		Enabled: false,
		Type:    CompressionNone,
	}

	if err := cfg.Validate(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestCompressionConfig_Validate_Nil(t *testing.T) {
	var cfg *CompressionConfig

	err := cfg.Validate()
	if err == nil {
		t.Error("expected error for nil config")
	}
}

func TestCompressionConfig_Validate_Gzip(t *testing.T) {
	cfg := &CompressionConfig{
		Enabled: true,
		Type:    CompressionGzip,
	}

	if err := cfg.Validate(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestCompressionConfig_Validate_InvalidType(t *testing.T) {
	cfg := &CompressionConfig{
		Enabled: true,
		Type:    CompressionType("invalid"),
	}

	err := cfg.Validate()
	if err == nil {
		t.Error("expected error for invalid compression type")
	}
}

func TestCompressionConfig_ApplyDialOptions_Disabled(t *testing.T) {
	cfg := &CompressionConfig{
		Enabled: false,
		Type:    CompressionNone,
	}

	opts := cfg.ApplyDialOptions()
	if len(opts) != 0 {
		t.Errorf("expected no dial options, got %d", len(opts))
	}
}

func TestCompressionConfig_ApplyDialOptions_Nil(t *testing.T) {
	var cfg *CompressionConfig

	opts := cfg.ApplyDialOptions()
	if len(opts) != 0 {
		t.Errorf("expected no dial options for nil config, got %d", len(opts))
	}
}

func TestCompressionConfig_ApplyDialOptions_Gzip(t *testing.T) {
	cfg := &CompressionConfig{
		Enabled: true,
		Type:    CompressionGzip,
	}

	opts := cfg.ApplyDialOptions()
	if len(opts) == 0 {
		t.Error("expected dial options for gzip compression")
	}

	// Verify gzip compressor name is correct
	if gzip.Name != "gzip" {
		t.Errorf("expected gzip compressor name to be 'gzip', got %s", gzip.Name)
	}
}
