package probe

import (
	"fmt"
	"strings"

	"google.golang.org/grpc/metadata"
)

// MetadataConfig holds key-value pairs to be sent as gRPC metadata with health check requests.
type MetadataConfig struct {
	// Headers is a list of "key=value" strings to include as gRPC metadata.
	Headers []string
}

// Validate checks that all header entries are properly formatted as "key=value".
func (m MetadataConfig) Validate() error {
	for _, h := range m.Headers {
		parts := strings.SplitN(h, "=", 2)
		if len(parts) != 2 || strings.TrimSpace(parts[0]) == "" {
			return fmt.Errorf("invalid metadata header %q: must be in key=value format", h)
		}
	}
	return nil
}

// ToGRPCMetadata converts the MetadataConfig into a gRPC metadata.MD map.
func (m MetadataConfig) ToGRPCMetadata() metadata.MD {
	md := metadata.MD{}
	for _, h := range m.Headers {
		parts := strings.SplitN(h, "=", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			val := strings.TrimSpace(parts[1])
			md[key] = append(md[key], val)
		}
	}
	return md
}

// DefaultMetadataConfig returns a MetadataConfig with no headers.
func DefaultMetadataConfig() MetadataConfig {
	return MetadataConfig{
		Headers: []string{},
	}
}
