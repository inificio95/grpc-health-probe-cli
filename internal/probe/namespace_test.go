package probe

import (
	"testing"
)

func TestDefaultNamespaceConfig(t *testing.T) {
	cfg := DefaultNamespaceConfig()
	if cfg == nil {
		t.Fatal("expected non-nil config")
	}
	if cfg.Enabled {
		t.Error("expected Enabled to be false")
	}
	if cfg.Namespace != "" {
		t.Errorf("expected empty Namespace, got %q", cfg.Namespace)
	}
	if cfg.Separator != "." {
		t.Errorf("expected Separator to be '.', got %q", cfg.Separator)
	}
}

func TestNamespaceConfig_Validate_Nil(t *testing.T) {
	var cfg *NamespaceConfig
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for nil config")
	}
}

func TestNamespaceConfig_Validate_Disabled(t *testing.T) {
	cfg := DefaultNamespaceConfig()
	if err := cfg.Validate(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestNamespaceConfig_Validate_EnabledValid(t *testing.T) {
	cfg := &NamespaceConfig{Enabled: true, Namespace: "my-service", Separator: "."}
	if err := cfg.Validate(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestNamespaceConfig_Validate_EnabledMissingNamespace(t *testing.T) {
	cfg := &NamespaceConfig{Enabled: true, Namespace: "", Separator: "."}
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for missing namespace")
	}
}

func TestNamespaceConfig_Validate_InvalidCharacters(t *testing.T) {
	cfg := &NamespaceConfig{Enabled: true, Namespace: "my namespace!", Separator: "."}
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for invalid namespace characters")
	}
}

func TestNamespaceConfig_Validate_EmptySeparator(t *testing.T) {
	cfg := &NamespaceConfig{Enabled: true, Namespace: "prod", Separator: ""}
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for empty separator")
	}
}

func TestNamespaceConfig_Qualify_Disabled(t *testing.T) {
	cfg := DefaultNamespaceConfig()
	result := cfg.Qualify("health")
	if result != "health" {
		t.Errorf("expected 'health', got %q", result)
	}
}

func TestNamespaceConfig_Qualify_Enabled(t *testing.T) {
	cfg := &NamespaceConfig{Enabled: true, Namespace: "prod", Separator: "."}
	result := cfg.Qualify("health")
	if result != "prod.health" {
		t.Errorf("expected 'prod.health', got %q", result)
	}
}

func TestNamespaceConfig_Qualify_Nil(t *testing.T) {
	var cfg *NamespaceConfig
	result := cfg.Qualify("health")
	if result != "health" {
		t.Errorf("expected 'health', got %q", result)
	}
}

func TestNamespaceConfig_Qualify_CustomSeparator(t *testing.T) {
	cfg := &NamespaceConfig{Enabled: true, Namespace: "staging", Separator: "/"}
	result := cfg.Qualify("grpc")
	if result != "staging/grpc" {
		t.Errorf("expected 'staging/grpc', got %q", result)
	}
}
