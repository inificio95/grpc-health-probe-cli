package probe

import (
	"errors"
	"time"
)

// DefaultWarmupConfig returns a WarmupConfig with warmup disabled.
func DefaultWarmupConfig() *WarmupConfig {
	return &WarmupConfig{
		Enabled:  false,
		Delay:    2 * time.Second,
		MaxDelay: 30 * time.Second,
	}
}

// WarmupConfig controls an initial delay before the first health probe
// is issued, allowing a service time to become ready.
type WarmupConfig struct {
	// Enabled controls whether warmup delay is applied.
	Enabled bool

	// Delay is the initial warmup duration to wait before probing.
	Delay time.Duration

	// MaxDelay caps the warmup delay to prevent excessive waits.
	MaxDelay time.Duration
}

// Validate checks that the WarmupConfig has consistent values.
func (c *WarmupConfig) Validate() error {
	if c == nil {
		return errors.New("warmup config must not be nil")
	}
	if !c.Enabled {
		return nil
	}
	if c.Delay <= 0 {
		return errors.New("warmup delay must be positive when enabled")
	}
	if c.MaxDelay <= 0 {
		return errors.New("warmup max delay must be positive when enabled")
	}
	if c.Delay > c.MaxDelay {
		return errors.New("warmup delay must not exceed max delay")
	}
	return nil
}

// EffectiveDelay returns the delay that should be applied, clamped to MaxDelay.
// Returns zero when warmup is disabled.
func (c *WarmupConfig) EffectiveDelay() time.Duration {
	if c == nil || !c.Enabled {
		return 0
	}
	if c.Delay > c.MaxDelay {
		return c.MaxDelay
	}
	return c.Delay
}
