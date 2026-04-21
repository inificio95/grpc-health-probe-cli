package probe

import (
	"context"
	"errors"
	"time"
)

// DefaultSemaphoreConfig returns a SemaphoreConfig with sensible defaults.
func DefaultSemaphoreConfig() *SemaphoreConfig {
	return &SemaphoreConfig{
		Enabled:    false,
		MaxTickets: 10,
		AcquireTimeout: 5 * time.Second,
	}
}

// SemaphoreConfig controls concurrency via a semaphore with a bounded ticket pool.
type SemaphoreConfig struct {
	Enabled        bool
	MaxTickets     int
	AcquireTimeout time.Duration

	ch chan struct{}
}

// Validate returns an error if the SemaphoreConfig is invalid.
func (c *SemaphoreConfig) Validate() error {
	if c == nil {
		return errors.New("semaphore config must not be nil")
	}
	if !c.Enabled {
		return nil
	}
	if c.MaxTickets <= 0 {
		return errors.New("semaphore max tickets must be greater than zero")
	}
	if c.AcquireTimeout <= 0 {
		return errors.New("semaphore acquire timeout must be greater than zero")
	}
	return nil
}

// Init initialises the internal channel used as the semaphore.
// It must be called before Acquire/Release when Enabled is true.
func (c *SemaphoreConfig) Init() {
	if c == nil || !c.Enabled {
		return
	}
	c.ch = make(chan struct{}, c.MaxTickets)
}

// Acquire attempts to obtain a ticket within AcquireTimeout.
// Returns an error if the context is cancelled or the timeout elapses.
func (c *SemaphoreConfig) Acquire(ctx context.Context) error {
	if c == nil || !c.Enabled {
		return nil
	}
	timeoutCtx, cancel := context.WithTimeout(ctx, c.AcquireTimeout)
	defer cancel()
	select {
	case c.ch <- struct{}{}:
		return nil
	case <-timeoutCtx.Done():
		return errors.New("semaphore: timed out waiting to acquire ticket")
	}
}

// Release returns a ticket to the pool.
func (c *SemaphoreConfig) Release() {
	if c == nil || !c.Enabled {
		return
	}
	select {
	case <-c.ch:
	default:
	}
}
