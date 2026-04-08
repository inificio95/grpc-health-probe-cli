package probe

import (
	"errors"
	"time"
)

// Config holds the configuration for a gRPC health probe check.
type Config struct {
	// Address is the gRPC server address in host:port format.
	Address string

	// Service is the gRPC service name to check. Empty string checks the overall server health.
	Service string

	// Timeout is the per-attempt connection and RPC timeout.
	Timeout time.Duration

	// RetryMax is the maximum number of retry attempts (0 means no retries).
	RetryMax int

	// RetryInterval is the duration to wait between retry attempts.
	RetryInterval time.Duration

	// TLS indicates whether to use TLS for the connection.
	TLS bool

	// TLSNoVerify skips TLS certificate verification when TLS is enabled.
	TLSNoVerify bool

	// TLSCACert is the path to a CA certificate file for TLS verification.
	TLSCACert string
}

// DefaultConfig returns a Config with sensible defaults.
func DefaultConfig() Config {
	return Config{
		Timeout:       5 * time.Second,
		RetryMax:      0,
		RetryInterval: 1 * time.Second,
	}
}

// Validate checks that the Config fields are valid.
func (c Config) Validate() error {
	if c.Address == "" {
		return errors.New("address must not be empty")
	}
	if c.Timeout <= 0 {
		return errors.New("timeout must be greater than zero")
	}
	if c.RetryMax < 0 {
		return errors.New("retry-max must be non-negative")
	}
	if c.RetryInterval < 0 {
		return errors.New("retry-interval must be non-negative")
	}
	if c.TLSNoVerify && !c.TLS {
		return errors.New("tls-no-verify requires tls to be enabled")
	}
	if c.TLSCACert != "" && !c.TLS {
		return errors.New("tls-ca-cert requires tls to be enabled")
	}
	return nil
}
