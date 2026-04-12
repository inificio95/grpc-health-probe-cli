package probe

import (
	"errors"
	"net"
	"strings"
)

// SanitizeConfig controls address sanitization behavior before dialing.
type SanitizeConfig struct {
	// Enabled controls whether address sanitization is applied.
	Enabled bool

	// StripScheme removes schemes like "grpc://" or "dns://" before dialing.
	StripScheme bool

	// DefaultPort is appended when no port is present in the address.
	DefaultPort string
}

// DefaultSanitizeConfig returns a SanitizeConfig with sensible defaults.
func DefaultSanitizeConfig() *SanitizeConfig {
	return &SanitizeConfig{
		Enabled:     true,
		StripScheme: true,
		DefaultPort: "443",
	}
}

// Validate checks that the SanitizeConfig is valid.
func (c *SanitizeConfig) Validate() error {
	if c == nil {
		return errors.New("sanitize config must not be nil")
	}
	if c.Enabled && c.DefaultPort == "" {
		return errors.New("sanitize config: default port must not be empty when enabled")
	}
	return nil
}

// Sanitize applies address sanitization rules and returns the cleaned address.
func (c *SanitizeConfig) Sanitize(addr string) string {
	if c == nil || !c.Enabled {
		return addr
	}

	if c.StripScheme {
		if idx := strings.Index(addr, "://"); idx != -1 {
			addr = addr[idx+3:]
		}
	}

	// Append default port if missing.
	host, port, err := net.SplitHostPort(addr)
	if err != nil || port == "" {
		// SplitHostPort failed — likely no port present.
		if host == "" {
			host = addr
		}
		addr = net.JoinHostPort(host, c.DefaultPort)
	}

	return addr
}
