package probe

import "errors"

// Sentinel errors for BackoffConfig validation.
var (
	ErrInvalidBackoffStrategy = errors.New("backoff: invalid strategy, must be one of: fixed, exponential, linear")
	ErrInvalidInitialDelay    = errors.New("backoff: initial delay must be greater than zero")
	ErrMaxDelayTooSmall       = errors.New("backoff: max delay must be greater than or equal to initial delay")
	ErrInvalidMultiplier      = errors.New("backoff: exponential multiplier must be greater than 1.0")
)
