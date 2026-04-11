package main

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/your-org/grpc-health-probe-cli/internal/probe"
)

// addCacheFlags registers result-caching flags onto cmd.
func addCacheFlags(cmd *cobra.Command) {
	if cmd == nil {
		return
	}
	cmd.Flags().Bool("cache", false, "Enable result caching between probe attempts")
	cmd.Flags().Duration("cache-ttl", 5*time.Second, "TTL for cached probe results")
}

// parseCacheConfig reads cache-related flags from cmd and returns a CacheConfig.
func parseCacheConfig(cmd *cobra.Command) (*probe.CacheConfig, error) {
	if cmd == nil {
		return probe.DefaultCacheConfig(), nil
	}

	cfg := probe.DefaultCacheConfig()

	if f := cmd.Flags().Lookup("cache"); f != nil {
		enabled, err := cmd.Flags().GetBool("cache")
		if err != nil {
			return nil, fmt.Errorf("parsing --cache: %w", err)
		}
		cfg.Enabled = enabled
	}

	if f := cmd.Flags().Lookup("cache-ttl"); f != nil {
		ttl, err := cmd.Flags().GetDuration("cache-ttl")
		if err != nil {
			return nil, fmt.Errorf("parsing --cache-ttl: %w", err)
		}
		cfg.TTL = ttl
	}

	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("invalid cache config: %w", err)
	}
	return cfg, nil
}
