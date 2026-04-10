package probe

import (
	"testing"
	"time"
)

func TestDefaultRateLimitConfig(t *testing.T) {
	cfg := DefaultRateLimitConfig()
	if cfg == nil {
		t.Fatal("expected non-nil config")
	}
	if cfg.Enabled {
		t.Error("expected Enabled to be false")
	}
	if cfg.MaxRequests != 10 {
		t.Errorf("expected MaxRequests=10, got %d", cfg.MaxRequests)
	}
	if cfg.WindowSize != time.Second {
		t.Errorf("expected WindowSize=1s, got %v", cfg.WindowSize)
	}
}

func TestRateLimitConfig_Validate_Nil(t *testing.T) {
	var cfg *RateLimitConfig
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for nil config")
	}
}

func TestRateLimitConfig_Validate_Disabled(t *testing.T) {
	cfg := &RateLimitConfig{Enabled: false}
	if err := cfg.Validate(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestRateLimitConfig_Validate_InvalidMaxRequests(t *testing.T) {
	cfg := &RateLimitConfig{Enabled: true, MaxRequests: 0, WindowSize: time.Second}
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for zero MaxRequests")
	}
}

func TestRateLimitConfig_Validate_InvalidWindowSize(t *testing.T) {
	cfg := &RateLimitConfig{Enabled: true, MaxRequests: 5, WindowSize: 0}
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for zero WindowSize")
	}
}

func TestRateLimitConfig_Validate_Valid(t *testing.T) {
	cfg := &RateLimitConfig{Enabled: true, MaxRequests: 5, WindowSize: time.Second}
	if err := cfg.Validate(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestRateLimitConfig_Allow_Disabled(t *testing.T) {
	cfg := DefaultRateLimitConfig()
	for i := 0; i < 100; i++ {
		if !cfg.Allow() {
			t.Error("expected Allow=true when disabled")
		}
	}
}

func TestRateLimitConfig_Allow_WithinLimit(t *testing.T) {
	cfg := &RateLimitConfig{Enabled: true, MaxRequests: 3, WindowSize: time.Second}
	for i := 0; i < 3; i++ {
		if !cfg.Allow() {
			t.Errorf("expected Allow=true on request %d", i+1)
		}
	}
}

func TestRateLimitConfig_Allow_ExceedsLimit(t *testing.T) {
	cfg := &RateLimitConfig{Enabled: true, MaxRequests: 2, WindowSize: time.Second}
	cfg.Allow()
	cfg.Allow()
	if cfg.Allow() {
		t.Error("expected Allow=false after exceeding limit")
	}
}

func TestRateLimitConfig_Remaining_Disabled(t *testing.T) {
	cfg := DefaultRateLimitConfig()
	if r := cfg.Remaining(); r != -1 {
		t.Errorf("expected Remaining=-1 when disabled, got %d", r)
	}
}

func TestRateLimitConfig_Remaining_Counts(t *testing.T) {
	cfg := &RateLimitConfig{Enabled: true, MaxRequests: 5, WindowSize: time.Second}
	cfg.Allow()
	cfg.Allow()
	if r := cfg.Remaining(); r != 3 {
		t.Errorf("expected Remaining=3, got %d", r)
	}
}
