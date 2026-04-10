package probe

import (
	"errors"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

// KeepaliveConfig holds gRPC client keepalive parameters.
type KeepaliveConfig struct {
	Enabled             bool
	Time                time.Duration
	Timeout             time.Duration
	PermitWithoutStream bool
}

// DefaultKeepaliveConfig returns a KeepaliveConfig with sensible defaults.
func DefaultKeepaliveConfig() *KeepaliveConfig {
	return &KeepaliveConfig{
		Enabled:             false,
		Time:                30 * time.Second,
		Timeout:             10 * time.Second,
		PermitWithoutStream: false,
	}
}

// Validate checks that the KeepaliveConfig fields are valid.
func (c *KeepaliveConfig) Validate() error {
	if c == nil {
		return errors.New("keepalive config must not be nil")
	}
	if c.Time <= 0 {
		return errors.New("keepalive time must be positive")
	}
	if c.Timeout <= 0 {
		return errors.New("keepalive timeout must be positive")
	}
	return nil
}

// DialOption returns a grpc.DialOption for the keepalive parameters,
// or nil if keepalive is disabled.
func (c *KeepaliveConfig) DialOption() grpc.DialOption {
	if c == nil || !c.Enabled {
		return nil
	}
	return grpc.WithKeepaliveParams(keepalive.ClientParameters{
		Time:                c.Time,
		Timeout:             c.Timeout,
		PermitWithoutStream: c.PermitWithoutStream,
	})
}
