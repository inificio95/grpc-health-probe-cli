package probe

import (
	"testing"
)

func TestDefaultLabelsConfig(t *testing.T) {
	cfg := DefaultLabelsConfig()
	if cfg == nil {
		t.Fatal("expected non-nil config")
	}
	if cfg.Enabled {
		t.Error("expected Enabled to be false")
	}
	if cfg.Labels == nil {
		t.Error("expected Labels map to be initialized")
	}
}

func TestLabelsConfig_Validate_Nil(t *testing.T) {
	var cfg *LabelsConfig
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for nil config")
	}
}

func TestLabelsConfig_Validate_Disabled(t *testing.T) {
	cfg := DefaultLabelsConfig()
	if err := cfg.Validate(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestLabelsConfig_Validate_Valid(t *testing.T) {
	cfg := &LabelsConfig{
		Enabled: true,
		Labels:  map[string]string{"env": "prod", "region": "us-east-1"},
	}
	if err := cfg.Validate(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestLabelsConfig_Validate_EmptyKey(t *testing.T) {
	cfg := &LabelsConfig{
		Enabled: true,
		Labels:  map[string]string{"": "value"},
	}
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for empty key")
	}
}

func TestLabelsConfig_Validate_EmptyValue(t *testing.T) {
	cfg := &LabelsConfig{
		Enabled: true,
		Labels:  map[string]string{"env": ""},
	}
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for empty value")
	}
}

func TestLabelsConfig_AsMap_Disabled(t *testing.T) {
	cfg := DefaultLabelsConfig()
	m := cfg.AsMap()
	if len(m) != 0 {
		t.Errorf("expected empty map when disabled, got %v", m)
	}
}

func TestLabelsConfig_AsMap_Enabled(t *testing.T) {
	cfg := &LabelsConfig{
		Enabled: true,
		Labels:  map[string]string{"env": "staging"},
	}
	m := cfg.AsMap()
	if v, ok := m["env"]; !ok || v != "staging" {
		t.Errorf("expected env=staging in map, got %v", m)
	}
}

func TestLabelsConfig_AsMap_Nil(t *testing.T) {
	var cfg *LabelsConfig
	m := cfg.AsMap()
	if m == nil {
		t.Error("expected non-nil map from nil config")
	}
	if len(m) != 0 {
		t.Errorf("expected empty map, got %v", m)
	}
}
