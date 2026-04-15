package probe

import (
	"testing"
)

func TestDefaultConcurrencyConfig(t *testing.T) {
	cfg := DefaultConcurrencyConfig()
	if cfg == nil {
		t.Fatal("expected non-nil config")
	}
	if cfg.Enabled {
		t.Error("expected Enabled to be false by default")
	}
	if cfg.MaxWorkers != 1 {
		t.Errorf("expected MaxWorkers=1, got %d", cfg.MaxWorkers)
	}
}

func TestConcurrencyConfig_Validate_Nil(t *testing.T) {
	var cfg *ConcurrencyConfig
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for nil config")
	}
}

func TestConcurrencyConfig_Validate_Disabled(t *testing.T) {
	cfg := &ConcurrencyConfig{Enabled: false, MaxWorkers: 0}
	if err := cfg.Validate(); err != nil {
		t.Errorf("expected no error for disabled config, got %v", err)
	}
}

func TestConcurrencyConfig_Validate_EnabledValid(t *testing.T) {
	cfg := &ConcurrencyConfig{Enabled: true, MaxWorkers: 4}
	if err := cfg.Validate(); err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestConcurrencyConfig_Validate_InvalidMaxWorkers(t *testing.T) {
	cfg := &ConcurrencyConfig{Enabled: true, MaxWorkers: 0}
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for max_workers=0")
	}
}

func TestConcurrencyConfig_Validate_MaxWorkersExceeded(t *testing.T) {
	cfg := &ConcurrencyConfig{Enabled: true, MaxWorkers: 300}
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for max_workers > 256")
	}
}

func TestConcurrencyConfig_WorkerPool_Disabled(t *testing.T) {
	cfg := DefaultConcurrencyConfig()
	pool := cfg.WorkerPool()
	if cap(pool) != 1 {
		t.Errorf("expected pool capacity 1, got %d", cap(pool))
	}
	if len(pool) != 1 {
		t.Errorf("expected pool length 1, got %d", len(pool))
	}
}

func TestConcurrencyConfig_WorkerPool_Enabled(t *testing.T) {
	cfg := &ConcurrencyConfig{Enabled: true, MaxWorkers: 5}
	pool := cfg.WorkerPool()
	if cap(pool) != 5 {
		t.Errorf("expected pool capacity 5, got %d", cap(pool))
	}
	if len(pool) != 5 {
		t.Errorf("expected pool length 5, got %d", len(pool))
	}
}

func TestConcurrencyConfig_WorkerPool_Nil(t *testing.T) {
	var cfg *ConcurrencyConfig
	pool := cfg.WorkerPool()
	if cap(pool) != 1 {
		t.Errorf("expected pool capacity 1 for nil config, got %d", cap(pool))
	}
}

func TestConcurrencyConfig_String_Disabled(t *testing.T) {
	cfg := DefaultConcurrencyConfig()
	if s := cfg.String(); s != "concurrency(disabled)" {
		t.Errorf("unexpected string: %s", s)
	}
}

func TestConcurrencyConfig_String_Enabled(t *testing.T) {
	cfg := &ConcurrencyConfig{Enabled: true, MaxWorkers: 8}
	if s := cfg.String(); s != "concurrency(max_workers=8)" {
		t.Errorf("unexpected string: %s", s)
	}
}

func TestConcurrencyConfig_String_Nil(t *testing.T) {
	var cfg *ConcurrencyConfig
	if s := cfg.String(); s != "concurrency(nil)" {
		t.Errorf("unexpected string: %s", s)
	}
}
