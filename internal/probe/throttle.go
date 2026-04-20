package probe

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

// ThrottleConfig controls request throttling to prevent overwhelming a target service.
type ThrottleConfig struct {
	Enabled     bool
	MinInterval time.Duration // minimum duration between consecutive probe calls
	Burst       int           // number of requests allowed before throttling kicks in
}

// DefaultThrottleConfig returns a ThrottleConfig with throttling disabled.
func DefaultThrottleConfig() *ThrottleConfig {
	return &ThrottleConfig{
		Enabled:     false,
		MinInterval: 500 * time.Millisecond,
		Burst:       1,
	}
}

// Validate checks that the ThrottleConfig fields are consistent and valid.
func (c *ThrottleConfig) Validate() error {
	if c == nil {
		return errors.New("throttle config must not be nil")
	}
	if !c.Enabled {
		return nil
	}
	if c.MinInterval <= 0 {
		return fmt.Errorf("throttle min_interval must be positive, got %s", c.MinInterval)
	}
	if c.Burst < 1 {
		return fmt.Errorf("throttle burst must be at least 1, got %d", c.Burst)
	}
	return nil
}

// Throttler enforces the minimum interval between probe invocations.
type Throttler struct {
	cfg      *ThrottleConfig
	mu       sync.Mutex
	lastCall time.Time
	tokens   int
}

// NewThrottler creates a Throttler from the given config.
func NewThrottler(cfg *ThrottleConfig) *Throttler {
	if cfg == nil {
		cfg = DefaultThrottleConfig()
	}
	return &Throttler{
		cfg:    cfg,
		tokens: cfg.Burst,
	}
}

// Wait blocks until the throttle allows the next probe call.
// Returns immediately if throttling is disabled.
func (t *Throttler) Wait() {
	if !t.cfg.Enabled {
		return
	}
	t.mu.Lock()
	defer t.mu.Unlock()

	now := time.Now()
	if t.tokens > 0 {
		t.tokens--
		t.lastCall = now
		return
	}
	elapsed := now.Sub(t.lastCall)
	if elapsed < t.cfg.MinInterval {
		time.Sleep(t.cfg.MinInterval - elapsed)
	}
	t.lastCall = time.Now()
}
