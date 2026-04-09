package probe

import (
	"testing"
)

func TestDefaultVerbosityConfig(t *testing.T) {
	cfg := DefaultVerbosityConfig()
	if cfg == nil {
		t.Fatal("expected non-nil config")
	}
	if cfg.Level != VerbosityNormal {
		t.Errorf("expected VerbosityNormal, got %d", cfg.Level)
	}
}

func TestVerbosityConfig_Validate_Valid(t *testing.T) {
	tests := []struct {
		name  string
		level VerbosityLevel
	}{
		{"quiet", VerbosityQuiet},
		{"normal", VerbosityNormal},
		{"verbose", VerbosityVerbose},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := &VerbosityConfig{Level: tt.level}
			if err := cfg.Validate(); err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}

func TestVerbosityConfig_Validate_Nil(t *testing.T) {
	var cfg *VerbosityConfig
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for nil config")
	}
}

func TestVerbosityConfig_Validate_InvalidLevel(t *testing.T) {
	cfg := &VerbosityConfig{Level: VerbosityLevel(99)}
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for invalid verbosity level")
	}
}

func TestVerbosityConfig_IsQuiet(t *testing.T) {
	cfg := &VerbosityConfig{Level: VerbosityQuiet}
	if !cfg.IsQuiet() {
		t.Error("expected IsQuiet to return true")
	}
	cfg.Level = VerbosityNormal
	if cfg.IsQuiet() {
		t.Error("expected IsQuiet to return false for Normal")
	}
}

func TestVerbosityConfig_IsVerbose(t *testing.T) {
	cfg := &VerbosityConfig{Level: VerbosityVerbose}
	if !cfg.IsVerbose() {
		t.Error("expected IsVerbose to return true")
	}
	cfg.Level = VerbosityNormal
	if cfg.IsVerbose() {
		t.Error("expected IsVerbose to return false for Normal")
	}
}
