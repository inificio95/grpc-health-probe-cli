package probe

import (
	"context"
	"time"
)

// RetryConfig holds configuration for retry behaviour.
type RetryConfig struct {
	// MaxAttempts is the total number of attempts (including the first).
	MaxAttempts int
	// Delay is the duration to wait between attempts.
	Delay time.Duration
}

// DefaultRetryConfig returns a RetryConfig with sensible defaults.
func DefaultRetryConfig() RetryConfig {
	return RetryConfig{
		MaxAttempts: 3,
		Delay:       500 * time.Millisecond,
	}
}

// Validate returns an error if the RetryConfig is invalid.
func (r RetryConfig) Validate() error {
	if r.MaxAttempts < 1 {
		return ErrInvalidRetryAttempts
	}
	if r.Delay < 0 {
		return ErrInvalidRetryDelay
	}
	return nil
}

// ErrInvalidRetryAttempts is returned when MaxAttempts is less than 1.
var ErrInvalidRetryAttempts = errString("retry max attempts must be at least 1")

// ErrInvalidRetryDelay is returned when Delay is negative.
var ErrInvalidRetryDelay = errString("retry delay must be non-negative")

// errString is a simple error type backed by a string.
type errString string

func (e errString) Error() string { return string(e) }

// WithRetry executes fn up to cfg.MaxAttempts times, waiting cfg.Delay between
// attempts. It returns the last error if all attempts fail, or nil on success.
func WithRetry(ctx context.Context, cfg RetryConfig, fn func(ctx context.Context) error) error {
	var lastErr error
	for attempt := 0; attempt < cfg.MaxAttempts; attempt++ {
		if err := ctx.Err(); err != nil {
			return err
		}
		lastErr = fn(ctx)
		if lastErr == nil {
			return nil
		}
		if attempt < cfg.MaxAttempts-1 {
			select {
			case <-time.After(cfg.Delay):
			case <-ctx.Done():
				return ctx.Err()
			}
		}
	}
	return lastErr
}
