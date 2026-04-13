package probe

import (
	"errors"
	"time"
)

// TimeoutBudgetConfig controls a shared deadline budget across retry attempts.
type TimeoutBudgetConfig struct {
	Enabled       bool
	TotalBudget   time.Duration
	ReserveBuffer time.Duration
}

// DefaultTimeoutBudgetConfig returns a TimeoutBudgetConfig with sensible defaults.
func DefaultTimeoutBudgetConfig() *TimeoutBudgetConfig {
	return &TimeoutBudgetConfig{
		Enabled:       false,
		TotalBudget:   30 * time.Second,
		ReserveBuffer: 500 * time.Millisecond,
	}
}

// Validate checks that the TimeoutBudgetConfig is well-formed.
func (c *TimeoutBudgetConfig) Validate() error {
	if c == nil {
		return errors.New("timeout budget config must not be nil")
	}
	if !c.Enabled {
		return nil
	}
	if c.TotalBudget <= 0 {
		return errors.New("timeout budget: total budget must be positive")
	}
	if c.ReserveBuffer < 0 {
		return errors.New("timeout budget: reserve buffer must not be negative")
	}
	if c.ReserveBuffer >= c.TotalBudget {
		return errors.New("timeout budget: reserve buffer must be less than total budget")
	}
	return nil
}

// Budget tracks remaining time within a shared deadline budget.
type Budget struct {
	deadline time.Time
	buffer   time.Duration
}

// NewBudget creates a Budget from the given config, starting now.
func NewBudget(cfg *TimeoutBudgetConfig) *Budget {
	return &Budget{
		deadline: time.Now().Add(cfg.TotalBudget),
		buffer:   cfg.ReserveBuffer,
	}
}

// Remaining returns how much usable time is left (total minus buffer).
func (b *Budget) Remaining() time.Duration {
	usable := time.Until(b.deadline) - b.buffer
	if usable < 0 {
		return 0
	}
	return usable
}

// Exhausted reports whether the usable budget has been consumed.
func (b *Budget) Exhausted() bool {
	return b.Remaining() == 0
}
