package probe

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

// TraceIDConfig controls whether a unique trace ID is injected into
// outgoing gRPC metadata on each probe request.
type TraceIDConfig struct {
	// Enabled controls whether trace ID injection is active.
	Enabled bool

	// HeaderName is the metadata key used to carry the trace ID.
	// Defaults to "x-probe-trace-id".
	HeaderName string
}

// DefaultTraceIDConfig returns a TraceIDConfig with sensible defaults.
func DefaultTraceIDConfig() *TraceIDConfig {
	return &TraceIDConfig{
		Enabled:    false,
		HeaderName: "x-probe-trace-id",
	}
}

// Validate checks that the TraceIDConfig is internally consistent.
func (c *TraceIDConfig) Validate() error {
	if c == nil {
		return fmt.Errorf("trace ID config must not be nil")
	}
	if c.Enabled && c.HeaderName == "" {
		return fmt.Errorf("trace ID header name must not be empty when enabled")
	}
	return nil
}

// NewTraceID generates a cryptographically random 16-byte hex trace ID.
func NewTraceID() (string, error) {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return "", fmt.Errorf("generating trace ID: %w", err)
	}
	return hex.EncodeToString(b), nil
}

// InjectTraceID returns a new context with the trace ID stored under the
// key used by the metadata interceptor, if tracing is enabled.
func (c *TraceIDConfig) InjectTraceID(ctx context.Context) (context.Context, string, error) {
	if c == nil || !c.Enabled {
		return ctx, "", nil
	}
	id, err := NewTraceID()
	if err != nil {
		return ctx, "", err
	}
	type traceKey struct{}
	return context.WithValue(ctx, traceKey{}, id), id, nil
}
