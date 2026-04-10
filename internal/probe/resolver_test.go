package probe

import (
	"testing"
)

func TestDefaultResolverConfig(t *testing.T) {
	cfg := DefaultResolverConfig()
	if cfg == nil {
		t.Fatal("expected non-nil config")
	}
	if cfg.Enabled {
		t.Error("expected Enabled=false by default")
	}
	if cfg.PreferIPv6 {
		t.Error("expected PreferIPv6=false by default")
	}
	if cfg.CustomResolver != "" {
		t.Errorf("expected empty CustomResolver, got %q", cfg.CustomResolver)
	}
}

func TestResolverConfig_Validate_Nil(t *testing.T) {
	var cfg *ResolverConfig
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for nil config")
	}
}

func TestResolverConfig_Validate_Valid(t *testing.T) {
	cfg := DefaultResolverConfig()
	if err := cfg.Validate(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestResolverConfig_Validate_CustomResolver_Valid(t *testing.T) {
	cfg := &ResolverConfig{Enabled: true, CustomResolver: "8.8.8.8:53"}
	if err := cfg.Validate(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestResolverConfig_Validate_CustomResolver_Invalid(t *testing.T) {
	cfg := &ResolverConfig{Enabled: true, CustomResolver: "not-valid"}
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for invalid custom resolver address")
	}
}

func TestResolverConfig_Resolve_Disabled(t *testing.T) {
	cfg := &ResolverConfig{Enabled: false}
	got, err := cfg.Resolve("example.com:443")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != "example.com:443" {
		t.Errorf("expected original addr, got %q", got)
	}
}

func TestResolverConfig_Resolve_IPLiteral(t *testing.T) {
	cfg := &ResolverConfig{Enabled: true}
	got, err := cfg.Resolve("127.0.0.1:50051")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != "127.0.0.1:50051" {
		t.Errorf("expected unchanged addr for IP literal, got %q", got)
	}
}

func TestResolverConfig_Resolve_NilConfig(t *testing.T) {
	var cfg *ResolverConfig
	got, err := cfg.Resolve("example.com:443")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != "example.com:443" {
		t.Errorf("expected original addr for nil config, got %q", got)
	}
}

func TestPickAddress_PreferIPv4(t *testing.T) {
	addrs := []string{"2001:db8::1", "192.0.2.1"}
	got := pickAddress(addrs, false)
	if got != "192.0.2.1" {
		t.Errorf("expected IPv4, got %q", got)
	}
}

func TestPickAddress_PreferIPv6(t *testing.T) {
	addrs := []string{"192.0.2.1", "2001:db8::1"}
	got := pickAddress(addrs, true)
	if got != "2001:db8::1" {
		t.Errorf("expected IPv6, got %q", got)
	}
}

func TestPickAddress_FallbackFirst(t *testing.T) {
	addrs := []string{"192.0.2.1"}
	got := pickAddress(addrs, true) // no IPv6 available
	if got != "192.0.2.1" {
		t.Errorf("expected fallback to first address, got %q", got)
	}
}
