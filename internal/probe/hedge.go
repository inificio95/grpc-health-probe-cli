package probe

import (
	"errors"
	"fmt"
	"time"
)

// HedgeConfig configures hedged request behavior, where additional
// parallel requests are issued after a delay if the initial request
// has not yet completed.
type HedgeConfig struct {
	Enabled     bool
	Delay       time.Duration
	MaxHedges   int
}

// DefaultHedgeConfig returns a HedgeConfig with hedging disabled.
func DefaultHedgeConfig() *HedgeConfig {
	return &HedgeConfig{
		Enabled:   false,
		Delay:     100 * time.Millisecond,
		MaxHedges: 1,
	}
}

// Validate checks that the HedgeConfig fields are consistent and valid.
func (h *HedgeConfig) Validate() error {
	if h == nil {
		return errors.New("hedge config must not be nil")
	}
	if !h.Enabled {
		return nil
	}
	if h.Delay <= 0 {
		return fmt.Errorf("hedge delay must be positive, got %s", h.Delay)
	}
	if h.MaxHedges < 1 {
		return fmt.Errorf("hedge max_hedges must be at least 1, got %d", h.MaxHedges)
	}
	if h.MaxHedges > 10 {
		return fmt.Errorf("hedge max_hedges must not exceed 10, got %d", h.MaxHedges)
	}
	return nil
}

// String returns a human-readable description of the hedge configuration.
func (h *HedgeConfig) String() string {
	if h == nil {
		return "hedge(nil)"
	}
	if !h.Enabled {
		return "hedge(disabled)"
	}
	return fmt.Sprintf("hedge(delay=%s, max=%d)", h.Delay, h.MaxHedges)
}
