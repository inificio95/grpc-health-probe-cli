package probe

import (
	"testing"
	"time"
)

func TestDefaultKeepaliveConfig(t *testing.T) {
	cfg := DefaultKeepaliveConfig()
	if cfg == nil {
		t.Fatal("expected non-nil config")
	}
	if cfg.Enabled {
		t.Error("expected Enabled to be false by default")
	}
	if cfg.Time != 30*time.Second {
		t.Errorf("expected Time=30s, got %v", cfg.Time)
	}
	if cfg.Timeout != 10*time.Second {
		t.Errorf("expected Timeout=10s, got %v", cfg.Timeout)
	}
}

func TestKeepaliveConfig_Validate_Nil(t *testing.T) {
	var cfg *KeepaliveConfig
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for nil config")
	}
}

func TestKeepaliveConfig_Validate_Valid(t *testing.T) {
	cfg := DefaultKeepaliveConfig()
	if err := cfg.Validate(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestKeepaliveConfig_Validate_ZeroTime(t *testing.T) {
	cfg := DefaultKeepaliveConfig()
	cfg.Time = 0
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for zero Time")
	}
}

func TestKeepaliveConfig_Validate_ZeroTimeout(t *testing.T) {
	cfg := DefaultKeepaliveConfig()
	cfg.Timeout = 0
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for zero Timeout")
	}
}

func TestKeepaliveConfig_DialOption_Disabled(t *testing.T) {
	cfg := DefaultKeepaliveConfig()
	cfg.Enabled = false
	if opt := cfg.DialOption(); opt != nil {
		t.Error("expected nil DialOption when disabled")
	}
}

func TestKeepaliveConfig_DialOption_Enabled(t *testing.T) {
	cfg := DefaultKeepaliveConfig()
	cfg.Enabled = true
	if opt := cfg.DialOption(); opt == nil {
		t.Error("expected non-nil DialOption when enabled")
	}
}

func TestKeepaliveConfig_DialOption_NilConfig(t *testing.T) {
	var cfg *KeepaliveConfig
	if opt := cfg.DialOption(); opt != nil {
		t.Error("expected nil DialOption for nil config")
	}
}
