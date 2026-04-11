package probe

import (
	"runtime"
	"strings"
	"testing"
)

func TestDefaultUserAgentConfig(t *testing.T) {
	cfg := DefaultUserAgentConfig()
	if cfg == nil {
		t.Fatal("expected non-nil config")
	}
	if cfg.AppName != defaultAppName {
		t.Errorf("expected AppName %q, got %q", defaultAppName, cfg.AppName)
	}
	if cfg.AppVersion != defaultAppVersion {
		t.Errorf("expected AppVersion %q, got %q", defaultAppVersion, cfg.AppVersion)
	}
	if cfg.Disabled {
		t.Error("expected Disabled to be false")
	}
}

func TestUserAgentConfig_Validate_Valid(t *testing.T) {
	cfg := DefaultUserAgentConfig()
	if err := cfg.Validate(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestUserAgentConfig_Validate_Nil(t *testing.T) {
	var cfg *UserAgentConfig
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for nil config")
	}
}

func TestUserAgentConfig_Validate_EmptyName(t *testing.T) {
	cfg := DefaultUserAgentConfig()
	cfg.AppName = ""
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for empty AppName")
	}
}

func TestUserAgentConfig_Validate_EmptyVersion(t *testing.T) {
	cfg := DefaultUserAgentConfig()
	cfg.AppVersion = ""
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for empty AppVersion")
	}
}

func TestUserAgentConfig_Validate_Disabled(t *testing.T) {
	cfg := &UserAgentConfig{Disabled: true}
	if err := cfg.Validate(); err != nil {
		t.Errorf("unexpected error for disabled config: %v", err)
	}
}

func TestUserAgentConfig_String(t *testing.T) {
	cfg := DefaultUserAgentConfig()
	ua := cfg.String()
	if !strings.Contains(ua, defaultAppName) {
		t.Errorf("expected user-agent to contain app name, got %q", ua)
	}
	if !strings.Contains(ua, defaultAppVersion) {
		t.Errorf("expected user-agent to contain version, got %q", ua)
	}
	if !strings.Contains(ua, runtime.GOOS) {
		t.Errorf("expected user-agent to contain OS, got %q", ua)
	}
}

func TestUserAgentConfig_String_Disabled(t *testing.T) {
	cfg := &UserAgentConfig{Disabled: true}
	if ua := cfg.String(); ua != "" {
		t.Errorf("expected empty string for disabled config, got %q", ua)
	}
}

func TestUserAgentConfig_String_CustomAppName(t *testing.T) {
	cfg := &UserAgentConfig{
		AppName:    "my-custom-app",
		AppVersion: "2.0.0",
	}
	ua := cfg.String()
	if !strings.Contains(ua, "my-custom-app") {
		t.Errorf("expected user-agent to contain custom app name, got %q", ua)
	}
	if !strings.Contains(ua, "2.0.0") {
		t.Errorf("expected user-agent to contain custom version, got %q", ua)
	}
}
