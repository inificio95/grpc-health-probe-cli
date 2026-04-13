package probe

import (
	"errors"
	"fmt"
)

// SNIConfig holds configuration for TLS Server Name Indication override.
type SNIConfig struct {
	Enabled    bool
	ServerName string
}

// DefaultSNIConfig returns a default SNIConfig with SNI disabled.
func DefaultSNIConfig() *SNIConfig {
	return &SNIConfig{
		Enabled:    false,
		ServerName: "",
	}
}

// Validate checks that the SNIConfig is valid.
func (c *SNIConfig) Validate() error {
	if c == nil {
		return errors.New("sni config must not be nil")
	}
	if c.Enabled && c.ServerName == "" {
		return errors.New("sni server name must not be empty when sni is enabled")
	}
	return nil
}

// String returns a human-readable representation of the SNIConfig.
func (c *SNIConfig) String() string {
	if c == nil {
		return "SNIConfig(nil)"
	}
	if !c.Enabled {
		return "SNIConfig(disabled)"
	}
	return fmt.Sprintf("SNIConfig(server_name=%s)", c.ServerName)
}
