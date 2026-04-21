package probe

import (
	"testing"
	"time"
)

func TestDefaultBurstConfig(t *testing.T) {
	c := DefaultBurstConfig()
	if c == nil {
		t.Fatal("expected non-nil config")
	}
	if c.Enabled {
		t.Error("expected Enabled=false by default")
	}
	if c.MaxBurst != 5 {
		t.Errorf("expected MaxBurst=5, got %d", c.MaxBurst)
	}
	if c.BurstInterval != 500*time.Millisecond {
		t.Errorf("unexpected BurstInterval: %s", c.BurstInterval)
	}
	if c.Cooldown != 2*time.Second {
		t.Errorf("unexpected Cooldown: %s", c.Cooldown)
	}
}

func TestBurstConfig_Validate_Nil(t *testing.T) {
	var c *BurstConfig
	if err := c.Validate(); err == nil {
		t.Error("expected error for nil config")
	}
}

func TestBurstConfig_Validate_Disabled(t *testing.T) {
	c := &BurstConfig{Enabled: false}
	if err := c.Validate(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestBurstConfig_Validate_InvalidMaxBurst(t *testing.T) {
	c := &BurstConfig{Enabled: true, MaxBurst: 0, BurstInterval: time.Second, Cooldown: time.Second}
	if err := c.Validate(); err == nil {
		t.Error("expected error for MaxBurst=0")
	}
}

func TestBurstConfig_Validate_InvalidInterval(t *testing.T) {
	c := &BurstConfig{Enabled: true, MaxBurst: 3, BurstInterval: 0, Cooldown: time.Second}
	if err := c.Validate(); err == nil {
		t.Error("expected error for zero BurstInterval")
	}
}

func TestBurstConfig_Validate_InvalidCooldown(t *testing.T) {
	c := &BurstConfig{Enabled: true, MaxBurst: 3, BurstInterval: time.Second, Cooldown: 0}
	if err := c.Validate(); err == nil {
		t.Error("expected error for zero Cooldown")
	}
}

func TestBurstConfig_Allow_Disabled(t *testing.T) {
	c := &BurstConfig{Enabled: false}
	for i := 0; i < 20; i++ {
		if !c.Allow() {
			t.Error("disabled burst config should always allow")
		}
	}
}

func TestBurstConfig_Allow_NilConfig(t *testing.T) {
	var c *BurstConfig
	if !c.Allow() {
		t.Error("nil burst config should always allow")
	}
}

func TestBurstConfig_Allow_WithinBurst(t *testing.T) {
	c := &BurstConfig{
		Enabled:       true,
		MaxBurst:      3,
		BurstInterval: time.Second,
		Cooldown:      2 * time.Second,
	}
	for i := 0; i < 3; i++ {
		if !c.Allow() {
			t.Errorf("attempt %d should be allowed", i+1)
		}
	}
}

func TestBurstConfig_Allow_ExceedsBurst(t *testing.T) {
	c := &BurstConfig{
		Enabled:       true,
		MaxBurst:      3,
		BurstInterval: time.Second,
		Cooldown:      2 * time.Second,
	}
	for i := 0; i < 3; i++ {
		c.Allow()
	}
	if c.Allow() {
		t.Error("4th attempt should be suppressed after burst")
	}
}
