package probe

import (
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/encoding/gzip"
)

// CompressionType represents the type of compression to use
type CompressionType string

const (
	// CompressionNone disables compression
	CompressionNone CompressionType = "none"
	// CompressionGzip enables gzip compression
	CompressionGzip CompressionType = "gzip"
)

// CompressionConfig holds compression-related configuration
type CompressionConfig struct {
	// Enabled indicates whether compression is enabled
	Enabled bool
	// Type specifies the compression algorithm to use
	Type CompressionType
}

// DefaultCompressionConfig returns the default compression configuration
func DefaultCompressionConfig() *CompressionConfig {
	return &CompressionConfig{
		Enabled: false,
		Type:    CompressionNone,
	}
}

// Validate checks if the compression configuration is valid
func (c *CompressionConfig) Validate() error {
	if c == nil {
		return fmt.Errorf("compression config cannot be nil")
	}

	if !c.Enabled {
		return nil
	}

	switch c.Type {
	case CompressionNone, CompressionGzip:
		return nil
	default:
		return fmt.Errorf("unsupported compression type: %s (supported: none, gzip)", c.Type)
	}
}

// ApplyDialOptions returns gRPC dial options for compression
func (c *CompressionConfig) ApplyDialOptions() []grpc.DialOption {
	if c == nil || !c.Enabled || c.Type == CompressionNone {
		return nil
	}

	var opts []grpc.DialOption

	switch c.Type {
	case CompressionGzip:
		// Register gzip compressor (it's registered by default in grpc-go)
		// Set default compression for outgoing messages
		opts = append(opts, grpc.WithDefaultCallOptions(grpc.UseCompressor(gzip.Name)))
	}

	return opts
}
