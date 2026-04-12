package probe

import (
	"testing"
)

func TestDefaultTagsConfig(t *testing.T) {
	cfg := DefaultTagsConfig()
	if cfg == nil {
		t.Fatal("expected non-nil config")
	}
	if cfg.Enabled {
		t.Error("expected Enabled to be false by default")
	}
	if cfg.Tags == nil {
		t.Error("expected Tags map to be initialised")
	}
}

func TestTagsConfig_Validate_Nil(t *testing.T) {
	var cfg *TagsConfig
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for nil config")
	}
}

func TestTagsConfig_Validate_Disabled(t *testing.T) {
	cfg := DefaultTagsConfig()
	if err := cfg.Validate(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestTagsConfig_Validate_Valid(t *testing.T) {
	cfg := &TagsConfig{
		Enabled: true,
		Tags:    map[string]string{"env": "prod", "region": "us-east-1"},
	}
	if err := cfg.Validate(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestTagsConfig_Validate_EmptyKey(t *testing.T) {
	cfg := &TagsConfig{
		Enabled: true,
		Tags:    map[string]string{"  ": "value"},
	}
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for empty tag key")
	}
}

func TestTagsConfig_Validate_EmptyValue(t *testing.T) {
	cfg := &TagsConfig{
		Enabled: true,
		Tags:    map[string]string{"key": "   "},
	}
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for empty tag value")
	}
}

func TestTagsConfig_AsLabels_Disabled(t *testing.T) {
	cfg := DefaultTagsConfig()
	if labels := cfg.AsLabels(); labels != nil {
		t.Errorf("expected nil labels when disabled, got %v", labels)
	}
}

func TestTagsConfig_AsLabels_Enabled(t *testing.T) {
	cfg := &TagsConfig{
		Enabled: true,
		Tags:    map[string]string{"env": "staging"},
	}
	labels := cfg.AsLabels()
	if len(labels) != 1 {
		t.Fatalf("expected 1 label, got %d", len(labels))
	}
	if labels[0] != "env=staging" {
		t.Errorf("unexpected label: %s", labels[0])
	}
}

func TestTagsConfig_AsLabels_Nil(t *testing.T) {
	var cfg *TagsConfig
	if labels := cfg.AsLabels(); labels != nil {
		t.Errorf("expected nil for nil config, got %v", labels)
	}
}
