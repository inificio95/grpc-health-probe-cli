package probe

import (
	"errors"
	"testing"
	"time"
)

func TestStatusString(t *testing.T) {
	tests := []struct {
		status Status
		want   string
	}{
		{StatusServing, "SERVING"},
		{StatusNotServing, "NOT_SERVING"},
		{StatusUnknown, "UNKNOWN"},
		{Status(99), "UNKNOWN"},
	}
	for _, tc := range tests {
		t.Run(tc.want, func(t *testing.T) {
			if got := tc.status.String(); got != tc.want {
				t.Errorf("Status.String() = %q, want %q", got, tc.want)
			}
		})
	}
}

func TestResultOK(t *testing.T) {
	tests := []struct {
		name   string
		result Result
		want   bool
	}{
		{
			name:   "serving no error",
			result: Result{Status: StatusServing, Err: nil},
			want:   true,
		},
		{
			name:   "not serving no error",
			result: Result{Status: StatusNotServing, Err: nil},
			want:   false,
		},
		{
			name:   "serving with error",
			result: Result{Status: StatusServing, Err: errors.New("oops")},
			want:   false,
		},
		{
			name:   "unknown no error",
			result: Result{Status: StatusUnknown, Err: nil},
			want:   false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.result.OK(); got != tc.want {
				t.Errorf("Result.OK() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestResultFields(t *testing.T) {
	r := Result{
		Address:  "localhost:50051",
		Service:  "my.Service",
		Status:   StatusServing,
		Attempts: 2,
		Duration: 150 * time.Millisecond,
		Err:      nil,
	}
	if r.Address != "localhost:50051" {
		t.Errorf("unexpected Address: %s", r.Address)
	}
	if r.Attempts != 2 {
		t.Errorf("unexpected Attempts: %d", r.Attempts)
	}
	if r.Duration != 150*time.Millisecond {
		t.Errorf("unexpected Duration: %v", r.Duration)
	}
}
