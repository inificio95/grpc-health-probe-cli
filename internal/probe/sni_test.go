package probe

import (
	"testing"
)

func TestDefaultSNIConfig(t *testing.T) {
	cfg := DefaultSNIConfig()
	if cfg == nil {
		t.Fatal("expected non-nil config")
	}
	if cfg.Enabled {
		t.Error("expected Enabled to be false")
	}
	if cfg.ServerName != "" {
		t.Errorf("expected empty ServerName, got %q", cfg.ServerName)
	}
}

func TestSNIConfig_Validate_Nil(t *testing.T) {
	var cfg *SNIConfig
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for nil config")
	}
}

func TestSNIConfig_Validate_Disabled(t *testing.T) {
	cfg := &SNIConfig{Enabled: false, ServerName: ""}
	if err := cfg.Validate(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestSNIConfig_Validate_EnabledMissingServerName(t *testing.T) {
	cfg := &SNIConfig{Enabled: true, ServerName: ""}
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for missing server name")
	}
}

func TestSNIConfig_Validate_EnabledValid(t *testing.T) {
	cfg := &SNIConfig{Enabled: true, ServerName: "example.com"}
	if err := cfg.Validate(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestSNIConfig_String_Nil(t *testing.T) {
	var cfg *SNIConfig
	got := cfg.String()
	if got != "SNIConfig(nil)" {
		t.Errorf("expected SNIConfig(nil), got %q", got)
	}
}

func TestSNIConfig_String_Disabled(t *testing.T) {
	cfg := DefaultSNIConfig()
	got := cfg.String()
	if got != "SNIConfig(disabled)" {
		t.Errorf("expected SNIConfig(disabled), got %q", got)
	}
}

func TestSNIConfig_String_Enabled(t *testing.T) {
	cfg := &SNIConfig{Enabled: true, ServerName: "myservice.internal"}
	got := cfg.String()
	expected := "SNIConfig(server_name=myservice.internal)"
	if got != expected {
		t.Errorf("expected %q, got %q", expected, got)
	}
}
