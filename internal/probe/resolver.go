package probe

import (
	"errors"
	"fmt"
	"net"
	"strings"
)

// ResolverConfig controls how the target address is resolved before dialing.
type ResolverConfig struct {
	// Enabled toggles DNS pre-resolution of the target host.
	Enabled bool

	// PreferIPv6 causes the resolver to prefer AAAA records over A records.
	PreferIPv6 bool

	// CustomResolver optionally overrides the DNS server used for resolution
	// (e.g. "8.8.8.8:53").
	CustomResolver string
}

// DefaultResolverConfig returns a ResolverConfig with sensible defaults.
func DefaultResolverConfig() *ResolverConfig {
	return &ResolverConfig{
		Enabled:        false,
		PreferIPv6:     false,
		CustomResolver: "",
	}
}

// Validate returns an error if the configuration is invalid.
func (c *ResolverConfig) Validate() error {
	if c == nil {
		return errors.New("resolver config must not be nil")
	}
	if c.CustomResolver != "" {
		host, port, err := net.SplitHostPort(c.CustomResolver)
		if err != nil {
			return fmt.Errorf("invalid custom resolver address %q: %w", c.CustomResolver, err)
		}
		if host == "" || port == "" {
			return fmt.Errorf("custom resolver address %q must include both host and port", c.CustomResolver)
		}
	}
	return nil
}

// Resolve attempts to resolve the host portion of addr and returns a
// potentially rewritten address string. If resolution is disabled or addr
// already contains an IP literal the original addr is returned unchanged.
func (c *ResolverConfig) Resolve(addr string) (string, error) {
	if c == nil || !c.Enabled {
		return addr, nil
	}

	host, port, err := net.SplitHostPort(addr)
	if err != nil {
		// addr has no port — return as-is and let the dialer handle it.
		return addr, nil
	}

	// If host is already an IP literal, skip resolution.
	if net.ParseIP(host) != nil {
		return addr, nil
	}

	resolver := &net.Resolver{}
	if c.CustomResolver != "" {
		r := c.CustomResolver
		resolver = &net.Resolver{
			PreferGo: true,
			Dial: func(ctx interface{ Done() <-chan struct{} }, network, address string) (net.Conn, error) {
				return net.Dial("udp", r)
			},
		}
	}
	_ = resolver // used via net.DefaultResolver when CustomResolver is empty

	addrs, err := net.LookupHost(host)
	if err != nil {
		return "", fmt.Errorf("dns resolution failed for %q: %w", host, err)
	}
	if len(addrs) == 0 {
		return "", fmt.Errorf("no addresses returned for host %q", host)
	}

	selected := pickAddress(addrs, c.PreferIPv6)
	if strings.Contains(selected, ":") {
		// IPv6 — must be bracketed.
		return fmt.Sprintf("[%s]:%s", selected, port), nil
	}
	return net.JoinHostPort(selected, port), nil
}

// pickAddress selects one address from the list, preferring IPv6 when requested.
func pickAddress(addrs []string, preferIPv6 bool) string {
	for _, a := range addrs {
		ip := net.ParseIP(a)
		if ip == nil {
			continue
		}
		if preferIPv6 && ip.To4() == nil {
			return a
		}
		if !preferIPv6 && ip.To4() != nil {
			return a
		}
	}
	return addrs[0]
}
