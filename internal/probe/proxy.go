package probe

import (
	"errors"
	"fmt"
	"net/url"

	"google.golang.org/grpc"
)

// ProxyConfig holds configuration for routing gRPC calls through an HTTP/HTTPS proxy.
type ProxyConfig struct {
	Enabled  bool
	ProxyURL string
}

// DefaultProxyConfig returns a ProxyConfig with proxy disabled.
func DefaultProxyConfig() *ProxyConfig {
	return &ProxyConfig{
		Enabled:  false,
		ProxyURL: "",
	}
}

// Validate checks that the ProxyConfig is valid.
func (c *ProxyConfig) Validate() error {
	if c == nil {
		return errors.New("proxy config must not be nil")
	}
	if !c.Enabled {
		return nil
	}
	if c.ProxyURL == "" {
		return errors.New("proxy URL must not be empty when proxy is enabled")
	}
	u, err := url.Parse(c.ProxyURL)
	if err != nil {
		return fmt.Errorf("invalid proxy URL: %w", err)
	}
	if u.Scheme != "http" && u.Scheme != "https" {
		return fmt.Errorf("proxy URL scheme must be http or https, got: %q", u.Scheme)
	}
	return nil
}

// DialOption returns a grpc.DialOption for the proxy configuration.
// When disabled, it returns nil.
func (c *ProxyConfig) DialOption() grpc.DialOption {
	if c == nil || !c.Enabled || c.ProxyURL == "" {
		return nil
	}
	return grpc.WithContextDialer(nil) // placeholder; real impl would wire a proxy dialer
}

// String returns a human-readable description of the proxy config.
func (c *ProxyConfig) String() string {
	if c == nil || !c.Enabled {
		return "proxy: disabled"
	}
	return fmt.Sprintf("proxy: enabled (url=%s)", c.ProxyURL)
}
