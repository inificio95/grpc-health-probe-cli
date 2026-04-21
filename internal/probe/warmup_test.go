package probe

import (
	"testing"
	"time"
)

func TestDefaultWarmupConfig(t *testing.T) {
	cfg := DefaultWarmupConfig()
	if cfg == nil {
		t.Fatal("expected non-nil config")
	}
	if cfg.Enabled {
		t.Error("expected Enabled to be false by default")
	}
	if cfg.Delay != 2*time.Second {
		t.Errorf("expected default Delay 2s, got %v", cfg.Delay)
	}
	if cfg.MaxDelay != 30*time.Second {
		t.Errorf("expected default MaxDelay 30s, got %v", cfg.MaxDelay)
	}
}

func TestWarmupConfig_Validate_Nil(t *testing.T) {
	var cfg *WarmupConfig
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for nil config")
	}
}

func TestWarmupConfig_Validate_Disabled(t *testing.T) {
	cfg := &WarmupConfig{Enabled: false}
	if err := cfg.Validate(); err != nil {
		t.Errorf("unexpected error for disabled config: %v", err)
	}
}

func TestWarmupConfig_Validate_EnabledValid(t *testing.T) {
	cfg := &WarmupConfig{
		Enabled:  true,
		Delay:    5 * time.Second,
		MaxDelay: 30 * time.Second,
	}
	if err := cfg.Validate(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestWarmupConfig_Validate_ZeroDelay(t *testing.T) {
	cfg := &WarmupConfig{
		Enabled:  true,
		Delay:    0,
		MaxDelay: 30 * time.Second,
	}
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for zero delay")
	}
}

func TestWarmupConfig_Validate_ZeroMaxDelay(t *testing.T) {
	cfg := &WarmupConfig{
		Enabled:  true,
		Delay:    5 * time.Second,
		MaxDelay: 0,
	}
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for zero max delay")
	}
}

func TestWarmupConfig_Validate_DelayExceedsMax(t *testing.T) {
	cfg := &WarmupConfig{
		Enabled:  true,
		Delay:    60 * time.Second,
		MaxDelay: 30 * time.Second,
	}
	if err := cfg.Validate(); err == nil {
		t.Error("expected error when delay exceeds max delay")
	}
}

func TestWarmupConfig_EffectiveDelay_Disabled(t *testing.T) {
	cfg := &WarmupConfig{Enabled: false, Delay: 5 * time.Second, MaxDelay: 30 * time.Second}
	if d := cfg.EffectiveDelay(); d != 0 {
		t.Errorf("expected 0 for disabled, got %v", d)
	}
}

func TestWarmupConfig_EffectiveDelay_Nil(t *testing.T) {
	var cfg *WarmupConfig
	if d := cfg.EffectiveDelay(); d != 0 {
		t.Errorf("expected 0 for nil config, got %v", d)
	}
}

func TestWarmupConfig_EffectiveDelay_Enabled(t *testing.T) {
	cfg := &WarmupConfig{
		Enabled:  true,
		Delay:    5 * time.Second,
		MaxDelay: 30 * time.Second,
	}
	if d := cfg.EffectiveDelay(); d != 5*time.Second {
		t.Errorf("expected 5s, got %v", d)
	}
}

func TestWarmupConfig_EffectiveDelay_ClampedToMax(t *testing.T) {
	cfg := &WarmupConfig{
		Enabled:  true,
		Delay:    60 * time.Second,
		MaxDelay: 30 * time.Second,
	}
	if d := cfg.EffectiveDelay(); d != 30*time.Second {
		t.Errorf("expected clamped value 30s, got %v", d)
	}
}
