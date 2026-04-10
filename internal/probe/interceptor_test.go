package probe

import (
	"testing"
)

func TestDefaultInterceptorConfig(t *testing.T) {
	cfg := DefaultInterceptorConfig()
	if cfg == nil {
		t.Fatal("expected non-nil config")
	}
	if cfg.EnableLogging {
		t.Errorf("expected EnableLogging=false, got true")
	}
	if cfg.LogPrefix == "" {
		t.Errorf("expected non-empty LogPrefix")
	}
}

func TestInterceptorConfig_Validate_Valid(t *testing.T) {
	cfg := DefaultInterceptorConfig()
	if err := cfg.Validate(); err != nil {
		t.Errorf("expected no error, got: %v", err)
	}
}

func TestInterceptorConfig_Validate_Nil(t *testing.T) {
	var cfg *InterceptorConfig
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for nil config, got nil")
	}
}

func TestInterceptorConfig_Validate_EmptyPrefix(t *testing.T) {
	cfg := DefaultInterceptorConfig()
	cfg.LogPrefix = ""
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for empty LogPrefix, got nil")
	}
}

func TestInterceptorConfig_UnaryClientInterceptor_Disabled(t *testing.T) {
	cfg := DefaultInterceptorConfig()
	cfg.EnableLogging = false
	if interceptor := cfg.UnaryClientInterceptor(); interceptor != nil {
		t.Error("expected nil interceptor when logging disabled")
	}
}

func TestInterceptorConfig_UnaryClientInterceptor_Enabled(t *testing.T) {
	cfg := DefaultInterceptorConfig()
	cfg.EnableLogging = true
	if interceptor := cfg.UnaryClientInterceptor(); interceptor == nil {
		t.Error("expected non-nil interceptor when logging enabled")
	}
}

func TestInterceptorConfig_UnaryClientInterceptor_Nil(t *testing.T) {
	var cfg *InterceptorConfig
	if interceptor := cfg.UnaryClientInterceptor(); interceptor != nil {
		t.Error("expected nil interceptor for nil config")
	}
}

func TestInterceptorConfig_DialOptions_Disabled(t *testing.T) {
	cfg := DefaultInterceptorConfig()
	cfg.EnableLogging = false
	opts := cfg.DialOptions()
	if len(opts) != 0 {
		t.Errorf("expected 0 dial options when logging disabled, got %d", len(opts))
	}
}

func TestInterceptorConfig_DialOptions_Enabled(t *testing.T) {
	cfg := DefaultInterceptorConfig()
	cfg.EnableLogging = true
	opts := cfg.DialOptions()
	if len(opts) != 1 {
		t.Errorf("expected 1 dial option when logging enabled, got %d", len(opts))
	}
}

func TestInterceptorConfig_DialOptions_Nil(t *testing.T) {
	var cfg *InterceptorConfig
	if opts := cfg.DialOptions(); opts != nil {
		t.Error("expected nil dial options for nil config")
	}
}
