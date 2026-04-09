package probe

import "fmt"

// Config holds all configuration for the Prober.
type Config struct {
	// Address is the host:port of the gRPC server to probe.
	Address string
	// Service is the gRPC service name to check (empty means server-level check).
	Service string
	// Timeout configures dial and per-request timeouts.
	Timeout TimeoutConfig
	// Retry configures retry behaviour.
	Retry RetryConfig
	// TLS configures transport-layer security.
	TLS TLSConfig
	// Format is the output format: "text" or "json".
	Format string
}

// DefaultConfig returns a Config populated with sensible defaults.
func DefaultConfig() Config {
	return Config{
		Timeout: DefaultTimeoutConfig(),
		Retry:   DefaultRetryConfig(),
		TLS:     TLSConfig{Enabled: false},
		Format:  "text",
	}
}

// Validate checks that the Config is consistent and complete.
func (c *Config) Validate() error {
	if c.Address == "" {
		return fmt.Errorf("config: address must not be empty")
	}
	if c.Format != "text" && c.Format != "json" {
		return fmt.Errorf("config: format must be \"text\" or \"json\", got %q", c.Format)
	}
	if err := c.Timeout.Validate(); err != nil {
		return fmt.Errorf("config: %w", err)
	}
	if err := c.Retry.Validate(); err != nil {
		return fmt.Errorf("config: %w", err)
	}
	if err := c.TLS.Validate(); err != nil {
		return fmt.Errorf("config: %w", err)
	}
	return nil
}
