package probe

import (
	"testing"
	"time"
)

func TestDefaultDrainConfig(t *testing.T) {
	cfg := DefaultDrainConfig()
	if cfg == nil {
		t.Fatal("expected non-nil config")
	}
	if cfg.Enabled {
		t.Error("expected Enabled to be false by default")
	}
	if cfg.DrainTimeout != 5*time.Second {
		t.Errorf("expected DrainTimeout 5s, got %s", cfg.DrainTimeout)
	}
}

func TestDrainConfig_Validate_Nil(t *testing.T) {
	var cfg *DrainConfig
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for nil config")
	}
}

func TestDrainConfig_Validate_Disabled(t *testing.T) {
	cfg := &DrainConfig{Enabled: false, DrainTimeout: 0}
	if err := cfg.Validate(); err != nil {
		t.Errorf("unexpected error for disabled config: %v", err)
	}
}

func TestDrainConfig_Validate_EnabledValid(t *testing.T) {
	cfg := &DrainConfig{Enabled: true, DrainTimeout: 3 * time.Second}
	if err := cfg.Validate(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestDrainConfig_Validate_EnabledZeroDuration(t *testing.T) {
	cfg := &DrainConfig{Enabled: true, DrainTimeout: 0}
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for zero drain timeout when enabled")
	}
}

func TestDrainConfig_Wait_Disabled(t *testing.T) {
	cfg := &DrainConfig{Enabled: false, DrainTimeout: 10 * time.Millisecond}
	done := make(chan struct{})
	// should return immediately without reading done
	if !cfg.Wait(done) {
		t.Error("expected Wait to return true when disabled")
	}
}

func TestDrainConfig_Wait_NilConfig(t *testing.T) {
	var cfg *DrainConfig
	done := make(chan struct{})
	if !cfg.Wait(done) {
		t.Error("expected Wait to return true for nil config")
	}
}

func TestDrainConfig_Wait_CompletesBeforeTimeout(t *testing.T) {
	cfg := &DrainConfig{Enabled: true, DrainTimeout: 500 * time.Millisecond}
	done := make(chan struct{})
	go func() {
		time.Sleep(10 * time.Millisecond)
		close(done)
	}()
	if !cfg.Wait(done) {
		t.Error("expected drain to complete before timeout")
	}
}

func TestDrainConfig_Wait_TimesOut(t *testing.T) {
	cfg := &DrainConfig{Enabled: true, DrainTimeout: 20 * time.Millisecond}
	done := make(chan struct{}) // never closed
	if cfg.Wait(done) {
		t.Error("expected drain to time out")
	}
}
