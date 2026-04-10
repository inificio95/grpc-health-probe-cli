package probe

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
)

// InterceptorConfig holds configuration for gRPC client interceptors.
type InterceptorConfig struct {
	// EnableLogging enables a logging interceptor that records each RPC call.
	EnableLogging bool
	// LogPrefix is the prefix used in log messages.
	LogPrefix string
}

// DefaultInterceptorConfig returns an InterceptorConfig with sensible defaults.
func DefaultInterceptorConfig() *InterceptorConfig {
	return &InterceptorConfig{
		EnableLogging: false,
		LogPrefix:     "grpc-health-probe",
	}
}

// Validate checks that the InterceptorConfig is valid.
func (c *InterceptorConfig) Validate() error {
	if c == nil {
		return fmt.Errorf("interceptor config must not be nil")
	}
	if c.LogPrefix == "" {
		return fmt.Errorf("interceptor log prefix must not be empty")
	}
	return nil
}

// UnaryClientInterceptor returns a gRPC UnaryClientInterceptor based on the config.
// If logging is disabled, it returns nil.
func (c *InterceptorConfig) UnaryClientInterceptor() grpc.UnaryClientInterceptor {
	if c == nil || !c.EnableLogging {
		return nil
	}
	prefix := c.LogPrefix
	return func(
		ctx context.Context,
		method string,
		req, reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		fmt.Printf("[%s] calling method: %s\n", prefix, method)
		err := invoker(ctx, method, req, reply, cc, opts...)
		if err != nil {
			fmt.Printf("[%s] method %s returned error: %v\n", prefix, method, err)
		} else {
			fmt.Printf("[%s] method %s succeeded\n", prefix, method)
		}
		return err
	}
}

// DialOptions returns gRPC dial options derived from the interceptor config.
func (c *InterceptorConfig) DialOptions() []grpc.DialOption {
	if c == nil {
		return nil
	}
	interceptor := c.UnaryClientInterceptor()
	if interceptor == nil {
		return nil
	}
	return []grpc.DialOption{
		grpc.WithUnaryInterceptor(interceptor),
	}
}
