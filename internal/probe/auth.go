package probe

import (
	"fmt"

	"google.golang.org/grpc/credentials"
)

// AuthType represents the authentication mechanism to use.
type AuthType string

const (
	AuthTypeNone  AuthType = "none"
	AuthTypeBearer AuthType = "bearer"
	AuthTypeBasic  AuthType = "basic"
)

// AuthConfig holds authentication configuration for gRPC calls.
type AuthConfig struct {
	Type     AuthType
	Token    string
	Username string
	Password string
}

// DefaultAuthConfig returns an AuthConfig with sensible defaults.
func DefaultAuthConfig() AuthConfig {
	return AuthConfig{
		Type: AuthTypeNone,
	}
}

// Validate checks that the AuthConfig is consistent and complete.
func (a AuthConfig) Validate() error {
	switch a.Type {
	case AuthTypeNone:
		return nil
	case AuthTypeBearer:
		if a.Token == "" {
			return fmt.Errorf("bearer auth requires a non-empty token")
		}
	case AuthTypeBasic:
		if a.Username == "" {
			return fmt.Errorf("basic auth requires a non-empty username")
		}
		if a.Password == "" {
			return fmt.Errorf("basic auth requires a non-empty password")
		}
	default:
		return fmt.Errorf("unknown auth type %q", a.Type)
	}
	return nil
}

// PerRPCCredentials returns gRPC per-RPC credentials for the configured auth type.
// Returns nil if auth type is none.
func (a AuthConfig) PerRPCCredentials() credentials.PerRPCCredentials {
	switch a.Type {
	case AuthTypeBearer:
		return &bearerCredentials{token: a.Token}
	case AuthTypeBasic:
		return &basicCredentials{username: a.Username, password: a.Password}
	default:
		return nil
	}
}
