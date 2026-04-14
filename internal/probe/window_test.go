package probe

import (
	"testing"
	"time"
)

func TestDefaultWindowConfig(t *testing.T) {
	cfg := DefaultWindowConfig()
	if cfg == nil {
		t.Fatal("expected non-nil config")
	}
	if cfg.Enabled {
		t.Error("expected Enabled to be false by default")
	}
	if cfg.Size != 10 {
		t.Errorf("expected Size 10, got %d", cfg.Size)
	}
	if cfg.Duration != 30*time.Second {
		t.Errorf("expected Duration 30s, got %v", cfg.Duration)
	}
}

func TestWindowConfig_Validate_Nil(t *testing.T) {
	var cfg *WindowConfig
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for nil config")
	}
}

func TestWindowConfig_Validate_Disabled(t *testing.T) {
	cfg := &WindowConfig{Enabled: false}
	if err := cfg.Validate(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestWindowConfig_Validate_InvalidSize(t *testing.T) {
	cfg := &WindowConfig{Enabled: true, Size: 0, Duration: time.Second}
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for zero size")
	}
}

func TestWindowConfig_Validate_InvalidDuration(t *testing.T) {
	cfg := &WindowConfig{Enabled: true, Size: 5, Duration: 0}
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for zero duration")
	}
}

func TestWindowConfig_Validate_Valid(t *testing.T) {
	cfg := &WindowConfig{Enabled: true, Size: 5, Duration: 10 * time.Second}
	if err := cfg.Validate(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestNewWindow_Disabled(t *testing.T) {
	cfg := &WindowConfig{Enabled: false, Size: 5, Duration: time.Second}
	w := NewWindow(cfg)
	w.Record(true)
	if w.Count() != 0 {
		t.Errorf("expected count 0 for disabled window, got %d", w.Count())
	}
	if r := w.SuccessRate(); r != 1.0 {
		t.Errorf("expected success rate 1.0 for empty window, got %f", r)
	}
}

func TestWindow_RecordAndSuccessRate(t *testing.T) {
	cfg := &WindowConfig{Enabled: true, Size: 4, Duration: time.Second}
	w := NewWindow(cfg)

	w.Record(true)
	w.Record(true)
	w.Record(false)
	w.Record(false)

	if w.Count() != 4 {
		t.Errorf("expected count 4, got %d", w.Count())
	}
	rate := w.SuccessRate()
	if rate != 0.5 {
		t.Errorf("expected success rate 0.5, got %f", rate)
	}
}

func TestWindow_SlidingOverwrite(t *testing.T) {
	cfg := &WindowConfig{Enabled: true, Size: 3, Duration: time.Second}
	w := NewWindow(cfg)

	w.Record(false)
	w.Record(false)
	w.Record(false)
	// overwrite all three with successes
	w.Record(true)
	w.Record(true)
	w.Record(true)

	if rate := w.SuccessRate(); rate != 1.0 {
		t.Errorf("expected success rate 1.0 after overwrite, got %f", rate)
	}
}
