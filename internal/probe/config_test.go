package probe

import (
	"testing"
	"time"
)

func TestDefaultConfig(t *testing.T) {
	cfg := DefaultConfig()

	if cfg.Timeout != 5*time.Second {
		t.Errorf("expected default timeout 5s, got %v", cfg.Timeout)
	}
	if cfg.RetryMax != 0 {
		t.Errorf("expected default retry-max 0, got %d", cfg.RetryMax)
	}
	if cfg.RetryInterval != 1*time.Second {
		t.Errorf("expected default retry-interval 1s, got %v", cfg.RetryInterval)
	}
}

func TestConfigValidate(t *testing.T) {
	tests := []struct {
		name    string
		cfg     Config
		wantErr bool
	}{
		{
			name:    "valid minimal config",
			cfg:     Config{Address: "localhost:50051", Timeout: 5 * time.Second, RetryInterval: time.Second},
			wantErr: false,
		},
		{
			name:    "missing address",
			cfg:     Config{Timeout: 5 * time.Second},
			wantErr: true,
		},
		{
			name:    "zero timeout",
			cfg:     Config{Address: "localhost:50051", Timeout: 0},
			wantErr: true,
		},
		{
			name:    "negative retry-max",
			cfg:     Config{Address: "localhost:50051", Timeout: time.Second, RetryMax: -1},
			wantErr: true,
		},
		{
			name:    "tls-no-verify without tls",
			cfg:     Config{Address: "localhost:50051", Timeout: time.Second, TLSNoVerify: true},
			wantErr: true,
		},
		{
			name:    "tls-ca-cert without tls",
			cfg:     Config{Address: "localhost:50051", Timeout: time.Second, TLSCACert: "/path/to/ca.crt"},
			wantErr: true,
		},
		{
			name:    "valid tls config",
			cfg:     Config{Address: "localhost:50051", Timeout: time.Second, TLS: true, TLSNoVerify: true},
			wantErr: false,
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
