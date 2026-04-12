package probe

import "fmt"

// HookEvent represents the lifecycle event at which a hook is triggered.
type HookEvent string

const (
	HookEventPreCheck  HookEvent = "pre_check"
	HookEventPostCheck HookEvent = "post_check"
)

// HookFunc is a function called at a lifecycle event during a probe check.
type HookFunc func(event HookEvent, result *Result)

// HooksConfig holds optional lifecycle hook functions for probe execution.
type HooksConfig struct {
	Enabled    bool
	PreCheck   HookFunc
	PostCheck  HookFunc
}

// DefaultHooksConfig returns a HooksConfig with hooks disabled.
func DefaultHooksConfig() *HooksConfig {
	return &HooksConfig{
		Enabled: false,
	}
}

// Validate checks that the HooksConfig is in a valid state.
func (h *HooksConfig) Validate() error {
	if h == nil {
		return fmt.Errorf("hooks config must not be nil")
	}
	if h.Enabled && h.PreCheck == nil && h.PostCheck == nil {
		return fmt.Errorf("hooks enabled but no hook functions provided")
	}
	return nil
}

// RunPreCheck invokes the PreCheck hook if hooks are enabled and the hook is set.
func (h *HooksConfig) RunPreCheck(result *Result) {
	if h == nil || !h.Enabled || h.PreCheck == nil {
		return
	}
	h.PreCheck(HookEventPreCheck, result)
}

// RunPostCheck invokes the PostCheck hook if hooks are enabled and the hook is set.
func (h *HooksConfig) RunPostCheck(result *Result) {
	if h == nil || !h.Enabled || h.PostCheck == nil {
		return
	}
	h.PostCheck(HookEventPostCheck, result)
}
