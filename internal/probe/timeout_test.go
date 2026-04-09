package probe

import (
	"context"
	"testing"
	"time"
)

func TestDefaultTimeoutConfig(t *testing.T) {
	cfg := DefaultTimeoutConfig()

	if cfg.DialTimeout != 5*time.Second {
		t.Errorf("expected DialTimeout 5s, got %s", cfg.DialTimeout)
	}
	if cfg.RequestTimeout != 10*time.Second {
		t.Errorf("expected RequestTimeout 10s, got %s", cfg.RequestTimeout)
	}
	if err := cfg.Validate(); err != nil {
		t.Errorf("default config should be valid, got: %v", err)
	}
}

func TestTimeoutConfigValidate(t *testing.T) {
	tests := []struct {
		name    string
		cfg     TimeoutConfig
		wantErr bool
	}{
		{
			name:    "valid config",
			cfg:     TimeoutConfig{DialTimeout: time.Second, RequestTimeout: time.Second},
			wantErr: false,
		},
		{
			name:    "zero dial timeout",
			cfg:     TimeoutConfig{DialTimeout: 0, RequestTimeout: time.Second},
			wantErr: true,
		},
		{
			name:    "negative dial timeout",
			cfg:     TimeoutConfig{DialTimeout: -1 * time.Second, RequestTimeout: time.Second},
			wantErr: true,
		},
		{
			name:    "zero request timeout",
			cfg:     TimeoutConfig{DialTimeout: time.Second, RequestTimeout: 0},
			wantErr: true,
		},
		{
			name:    "negative request timeout",
			cfg:     TimeoutConfig{DialTimeout: time.Second, RequestTimeout: -1 * time.Second},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.cfg.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTimeoutConfig_WithRequestTimeout(t *testing.T) {
	cfg := TimeoutConfig{DialTimeout: time.Second, RequestTimeout: 50 * time.Millisecond}
	ctx, cancel := cfg.WithRequestTimeout(context.Background())
	defer cancel()

	deadline, ok := ctx.Deadline()
	if !ok {
		t.Fatal("expected context to have a deadline")
	}
	remaining := time.Until(deadline)
	if remaining <= 0 || remaining > 50*time.Millisecond {
		t.Errorf("unexpected deadline remaining: %s", remaining)
	}
}

func TestTimeoutConfig_WithDialTimeout(t *testing.T) {
	cfg := TimeoutConfig{DialTimeout: 50 * time.Millisecond, RequestTimeout: time.Second}
	ctx, cancel := cfg.WithDialTimeout(context.Background())
	defer cancel()

	deadline, ok := ctx.Deadline()
	if !ok {
		t.Fatal("expected context to have a deadline")
	}
	remaining := time.Until(deadline)
	if remaining <= 0 || remaining > 50*time.Millisecond {
		t.Errorf("unexpected deadline remaining: %s", remaining)
	}
}
