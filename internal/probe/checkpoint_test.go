package probe

import (
	"testing"
	"time"
)

func TestDefaultCheckpointConfig(t *testing.T) {
	cfg := DefaultCheckpointConfig()
	if cfg == nil {
		t.Fatal("expected non-nil config")
	}
	if cfg.Enabled {
		t.Error("expected Enabled to be false")
	}
	if cfg.Path != "" {
		t.Errorf("expected empty path, got %q", cfg.Path)
	}
	if cfg.Interval != 30*time.Second {
		t.Errorf("expected 30s interval, got %v", cfg.Interval)
	}
	if cfg.Format != "text" {
		t.Errorf("expected format \"text\", got %q", cfg.Format)
	}
}

func TestCheckpointConfig_Validate_Nil(t *testing.T) {
	var cfg *CheckpointConfig
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for nil config")
	}
}

func TestCheckpointConfig_Validate_Disabled(t *testing.T) {
	cfg := DefaultCheckpointConfig()
	if err := cfg.Validate(); err != nil {
		t.Errorf("unexpected error for disabled config: %v", err)
	}
}

func TestCheckpointConfig_Validate_EnabledValid(t *testing.T) {
	cfg := &CheckpointConfig{
		Enabled:  true,
		Path:     "/tmp/probe.checkpoint",
		Interval: 10 * time.Second,
		Format:   "json",
	}
	if err := cfg.Validate(); err != nil {
		t.Errorf("unexpected error for valid config: %v", err)
	}
}

func TestCheckpointConfig_Validate_MissingPath(t *testing.T) {
	cfg := &CheckpointConfig{
		Enabled:  true,
		Path:     "",
		Interval: 10 * time.Second,
		Format:   "text",
	}
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for missing path")
	}
}

func TestCheckpointConfig_Validate_ZeroInterval(t *testing.T) {
	cfg := &CheckpointConfig{
		Enabled:  true,
		Path:     "/tmp/probe.checkpoint",
		Interval: 0,
		Format:   "text",
	}
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for zero interval")
	}
}

func TestCheckpointConfig_Validate_InvalidFormat(t *testing.T) {
	cfg := &CheckpointConfig{
		Enabled:  true,
		Path:     "/tmp/probe.checkpoint",
		Interval: 5 * time.Second,
		Format:   "yaml",
	}
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for invalid format")
	}
}

func TestCheckpointConfig_IsEnabled(t *testing.T) {
	cfg := DefaultCheckpointConfig()
	if cfg.IsEnabled() {
		t.Error("expected IsEnabled to return false for default config")
	}

	cfg.Enabled = true
	if !cfg.IsEnabled() {
		t.Error("expected IsEnabled to return true after enabling")
	}

	var nilCfg *CheckpointConfig
	if nilCfg.IsEnabled() {
		t.Error("expected IsEnabled to return false for nil config")
	}
}
