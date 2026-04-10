package probe

import (
	"strings"
	"testing"
	"time"
)

func TestKeepaliveConfig_DialOption_Disabled(t *testing.T) {
	cfg := &KeepaliveConfig{Enabled: false}
	_, ok := cfg.DialOption()
	if ok {
		t.Error("expected DialOption to return false when disabled")
	}
}

func TestKeepaliveConfig_DialOption_Nil(t *testing.T) {
	var cfg *KeepaliveConfig
	_, ok := cfg.DialOption()
	if ok {
		t.Error("expected DialOption to return false for nil config")
	}
}

func TestKeepaliveConfig_DialOption_Enabled(t *testing.T) {
	cfg := &KeepaliveConfig{
		Enabled:             true,
		Time:                30 * time.Second,
		Timeout:             10 * time.Second,
		PermitWithoutStream: true,
	}
	params, ok := cfg.DialOption()
	if !ok {
		t.Fatal("expected DialOption to return true when enabled")
	}
	if params.Time != 30*time.Second {
		t.Errorf("expected Time=30s, got %s", params.Time)
	}
	if params.Timeout != 10*time.Second {
		t.Errorf("expected Timeout=10s, got %s", params.Timeout)
	}
	if !params.PermitWithoutStream {
		t.Error("expected PermitWithoutStream=true")
	}
}

func TestKeepaliveConfig_String_Disabled(t *testing.T) {
	cfg := &KeepaliveConfig{Enabled: false}
	s := cfg.String()
	if !strings.Contains(s, "disabled") {
		t.Errorf("expected 'disabled' in string, got: %s", s)
	}
}

func TestKeepaliveConfig_String_Nil(t *testing.T) {
	var cfg *KeepaliveConfig
	s := cfg.String()
	if !strings.Contains(s, "disabled") {
		t.Errorf("expected 'disabled' in string for nil config, got: %s", s)
	}
}

func TestKeepaliveConfig_String_Enabled(t *testing.T) {
	cfg := &KeepaliveConfig{
		Enabled:             true,
		Time:                20 * time.Second,
		Timeout:             5 * time.Second,
		PermitWithoutStream: false,
	}
	s := cfg.String()
	if !strings.Contains(s, "20s") {
		t.Errorf("expected '20s' in string, got: %s", s)
	}
	if !strings.Contains(s, "5s") {
		t.Errorf("expected '5s' in string, got: %s", s)
	}
}
