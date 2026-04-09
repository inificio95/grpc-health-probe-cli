package probe

import (
	"testing"
	"time"
)

func TestDefaultWatchConfig(t *testing.T) {
	cfg := DefaultWatchConfig()
	if cfg == nil {
		t.Fatal("expected non-nil default watch config")
	}
	if cfg.Mode != WatchModeDisabled {
		t.Errorf("expected mode Disabled, got %v", cfg.Mode)
	}
	if cfg.Interval != 5*time.Second {
		t.Errorf("expected interval 5s, got %v", cfg.Interval)
	}
	if cfg.MaxChecks != 0 {
		t.Errorf("expected max checks 0, got %d", cfg.MaxChecks)
	}
}

func TestWatchConfig_Validate_Nil(t *testing.T) {
	var cfg *WatchConfig
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for nil config")
	}
}

func TestWatchConfig_Validate_Disabled(t *testing.T) {
	cfg := DefaultWatchConfig()
	if err := cfg.Validate(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestWatchConfig_Validate_PollingValid(t *testing.T) {
	cfg := &WatchConfig{
		Mode:      WatchModePolling,
		Interval:  2 * time.Second,
		MaxChecks: 10,
	}
	if err := cfg.Validate(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestWatchConfig_Validate_PollingZeroInterval(t *testing.T) {
	cfg := &WatchConfig{
		Mode:     WatchModePolling,
		Interval: 0,
	}
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for zero interval in polling mode")
	}
}

func TestWatchConfig_Validate_NegativeMaxChecks(t *testing.T) {
	cfg := &WatchConfig{
		Mode:      WatchModePolling,
		Interval:  time.Second,
		MaxChecks: -1,
	}
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for negative max checks")
	}
}

func TestWatchConfig_Validate_UnknownMode(t *testing.T) {
	cfg := &WatchConfig{
		Mode:     WatchMode(99),
		Interval: time.Second,
	}
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for unknown watch mode")
	}
}

func TestWatchConfig_IsEnabled(t *testing.T) {
	disabled := DefaultWatchConfig()
	if disabled.IsEnabled() {
		t.Error("expected IsEnabled false for disabled mode")
	}

	polling := &WatchConfig{Mode: WatchModePolling, Interval: time.Second}
	if !polling.IsEnabled() {
		t.Error("expected IsEnabled true for polling mode")
	}

	var nilCfg *WatchConfig
	if nilCfg.IsEnabled() {
		t.Error("expected IsEnabled false for nil config")
	}
}
