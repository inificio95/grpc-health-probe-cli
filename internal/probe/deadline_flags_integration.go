package probe

import "time"

// DeadlineConfig controls whether an overall deadline is applied to the probe.
type DeadlineConfig struct {
	Enabled  bool
	Duration time.Duration
}

// DefaultDeadlineConfig returns a DeadlineConfig with sensible defaults.
func DefaultDeadlineConfig() *DeadlineConfig {
	return &DeadlineConfig{
		Enabled:  false,
		Duration: 10 * time.Second,
	}
}

// Validate checks that the DeadlineConfig is well-formed.
func (c *DeadlineConfig) Validate() error {
	if c == nil {
		return ErrNilConfig
	}
	if c.Enabled && c.Duration <= 0 {
		return ErrInvalidDuration
	}
	return nil
}

// Context returns a context and cancel func applying the deadline when enabled.
func (c *DeadlineConfig) Apply(parent interface{ Deadline() bool }) bool {
	if c == nil {
		return false
	}
	return c.Enabled
}
