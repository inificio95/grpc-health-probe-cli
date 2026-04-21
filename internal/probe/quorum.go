package probe

import "fmt"

// QuorumConfig defines the minimum number of successful health checks
// required before a service is considered healthy.
type QuorumConfig struct {
	Enabled    bool
	MinSuccess int
	Total      int
}

// DefaultQuorumConfig returns a QuorumConfig with quorum disabled.
func DefaultQuorumConfig() *QuorumConfig {
	return &QuorumConfig{
		Enabled:    false,
		MinSuccess: 1,
		Total:      1,
	}
}

// Validate checks that the QuorumConfig fields are consistent.
func (c *QuorumConfig) Validate() error {
	if c == nil {
		return fmt.Errorf("quorum config must not be nil")
	}
	if !c.Enabled {
		return nil
	}
	if c.Total <= 0 {
		return fmt.Errorf("quorum total must be greater than zero, got %d", c.Total)
	}
	if c.MinSuccess <= 0 {
		return fmt.Errorf("quorum min_success must be greater than zero, got %d", c.MinSuccess)
	}
	if c.MinSuccess > c.Total {
		return fmt.Errorf("quorum min_success (%d) must not exceed total (%d)", c.MinSuccess, c.Total)
	}
	return nil
}

// IsMet reports whether the given number of successes satisfies the quorum.
// If quorum is disabled, any non-negative success count is considered met.
func (c *QuorumConfig) IsMet(successes int) bool {
	if c == nil || !c.Enabled {
		return successes >= 1
	}
	return successes >= c.MinSuccess
}

// String returns a human-readable description of the quorum configuration.
func (c *QuorumConfig) String() string {
	if c == nil {
		return "quorum(nil)"
	}
	if !c.Enabled {
		return "quorum(disabled)"
	}
	return fmt.Sprintf("quorum(min=%d/%d)", c.MinSuccess, c.Total)
}
