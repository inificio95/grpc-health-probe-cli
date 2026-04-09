package probe

import (
	"context"
	"testing"
)

func TestDefaultConnectConfig(t *testing.T) {
	cfg := DefaultConnectConfig()
	if cfg == nil {
		t.Fatal("expected non-nil config")
	}
	if cfg.UserAgent == "" {
		t.Error("expected non-empty UserAgent")
	}
	if cfg.Block {
		t.Error("expected Block to be false by default")
	}
	if cfg.FailOnNonTempDialError {
		t.Error("expected FailOnNonTempDialError to be false by default")
	}
}

func TestConnectConfig_Validate_Nil(t *testing.T) {
	var cfg *ConnectConfig
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for nil config")
	}
}

func TestConnectConfig_Validate_Valid(t *testing.T) {
	cfg := DefaultConnectConfig()
	if err := cfg.Validate(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestConnectConfig_DialOptions_NoTLS(t *testing.T) {
	cfg := DefaultConnectConfig()
	tls := DefaultTLSConfig()

	opts, err := cfg.DialOptions(context.Background(), tls)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(opts) == 0 {
		t.Error("expected at least one dial option")
	}
}

func TestConnectConfig_DialOptions_WithBlock(t *testing.T) {
	cfg := DefaultConnectConfig()
	cfg.Block = true
	tls := DefaultTLSConfig()

	opts, err := cfg.DialOptions(context.Background(), tls)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// UserAgent + Block + insecure transport = 3 options
	if len(opts) < 3 {
		t.Errorf("expected at least 3 dial options, got %d", len(opts))
	}
}

func TestConnectConfig_DialOptions_EmptyUserAgent(t *testing.T) {
	cfg := DefaultConnectConfig()
	cfg.UserAgent = ""
	tls := DefaultTLSConfig()

	opts, err := cfg.DialOptions(context.Background(), tls)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// Only insecure transport when no user-agent and no block
	if len(opts) != 1 {
		t.Errorf("expected 1 dial option, got %d", len(opts))
	}
}
