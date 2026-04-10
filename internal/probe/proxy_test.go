package probe

import (
	"testing"
)

func TestDefaultProxyConfig(t *testing.T) {
	cfg := DefaultProxyConfig()
	if cfg == nil {
		t.Fatal("expected non-nil config")
	}
	if cfg.Enabled {
		t.Error("expected proxy to be disabled by default")
	}
	if cfg.ProxyURL != "" {
		t.Errorf("expected empty proxy URL, got %q", cfg.ProxyURL)
	}
}

func TestProxyConfig_Validate_Nil(t *testing.T) {
	var cfg *ProxyConfig
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for nil config")
	}
}

func TestProxyConfig_Validate_Disabled(t *testing.T) {
	cfg := DefaultProxyConfig()
	if err := cfg.Validate(); err != nil {
		t.Errorf("unexpected error for disabled proxy: %v", err)
	}
}

func TestProxyConfig_Validate_EnabledMissingURL(t *testing.T) {
	cfg := &ProxyConfig{Enabled: true, ProxyURL: ""}
	if err := cfg.Validate(); err == nil {
		t.Error("expected error when proxy enabled but URL empty")
	}
}

func TestProxyConfig_Validate_InvalidURL(t *testing.T) {
	cfg := &ProxyConfig{Enabled: true, ProxyURL: "://bad-url"}
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for invalid proxy URL")
	}
}

func TestProxyConfig_Validate_InvalidScheme(t *testing.T) {
	cfg := &ProxyConfig{Enabled: true, ProxyURL: "socks5://proxy.example.com:1080"}
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for non-http/https scheme")
	}
}

func TestProxyConfig_Validate_ValidHTTP(t *testing.T) {
	cfg := &ProxyConfig{Enabled: true, ProxyURL: "http://proxy.example.com:8080"}
	if err := cfg.Validate(); err != nil {
		t.Errorf("unexpected error for valid http proxy URL: %v", err)
	}
}

func TestProxyConfig_Validate_ValidHTTPS(t *testing.T) {
	cfg := &ProxyConfig{Enabled: true, ProxyURL: "https://proxy.example.com:8443"}
	if err := cfg.Validate(); err != nil {
		t.Errorf("unexpected error for valid https proxy URL: %v", err)
	}
}

func TestProxyConfig_String_Disabled(t *testing.T) {
	cfg := DefaultProxyConfig()
	s := cfg.String()
	if s != "proxy: disabled" {
		t.Errorf("unexpected string for disabled proxy: %q", s)
	}
}

func TestProxyConfig_String_Enabled(t *testing.T) {
	cfg := &ProxyConfig{Enabled: true, ProxyURL: "http://proxy.example.com:8080"}
	s := cfg.String()
	if s == "proxy: disabled" {
		t.Error("expected non-disabled string for enabled proxy")
	}
}

func TestProxyConfig_String_Nil(t *testing.T) {
	var cfg *ProxyConfig
	if cfg.String() != "proxy: disabled" {
		t.Error("expected 'proxy: disabled' for nil config")
	}
}
