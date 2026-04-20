package probe

import (
	"testing"
	"time"
)

func TestDefaultThrottleConfig(t *testing.T) {
	cfg := DefaultThrottleConfig()
	if cfg == nil {
		t.Fatal("expected non-nil config")
	}
	if cfg.Enabled {
		t.Error("expected throttle to be disabled by default")
	}
	if cfg.MinInterval != 500*time.Millisecond {
		t.Errorf("unexpected default min_interval: %s", cfg.MinInterval)
	}
	if cfg.Burst != 1 {
		t.Errorf("unexpected default burst: %d", cfg.Burst)
	}
}

func TestThrottleConfig_Validate_Nil(t *testing.T) {
	var cfg *ThrottleConfig
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for nil config")
	}
}

func TestThrottleConfig_Validate_Disabled(t *testing.T) {
	cfg := &ThrottleConfig{Enabled: false}
	if err := cfg.Validate(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestThrottleConfig_Validate_InvalidMinInterval(t *testing.T) {
	cfg := &ThrottleConfig{Enabled: true, MinInterval: 0, Burst: 1}
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for zero min_interval")
	}
}

func TestThrottleConfig_Validate_InvalidBurst(t *testing.T) {
	cfg := &ThrottleConfig{Enabled: true, MinInterval: 100 * time.Millisecond, Burst: 0}
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for burst < 1")
	}
}

func TestThrottleConfig_Validate_Valid(t *testing.T) {
	cfg := &ThrottleConfig{Enabled: true, MinInterval: 200 * time.Millisecond, Burst: 2}
	if err := cfg.Validate(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestNewThrottler_NilConfig(t *testing.T) {
	th := NewThrottler(nil)
	if th == nil {
		t.Fatal("expected non-nil throttler")
	}
	if th.cfg.Enabled {
		t.Error("expected throttler to be disabled when config is nil")
	}
}

func TestThrottler_Wait_Disabled(t *testing.T) {
	cfg := &ThrottleConfig{Enabled: false, MinInterval: 10 * time.Second, Burst: 1}
	th := NewThrottler(cfg)
	start := time.Now()
	th.Wait()
	th.Wait()
	if time.Since(start) > 50*time.Millisecond {
		t.Error("Wait should return immediately when throttling is disabled")
	}
}

func TestThrottler_Wait_BurstAllowsImmediateCall(t *testing.T) {
	cfg := &ThrottleConfig{Enabled: true, MinInterval: 5 * time.Second, Burst: 2}
	th := NewThrottler(cfg)
	start := time.Now()
	th.Wait() // consumes burst token 1
	th.Wait() // consumes burst token 2
	if time.Since(start) > 50*time.Millisecond {
		t.Error("burst tokens should allow immediate calls")
	}
}

func TestThrottler_Wait_EnforcesInterval(t *testing.T) {
	cfg := &ThrottleConfig{Enabled: true, MinInterval: 100 * time.Millisecond, Burst: 1}
	th := NewThrottler(cfg)
	th.Wait() // use burst token
	start := time.Now()
	th.Wait() // should throttle
	elapsed := time.Since(start)
	if elapsed < 80*time.Millisecond {
		t.Errorf("expected throttle delay, got %s", elapsed)
	}
}
