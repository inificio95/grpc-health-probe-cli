package probe

import (
	"testing"

	"google.golang.org/grpc"
)

func TestDefaultFailFastConfig(t *testing.T) {
	cfg := DefaultFailFastConfig()
	if cfg == nil {
		t.Fatal("expected non-nil config")
	}
	if !cfg.Enabled {
		t.Error("expected Enabled to be true by default")
	}
}

func TestFailFastConfig_Validate_Nil(t *testing.T) {
	var cfg *FailFastConfig
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for nil config")
	}
}

func TestFailFastConfig_Validate_Valid(t *testing.T) {
	cfg := DefaultFailFastConfig()
	if err := cfg.Validate(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestFailFastConfig_CallOption_Enabled(t *testing.T) {
	cfg := &FailFastConfig{Enabled: true}
	opt := cfg.CallOption()
	// grpc.WaitForReady(false) should equal fail-fast behaviour.
	// We verify the option is of the correct type by comparing with the
	// expected option value.
	want := grpc.WaitForReady(false)
	if opt != want {
		t.Errorf("expected WaitForReady(false), got %v", opt)
	}
}

func TestFailFastConfig_CallOption_Disabled(t *testing.T) {
	cfg := &FailFastConfig{Enabled: false}
	opt := cfg.CallOption()
	want := grpc.WaitForReady(true)
	if opt != want {
		t.Errorf("expected WaitForReady(true), got %v", opt)
	}
}

func TestFailFastConfig_CallOption_Nil(t *testing.T) {
	var cfg *FailFastConfig
	// nil config should behave as fail-fast enabled.
	opt := cfg.CallOption()
	want := grpc.WaitForReady(false)
	if opt != want {
		t.Errorf("expected WaitForReady(false) for nil config, got %v", opt)
	}
}

func TestFailFastConfig_String_Enabled(t *testing.T) {
	cfg := &FailFastConfig{Enabled: true}
	if s := cfg.String(); s != "failfast=enabled" {
		t.Errorf("unexpected string: %q", s)
	}
}

func TestFailFastConfig_String_Disabled(t *testing.T) {
	cfg := &FailFastConfig{Enabled: false}
	if s := cfg.String(); s != "failfast=disabled" {
		t.Errorf("unexpected string: %q", s)
	}
}

func TestFailFastConfig_String_Nil(t *testing.T) {
	var cfg *FailFastConfig
	if s := cfg.String(); s != "failfast=nil" {
		t.Errorf("unexpected string: %q", s)
	}
}
