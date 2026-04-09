package probe

import "fmt"

// VerbosityLevel represents how much detail to include in output.
type VerbosityLevel int

const (
	// VerbosityQuiet suppresses all output except errors.
	VerbosityQuiet VerbosityLevel = iota
	// VerbosityNormal is the default output level.
	VerbosityNormal
	// VerbosityVerbose includes additional diagnostic information.
	VerbosityVerbose
)

// VerbosityConfig holds verbosity settings for the probe.
type VerbosityConfig struct {
	Level VerbosityLevel
}

// DefaultVerbosityConfig returns a VerbosityConfig with sensible defaults.
func DefaultVerbosityConfig() *VerbosityConfig {
	return &VerbosityConfig{
		Level: VerbosityNormal,
	}
}

// Validate checks that the VerbosityConfig is valid.
func (v *VerbosityConfig) Validate() error {
	if v == nil {
		return fmt.Errorf("verbosity config must not be nil")
	}
	if v.Level < VerbosityQuiet || v.Level > VerbosityVerbose {
		return fmt.Errorf("invalid verbosity level: %d", v.Level)
	}
	return nil
}

// IsQuiet returns true when output should be suppressed.
func (v *VerbosityConfig) IsQuiet() bool {
	return v.Level == VerbosityQuiet
}

// IsVerbose returns true when verbose output is enabled.
func (v *VerbosityConfig) IsVerbose() bool {
	return v.Level == VerbosityVerbose
}
