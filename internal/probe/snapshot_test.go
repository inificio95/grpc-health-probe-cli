package probe

import (
	"testing"
)

func TestDefaultSnapshotConfig(t *testing.T) {
	cfg := DefaultSnapshotConfig()
	if cfg == nil {
		t.Fatal("expected non-nil config")
	}
	if cfg.Enabled {
		t.Error("expected Enabled to be false")
	}
	if cfg.FilePath != "" {
		t.Errorf("expected empty FilePath, got %q", cfg.FilePath)
	}
	if cfg.Format != "json" {
		t.Errorf("expected Format \"json\", got %q", cfg.Format)
	}
}

func TestSnapshotConfig_Validate_Nil(t *testing.T) {
	var cfg *SnapshotConfig
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for nil config")
	}
}

func TestSnapshotConfig_Validate_Disabled(t *testing.T) {
	cfg := DefaultSnapshotConfig()
	if err := cfg.Validate(); err != nil {
		t.Errorf("unexpected error for disabled snapshot: %v", err)
	}
}

func TestSnapshotConfig_Validate_EnabledMissingPath(t *testing.T) {
	cfg := &SnapshotConfig{Enabled: true, FilePath: "", Format: "json"}
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for missing file path")
	}
}

func TestSnapshotConfig_Validate_InvalidFormat(t *testing.T) {
	cfg := &SnapshotConfig{Enabled: true, FilePath: "/tmp/snap.out", Format: "xml"}
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for invalid format")
	}
}

func TestSnapshotConfig_Validate_ValidJSON(t *testing.T) {
	cfg := &SnapshotConfig{Enabled: true, FilePath: "/tmp/snap.json", Format: "json"}
	if err := cfg.Validate(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestSnapshotConfig_Validate_ValidText(t *testing.T) {
	cfg := &SnapshotConfig{Enabled: true, FilePath: "/tmp/snap.txt", Format: "text"}
	if err := cfg.Validate(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}
