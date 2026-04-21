package probe

import (
	"errors"
	"time"
)

// CheckpointConfig controls periodic checkpoint persistence of probe state.
// When enabled, the prober writes a lightweight status file at a configurable
// interval so external tooling can observe liveness without polling gRPC.
type CheckpointConfig struct {
	Enabled  bool
	Path     string
	Interval time.Duration
	Format   string // "text" or "json"
}

// DefaultCheckpointConfig returns a CheckpointConfig with checkpointing disabled.
func DefaultCheckpointConfig() *CheckpointConfig {
	return &CheckpointConfig{
		Enabled:  false,
		Path:     "",
		Interval: 30 * time.Second,
		Format:   "text",
	}
}

// Validate returns an error if the CheckpointConfig is invalid.
func (c *CheckpointConfig) Validate() error {
	if c == nil {
		return errors.New("checkpoint config must not be nil")
	}
	if !c.Enabled {
		return nil
	}
	if c.Path == "" {
		return errors.New("checkpoint path must not be empty when checkpointing is enabled")
	}
	if c.Interval <= 0 {
		return errors.New("checkpoint interval must be a positive duration")
	}
	if c.Format != "text" && c.Format != "json" {
		return errors.New("checkpoint format must be \"text\" or \"json\"")
	}
	return nil
}

// IsEnabled returns true when checkpointing is active and the config is valid.
func (c *CheckpointConfig) IsEnabled() bool {
	return c != nil && c.Enabled
}
