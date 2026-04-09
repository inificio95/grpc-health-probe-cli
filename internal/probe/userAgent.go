package probe

import (
	"fmt"
	"runtime"
)

const (
	defaultAppName    = "grpc-health-probe-cli"
	defaultAppVersion = "0.1.0"
)

// UserAgentConfig holds configuration for the gRPC user-agent header.
type UserAgentConfig struct {
	// AppName is the application name included in the user-agent string.
	AppName string
	// AppVersion is the application version included in the user-agent string.
	AppVersion string
	// Disabled suppresses the custom user-agent header when true.
	Disabled bool
}

// DefaultUserAgentConfig returns a UserAgentConfig with sensible defaults.
func DefaultUserAgentConfig() *UserAgentConfig {
	return &UserAgentConfig{
		AppName:    defaultAppName,
		AppVersion: defaultAppVersion,
		Disabled:   false,
	}
}

// Validate checks that the UserAgentConfig is consistent.
func (c *UserAgentConfig) Validate() error {
	if c == nil {
		return fmt.Errorf("user-agent config must not be nil")
	}
	if !c.Disabled {
		if c.AppName == "" {
			return fmt.Errorf("user-agent app name must not be empty")
		}
		if c.AppVersion == "" {
			return fmt.Errorf("user-agent app version must not be empty")
		}
	}
	return nil
}

// String returns the formatted user-agent string.
func (c *UserAgentConfig) String() string {
	if c == nil || c.Disabled {
		return ""
	}
	return fmt.Sprintf("%s/%s (%s/%s)", c.AppName, c.AppVersion, runtime.GOOS, runtime.GOARCH)
}
