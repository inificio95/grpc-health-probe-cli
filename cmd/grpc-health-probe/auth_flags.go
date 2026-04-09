package main

import (
	"github.com/spf13/cobra"

	"github.com/your-org/grpc-health-probe-cli/internal/probe"
)

// addAuthFlags registers authentication-related flags onto the given command.
func addAuthFlags(cmd *cobra.Command) {
	if cmd == nil {
		return
	}

	cmd.Flags().String("auth-type", "none", `Authentication type: "none", "bearer", or "basic"`)
	cmd.Flags().String("auth-token", "", "Bearer token for bearer authentication")
	cmd.Flags().String("auth-username", "", "Username for basic authentication")
	cmd.Flags().String("auth-password", "", "Password for basic authentication")
}

// parseAuthConfig builds a probe.AuthConfig from the flags set on cmd.
func parseAuthConfig(cmd *cobra.Command) (*probe.AuthConfig, error) {
	if cmd == nil {
		return probe.DefaultAuthConfig(), nil
	}

	authType, err := cmd.Flags().GetString("auth-type")
	if err != nil {
		return nil, err
	}

	token, err := cmd.Flags().GetString("auth-token")
	if err != nil {
		return nil, err
	}

	username, err := cmd.Flags().GetString("auth-username")
	if err != nil {
		return nil, err
	}

	password, err := cmd.Flags().GetString("auth-password")
	if err != nil {
		return nil, err
	}

	cfg := &probe.AuthConfig{
		Type:     authType,
		Token:    token,
		Username: username,
		Password: password,
	}

	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	return cfg, nil
}
