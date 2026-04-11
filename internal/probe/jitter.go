package probe

import (
	"errors"
	"math/rand"
	"time"
)

// JitterConfig controls random jitter applied to retry delays.
type JitterConfig struct {
	// Enabled indicates whether jitter is applied to retry delays.
	Enabled bool

	// Factor is the fraction of the delay to use as the jitter range.
	// Must be between 0.0 and 1.0 when enabled.
	Factor float64

	// Seed is used to initialize the random source. Zero means use default.
	Seed int64

	rng *rand.Rand
}

// DefaultJitterConfig returns a JitterConfig with sensible defaults.
func DefaultJitterConfig() *JitterConfig {
	return &JitterConfig{
		Enabled: false,
		Factor:  0.25,
		Seed:    0,
	}
}

// Validate checks the JitterConfig for correctness.
func (j *JitterConfig) Validate() error {
	if j == nil {
		return errors.New("jitter config must not be nil")
	}
	if j.Enabled {
		if j.Factor <= 0.0 || j.Factor > 1.0 {
			return errors.New("jitter factor must be between 0.0 (exclusive) and 1.0 (inclusive) when enabled")
		}
	}
	return nil
}

// init initialises the internal RNG if not already set.
func (j *JitterConfig) init() {
	if j.rng == nil {
		seed := j.Seed
		if seed == 0 {
			seed = time.Now().UnixNano()
		}
		//nolint:gosec // non-cryptographic use
		j.rng = rand.New(rand.NewSource(seed))
	}
}

// Apply returns the delay with jitter added when enabled.
// The jitter is a random value in [0, delay*Factor).
func (j *JitterConfig) Apply(delay time.Duration) time.Duration {
	if j == nil || !j.Enabled || delay <= 0 {
		return delay
	}
	j.init()
	jitterRange := float64(delay) * j.Factor
	//nolint:gosec
	jitter := time.Duration(j.rng.Float64() * jitterRange)
	return delay + jitter
}
