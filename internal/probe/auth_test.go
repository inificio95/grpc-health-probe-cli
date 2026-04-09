package probe

import (
	"context"
	"testing"
)

func TestDefaultAuthConfig(t *testing.T) {
	cfg := DefaultAuthConfig()
	if cfg.Type != AuthTypeNone {
		t.Errorf("expected type %q, got %q", AuthTypeNone, cfg.Type)
	}
}

func TestAuthConfig_Validate_None(t *testing.T) {
	cfg := AuthConfig{Type: AuthTypeNone}
	if err := cfg.Validate(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestAuthConfig_Validate_BearerMissingToken(t *testing.T) {
	cfg := AuthConfig{Type: AuthTypeBearer}
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for bearer with no token")
	}
}

func TestAuthConfig_Validate_BearerValid(t *testing.T) {
	cfg := AuthConfig{Type: AuthTypeBearer, Token: "my-token"}
	if err := cfg.Validate(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestAuthConfig_Validate_BasicMissingUsername(t *testing.T) {
	cfg := AuthConfig{Type: AuthTypeBasic, Password: "pass"}
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for basic auth with no username")
	}
}

func TestAuthConfig_Validate_BasicMissingPassword(t *testing.T) {
	cfg := AuthConfig{Type: AuthTypeBasic, Username: "user"}
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for basic auth with no password")
	}
}

func TestAuthConfig_Validate_BasicValid(t *testing.T) {
	cfg := AuthConfig{Type: AuthTypeBasic, Username: "user", Password: "pass"}
	if err := cfg.Validate(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestAuthConfig_Validate_UnknownType(t *testing.T) {
	cfg := AuthConfig{Type: AuthType("oauth2")}
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for unknown auth type")
	}
}

func TestAuthConfig_PerRPCCredentials_None(t *testing.T) {
	cfg := DefaultAuthConfig()
	if creds := cfg.PerRPCCredentials(); creds != nil {
		t.Errorf("expected nil credentials for auth type none")
	}
}

func TestBearerCredentials_GetRequestMetadata(t *testing.T) {
	creds := &bearerCredentials{token: "abc123"}
	md, err := creds.GetRequestMetadata(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if md["authorization"] != "Bearer abc123" {
		t.Errorf("unexpected authorization header: %q", md["authorization"])
	}
}

func TestBasicCredentials_GetRequestMetadata(t *testing.T) {
	creds := &basicCredentials{username: "user", password: "pass"}
	md, err := creds.GetRequestMetadata(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// base64("user:pass") = "dXNlcjpwYXNz"
	expected := "Basic dXNlcjpwYXNz"
	if md["authorization"] != expected {
		t.Errorf("expected %q, got %q", expected, md["authorization"])
	}
}
