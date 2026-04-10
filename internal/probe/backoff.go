package probe

import (
	"math"
	"time"
)

// BackoffStrategy defines how retry delays are calculated.
type BackoffStrategy string

const (
	// BackoffFixed uses a constant delay between retries.
	BackoffFixed BackoffStrategy = "fixed"
	// BackoffExponential doubles the delay on each retry attempt.
	BackoffExponential BackoffStrategy = "exponential"
	// BackoffLinear increases the delay linearly on each retry attempt.
	BackoffLinear BackoffStrategy = "linear"
)

// BackoffConfig holds configuration for retry backoff behavior.
type BackoffConfig struct {
	Strategy    BackoffStrategy
	InitialDelay time.Duration
	MaxDelay     time.Duration
	Multiplier   float64
}

// DefaultBackoffConfig returns a BackoffConfig with sensible defaults.
func DefaultBackoffConfig() *BackoffConfig {
	return &BackoffConfig{
		Strategy:     BackoffFixed,
		InitialDelay: 500 * time.Millisecond,
		MaxDelay:     30 * time.Second,
		Multiplier:   2.0,
	}
}

// Validate checks that the BackoffConfig fields are valid.
func (c *BackoffConfig) Validate() error {
	if c == nil {
		return ErrNilConfig
	}
	switch c.Strategy {
	case BackoffFixed, BackoffExponential, BackoffLinear:
	default:
		return ErrInvalidBackoffStrategy
	}
	if c.InitialDelay <= 0 {
		return ErrInvalidInitialDelay
	}
	if c.MaxDelay < c.InitialDelay {
		return ErrMaxDelayTooSmall
	}
	if c.Strategy == BackoffExponential && c.Multiplier <= 1.0 {
		return ErrInvalidMultiplier
	}
	return nil
}

// Delay returns the computed delay for the given attempt (0-indexed).
func (c *BackoffConfig) Delay(attempt int) time.Duration {
	var d time.Duration
	switch c.Strategy {
	case BackoffExponential:
		d = time.Duration(float64(c.InitialDelay) * math.Pow(c.Multiplier, float64(attempt)))
	case BackoffLinear:
		d = c.InitialDelay * time.Duration(attempt+1)
	default:
		d = c.InitialDelay
	}
	if d > c.MaxDelay {
		return c.MaxDelay
	}
	return d
}
