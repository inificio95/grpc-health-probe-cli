package probe

import (
	"testing"
	"time"
)

func TestDefaultCircuitBreakerConfig(t *testing.T) {
	cfg := DefaultCircuitBreakerConfig()
	if cfg == nil {
		t.Fatal("expected non-nil config")
	}
	if cfg.Enabled {
		t.Error("expected Enabled=false by default")
	}
	if cfg.MaxFailures != 5 {
		t.Errorf("expected MaxFailures=5, got %d", cfg.MaxFailures)
	}
	if cfg.OpenDuration != 30*time.Second {
		t.Errorf("expected OpenDuration=30s, got %s", cfg.OpenDuration)
	}
}

func TestCircuitBreakerConfig_Validate_Nil(t *testing.T) {
	var cfg *CircuitBreakerConfig
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for nil config")
	}
}

func TestCircuitBreakerConfig_Validate_Disabled(t *testing.T) {
	cfg := DefaultCircuitBreakerConfig()
	if err := cfg.Validate(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestCircuitBreakerConfig_Validate_InvalidMaxFailures(t *testing.T) {
	cfg := DefaultCircuitBreakerConfig()
	cfg.Enabled = true
	cfg.MaxFailures = 0
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for MaxFailures=0")
	}
}

func TestCircuitBreakerConfig_Validate_InvalidOpenDuration(t *testing.T) {
	cfg := DefaultCircuitBreakerConfig()
	cfg.Enabled = true
	cfg.OpenDuration = 0
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for OpenDuration=0")
	}
}

func TestCircuitBreakerConfig_Validate_Valid(t *testing.T) {
	cfg := DefaultCircuitBreakerConfig()
	cfg.Enabled = true
	if err := cfg.Validate(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestCircuitBreaker_AllowWhenClosed(t *testing.T) {
	cfg := DefaultCircuitBreakerConfig()
	cfg.Enabled = true
	if !cfg.Allow() {
		t.Error("expected Allow=true when circuit is closed")
	}
}

func TestCircuitBreaker_OpensAfterMaxFailures(t *testing.T) {
	cfg := DefaultCircuitBreakerConfig()
	cfg.Enabled = true
	cfg.MaxFailures = 3
	for i := 0; i < 3; i++ {
		cfg.RecordFailure()
	}
	if cfg.State() != CircuitOpen {
		t.Errorf("expected CircuitOpen, got %v", cfg.State())
	}
	if cfg.Allow() {
		t.Error("expected Allow=false when circuit is open")
	}
}

func TestCircuitBreaker_HalfOpenAfterDuration(t *testing.T) {
	cfg := DefaultCircuitBreakerConfig()
	cfg.Enabled = true
	cfg.MaxFailures = 1
	cfg.OpenDuration = 10 * time.Millisecond
	cfg.HalfOpenRequests = 1
	cfg.RecordFailure()
	time.Sleep(20 * time.Millisecond)
	if !cfg.Allow() {
		t.Error("expected Allow=true after open duration elapsed")
	}
	if cfg.State() != CircuitHalfOpen {
		t.Errorf("expected CircuitHalfOpen, got %v", cfg.State())
	}
}

func TestCircuitBreaker_ClosesOnSuccess(t *testing.T) {
	cfg := DefaultCircuitBreakerConfig()
	cfg.Enabled = true
	cfg.MaxFailures = 1
	cfg.OpenDuration = 10 * time.Millisecond
	cfg.RecordFailure()
	time.Sleep(20 * time.Millisecond)
	cfg.Allow()
	cfg.RecordSuccess()
	if cfg.State() != CircuitClosed {
		t.Errorf("expected CircuitClosed after success, got %v", cfg.State())
	}
}

func TestCircuitBreaker_DisabledAlwaysAllows(t *testing.T) {
	cfg := DefaultCircuitBreakerConfig()
	for i := 0; i < 100; i++ {
		cfg.RecordFailure()
	}
	if !cfg.Allow() {
		t.Error("expected Allow=true when circuit breaker is disabled")
	}
}
