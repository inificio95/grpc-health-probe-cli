package probe

import (
	"errors"

	"google.golang.org/grpc"
)

// FailFastConfig controls whether gRPC calls fail immediately on a
// non-transient error instead of waiting/retrying.
type FailFastConfig struct {
	// Enabled determines whether fail-fast behaviour is active.
	// When true, the call returns immediately if the connection is not ready.
	Enabled bool
}

// DefaultFailFastConfig returns a FailFastConfig with sensible defaults.
// Fail-fast is enabled by default to surface connectivity issues quickly.
func DefaultFailFastConfig() *FailFastConfig {
	return &FailFastConfig{
		Enabled: true,
	}
}

// Validate checks that the FailFastConfig is in a consistent state.
func (c *FailFastConfig) Validate() error {
	if c == nil {
		return errors.New("failfast config must not be nil")
	}
	return nil
}

// CallOption returns the appropriate grpc.CallOption based on the config.
// When Enabled is true it returns grpc.WaitForReady(false) (fail-fast);
// otherwise it returns grpc.WaitForReady(true) so the call blocks until
// the connection becomes ready.
func (c *FailFastConfig) CallOption() grpc.CallOption {
	if c == nil || c.Enabled {
		return grpc.WaitForReady(false)
	}
	return grpc.WaitForReady(true)
}

// String returns a human-readable representation of the config.
func (c *FailFastConfig) String() string {
	if c == nil {
		return "failfast=nil"
	}
	if c.Enabled {
		return "failfast=enabled"
	}
	return "failfast=disabled"
}
