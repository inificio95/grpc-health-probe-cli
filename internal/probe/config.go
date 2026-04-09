package probe

import (
	"errors"
	"time"
)

// Config holds all parameters for a health probe run.
type Config struct {
	// Address is the host:port of the gRPC server.
	Address string
	// Service is the gRPC service name to check. Empty string checks overall server health.
	Service string
	// Timeout is the per-attempt deadline.
	Timeout time.Duration
	// MaxRetries is the number of additional attempts after the first failure.
	MaxRetries int
	// RetryInterval is the wait time between retry attempts.
	RetryInterval time.Duration
	// UserAgent is the gRPC user-agent header value.
	UserAgent string
}

// DefaultConfig returns a Config populated with sensible defaults.
func DefaultConfig() Config {
	return Config{
		Address:       "localhost:50051",
		Service:       "",
		Timeout:       5 * time.Second,
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		UserAgent:     "grpc-health-probe-cli",
	}
}

// Validate checks that the Config has all required fields set correctly.
func (c Config) Validate() error {
	if c.Address == "" {
		return errors.New("address must not be empty")
	}
	if c.Timeout <= 0 {
		return errors.New("timeout must be positive")
	}
	if c.MaxRetries < 0 {
		return errors.New("max-retries must be non-negative")
	}
	if c.RetryInterval < 0 {
		return errors.New("retry-interval must be non-negative")
	}
	return nil
}
