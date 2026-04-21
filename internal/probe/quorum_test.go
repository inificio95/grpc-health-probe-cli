package probe

import (
	"testing"
)

func TestDefaultQuorumConfig(t *testing.T) {
	cfg := DefaultQuorumConfig()
	if cfg == nil {
		t.Fatal("expected non-nil config")
	}
	if cfg.Enabled {
		t.Error("expected Enabled to be false by default")
	}
	if cfg.MinSuccess != 1 {
		t.Errorf("expected MinSuccess=1, got %d", cfg.MinSuccess)
	}
	if cfg.Total != 1 {
		t.Errorf("expected Total=1, got %d", cfg.Total)
	}
}

func TestQuorumConfig_Validate_Nil(t *testing.T) {
	var cfg *QuorumConfig
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for nil config")
	}
}

func TestQuorumConfig_Validate_Disabled(t *testing.T) {
	cfg := DefaultQuorumConfig()
	if err := cfg.Validate(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestQuorumConfig_Validate_Valid(t *testing.T) {
	cfg := &QuorumConfig{Enabled: true, MinSuccess: 2, Total: 3}
	if err := cfg.Validate(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestQuorumConfig_Validate_ZeroTotal(t *testing.T) {
	cfg := &QuorumConfig{Enabled: true, MinSuccess: 1, Total: 0}
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for zero total")
	}
}

func TestQuorumConfig_Validate_MinSuccessExceedsTotal(t *testing.T) {
	cfg := &QuorumConfig{Enabled: true, MinSuccess: 4, Total: 3}
	if err := cfg.Validate(); err == nil {
		t.Error("expected error when min_success > total")
	}
}

func TestQuorumConfig_IsMet_Disabled(t *testing.T) {
	cfg := DefaultQuorumConfig()
	if !cfg.IsMet(1) {
		t.Error("expected quorum to be met with 1 success when disabled")
	}
	if cfg.IsMet(0) {
		t.Error("expected quorum not met with 0 successes")
	}
}

func TestQuorumConfig_IsMet_Enabled(t *testing.T) {
	cfg := &QuorumConfig{Enabled: true, MinSuccess: 2, Total: 3}
	if cfg.IsMet(1) {
		t.Error("expected quorum not met with only 1 success")
	}
	if !cfg.IsMet(2) {
		t.Error("expected quorum met with 2 successes")
	}
	if !cfg.IsMet(3) {
		t.Error("expected quorum met with 3 successes")
	}
}

func TestQuorumConfig_IsMet_Nil(t *testing.T) {
	var cfg *QuorumConfig
	if !cfg.IsMet(1) {
		t.Error("expected nil config to treat 1 success as met")
	}
}

func TestQuorumConfig_String_Disabled(t *testing.T) {
	cfg := DefaultQuorumConfig()
	if s := cfg.String(); s != "quorum(disabled)" {
		t.Errorf("unexpected string: %q", s)
	}
}

func TestQuorumConfig_String_Enabled(t *testing.T) {
	cfg := &QuorumConfig{Enabled: true, MinSuccess: 2, Total: 3}
	if s := cfg.String(); s != "quorum(min=2/3)" {
		t.Errorf("unexpected string: %q", s)
	}
}

func TestQuorumConfig_String_Nil(t *testing.T) {
	var cfg *QuorumConfig
	if s := cfg.String(); s != "quorum(nil)" {
		t.Errorf("unexpected string: %q", s)
	}
}
