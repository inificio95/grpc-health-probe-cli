package probe

// ServiceConfig holds configuration for targeting a specific gRPC service
// and optional named health check service.
type ServiceConfig struct {
	// Service is the fully-qualified gRPC service name to check.
	// If empty, the overall server health is checked.
	Service string

	// WatchMode enables streaming health watch instead of a single check.
	WatchMode bool
}

// DefaultServiceConfig returns a ServiceConfig with default values.
func DefaultServiceConfig() ServiceConfig {
	return ServiceConfig{
		Service:   "",
		WatchMode: false,
	}
}

// Validate checks that the ServiceConfig is valid.
func (c ServiceConfig) Validate() error {
	// Service name can be empty (checks overall server health) or a non-empty string.
	// No additional constraints are enforced here; the gRPC server will reject
	// unknown service names with an appropriate status.
	return nil
}

// IsServerLevel reports whether the config targets overall server health
// rather than a specific named service.
func (c ServiceConfig) IsServerLevel() bool {
	return c.Service == ""
}
