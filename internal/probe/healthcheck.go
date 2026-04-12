package probe

import (
	"errors"
	"time"
)

// HealthCheckConfig controls per-request health check behaviour such as
// whether to include the service name in the request and how long to wait
// before considering the check timed out at the application level.
type HealthCheckConfig struct {
	// Enabled controls whether health-check specific logic is applied.
	Enabled bool

	// IncludeServiceName appends the service name to the health check request.
	IncludeServiceName bool

	// CheckInterval is the minimum duration between consecutive health checks
	// when running in watch mode.
	CheckInterval time.Duration

	// MaxConsecutiveFailures is the number of consecutive failures before the
	// prober reports a hard failure. 0 means no limit.
	MaxConsecutiveFailures int
}

// DefaultHealthCheckConfig returns a HealthCheckConfig populated with
// sensible defaults.
func DefaultHealthCheckConfig() *HealthCheckConfig {
	return &HealthCheckConfig{
		Enabled:                true,
		IncludeServiceName:     true,
		CheckInterval:          5 * time.Second,
		MaxConsecutiveFailures: 0,
	}
}

// Validate returns an error if the HealthCheckConfig contains invalid values.
func (c *HealthCheckConfig) Validate() error {
	if c == nil {
		return errors.New("healthcheck config must not be nil")
	}
	if c.CheckInterval < 0 {
		return errors.New("healthcheck check_interval must be non-negative")
	}
	if c.MaxConsecutiveFailures < 0 {
		return errors.New("healthcheck max_consecutive_failures must be non-negative")
	}
	return nil
}

// IsLimited reports whether a maximum consecutive failure limit is set.
func (c *HealthCheckConfig) IsLimited() bool {
	return c != nil && c.MaxConsecutiveFailures > 0
}
