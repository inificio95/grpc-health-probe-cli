package probe

import (
	"errors"
	"time"
)

// RateLimitConfig configures request rate limiting for the health probe.
type RateLimitConfig struct {
	Enabled      bool
	MaxRequests  int
	WindowSize   time.Duration
	lastWindow   time.Time
	requestCount int
}

// DefaultRateLimitConfig returns a RateLimitConfig with rate limiting disabled.
func DefaultRateLimitConfig() *RateLimitConfig {
	return &RateLimitConfig{
		Enabled:     false,
		MaxRequests: 10,
		WindowSize:  time.Second,
	}
}

// Validate checks that the RateLimitConfig is valid.
func (r *RateLimitConfig) Validate() error {
	if r == nil {
		return errors.New("rate limit config must not be nil")
	}
	if !r.Enabled {
		return nil
	}
	if r.MaxRequests <= 0 {
		return errors.New("rate limit max requests must be greater than zero")
	}
	if r.WindowSize <= 0 {
		return errors.New("rate limit window size must be greater than zero")
	}
	return nil
}

// Allow returns true if a request is permitted under the current rate limit.
// It resets the window counter when the window has elapsed.
func (r *RateLimitConfig) Allow() bool {
	if r == nil || !r.Enabled {
		return true
	}
	now := time.Now()
	if now.Sub(r.lastWindow) >= r.WindowSize {
		r.lastWindow = now
		r.requestCount = 0
	}
	if r.requestCount >= r.MaxRequests {
		return false
	}
	r.requestCount++
	return true
}

// Remaining returns the number of requests remaining in the current window.
func (r *RateLimitConfig) Remaining() int {
	if r == nil || !r.Enabled {
		return -1
	}
	remaining := r.MaxRequests - r.requestCount
	if remaining < 0 {
		return 0
	}
	return remaining
}
