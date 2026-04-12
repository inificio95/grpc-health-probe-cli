package main

import (
	"strings"

	"github.com/spf13/cobra"

	"github.com/andrewstucki/grpc-health-probe-cli/internal/probe"
)

// addTagsFlags registers tag-related CLI flags on cmd.
func addTagsFlags(cmd *cobra.Command) {
	if cmd == nil {
		return
	}
	cmd.Flags().StringSlice(
		"tag",
		nil,
		"Key=value label to attach to probe results (repeatable, e.g. --tag env=prod)",
	)
}

// parseTagsConfig builds a TagsConfig from the flags bound to cmd.
func parseTagsConfig(cmd *cobra.Command) *probe.TagsConfig {
	if cmd == nil {
		return probe.DefaultTagsConfig()
	}

	raw, err := cmd.Flags().GetStringSlice("tag")
	if err != nil || len(raw) == 0 {
		return probe.DefaultTagsConfig()
	}

	tags := make(map[string]string, len(raw))
	for _, entry := range raw {
		parts := strings.SplitN(entry, "=", 2)
		if len(parts) != 2 {
			continue
		}
		tags[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
	}

	return &probe.TagsConfig{
		Enabled: len(tags) > 0,
		Tags:    tags,
	}
}
