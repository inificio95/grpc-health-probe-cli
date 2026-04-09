package probe

import "testing"

func TestDefaultConfig(t *testing.T) {
	cfg := DefaultConfig()
	if cfg.Format != "text" {
		t.Errorf("expected format \"text\", got %q", cfg.Format)
	}
	if cfg.TLS.Enabled {
		t.Error("expected TLS to be disabled by default")
	}
}

func TestConfigValidate(t *testing.T) {
	tests := []struct {
		name    string
		mutate  func(*Config)
		wantErr bool
	}{
		{
			name:    "valid config",
			mutate:  func(c *Config) { c.Address = "localhost:50051" },
			wantErr: false,
		},
		{
			name:    "missing address",
			mutate:  func(c *Config) {},
			wantErr: true,
		},
		{
			name:    "invalid format",
			mutate:  func(c *Config) { c.Address = "localhost:50051"; c.Format = "xml" },
			wantErr: true,
		},
		{
			name: "tls mismatch propagated",
			mutate: func(c *Config) {
				c.Address = "localhost:50051"
				c.TLS = TLSConfig{Enabled: true, ClientCertFile: "cert.pem"}
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := DefaultConfig()
			tt.mutate(&cfg)
			err := cfg.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
