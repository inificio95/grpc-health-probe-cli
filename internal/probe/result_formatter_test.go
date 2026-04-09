package probe_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/yourorg/grpc-health-probe-cli/internal/probe"
)

func makeResult(status probe.Status, attempt int, dur time.Duration, err error) probe.Result {
	return probe.Result{
		Status:   status,
		Attempt:  attempt,
		Duration: dur,
		Err:      err,
	}
}

func TestFormatResult_Text_Serving(t *testing.T) {
	var buf bytes.Buffer
	r := makeResult(probe.StatusServing, 1, 42*time.Millisecond, nil)
	if err := probe.FormatResult(&buf, r, probe.OutputFormatText); err != nil {
		t.Fatalf("FormatResult: %v", err)
	}
	out := buf.String()
	if !strings.Contains(out, "SERVING") {
		t.Errorf("expected SERVING in output, got: %s", out)
	}
	if strings.Contains(out, "Error") {
		t.Errorf("unexpected Error field in output: %s", out)
	}
}

func TestFormatResult_Text_WithError(t *testing.T) {
	var buf bytes.Buffer
	r := makeResult(probe.StatusUnknown, 2, 10*time.Millisecond, errors.New("connection refused"))
	if err := probe.FormatResult(&buf, r, probe.OutputFormatText); err != nil {
		t.Fatalf("FormatResult: %v", err)
	}
	out := buf.String()
	if !strings.Contains(out, "connection refused") {
		t.Errorf("expected error message in output, got: %s", out)
	}
}

func TestFormatResult_JSON_Serving(t *testing.T) {
	var buf bytes.Buffer
	r := makeResult(probe.StatusServing, 1, 55*time.Millisecond, nil)
	if err := probe.FormatResult(&buf, r, probe.OutputFormatJSON); err != nil {
		t.Fatalf("FormatResult: %v", err)
	}
	var out map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &out); err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}
	if out["status"] != "SERVING" {
		t.Errorf("expected status SERVING, got %v", out["status"])
	}
	if _, ok := out["error"]; ok {
		t.Error("error field should be absent on success")
	}
}

func TestFormatResult_JSON_WithError(t *testing.T) {
	var buf bytes.Buffer
	r := makeResult(probe.StatusUnknown, 3, 5*time.Millisecond, errors.New("timeout"))
	if err := probe.FormatResult(&buf, r, probe.OutputFormatJSON); err != nil {
		t.Fatalf("FormatResult: %v", err)
	}
	var out map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &out); err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}
	if out["error"] != "timeout" {
		t.Errorf("expected error 'timeout', got %v", out["error"])
	}
}

func TestFormatResult_DefaultIsText(t *testing.T) {
	var buf bytes.Buffer
	r := makeResult(probe.StatusNotServing, 1, 1*time.Millisecond, nil)
	if err := probe.FormatResult(&buf, r, probe.OutputFormat("unknown")); err != nil {
		t.Fatalf("FormatResult: %v", err)
	}
	if !strings.Contains(buf.String(), "NOT_SERVING") {
		t.Errorf("expected NOT_SERVING in text output, got: %s", buf.String())
	}
}
