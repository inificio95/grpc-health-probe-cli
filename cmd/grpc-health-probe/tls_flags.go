package main

import (
	"github.com/spf13/cobra"
	"github.com/your-org/grpc-health-probe-cli/internal/probe"
)

// addTLSFlags registers TLS-related flags onto the given command.
func addTLSFlags(cmd *cobra.Command) {
	if cmd == nil {
		return
	}
	cmd.Flags().Bool("tls", false, "Enable TLS for the connection")
	cmd.Flags().Bool("tls-insecure-skip-verify", false, "Skip TLS certificate verification (insecure)")
	cmd.Flags().String("tls-ca-cert", "", "Path to CA certificate file for verifying the server")
	cmd.Flags().String("tls-client-cert", "", "Path to client certificate file for mTLS")
	cmd.Flags().String("tls-client-key", "", "Path to client key file for mTLS")
	cmd.Flags().String("tls-server-name", "", "Override the server name used for TLS SNI")
}

// parseTLSConfig builds a probe.TLSConfig from the flags on the given command.
func parseTLSConfig(cmd *cobra.Command) (*probe.TLSConfig, error) {
	if cmd == nil {
		return probe.DefaultTLSConfig(), nil
	}

	enabled, err := cmd.Flags().GetBool("tls")
	if err != nil {
		return nil, err
	}
	insecure, err := cmd.Flags().GetBool("tls-insecure-skip-verify")
	if err != nil {
		return nil, err
	}
	caCert, err := cmd.Flags().GetString("tls-ca-cert")
	if err != nil {
		return nil, err
	}
	clientCert, err := cmd.Flags().GetString("tls-client-cert")
	if err != nil {
		return nil, err
	}
	clientKey, err := cmd.Flags().GetString("tls-client-key")
	if err != nil {
		return nil, err
	}
	serverName, err := cmd.Flags().GetString("tls-server-name")
	if err != nil {
		return nil, err
	}

	return &probe.TLSConfig{
		Enabled:            enabled,
		InsecureSkipVerify: insecure,
		CACertFile:         caCert,
		ClientCertFile:     clientCert,
		ClientKeyFile:      clientKey,
		ServerName:         serverName,
	}, nil
}
