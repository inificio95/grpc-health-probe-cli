package probe

import (
	"testing"
	"time"
)

func TestDefaultBackoffConfig(t *testing.T) {
	c := DefaultBackoffConfig()
	if c == nil {
		t.Fatal("expected non-nil config")
	}
	if c.Strategy != BackoffFixed {
		t.Errorf("expected strategy %q, got %q", BackoffFixed, c.Strategy)
	}
	if c.InitialDelay != 500*time.Millisecond {
		t.Errorf("unexpected initial delay: %v", c.InitialDelay)
	}
	if c.MaxDelay != 30*time.Second {
		t.Errorf("unexpected max delay: %v", c.MaxDelay)
	}
	if c.Multiplier != 2.0 {
		t.Errorf("unexpected multiplier: %v", c.Multiplier)
	}
}

func TestBackoffConfig_Validate_Nil(t *testing.T) {
	var c *BackoffConfig
	if err := c.Validate(); err != ErrNilConfig {
		t.Errorf("expected ErrNilConfig, got %v", err)
	}
}

func TestBackoffConfig_Validate_InvalidStrategy(t *testing.T) {
	c := DefaultBackoffConfig()
	c.Strategy = "random"
	if err := c.Validate(); err != ErrInvalidBackoffStrategy {
		t.Errorf("expected ErrInvalidBackoffStrategy, got %v", err)
	}
}

func TestBackoffConfig_Validate_InvalidInitialDelay(t *testing.T) {
	c := DefaultBackoffConfig()
	c.InitialDelay = 0
	if err := c.Validate(); err != ErrInvalidInitialDelay {
		t.Errorf("expected ErrInvalidInitialDelay, got %v", err)
	}
}

func TestBackoffConfig_Validate_MaxDelayTooSmall(t *testing.T) {
	c := DefaultBackoffConfig()
	c.MaxDelay = c.InitialDelay - time.Millisecond
	if err := c.Validate(); err != ErrMaxDelayTooSmall {
		t.Errorf("expected ErrMaxDelayTooSmall, got %v", err)
	}
}

func TestBackoffConfig_Validate_ExponentialInvalidMultiplier(t *testing.T) {
	c := DefaultBackoffConfig()
	c.Strategy = BackoffExponential
	c.Multiplier = 1.0
	if err := c.Validate(); err != ErrInvalidMultiplier {
		t.Errorf("expected ErrInvalidMultiplier, got %v", err)
	}
}

func TestBackoffConfig_Delay_Fixed(t *testing.T) {
	c := &BackoffConfig{Strategy: BackoffFixed, InitialDelay: 100 * time.Millisecond, MaxDelay: 10 * time.Second, Multiplier: 2.0}
	for i := 0; i < 5; i++ {
		if d := c.Delay(i); d != 100*time.Millisecond {
			t.Errorf("attempt %d: expected 100ms, got %v", i, d)
		}
	}
}

func TestBackoffConfig_Delay_Exponential(t *testing.T) {
	c := &BackoffConfig{Strategy: BackoffExponential, InitialDelay: 100 * time.Millisecond, MaxDelay: 10 * time.Second, Multiplier: 2.0}
	expected := []time.Duration{100 * time.Millisecond, 200 * time.Millisecond, 400 * time.Millisecond}
	for i, want := range expected {
		if got := c.Delay(i); got != want {
			t.Errorf("attempt %d: expected %v, got %v", i, want, got)
		}
	}
}

func TestBackoffConfig_Delay_Linear(t *testing.T) {
	c := &BackoffConfig{Strategy: BackoffLinear, InitialDelay: 100 * time.Millisecond, MaxDelay: 10 * time.Second, Multiplier: 2.0}
	expected := []time.Duration{100 * time.Millisecond, 200 * time.Millisecond, 300 * time.Millisecond}
	for i, want := range expected {
		if got := c.Delay(i); got != want {
			t.Errorf("attempt %d: expected %v, got %v", i, want, got)
		}
	}
}

func TestBackoffConfig_Delay_CappedAtMax(t *testing.T) {
	c := &BackoffConfig{Strategy: BackoffExponential, InitialDelay: 1 * time.Second, MaxDelay: 3 * time.Second, Multiplier: 2.0}
	if got := c.Delay(5); got != 3*time.Second {
		t.Errorf("expected delay capped at 3s, got %v", got)
	}
}
