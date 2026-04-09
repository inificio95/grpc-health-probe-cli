package probe

import (
	"testing"
)

func TestDefaultMetadataConfig(t *testing.T) {
	cfg := DefaultMetadataConfig()
	if len(cfg.Headers) != 0 {
		t.Errorf("expected empty headers, got %v", cfg.Headers)
	}
}

func TestMetadataConfig_Validate_Valid(t *testing.T) {
	cfg := MetadataConfig{
		Headers: []string{"x-request-id=abc123", "authorization=Bearer token"},
	}
	if err := cfg.Validate(); err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestMetadataConfig_Validate_MissingValue(t *testing.T) {
	cfg := MetadataConfig{
		Headers: []string{"x-request-id"},
	}
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for missing value, got nil")
	}
}

func TestMetadataConfig_Validate_EmptyKey(t *testing.T) {
	cfg := MetadataConfig{
		Headers: []string{"=value"},
	}
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for empty key, got nil")
	}
}

func TestMetadataConfig_Validate_Empty(t *testing.T) {
	cfg := MetadataConfig{Headers: []string{}}
	if err := cfg.Validate(); err != nil {
		t.Errorf("expected no error for empty headers, got %v", err)
	}
}

func TestMetadataConfig_ToGRPCMetadata(t *testing.T) {
	cfg := MetadataConfig{
		Headers: []string{"x-trace-id=xyz", "x-env=production"},
	}
	md := cfg.ToGRPCMetadata()

	if vals, ok := md["x-trace-id"]; !ok || len(vals) != 1 || vals[0] != "xyz" {
		t.Errorf("expected x-trace-id=xyz, got %v", md["x-trace-id"])
	}
	if vals, ok := md["x-env"]; !ok || len(vals) != 1 || vals[0] != "production" {
		t.Errorf("expected x-env=production, got %v", md["x-env"])
	}
}

func TestMetadataConfig_ToGRPCMetadata_DuplicateKeys(t *testing.T) {
	cfg := MetadataConfig{
		Headers: []string{"x-tag=foo", "x-tag=bar"},
	}
	md := cfg.ToGRPCMetadata()

	if len(md["x-tag"]) != 2 {
		t.Errorf("expected 2 values for x-tag, got %d", len(md["x-tag"]))
	}
}

func TestMetadataConfig_ToGRPCMetadata_Empty(t *testing.T) {
	cfg := DefaultMetadataConfig()
	md := cfg.ToGRPCMetadata()
	if len(md) != 0 {
		t.Errorf("expected empty metadata, got %v", md)
	}
}
