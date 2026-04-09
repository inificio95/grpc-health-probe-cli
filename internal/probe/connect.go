package probe

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// DefaultConnectConfig returns a ConnectConfig with sensible defaults.
func DefaultConnectConfig() *ConnectConfig {
	return &ConnectConfig{
		UserAgent:   "grpc-health-probe-cli/1.0",
		Block:       false,
		FailOnNonTempDialError: false,
	}
}

// ConnectConfig holds options that control how the gRPC connection is established.
type ConnectConfig struct {
	// UserAgent is sent as the gRPC user-agent header.
	UserAgent string

	// Block causes Dial to block until the connection is ready.
	Block bool

	// FailOnNonTempDialError causes Dial to fail fast on non-temporary errors.
	FailOnNonTempDialError bool
}

// Validate returns an error if the ConnectConfig is invalid.
func (c *ConnectConfig) Validate() error {
	if c == nil {
		return fmt.Errorf("connect config must not be nil")
	}
	return nil
}

// DialOptions converts the ConnectConfig into a slice of grpc.DialOption.
func (c *ConnectConfig) DialOptions(ctx context.Context, tls *TLSConfig) ([]grpc.DialOption, error) {
	var opts []grpc.DialOption

	if c.UserAgent != "" {
		opts = append(opts, grpc.WithUserAgent(c.UserAgent))
	}

	if c.Block {
		opts = append(opts, grpc.WithBlock())
	}

	if c.FailOnNonTempDialError {
		opts = append(opts, grpc.FailOnNonTempDialError(true))
	}

	creds, err := tls.BuildCredentials()
	if err != nil {
		return nil, fmt.Errorf("building TLS credentials: %w", err)
	}
	if creds != nil {
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	return opts, nil
}
