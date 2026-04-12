package probe

import "fmt"

// ExitCode represents a CLI exit code returned after a probe attempt.
type ExitCode int

const (
	// ExitOK indicates the service is healthy.
	ExitOK ExitCode = 0
	// ExitNotServing indicates the service is not serving.
	ExitNotServing ExitCode = 1
	// ExitUnknown indicates the health status is unknown.
	ExitUnknown ExitCode = 2
	// ExitConnectionFailure indicates a connection or transport error.
	ExitConnectionFailure ExitCode = 3
	// ExitTimeout indicates the probe timed out.
	ExitTimeout ExitCode = 4
	// ExitConfigError indicates invalid configuration was provided.
	ExitConfigError ExitCode = 5
)

// DefaultExitCodeConfig returns an ExitCodeConfig with sensible defaults.
func DefaultExitCodeConfig() *ExitCodeConfig {
	return &ExitCodeConfig{
		Enabled: true,
		NotServingCode: int(ExitNotServing),
		UnknownCode:    int(ExitUnknown),
		TimeoutCode:    int(ExitTimeout),
	}
}

// ExitCodeConfig controls how exit codes are mapped from probe results.
type ExitCodeConfig struct {
	Enabled        bool
	NotServingCode int
	UnknownCode    int
	TimeoutCode    int
}

// Validate checks that the ExitCodeConfig is valid.
func (c *ExitCodeConfig) Validate() error {
	if c == nil {
		return fmt.Errorf("exitcode config must not be nil")
	}
	for name, code := range map[string]int{
		"not-serving": c.NotServingCode,
		"unknown":     c.UnknownCode,
		"timeout":     c.TimeoutCode,
	} {
		if code < 0 || code > 125 {
			return fmt.Errorf("exitcode config: %s code %d is out of range [0, 125]", name, code)
		}
	}
	return nil
}

// Resolve maps a Result to the appropriate exit code.
func (c *ExitCodeConfig) Resolve(r Result) ExitCode {
	if !c.Enabled {
		return ExitOK
	}
	if r.Error != nil {
		return ExitConnectionFailure
	}
	switch r.Status {
	case StatusServing:
		return ExitOK
	case StatusNotServing:
		return ExitCode(c.NotServingCode)
	case StatusUnknown:
		return ExitCode(c.UnknownCode)
	default:
		return ExitConnectionFailure
	}
}
