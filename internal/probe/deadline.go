package probe

import (
	"context"
	"errors"
	"fmt"
	"time"
)

// DeadlineConfig holds configuration for overall probe deadline enforcement.
type DeadlineConfig struct {
	// Enabled controls whether a global deadline is applied to the probe.
	Enabled bool

	// Duration is the maximum total time allowed for the probe to complete.
	Duration time.Duration
}

// DefaultDeadlineConfig returns a DeadlineConfig with sensible defaults.
func DefaultDeadlineConfig() *DeadlineConfig {
	return &DeadlineConfig{
		Enabled:  false,
		Duration: 30 * time.Second,
	}
}

// Validate checks that the DeadlineConfig is valid.
func (c *DeadlineConfig) Validate() error {
	if c == nil {
		return errors.New("deadline config must not be nil")
	}
	if c.Enabled && c.Duration <= 0 {
		return fmt.Errorf("deadline duration must be positive when enabled, got %v", c.Duration)
	}
	return nil
}

// Apply wraps the given context with a deadline if enabled.
// It returns the (possibly wrapped) context and a cancel function.
func (c *DeadlineConfig) Apply(ctx context.Context) (context.Context, context.CancelFunc) {
	if c == nil || !c.Enabled {
		return ctx, func() {}
	}
	return context.WithTimeout(ctx, c.Duration)
}
