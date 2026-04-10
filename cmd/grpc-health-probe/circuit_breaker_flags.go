package main

import (
	"github.com/spf13/cobra"

	"github.com/your-org/grpc-health-probe-cli/internal/probe"
)

// addCircuitBreakerFlags registers circuit breaker flags on the given command.
func addCircuitBreakerFlags(cmd *cobra.Command) {
	if cmd == nil {
		return
	}
	flags := cmd.Flags()
	flags.Bool("circuit-breaker", false, "Enable circuit breaker")
	flags.Int("circuit-breaker-max-failures", 5, "Number of failures before opening the circuit")
	flags.Duration("circuit-breaker-open-duration", probe.DefaultCircuitBreakerConfig().OpenDuration, "Duration to keep the circuit open before attempting half-open")
	flags.Int("circuit-breaker-half-open-requests", 1, "Number of requests allowed in half-open state")
}

// parseCircuitBreakerConfig builds a CircuitBreakerConfig from command flags.
func parseCircuitBreakerConfig(cmd *cobra.Command) *probe.CircuitBreakerConfig {
	if cmd == nil {
		return probe.DefaultCircuitBreakerConfig()
	}
	cfg := probe.DefaultCircuitBreakerConfig()
	flags := cmd.Flags()

	if v, err := flags.GetBool("circuit-breaker"); err == nil {
		cfg.Enabled = v
	}
	if v, err := flags.GetInt("circuit-breaker-max-failures"); err == nil {
		cfg.MaxFailures = v
	}
	if v, err := flags.GetDuration("circuit-breaker-open-duration"); err == nil {
		cfg.OpenDuration = v
	}
	if v, err := flags.GetInt("circuit-breaker-half-open-requests"); err == nil {
		cfg.HalfOpenRequests = v
	}
	return cfg
}
