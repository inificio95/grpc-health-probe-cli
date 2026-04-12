package probe

import (
	"testing"
)

func TestDefaultSanitizeConfig(t *testing.T) {
	cfg := DefaultSanitizeConfig()
	if cfg == nil {
		t.Fatal("expected non-nil config")
	}
	if !cfg.Enabled {
		t.Error("expected Enabled to be true")
	}
	if !cfg.StripScheme {
		t.Error("expected StripScheme to be true")
	}
	if cfg.DefaultPort != "443" {
		t.Errorf("expected DefaultPort '443', got '%s'", cfg.DefaultPort)
	}
}

func TestSanitizeConfig_Validate_Nil(t *testing.T) {
	var cfg *SanitizeConfig
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for nil config")
	}
}

func TestSanitizeConfig_Validate_Valid(t *testing.T) {
	cfg := DefaultSanitizeConfig()
	if err nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestSanitizeConfig_Validate_EmptyDefaultPort(t *testing.T) {
	cfg := &SanitizeConfig{Enabled: true, DefaultPort: ""}
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for empty default port when enabled")
	}
}

func TestSanitizeConfig_Sanitize_Disabled(t *testing.T) {
	cfg := &SanitizeConfig{Enabled: false}
	input := "grpc://localhost:50051"
	if got := cfg.Sanitize(input); got != input {
		t.Errorf("expected unchanged address '%s', got '%s'", input, got)
	}
}

func TestSanitizeConfig_Sanitize_StripScheme(t *testing.T) {
	cfg := &SanitizeConfig{Enabled: true, StripScheme: true, DefaultPort: "443"}
	got := cfg.Sanitize("grpc://myservice:50051")
	if got != "myservice:50051" {
		t.Errorf("expected 'myservice:50051', got '%s'", got)
	}
}

func TestSanitizeConfig_Sanitize_AppendDefaultPort(t *testing.T) {
	cfg := &SanitizeConfig{Enabled: true, StripScheme: false, DefaultPort: "443"}
	got := cfg.Sanitize("myservice.example.com")
	if got != "myservice.example.com:443" {
		t.Errorf("expected 'myservice.example.com:443', got '%s'", got)
	}
}

func TestSanitizeConfig_Sanitize_PreservesExistingPort(t *testing.T) {
	cfg := DefaultSanitizeConfig()
	got := cfg.Sanitize("localhost:9090")
	if got != "localhost:9090" {
		t.Errorf("expected 'localhost:9090', got '%s'", got)
	}
}

func TestSanitizeConfig_Sanitize_NilConfig(t *testing.T) {
	var cfg *SanitizeConfig
	input := "grpc://localhost:50051"
	if got := cfg.Sanitize(input); got != input {
		t.Errorf("expected unchanged address for nil config, got '%s'", got)
	}
}
