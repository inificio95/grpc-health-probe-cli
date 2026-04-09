package main

import (
	"github.com/spf13/cobra"
	"grpc-health-probe-cli/internal/probe"
)

// compressionFlags holds compression-related CLI flags
type compressionFlags struct {
	enabled bool
	type_   string
}

// addCompressionFlags adds compression flags to the command
func addCompressionFlags(cmd *cobra.Command, flags *compressionFlags) {
	cmd.Flags().BoolVar(
		&flags.enabled,
		"compression-enabled",
		false,
		"Enable compression for gRPC requests",
	)

	cmd.Flags().StringVar(
		&flags.type_,
		"compression",
		"none",
		"Compression type to use (none, gzip)",
	)
}

// toCompressionConfig converts CLI flags to CompressionConfig
func (f *compressionFlags) toCompressionConfig() *probe.CompressionConfig {
	// If compression type is specified and not "none", enable compression
	enabled := f.enabled || (f.type_ != "" && f.type_ != "none")

	compressionType := probe.CompressionType(f.type_)
	if f.type_ == "" {
		compressionType = probe.CompressionNone
	}

	return &probe.CompressionConfig{
		Enabled: enabled,
		Type:    compressionType,
	}
}

// Example usage in main command:
// var compFlags compressionFlags
// addCompressionFlags(rootCmd, &compFlags)
// compressionConfig := compFlags.toCompressionConfig()
