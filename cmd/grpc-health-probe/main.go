package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/example/grpc-health-probe-cli/internal/probe"
)

func newRootCmd() *cobra.Command {
	var (
		addr       string
		service    string
		format     string
		timeout    int
		dialTimeout int
		maxRetries int
		metadata   []string
		tlsEnabled bool
		tlsSkipVerify bool
		tlsCACert  string
		tlsCert    string
		tlsKey     string
		authType   string
		authToken  string
		authUser   string
		authPass   string
	)

	cmd := &cobra.Command{
		Use:   "grpc-health-probe",
		Short: "Check gRPC service health endpoints",
		RunE: func(cmd *cobra.Command, args []string) error {
			return run(addr, service, format, timeout, dialTimeout, maxRetries,
				metadata, tlsEnabled, tlsSkipVerify, tlsCACert, tlsCert, tlsKey,
				authType, authToken, authUser, authPass)
		},
	}

	cmd.Flags().StringVar(&addr, "addr", "", "gRPC server address (required)")
	cmd.Flags().StringVar(&service, "service", "", "Service name to check (empty for overall health)")
	cmd.Flags().StringVar(&format, "format", "text", "Output format: text or json")
	cmd.Flags().IntVar(&timeout, "timeout", 5, "Request timeout in seconds")
	cmd.Flags().IntVar(&dialTimeout, "dial-timeout", 5, "Dial timeout in seconds")
	cmd.Flags().IntVar(&maxRetries, "max-retries", 0, "Maximum number of retries")
	cmd.Flags().StringArrayVar(&metadata, "metadata", nil, "Metadata key=value pairs")
	cmd.Flags().BoolVar(&tlsEnabled, "tls", false, "Enable TLS")
	cmd.Flags().BoolVar(&tlsSkipVerify, "tls-skip-verify", false, "Skip TLS certificate verification")
	cmd.Flags().StringVar(&tlsCACert, "tls-ca-cert", "", "Path to CA certificate")
	cmd.Flags().StringVar(&tlsCert, "tls-cert", "", "Path to client certificate")
	cmd.Flags().StringVar(&tlsKey, "tls-key", "", "Path to client key")
	cmd.Flags().StringVar(&authType, "auth-type", "none", "Auth type: none, bearer, or basic")
	cmd.Flags().StringVar(&authToken, "auth-token", "", "Bearer token")
	cmd.Flags().StringVar(&authUser, "auth-username", "", "Basic auth username")
	cmd.Flags().StringVar(&authPass, "auth-password", "", "Basic auth password")

	_ = cmd.MarkFlagRequired("addr")
	return cmd
}

func run(addr, service, format string, timeoutSec, dialTimeoutSec, maxRetries int,
	metadata []string, tlsEnabled, tlsSkipVerify bool, tlsCACert, tlsCert, tlsKey string,
	authType, authToken, authUser, authPass string) error {

	cfg := probe.DefaultConfig()
	cfg.Address = addr
	cfg.Service = service

	cfg.Auth = probe.AuthConfig{
		Type:     probe.AuthType(authType),
		Token:    authToken,
		Username: authUser,
		Password: authPass,
	}

	if err := cfg.Auth.Validate(); err != nil {
		return fmt.Errorf("auth config: %w", err)
	}

	p, err := probe.New(cfg)
	if err != nil {
		return fmt.Errorf("creating prober: %w", err)
	}

	result := p.Check()
	fmt.Print(probe.FormatResult(result, format))

	if !result.OK {
		os.Exit(1)
	}
	return nil
}

func main() {
	if err := newRootCmd().Execute(); err != nil {
		os.Exit(1)
	}
}
