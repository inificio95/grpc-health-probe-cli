package probe

import (
	"context"
	"fmt"
	"time"
)

// DefaultTimeoutConfig returns a TimeoutConfig with sensible defaults.
func DefaultTimeoutConfig() TimeoutConfig {
	return TimeoutConfig{
		DialTimeout:    5 * time.Second,
		RequestTimeout: 10 * time.Second,
	}
}

// TimeoutConfig holds timeout settings for the probe.
type TimeoutConfig struct {
	// DialTimeout is the maximum time allowed to establish a connection.
	DialTimeout time.Duration
	// RequestTimeout is the maximum time allowed for a single health check RPC.
	RequestTimeout time.Duration
}

// Validate checks that all timeout values are positive.
func (c TimeoutConfig) Validate() error {
	if c.DialTimeout <= 0 {
		return fmt.Errorf("dial_timeout must be greater than zero, got %s", c.DialTimeout)
	}
	if c.RequestTimeout <= 0 {
		return fmt.Errorf("request_timeout must be greater than zero, got %s", c.RequestTimeout)
	}
	return nil
}

// WithRequestTimeout returns a derived context that is cancelled after
// the configured RequestTimeout duration.
func (c TimeoutConfig) WithRequestTimeout(parent context.Context) (context.Context, context.CancelFunc) {
	return context.WithTimeout(parent, c.RequestTimeout)
}

// WithDialTimeout returns a derived context that is cancelled after
// the configured DialTimeout duration.
func (c TimeoutConfig) WithDialTimeout(parent context.Context) (context.Context, context.CancelFunc) {
	return context.WithTimeout(parent, c.DialTimeout)
}
