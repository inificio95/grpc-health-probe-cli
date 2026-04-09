package probe

import (
	"errors"
	"time"
)

// WatchMode defines how the prober should behave in watch mode.
type WatchMode int

const (
	// WatchModeDisabled means no watching; single-shot probe.
	WatchModeDisabled WatchMode = iota
	// WatchModePolling means poll the endpoint at a fixed interval.
	WatchModePolling
)

// WatchConfig holds configuration for watch/polling mode.
type WatchConfig struct {
	// Mode controls whether watching is enabled.
	Mode WatchMode
	// Interval is the duration between successive health checks.
	Interval time.Duration
	// MaxChecks is the maximum number of checks to perform (0 = unlimited).
	MaxChecks int
}

// DefaultWatchConfig returns a WatchConfig with sensible defaults.
func DefaultWatchConfig() *WatchConfig {
	return &WatchConfig{
		Mode:      WatchModeDisabled,
		Interval:  5 * time.Second,
		MaxChecks: 0,
	}
}

// Validate checks that the WatchConfig fields are consistent and valid.
func (w *WatchConfig) Validate() error {
	if w == nil {
		return errors.New("watch config must not be nil")
	}
	if w.Mode != WatchModeDisabled && w.Mode != WatchModePolling {
		return errors.New("watch config: unknown watch mode")
	}
	if w.Mode == WatchModePolling && w.Interval <= 0 {
		return errors.New("watch config: interval must be positive when polling is enabled")
	}
	if w.MaxChecks < 0 {
		return errors.New("watch config: max checks must be non-negative")
	}
	return nil
}

// IsEnabled returns true when watch mode is active.
func (w *WatchConfig) IsEnabled() bool {
	return w != nil && w.Mode == WatchModePolling
}
