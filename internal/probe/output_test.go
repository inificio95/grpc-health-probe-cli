package probe

import (
	"bytes"
	"strings"
	"testing"
	"time"

	grpc_health_v1 "google.golang.org/grpc/health/grpc_health_v1"
)

func TestDefaultOutputConfig(t *testing.T) {
	cfg := DefaultOutputConfig()
	if cfg.Writer == nil {
		t.Error("expected non-nil Writer")
	}
	if cfg.ErrWriter == nil {
		t.Error("expected non-nil ErrWriter")
	}
	if cfg.Format != "text" {
		t.Errorf("expected format \"text\", got %q", cfg.Format)
	}
	if cfg.Verbose {
		t.Error("expected Verbose to be false by default")
	}
}

func TestOutputConfig_Validate_Valid(t *testing.T) {
	cfg := DefaultOutputConfig()
	if err := cfg.Validate(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestOutputConfig_Validate_NilWriter(t *testing.T) {
	cfg := DefaultOutputConfig()
	cfg.Writer = nil
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for nil Writer")
	}
}

func TestOutputConfig_Validate_InvalidFormat(t *testing.T) {
	cfg := DefaultOutputConfig()
	cfg.Format = "xml"
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for unsupported format")
	}
}

func TestOutputConfig_Write_Text(t *testing.T) {
	var buf bytes.Buffer
	cfg := DefaultOutputConfig()
	cfg.Writer = &buf

	r := Result{
		Address:  "localhost:50051",
		Status:   grpc_health_v1.HealthCheckResponse_SERVING,
		Duration: 5 * time.Millisecond,
	}
	if err := cfg.Write(r); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(buf.String(), "SERVING") {
		t.Errorf("expected output to contain SERVING, got: %s", buf.String())
	}
}

func TestOutputConfig_Write_JSON(t *testing.T) {
	var buf bytes.Buffer
	cfg := DefaultOutputConfig()
	cfg.Writer = &buf
	cfg.Format = "json"

	r := Result{
		Address:  "localhost:50051",
		Status:   grpc_health_v1.HealthCheckResponse_SERVING,
		Duration: 5 * time.Millisecond,
	}
	if err := cfg.Write(r); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(buf.String(), `"status"`) {
		t.Errorf("expected JSON output, got: %s", buf.String())
	}
}

func TestOutputConfig_Write_Verbose(t *testing.T) {
	var out, errBuf bytes.Buffer
	cfg := DefaultOutputConfig()
	cfg.Writer = &out
	cfg.ErrWriter = &errBuf
	cfg.Verbose = true

	r := Result{
		Address:  "localhost:50051",
		Status:   grpc_health_v1.HealthCheckResponse_NOT_SERVING,
		Duration: 2 * time.Millisecond,
		Error:    fmt.Errorf("service unavailable"),
	}
	_ = cfg.Write(r)
	if !strings.Contains(errBuf.String(), "service unavailable") {
		t.Errorf("expected verbose error detail in errWriter, got: %s", errBuf.String())
	}
}
