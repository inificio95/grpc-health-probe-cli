package probe

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

// CircuitState represents the state of the circuit breaker.
type CircuitState int

const (
	CircuitClosed CircuitState = iota
	CircuitOpen
	CircuitHalfOpen
)

// DefaultCircuitBreakerConfig returns a CircuitBreakerConfig with sensible defaults.
func DefaultCircuitBreakerConfig() *CircuitBreakerConfig {
	return &CircuitBreakerConfig{
		Enabled:          false,
		MaxFailures:      5,
		OpenDuration:     30 * time.Second,
		HalfOpenRequests: 1,
	}
}

// CircuitBreakerConfig configures the circuit breaker behaviour.
type CircuitBreakerConfig struct {
	Enabled          bool
	MaxFailures      int
	OpenDuration     time.Duration
	HalfOpenRequests int

	mu           sync.Mutex
	state        CircuitState
	failureCount int
	openedAt     time.Time
	halfOpenSent int
}

// Validate checks that the CircuitBreakerConfig is valid.
func (c *CircuitBreakerConfig) Validate() error {
	if c == nil {
		return errors.New("circuit breaker config must not be nil")
	}
	if !c.Enabled {
		return nil
	}
	if c.MaxFailures <= 0 {
		return fmt.Errorf("max_failures must be > 0, got %d", c.MaxFailures)
	}
	if c.OpenDuration <= 0 {
		return fmt.Errorf("open_duration must be > 0, got %s", c.OpenDuration)
	}
	if c.HalfOpenRequests <= 0 {
		return fmt.Errorf("half_open_requests must be > 0, got %d", c.HalfOpenRequests)
	}
	return nil
}

// Allow reports whether the circuit breaker permits a request.
func (c *CircuitBreakerConfig) Allow() bool {
	if !c.Enabled {
		return true
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	switch c.state {
	case CircuitOpen:
		if time.Since(c.openedAt) >= c.OpenDuration {
			c.state = CircuitHalfOpen
			c.halfOpenSent = 0
			return c.allowHalfOpen()
		}
		return false
	case CircuitHalfOpen:
		return c.allowHalfOpen()
	default:
		return true
	}
}

func (c *CircuitBreakerConfig) allowHalfOpen() bool {
	if c.halfOpenSent < c.HalfOpenRequests {
		c.halfOpenSent++
		return true
	}
	return false
}

// RecordSuccess records a successful request.
func (c *CircuitBreakerConfig) RecordSuccess() {
	if !c.Enabled {
		return
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	c.failureCount = 0
	c.state = CircuitClosed
}

// RecordFailure records a failed request.
func (c *CircuitBreakerConfig) RecordFailure() {
	if !c.Enabled {
		return
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	c.failureCount++
	if c.state == CircuitHalfOpen || c.failureCount >= c.MaxFailures {
		c.state = CircuitOpen
		c.openedAt = time.Now()
	}
}

// State returns the current circuit state.
func (c *CircuitBreakerConfig) State() CircuitState {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.state
}
