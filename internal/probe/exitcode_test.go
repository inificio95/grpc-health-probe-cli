package probe

import (
	"errors"
	"testing"
)

func TestDefaultExitCodeConfig(t *testing.T) {
	cfg := DefaultExitCodeConfig()
	if cfg == nil {
		t.Fatal("expected non-nil config")
	}
	if !cfg.Enabled {
		t.Error("expected Enabled to be true")
	}
	if cfg.NotServingCode != 1 {
		t.Errorf("expected NotServingCode 1, got %d", cfg.NotServingCode)
	}
	if cfg.UnknownCode != 2 {
		t.Errorf("expected UnknownCode 2, got %d", cfg.UnknownCode)
	}
	if cfg.TimeoutCode != 4 {
		t.Errorf("expected TimeoutCode 4, got %d", cfg.TimeoutCode)
	}
}

func TestExitCodeConfig_Validate_Nil(t *testing.T) {
	var cfg *ExitCodeConfig
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for nil config")
	}
}

func TestExitCodeConfig_Validate_Valid(t *testing.T) {
	cfg := DefaultExitCodeConfig()
	if err := cfg.Validate(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestExitCodeConfig_Validate_OutOfRange(t *testing.T) {
	cfg := DefaultExitCodeConfig()
	cfg.NotServingCode = 200
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for out-of-range code")
	}
}

func TestExitCodeConfig_Resolve_Serving(t *testing.T) {
	cfg := DefaultExitCodeConfig()
	r := Result{Status: StatusServing}
	if code := cfg.Resolve(r); code != ExitOK {
		t.Errorf("expected ExitOK, got %d", code)
	}
}

func TestExitCodeConfig_Resolve_NotServing(t *testing.T) {
	cfg := DefaultExitCodeConfig()
	r := Result{Status: StatusNotServing}
	if code := cfg.Resolve(r); code != ExitNotServing {
		t.Errorf("expected ExitNotServing, got %d", code)
	}
}

func TestExitCodeConfig_Resolve_Unknown(t *testing.T) {
	cfg := DefaultExitCodeConfig()
	r := Result{Status: StatusUnknown}
	if code := cfg.Resolve(r); code != ExitUnknown {
		t.Errorf("expected ExitUnknown, got %d", code)
	}
}

func TestExitCodeConfig_Resolve_Error(t *testing.T) {
	cfg := DefaultExitCodeConfig()
	r := Result{Error: errors.New("dial failed")}
	if code := cfg.Resolve(r); code != ExitConnectionFailure {
		t.Errorf("expected ExitConnectionFailure, got %d", code)
	}
}

func TestExitCodeConfig_Resolve_Disabled(t *testing.T) {
	cfg := DefaultExitCodeConfig()
	cfg.Enabled = false
	r := Result{Status: StatusNotServing}
	if code := cfg.Resolve(r); code != ExitOK {
		t.Errorf("expected ExitOK when disabled, got %d", code)
	}
}
