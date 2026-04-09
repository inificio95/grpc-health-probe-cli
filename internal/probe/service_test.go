package probe

import (
	"testing"
)

func TestDefaultServiceConfig(t *testing.T) {
	cfg := DefaultServiceConfig()

	if cfg.Service != "" {
		t.Errorf("expected empty Service, got %q", cfg.Service)
	}
	if cfg.WatchMode {
		t.Error("expected WatchMode to be false by default")
	}
}

func TestServiceConfig_Validate_Empty(t *testing.T) {
	cfg := DefaultServiceConfig()
	if err := cfg.Validate(); err != nil {
		t.Errorf("expected no error for empty service name, got %v", err)
	}
}

func TestServiceConfig_Validate_Named(t *testing.T) {
	cfg := ServiceConfig{Service: "grpc.health.v1.Health"}
	if err := cfg.Validate(); err != nil {
		t.Errorf("expected no error for named service, got %v", err)
	}
}

func TestServiceConfig_IsServerLevel_Empty(t *testing.T) {
	cfg := DefaultServiceConfig()
	if !cfg.IsServerLevel() {
		t.Error("expected IsServerLevel to return true for empty Service")
	}
}

func TestServiceConfig_IsServerLevel_Named(t *testing.T) {
	cfg := ServiceConfig{Service: "my.package.MyService"}
	if cfg.IsServerLevel() {
		t.Error("expected IsServerLevel to return false for named service")
	}
}

func TestServiceConfig_WatchMode(t *testing.T) {
	cfg := ServiceConfig{
		Service:   "my.package.MyService",
		WatchMode: true,
	}
	if !cfg.WatchMode {
		t.Error("expected WatchMode to be true")
	}
	if err := cfg.Validate(); err != nil {
		t.Errorf("unexpected validation error: %v", err)
	}
}
