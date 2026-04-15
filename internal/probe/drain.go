package probe

import (
	"errors"
	"time"
)

// DrainConfig controls graceful drain behaviour before the probe exits.
// When enabled the prober waits up to DrainTimeout for in-flight requests to
// complete before returning results.
type DrainConfig struct {
	Enabled      bool
	DrainTimeout time.Duration
}

// DefaultDrainConfig returns a DrainConfig with sensible defaults.
func DefaultDrainConfig() *DrainConfig {
	return &DrainConfig{
		Enabled:      false,
		DrainTimeout: 5 * time.Second,
	}
}

// Validate checks that the DrainConfig is self-consistent.
func (c *DrainConfig) Validate() error {
	if c == nil {
		return errors.New("drain config must not be nil")
	}
	if !c.Enabled {
		return nil
	}
	if c.DrainTimeout <= 0 {
		return errors.New("drain timeout must be a positive duration when drain is enabled")
	}
	return nil
}

// Wait blocks for at most DrainTimeout when drain is enabled.
// done should be closed (or receive) once all in-flight work has finished.
// Returns true if the drain completed cleanly, false if it timed out.
func (c *DrainConfig) Wait(done <-chan struct{}) bool {
	if c == nil || !c.Enabled {
		return true
	}
	select {
	case <-done:
		return true
	case <-time.After(c.DrainTimeout):
		return false
	}
}
