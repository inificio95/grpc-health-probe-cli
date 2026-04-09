package main

import (
	"strings"

	"github.com/spf13/cobra"

	"github.com/your-org/grpc-health-probe-cli/internal/probe"
)

// addMetadataFlags registers metadata-related CLI flags on the given command.
func addMetadataFlags(cmd *cobra.Command) {
	if cmd == nil {
		return
	}
	cmd.Flags().StringArrayP(
		"header", "H", nil,
		`Metadata header to attach to the request in key=value format (may be repeated)`,
	)
}

// parseMetadataConfig builds a MetadataConfig from the flags registered by
// addMetadataFlags. It returns the default config when no headers are provided.
func parseMetadataConfig(cmd *cobra.Command) (*probe.MetadataConfig, error) {
	if cmd == nil {
		return probe.DefaultMetadataConfig(), nil
	}

	headers, err := cmd.Flags().GetStringArray("header")
	if err != nil || len(headers) == 0 {
		return probe.DefaultMetadataConfig(), nil
	}

	entries := make([]probe.MetadataEntry, 0, len(headers))
	for _, h := range headers {
		parts := strings.SplitN(h, "=", 2)
		if len(parts) != 2 {
			continue
		}
		entries = append(entries, probe.MetadataEntry{
			Key:   strings.TrimSpace(parts[0]),
			Value: strings.TrimSpace(parts[1]),
		})
	}

	return &probe.MetadataConfig{Entries: entries}, nil
}
