package probe

import (
	"testing"
)

func TestDefaultHooksConfig(t *testing.T) {
	cfg := DefaultHooksConfig()
	if cfg == nil {
		t.Fatal("expected non-nil config")
	}
	if cfg.Enabled {
		t.Error("expected hooks to be disabled by default")
	}
	if cfg.PreCheck != nil || cfg.PostCheck != nil {
		t.Error("expected nil hook functions by default")
	}
}

func TestHooksConfig_Validate_Nil(t *testing.T) {
	var cfg *HooksConfig
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for nil config")
	}
}

func TestHooksConfig_Validate_DisabledNoHooks(t *testing.T) {
	cfg := DefaultHooksConfig()
	if err := cfg.Validate(); err != nil {
		t.Errorf("unexpected error for disabled hooks: %v", err)
	}
}

func TestHooksConfig_Validate_EnabledNoHooks(t *testing.T) {
	cfg := &HooksConfig{Enabled: true}
	if err := cfg.Validate(); err == nil {
		t.Error("expected error when enabled but no hooks provided")
	}
}

func TestHooksConfig_Validate_EnabledWithHooks(t *testing.T) {
	cfg := &HooksConfig{
		Enabled:   true,
		PreCheck:  func(_ HookEvent, _ *Result) {},
		PostCheck: func(_ HookEvent, _ *Result) {},
	}
	if err := cfg.Validate(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestHooksConfig_RunPreCheck_Invoked(t *testing.T) {
	called := false
	cfg := &HooksConfig{
		Enabled: true,
		PreCheck: func(event HookEvent, _ *Result) {
			if event != HookEventPreCheck {
				t.Errorf("expected pre_check event, got %s", event)
			}
			called = true
		},
	}
	cfg.RunPreCheck(nil)
	if !called {
		t.Error("expected PreCheck hook to be called")
	}
}

func TestHooksConfig_RunPostCheck_Invoked(t *testing.T) {
	called := false
	cfg := &HooksConfig{
		Enabled: true,
		PostCheck: func(event HookEvent, _ *Result) {
			if event != HookEventPostCheck {
				t.Errorf("expected post_check event, got %s", event)
			}
			called = true
		},
	}
	cfg.RunPostCheck(nil)
	if !called {
		t.Error("expected PostCheck hook to be called")
	}
}

func TestHooksConfig_RunPreCheck_Disabled(t *testing.T) {
	called := false
	cfg := &HooksConfig{
		Enabled:  false,
		PreCheck: func(_ HookEvent, _ *Result) { called = true },
	}
	cfg.RunPreCheck(nil)
	if called {
		t.Error("expected PreCheck hook NOT to be called when disabled")
	}
}

func TestHooksConfig_RunPreCheck_NilConfig(t *testing.T) {
	var cfg *HooksConfig
	// should not panic
	cfg.RunPreCheck(nil)
	cfg.RunPostCheck(nil)
}
