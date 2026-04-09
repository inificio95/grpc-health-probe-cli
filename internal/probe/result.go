package probe

import "time"

// Status represents the health status returned by a gRPC health check.
type Status int

const (
	// StatusUnknown indicates the health status could not be determined.
	StatusUnknown Status = iota
	// StatusServing indicates the service is healthy and serving.
	StatusServing
	// StatusNotServing indicates the service is unhealthy.
	StatusNotServing
)

// String returns a human-readable representation of the Status.
func (s Status) String() string {
	switch s {
	case StatusServing:
		return "SERVING"
	case StatusNotServing:
		return "NOT_SERVING"
	default:
		return "UNKNOWN"
	}
}

// Result holds the outcome of a single health probe attempt.
type Result struct {
	// Address is the target that was probed.
	Address string
	// Service is the gRPC service name that was queried.
	Service string
	// Status is the resolved health status.
	Status Status
	// Attempts is the number of probe attempts made (including retries).
	Attempts int
	// Duration is the total elapsed time for all attempts.
	Duration time.Duration
	// Err holds any error encountered on the final attempt.
	Err error
}

// OK returns true when the result represents a healthy, serving endpoint.
func (r Result) OK() bool {
	return r.Err == nil && r.Status == StatusServing
}
