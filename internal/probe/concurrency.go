package probe

import (
	"errors"
	"fmt"
)

// ConcurrencyConfig controls the maximum number of concurrent probe requests.
type ConcurrencyConfig struct {
	Enabled    bool
	MaxWorkers int
}

// DefaultConcurrencyConfig returns a ConcurrencyConfig with sensible defaults.
func DefaultConcurrencyConfig() *ConcurrencyConfig {
	return &ConcurrencyConfig{
		Enabled:    false,
		MaxWorkers: 1,
	}
}

// Validate checks that the ConcurrencyConfig is valid.
func (c *ConcurrencyConfig) Validate() error {
	if c == nil {
		return errors.New("concurrency config must not be nil")
	}
	if !c.Enabled {
		return nil
	}
	if c.MaxWorkers <= 0 {
		return fmt.Errorf("concurrency max_workers must be greater than 0, got %d", c.MaxWorkers)
	}
	if c.MaxWorkers > 256 {
		return fmt.Errorf("concurrency max_workers must not exceed 256, got %d", c.MaxWorkers)
	}
	return nil
}

// WorkerPool returns a buffered channel acting as a semaphore for MaxWorkers slots.
// If concurrency is disabled, a channel of size 1 is returned.
func (c *ConcurrencyConfig) WorkerPool() chan struct{} {
	if c == nil || !c.Enabled {
		ch := make(chan struct{}, 1)
		ch <- struct{}{}
		return ch
	}
	ch := make(chan struct{}, c.MaxWorkers)
	for i := 0; i < c.MaxWorkers; i++ {
		ch <- struct{}{}
	}
	return ch
}

// String returns a human-readable description of the config.
func (c *ConcurrencyConfig) String() string {
	if c == nil {
		return "concurrency(nil)"
	}
	if !c.Enabled {
		return "concurrency(disabled)"
	}
	return fmt.Sprintf("concurrency(max_workers=%d)", c.MaxWorkers)
}
