package probe

import (
	"testing"
	"time"
)

func TestDefaultTimeoutBudgetConfig(t *testing.T) {
	cfg := DefaultTimeoutBudgetConfig()
	if cfg == nil {
		t.Fatal("expected non-nil config")
	}
	if cfg.Enabled {
		t.Error("expected Enabled to be false by default")
	}
	if cfg.TotalBudget != 30*time.Second {
		t.Errorf("expected TotalBudget 30s, got %v", cfg.TotalBudget)
	}
	if cfg.ReserveBuffer != 500*time.Millisecond {
		t.Errorf("expected ReserveBuffer 500ms, got %v", cfg.ReserveBuffer)
	}
}

func TestTimeoutBudgetConfig_Validate_Nil(t *testing.T) {
	var cfg *TimeoutBudgetConfig
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for nil config")
	}
}

func TestTimeoutBudgetConfig_Validate_Disabled(t *testing.T) {
	cfg := &TimeoutBudgetConfig{Enabled: false}
	if err := cfg.Validate(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestTimeoutBudgetConfig_Validate_Valid(t *testing.T) {
	cfg := &TimeoutBudgetConfig{
		Enabled:       true,
		TotalBudget:   10 * time.Second,
		ReserveBuffer: 200 * time.Millisecond,
	}
	if err := cfg.Validate(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestTimeoutBudgetConfig_Validate_ZeroBudget(t *testing.T) {
	cfg := &TimeoutBudgetConfig{Enabled: true, TotalBudget: 0, ReserveBuffer: 0}
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for zero total budget")
	}
}

func TestTimeoutBudgetConfig_Validate_BufferExceedsBudget(t *testing.T) {
	cfg := &TimeoutBudgetConfig{
		Enabled:       true,
		TotalBudget:   1 * time.Second,
		ReserveBuffer: 2 * time.Second,
	}
	if err := cfg.Validate(); err == nil {
		t.Error("expected error when buffer >= budget")
	}
}

func TestBudget_Remaining_Fresh(t *testing.T) {
	cfg := &TimeoutBudgetConfig{
		Enabled:       true,
		TotalBudget:   10 * time.Second,
		ReserveBuffer: 500 * time.Millisecond,
	}
	b := NewBudget(cfg)
	remaining := b.Remaining()
	if remaining <= 0 {
		t.Errorf("expected positive remaining, got %v", remaining)
	}
	if remaining > 10*time.Second {
		t.Errorf("remaining %v exceeds total budget", remaining)
	}
}

func TestBudget_Exhausted_AfterExpiry(t *testing.T) {
	cfg := &TimeoutBudgetConfig{
		Enabled:       true,
		TotalBudget:   1 * time.Millisecond,
		ReserveBuffer: 0,
	}
	b := NewBudget(cfg)
	time.Sleep(5 * time.Millisecond)
	if !b.Exhausted() {
		t.Error("expected budget to be exhausted")
	}
}

func TestBudget_Exhausted_False_WhenFresh(t *testing.T) {
	cfg := &TimeoutBudgetConfig{
		Enabled:       true,
		TotalBudget:   10 * time.Second,
		ReserveBuffer: 0,
	}
	b := NewBudget(cfg)
	if b.Exhausted() {
		t.Error("expected budget not to be exhausted immediately")
	}
}
