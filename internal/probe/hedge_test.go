package probe

import (
	"testing"
	"time"
)

func TestDefaultHedgeConfig(t *testing.T) {
	cfg := DefaultHedgeConfig()
	if cfg == nil {
		t.Fatal("expected non-nil config")
	}
	if cfg.Enabled {
		t.Error("expected Enabled to be false by default")
	}
	if cfg.Delay != 100*time.Millisecond {
		t.Errorf("expected default delay 100ms, got %s", cfg.Delay)
	}
	if cfg.MaxHedges != 1 {
		t.Errorf("expected default MaxHedges 1, got %d", cfg.MaxHedges)
	}
}

func TestHedgeConfig_Validate_Nil(t *testing.T) {
	var cfg *HedgeConfig
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for nil config")
	}
}

func TestHedgeConfig_Validate_Disabled(t *testing.T) {
	cfg := &HedgeConfig{Enabled: false}
	if err := cfg.Validate(); err != nil {
		t.Errorf("unexpected error for disabled config: %v", err)
	}
}

func TestHedgeConfig_Validate_Valid(t *testing.T) {
	cfg := &HedgeConfig{
		Enabled:   true,
		Delay:     50 * time.Millisecond,
		MaxHedges: 2,
	}
	if err := cfg.Validate(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestHedgeConfig_Validate_InvalidDelay(t *testing.T) {
	cfg := &HedgeConfig{
		Enabled:   true,
		Delay:     0,
		MaxHedges: 1,
	}
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for zero delay")
	}
}

func TestHedgeConfig_Validate_MaxHedgesTooLow(t *testing.T) {
	cfg := &HedgeConfig{
		Enabled:   true,
		Delay:     10 * time.Millisecond,
		MaxHedges: 0,
	}
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for MaxHedges < 1")
	}
}

func TestHedgeConfig_Validate_MaxHedgesTooHigh(t *testing.T) {
	cfg := &HedgeConfig{
		Enabled:   true,
		Delay:     10 * time.Millisecond,
		MaxHedges: 11,
	}
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for MaxHedges > 10")
	}
}

func TestHedgeConfig_String_Nil(t *testing.T) {
	var cfg *HedgeConfig
	if s := cfg.String(); s != "hedge(nil)" {
		t.Errorf("expected 'hedge(nil)', got %q", s)
	}
}

func TestHedgeConfig_String_Disabled(t *testing.T) {
	cfg := &HedgeConfig{Enabled: false}
	if s := cfg.String(); s != "hedge(disabled)" {
		t.Errorf("expected 'hedge(disabled)', got %q", s)
	}
}

func TestHedgeConfig_String_Enabled(t *testing.T) {
	cfg := &HedgeConfig{
		Enabled:   true,
		Delay:     200 * time.Millisecond,
		MaxHedges: 3,
	}
	expected := "hedge(delay=200ms, max=3)"
	if s := cfg.String(); s != expected {
		t.Errorf("expected %q, got %q", expected, s)
	}
}
