package probe

import (
	"errors"
	"time"
)

// WindowConfig defines a sliding window for tracking probe results over time.
type WindowConfig struct {
	Enabled  bool
	Size     int
	Duration time.Duration
}

// DefaultWindowConfig returns a WindowConfig with sensible defaults.
func DefaultWindowConfig() *WindowConfig {
	return &WindowConfig{
		Enabled:  false,
		Size:     10,
		Duration: 30 * time.Second,
	}
}

// Validate checks that the WindowConfig is valid.
func (c *WindowConfig) Validate() error {
	if c == nil {
		return errors.New("window config must not be nil")
	}
	if !c.Enabled {
		return nil
	}
	if c.Size <= 0 {
		return errors.New("window size must be greater than zero")
	}
	if c.Duration <= 0 {
		return errors.New("window duration must be greater than zero")
	}
	return nil
}

// Window tracks a fixed-size sliding window of boolean probe outcomes.
type Window struct {
	cfg     *WindowConfig
	results []bool
	pos     int
	count   int
}

// NewWindow creates a new Window from the given config.
func NewWindow(cfg *WindowConfig) *Window {
	if cfg == nil || !cfg.Enabled {
		return &Window{cfg: cfg}
	}
	return &Window{
		cfg:     cfg,
		results: make([]bool, cfg.Size),
	}
}

// Record adds a probe outcome to the window.
func (w *Window) Record(success bool) {
	if w.cfg == nil || !w.cfg.Enabled {
		return
	}
	w.results[w.pos] = success
	w.pos = (w.pos + 1) % w.cfg.Size
	if w.count < w.cfg.Size {
		w.count++
	}
}

// SuccessRate returns the fraction of successful probes in the window.
// Returns 1.0 if no results have been recorded yet.
func (w *Window) SuccessRate() float64 {
	if w.count == 0 {
		return 1.0
	}
	var successes int
	for i := 0; i < w.count; i++ {
		if w.results[i] {
			successes++
		}
	}
	return float64(successes) / float64(w.count)
}

// Count returns the number of results currently recorded in the window.
func (w *Window) Count() int {
	return w.count
}
