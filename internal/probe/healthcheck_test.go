package probe

import (
	"testing"
	"time"
)

func TestDefaultHealthCheckConfig(t *testing.T) {
	cfg := DefaultHealthCheckConfig()
	if cfg == nil {
		t.Fatal("expected non-nil config")
	}
	if !cfg.Enabled {
		t.Error("expected Enabled to be true")
	}
	if !cfg.IncludeServiceName {
		t.Error("expected IncludeServiceName to be true")
	}
	if cfg.CheckInterval != 5*time.Second {
		t.Errorf("expected CheckInterval 5s, got %v", cfg.CheckInterval)
	}
	if cfg.MaxConsecutiveFailures != 0 {
		t.Errorf("expected MaxConsecutiveFailures 0, got %d", cfg.MaxConsecutiveFailures)
	}
}

func TestHealthCheckConfig_Validate_Nil(t *testing.T) {
	var cfg *HealthCheckConfig
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for nil config")
	}
}

func TestHealthCheckConfig_Validate_Valid(t *testing.T) {
	cfg := DefaultHealthCheckConfig()
	if err := cfg.Validate(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestHealthCheckConfig_Validate_NegativeInterval(t *testing.T) {
	cfg := DefaultHealthCheckConfig()
	cfg.CheckInterval = -1 * time.Second
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for negative check_interval")
	}
}

func TestHealthCheckConfig_Validate_NegativeMaxFailures(t *testing.T) {
	cfg := DefaultHealthCheckConfig()
	cfg.MaxConsecutiveFailures = -3
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for negative max_consecutive_failures")
	}
}

func TestHealthCheckConfig_IsLimited_False(t *testing.T) {
	cfg := DefaultHealthCheckConfig()
	if cfg.IsLimited() {
		t.Error("expected IsLimited to return false when MaxConsecutiveFailures is 0")
	}
}

func TestHealthCheckConfig_IsLimited_True(t *testing.T) {
	cfg := DefaultHealthCheckConfig()
	cfg.MaxConsecutiveFailures = 5
	if !cfg.IsLimited() {
		t.Error("expected IsLimited to return true when MaxConsecutiveFailures > 0")
	}
}

func TestHealthCheckConfig_IsLimited_Nil(t *testing.T) {
	var cfg *HealthCheckConfig
	if cfg.IsLimited() {
		t.Error("expected IsLimited to return false for nil config")
	}
}
