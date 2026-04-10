package probe

import (
	"context"
	"testing"
	"time"
)

func TestDefaultDeadlineConfig(t *testing.T) {
	cfg := DefaultDeadlineConfig()
	if cfg == nil {
		t.Fatal("expected non-nil config")
	}
	if cfg.Enabled {
		t.Error("expected Enabled to be false by default")
	}
	if cfg.Duration != 30*time.Second {
		t.Errorf("expected Duration 30s, got %v", cfg.Duration)
	}
}

func TestDeadlineConfig_Validate_Nil(t *testing.T) {
	var cfg *DeadlineConfig
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for nil config")
	}
}

func TestDeadlineConfig_Validate_Disabled(t *testing.T) {
	cfg := &DeadlineConfig{Enabled: false, Duration: 0}
	if err := cfg.Validate(); err != nil {
		t.Errorf("expected no error for disabled config, got %v", err)
	}
}

func TestDeadlineConfig_Validate_EnabledPositive(t *testing.T) {
	cfg := &DeadlineConfig{Enabled: true, Duration: 10 * time.Second}
	if err := cfg.Validate(); err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestDeadlineConfig_Validate_EnabledZeroDuration(t *testing.T) {
	cfg := &DeadlineConfig{Enabled: true, Duration: 0}
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for enabled config with zero duration")
	}
}

func TestDeadlineConfig_Apply_Disabled(t *testing.T) {
	cfg := &DeadlineConfig{Enabled: false, Duration: 5 * time.Second}
	ctx := context.Background()
	newCtx, cancel := cfg.Apply(ctx)
	defer cancel()
	if _, ok := newCtx.Deadline(); ok {
		t.Error("expected no deadline on disabled config")
	}
}

func TestDeadlineConfig_Apply_Enabled(t *testing.T) {
	cfg := &DeadlineConfig{Enabled: true, Duration: 5 * time.Second}
	ctx := context.Background()
	newCtx, cancel := cfg.Apply(ctx)
	defer cancel()
	deadline, ok := newCtx.Deadline()
	if !ok {
		t.Fatal("expected deadline to be set")
	}
	if time.Until(deadline) > 5*time.Second || time.Until(deadline) <= 0 {
		t.Errorf("deadline out of expected range: %v", deadline)
	}
}

func TestDeadlineConfig_Apply_NilConfig(t *testing.T) {
	var cfg *DeadlineConfig
	ctx := context.Background()
	newCtx, cancel := cfg.Apply(ctx)
	defer cancel()
	if _, ok := newCtx.Deadline(); ok {
		t.Error("expected no deadline for nil config")
	}
}
