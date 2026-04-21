package probe

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// DefaultSignalConfig returns a SignalConfig with safe defaults.
// By default, signal handling is enabled and listens for SIGINT and SIGTERM.
func DefaultSignalConfig() *SignalConfig {
	return &SignalConfig{
		Enabled: true,
		Signals: []os.Signal{syscall.SIGINT, syscall.SIGTERM},
	}
}

// SignalConfig controls how the probe responds to OS signals.
type SignalConfig struct {
	// Enabled determines whether signal handling is active.
	Enabled bool

	// Signals is the list of OS signals that will trigger a graceful shutdown.
	Signals []os.Signal
}

// Validate checks that the SignalConfig is valid.
func (c *SignalConfig) Validate() error {
	if c == nil {
		return fmt.Errorf("signal config must not be nil")
	}
	if c.Enabled && len(c.Signals) == 0 {
		return fmt.Errorf("signal config: at least one signal must be specified when enabled")
	}
	return nil
}

// WithSignalContext returns a derived context that is cancelled when one of
// the configured signals is received. If signal handling is disabled, the
// original context is returned unchanged along with a no-op cancel function.
//
// Callers are responsible for calling the returned CancelFunc to release
// resources even if the signal is never received.
func (c *SignalConfig) WithSignalContext(ctx context.Context) (context.Context, context.CancelFunc) {
	if c == nil || !c.Enabled || len(c.Signals) == 0 {
		return ctx, func() {}
	}

	ctx, cancel := context.WithCancel(ctx)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, c.Signals...)

	go func() {
		defer signal.Stop(ch)
		select {
		case <-ch:
			cancel()
		case <-ctx.Done():
		}
	}()

	return ctx, cancel
}

// String returns a human-readable description of the SignalConfig.
func (c *SignalConfig) String() string {
	if c == nil {
		return "signal(nil)"
	}
	if !c.Enabled {
		return "signal(disabled)"
	}
	names := make([]string, 0, len(c.Signals))
	for _, s := range c.Signals {
		if sig, ok := s.(syscall.Signal); ok {
			names = append(names, sig.String())
		} else {
			names = append(names, fmt.Sprintf("%v", s))
		}
	}
	return fmt.Sprintf("signal(enabled, signals=%v)", names)
}
