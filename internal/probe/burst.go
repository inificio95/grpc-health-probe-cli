package probe

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

// DefaultBurstConfig returns a BurstConfig with sensible defaults.
func DefaultBurstConfig() *BurstConfig {
	return &BurstConfig{
		Enabled:       false,
		MaxBurst:      5,
		BurstInterval: 500 * time.Millisecond,
		Cooldown:      2 * time.Second,
	}
}

// BurstConfig controls burst-detection and suppression for health probes.
type BurstConfig struct {
	Enabled       bool
	MaxBurst      int
	BurstInterval time.Duration
	Cooldown      time.Duration

	mu       sync.Mutex
	times    []time.Time
	coolingUntil time.Time
}

// Validate checks that the BurstConfig is valid.
func (c *BurstConfig) Validate() error {
	if c == nil {
		return errors.New("burst config must not be nil")
	}
	if !c.Enabled {
		return nil
	}
	if c.MaxBurst <= 0 {
		return fmt.Errorf("burst max_burst must be > 0, got %d", c.MaxBurst)
	}
	if c.BurstInterval <= 0 {
		return fmt.Errorf("burst interval must be > 0, got %s", c.BurstInterval)
	}
	if c.Cooldown <= 0 {
		return fmt.Errorf("burst cooldown must be > 0, got %s", c.Cooldown)
	}
	return nil
}

// Allow reports whether the current probe attempt should be allowed through
// (i.e. not suppressed due to burst detection). It is safe for concurrent use.
func (c *BurstConfig) Allow() bool {
	if c == nil || !c.Enabled {
		return true
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	now := time.Now()
	if now.Before(c.coolingUntil) {
		return false
	}
	cutoff := now.Add(-c.BurstInterval)
	filtered := c.times[:0]
	for _, t := range c.times {
		if t.After(cutoff) {
			filtered = append(filtered, t)
		}
	}
	c.times = filtered
	c.times = append(c.times, now)
	if len(c.times) > c.MaxBurst {
		c.coolingUntil = now.Add(c.Cooldown)
		c.times = nil
		return false
	}
	return true
}
