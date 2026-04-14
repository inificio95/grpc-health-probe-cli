package probe

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestDefaultRetryConfig(t *testing.T) {
	cfg := DefaultRetryConfig()
	if cfg.MaxAttempts != 3 {
		t.Errorf("expected MaxAttempts=3, got %d", cfg.MaxAttempts)
	}
	if cfg.Delay != 500*time.Millisecond {
		t.Errorf("expected Delay=500ms, got %v", cfg.Delay)
	}
}

func TestRetryConfigValidate(t *testing.T) {
	tests := []struct {
		name    string
		cfg     RetryConfig
		wantErr bool
	}{
		{"valid", RetryConfig{MaxAttempts: 1, Delay: 0}, false},
		{"zero attempts", RetryConfig{MaxAttempts: 0, Delay: 0}, true},
		{"negative delay", RetryConfig{MaxAttempts: 1, Delay: -1}, true},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.cfg.Validate()
			if (err != nil) != tc.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}

func TestWithRetry_SuccessOnFirstAttempt(t *testing.T) {
	calls := 0
	err := WithRetry(context.Background(), RetryConfig{MaxAttempts: 3, Delay: 0}, func(_ context.Context) error {
		calls++
		return nil
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if calls != 1 {
		t.Errorf("expected 1 call, got %d", calls)
	}
}

func TestWithRetry_SuccessOnSecondAttempt(t *testing.T) {
	calls := 0
	sentinel := errors.New("temporary")
	err := WithRetry(context.Background(), RetryConfig{MaxAttempts: 3, Delay: 0}, func(_ context.Context) error {
		calls++
		if calls < 2 {
			return sentinel
		}
		return nil
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if calls != 2 {
		t.Errorf("expected 2 calls, got %d", calls)
	}
}

func TestWithRetry_AllAttemptsFail(t *testing.T) {
	sentinel := errors.New("always fails")
	calls := 0
	err := WithRetry(context.Background(), RetryConfig{MaxAttempts: 3, Delay: 0}, func(_ context.Context) error {
		calls++
		return sentinel
	})
	if !errors.Is(err, sentinel) {
		t.Fatalf("expected sentinel error, got %v", err)
	}
	if calls != 3 {
		t.Errorf("expected 3 calls, got %d", calls)
	}
}

func TestWithRetry_ContextCancelled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	err := WithRetry(ctx, RetryConfig{MaxAttempts: 3, Delay: 0}, func(_ context.Context) error {
		return nil
	})
	if !errors.Is(err, context.Canceled) {
		t.Errorf("expected context.Canceled, got %v", err)
	}
}

func TestWithRetry_ContextCancelledDuringDelay(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	calls := 0
	sentinel := errors.New("fail once")
	err := WithRetry(ctx, RetryConfig{MaxAttempts: 3, Delay: 100 * time.Millisecond}, func(_ context.Context) error {
		calls++
		cancel() // cancel after first failure to interrupt the retry delay
		return sentinel
	})
	if !errors.Is(err, context.Canceled) {
		t.Errorf("expected context.Canceled, got %v", err)
	}
	if calls != 1 {
		t.Errorf("expected 1 call before cancellation, got %d", calls)
	}
}
