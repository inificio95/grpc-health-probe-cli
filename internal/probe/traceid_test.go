package probe

import (
	"context"
	"testing"
)

func TestDefaultTraceIDConfig(t *testing.T) {
	cfg := DefaultTraceIDConfig()
	if cfg == nil {
		t.Fatal("expected non-nil config")
	}
	if cfg.Enabled {
		t.Error("expected Enabled to be false by default")
	}
	if cfg.HeaderName != "x-probe-trace-id" {
		t.Errorf("unexpected default HeaderName: %q", cfg.HeaderName)
	}
}

func TestTraceIDConfig_Validate_Nil(t *testing.T) {
	var cfg *TraceIDConfig
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for nil config")
	}
}

func TestTraceIDConfig_Validate_Disabled(t *testing.T) {
	cfg := &TraceIDConfig{Enabled: false, HeaderName: ""}
	if err := cfg.Validate(); err != nil {
		t.Errorf("unexpected error for disabled config: %v", err)
	}
}

func TestTraceIDConfig_Validate_EnabledMissingHeader(t *testing.T) {
	cfg := &TraceIDConfig{Enabled: true, HeaderName: ""}
	if err := cfg.Validate(); err == nil {
		t.Error("expected error when header name is empty and enabled")
	}
}

func TestTraceIDConfig_Validate_EnabledValid(t *testing.T) {
	cfg := &TraceIDConfig{Enabled: true, HeaderName: "x-trace-id"}
	if err := cfg.Validate(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestNewTraceID_Unique(t *testing.T) {
	id1, err := NewTraceID()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	id2, err := NewTraceID()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if id1 == id2 {
		t.Error("expected unique trace IDs")
	}
	if len(id1) != 32 {
		t.Errorf("expected 32-char hex string, got len %d", len(id1))
	}
}

func TestTraceIDConfig_InjectTraceID_Disabled(t *testing.T) {
	cfg := DefaultTraceIDConfig() // Enabled = false
	ctx, id, err := cfg.InjectTraceID(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if id != "" {
		t.Errorf("expected empty trace ID when disabled, got %q", id)
	}
	if ctx == nil {
		t.Error("expected non-nil context")
	}
}

func TestTraceIDConfig_InjectTraceID_Enabled(t *testing.T) {
	cfg := &TraceIDConfig{Enabled: true, HeaderName: "x-probe-trace-id"}
	_, id, err := cfg.InjectTraceID(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if id == "" {
		t.Error("expected non-empty trace ID when enabled")
	}
	if len(id) != 32 {
		t.Errorf("expected 32-char hex string, got len %d", len(id))
	}
}

func TestTraceIDConfig_InjectTraceID_Nil(t *testing.T) {
	var cfg *TraceIDConfig
	ctx, id, err := cfg.InjectTraceID(context.Background())
	if err != nil {
		t.Fatalf("unexpected error for nil config: %v", err)
	}
	if id != "" {
		t.Errorf("expected empty ID for nil config, got %q", id)
	}
	if ctx == nil {
		t.Error("expected non-nil context")
	}
}
