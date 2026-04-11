package probe

import (
	"testing"
	"time"
)

func TestDefaultJitterConfig(t *testing.T) {
	cfg := DefaultJitterConfig()
	if cfg == nil {
		t.Fatal("expected non-nil config")
	}
	if cfg.Enabled {
		t.Error("expected Enabled to be false by default")
	}
	if cfg.Factor != 0.25 {
		t.Errorf("expected Factor 0.25, got %v", cfg.Factor)
	}
}

func TestJitterConfig_Validate_Nil(t *testing.T) {
	var cfg *JitterConfig
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for nil config")
	}
}

func TestJitterConfig_Validate_Disabled(t *testing.T) {
	cfg := DefaultJitterConfig()
	if err := cfg.Validate(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestJitterConfig_Validate_EnabledValid(t *testing.T) {
	cfg := &JitterConfig{Enabled: true, Factor: 0.5}
	if err := cfg.Validate(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestJitterConfig_Validate_InvalidFactor(t *testing.T) {
	tests := []struct {
		name   string
		factor float64
	}{
		{"zero", 0.0},
		{"negative", -0.1},
		{"above one", 1.5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := &JitterConfig{Enabled: true, Factor: tt.factor}
			if err := cfg.Validate(); err == nil {
				t.Errorf("expected error for factor %v", tt.factor)
			}
		})
	}
}

func TestJitterConfig_Apply_Disabled(t *testing.T) {
	cfg := DefaultJitterConfig() // Enabled = false
	delay := 100 * time.Millisecond
	result := cfg.Apply(delay)
	if result != delay {
		t.Errorf("expected delay unchanged, got %v", result)
	}
}

func TestJitterConfig_Apply_NilConfig(t *testing.T) {
	var cfg *JitterConfig
	delay := 200 * time.Millisecond
	result := cfg.Apply(delay)
	if result != delay {
		t.Errorf("expected delay unchanged for nil config, got %v", result)
	}
}

func TestJitterConfig_Apply_Enabled(t *testing.T) {
	cfg := &JitterConfig{Enabled: true, Factor: 0.5, Seed: 42}
	delay := 100 * time.Millisecond
	result := cfg.Apply(delay)
	if result < delay {
		t.Errorf("jittered delay %v should be >= base delay %v", result, delay)
	}
	max := delay + time.Duration(float64(delay)*0.5)
	if result > max {
		t.Errorf("jittered delay %v exceeds expected max %v", result, max)
	}
}

func TestJitterConfig_Apply_ZeroDelay(t *testing.T) {
	cfg := &JitterConfig{Enabled: true, Factor: 0.5, Seed: 1}
	result := cfg.Apply(0)
	if result != 0 {
		t.Errorf("expected 0 for zero delay, got %v", result)
	}
}
