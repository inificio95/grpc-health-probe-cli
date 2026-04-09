package probe

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"

	"google.golang.org/grpc/credentials"
)

// TLSConfig holds TLS configuration for gRPC connections.
type TLSConfig struct {
	// Enabled indicates whether TLS should be used.
	Enabled bool
	// CACertFile is the path to the CA certificate file.
	CACertFile string
	// ClientCertFile is the path to the client certificate file.
	ClientCertFile string
	// ClientKeyFile is the path to the client key file.
	ClientKeyFile string
	// InsecureSkipVerify disables server certificate verification.
	InsecureSkipVerify bool
	// ServerName overrides the server name used for TLS verification.
	ServerName string
}

// Validate checks that the TLSConfig fields are consistent.
func (c *TLSConfig) Validate() error {
	if !c.Enabled {
		return nil
	}
	if (c.ClientCertFile == "") != (c.ClientKeyFile == "") {
		return fmt.Errorf("tls: client cert and key must both be provided")
	}
	return nil
}

// BuildCredentials constructs gRPC transport credentials from TLSConfig.
func (c *TLSConfig) BuildCredentials() (credentials.TransportCredentials, error) {
	if err := c.Validate(); err != nil {
		return nil, err
	}

	tlsCfg := &tls.Config{
		InsecureSkipVerify: c.InsecureSkipVerify, //nolint:gosec
		ServerName:         c.ServerName,
	}

	if c.CACertFile != "" {
		caCert, err := os.ReadFile(c.CACertFile)
		if err != nil {
			return nil, fmt.Errorf("tls: reading CA cert %q: %w", c.CACertFile, err)
		}
		pool := x509.NewCertPool()
		if !pool.AppendCertsFromPEM(caCert) {
			return nil, fmt.Errorf("tls: failed to append CA cert from %q", c.CACertFile)
		}
		tlsCfg.RootCAs = pool
	}

	if c.ClientCertFile != "" {
		cert, err := tls.LoadX509KeyPair(c.ClientCertFile, c.ClientKeyFile)
		if err != nil {
			return nil, fmt.Errorf("tls: loading client key pair: %w", err)
		}
		tlsCfg.Certificates = []tls.Certificate{cert}
	}

	return credentials.NewTLS(tlsCfg), nil
}
