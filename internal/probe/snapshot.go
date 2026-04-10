package probe

import (
	"fmt"
	"time"
)

// SnapshotConfig controls whether probe results are persisted to a file.
type SnapshotConfig struct {
	Enabled  bool
	FilePath string
	Format   string // "text" or "json"
}

// DefaultSnapshotConfig returns a SnapshotConfig with snapshots disabled.
func DefaultSnapshotConfig() *SnapshotConfig {
	return &SnapshotConfig{
		Enabled:  false,
		FilePath: "",
		Format:   "json",
	}
}

// Validate checks the SnapshotConfig for logical errors.
func (c *SnapshotConfig) Validate() error {
	if c == nil {
		return fmt.Errorf("snapshot config must not be nil")
	}
	if !c.Enabled {
		return nil
	}
	if c.FilePath == "" {
		return fmt.Errorf("snapshot file path must not be empty when snapshots are enabled")
	}
	if c.Format != "text" && c.Format != "json" {
		return fmt.Errorf("snapshot format %q is invalid: must be \"text\" or \"json\"", c.Format)
	}
	return nil
}

// SnapshotEntry represents a single persisted probe result.
type SnapshotEntry struct {
	Timestamp time.Time `json:"timestamp"`
	Address   string    `json:"address"`
	Service   string    `json:"service"`
	Status    string    `json:"status"`
	LatencyMs int64     `json:"latency_ms"`
	Error     string    `json:"error,omitempty"`
}
