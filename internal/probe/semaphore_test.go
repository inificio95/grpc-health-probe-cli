package probe

import (
	"context"
	"sync"
	"testing"
	"time"
)

func TestDefaultSemaphoreConfig(t *testing.T) {
	cfg := DefaultSemaphoreConfig()
	if cfg == nil {
		t.Fatal("expected non-nil config")
	}
	if cfg.Enabled {
		t.Error("expected Enabled=false by default")
	}
	if cfg.MaxTickets != 10 {
		t.Errorf("expected MaxTickets=10, got %d", cfg.MaxTickets)
	}
	if cfg.AcquireTimeout != 5*time.Second {
		t.Errorf("expected AcquireTimeout=5s, got %v", cfg.AcquireTimeout)
	}
}

func TestSemaphoreConfig_Validate_Nil(t *testing.T) {
	var cfg *SemaphoreConfig
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for nil config")
	}
}

func TestSemaphoreConfig_Validate_Disabled(t *testing.T) {
	cfg := &SemaphoreConfig{Enabled: false}
	if err := cfg.Validate(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestSemaphoreConfig_Validate_InvalidMaxTickets(t *testing.T) {
	cfg := &SemaphoreConfig{Enabled: true, MaxTickets: 0, AcquireTimeout: time.Second}
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for MaxTickets=0")
	}
}

func TestSemaphoreConfig_Validate_InvalidAcquireTimeout(t *testing.T) {
	cfg := &SemaphoreConfig{Enabled: true, MaxTickets: 5, AcquireTimeout: 0}
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for AcquireTimeout=0")
	}
}

func TestSemaphoreConfig_Validate_Valid(t *testing.T) {
	cfg := &SemaphoreConfig{Enabled: true, MaxTickets: 5, AcquireTimeout: time.Second}
	if err := cfg.Validate(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestSemaphoreConfig_Acquire_Disabled(t *testing.T) {
	cfg := &SemaphoreConfig{Enabled: false}
	if err := cfg.Acquire(context.Background()); err != nil {
		t.Errorf("unexpected error when disabled: %v", err)
	}
	cfg.Release() // should not panic
}

func TestSemaphoreConfig_Acquire_LimitsParallelism(t *testing.T) {
	cfg := &SemaphoreConfig{Enabled: true, MaxTickets: 2, AcquireTimeout: 100 * time.Millisecond}
	cfg.Init()

	var mu sync.Mutex
	var active, maxActive int

	var wg sync.WaitGroup
	for i := 0; i < 6; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := cfg.Acquire(context.Background()); err != nil {
				return
			}
			defer cfg.Release()
			mu.Lock()
			active++
			if active > maxActive {
				maxActive = active
			}
			mu.Unlock()
			time.Sleep(10 * time.Millisecond)
			mu.Lock()
			active--
			mu.Unlock()
		}()
	}
	wg.Wait()
	if maxActive > 2 {
		t.Errorf("parallelism exceeded MaxTickets: got %d active at once", maxActive)
	}
}

func TestSemaphoreConfig_Acquire_Timeout(t *testing.T) {
	cfg := &SemaphoreConfig{Enabled: true, MaxTickets: 1, AcquireTimeout: 50 * time.Millisecond}
	cfg.Init()

	// Fill the semaphore.
	if err := cfg.Acquire(context.Background()); err != nil {
		t.Fatalf("unexpected error on first acquire: %v", err)
	}

	// Second acquire should time out.
	if err := cfg.Acquire(context.Background()); err == nil {
		t.Error("expected timeout error on second acquire")
	}
	cfg.Release()
}
