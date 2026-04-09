package main

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"

	"github.com/example/grpc-health-probe-cli/internal/probe"
)

var (
	address    string
	service    string
	timeout    time.Duration
	retryCount int
	retryDelay time.Duration
	format     string
	tls        bool
)

func newRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "grpc-health-probe",
		Short: "Check gRPC service health endpoints",
		Long:  "A lightweight CLI tool for checking gRPC service health endpoints with configurable retry and timeout logic.",
		RunE:  run,
	}

	cmd.Flags().StringVarP(&address, "addr", "a", "", "gRPC server address (host:port) (required)")
	cmd.Flags().StringVarP(&service, "service", "s", "", "gRPC service name to check (empty for overall health)")
	cmd.Flags().DurationVarP(&timeout, "timeout", "t", 5*time.Second, "timeout per probe attempt")
	cmd.Flags().IntVarP(&retryCount, "retry", "r", 0, "number of retries on failure")
	cmd.Flags().DurationVar(&retryDelay, "retry-delay", time.Second, "delay between retries")
	cmd.Flags().StringVarP(&format, "format", "f", "text", "output format: text or json")
	cmd.Flags().BoolVar(&tls, "tls", false, "use TLS when connecting")

	_ = cmd.MarkFlagRequired("addr")

	return cmd
}

func run(cmd *cobra.Command, _ []string) error {
	cfg := probe.Config{
		Address:    address,
		Service:    service,
		Timeout:    timeout,
		RetryCount: retryCount,
		RetryDelay: retryDelay,
		TLS:        tls,
	}

	p, err := probe.New(cfg)
	if err != nil {
		return fmt.Errorf("invalid configuration: %w", err)
	}

	result := p.Check(cmd.Context())

	output, err := probe.FormatResult(result, format)
	if err != nil {
		return fmt.Errorf("formatting result: %w", err)
	}

	fmt.Fprintln(cmd.OutOrStdout(), output)

	if !result.Healthy {
		return fmt.Errorf("service is not healthy")
	}
	return nil
}

func main() {
	if err := newRootCmd().Execute(); err != nil {
		os.Exit(1)
	}
}
