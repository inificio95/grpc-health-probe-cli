package probe_test

import (
	"context"
	"net"
	"testing"
	"time"

	"github.com/yourorg/grpc-health-probe-cli/internal/probe"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

// startHealthServer starts an in-process gRPC health server and returns its address.
// The server is automatically stopped when the test completes.
func startHealthServer(t *testing.T, status grpc_health_v1.HealthCheckResponse_ServingStatus) string {
	t.Helper()
	lis, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("failed to listen: %v", err)
	}
	srv := grpc.NewServer()
	healthSrv := health.NewServer()
	healthSrv.SetServingStatus("", status)
	grpc_health_v1.RegisterHealthServer(srv, healthSrv)
	go srv.Serve(lis) //nolint:errcheck
	t.Cleanup(srv.GracefulStop)
	return lis.Addr().String()
}

func TestProberCheck_Serving(t *testing.T) {
	addr := startHealthServer(t, grpc_health_v1.HealthCheckResponse_SERVING)
	cfg := probe.DefaultConfig()
	cfg.Address = addr
	p, err := probe.New(cfg)
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	res := p.Check(context.Background())
	if res.Err != nil {
		t.Fatalf("unexpected error: %v", res.Err)
	}
	if res.Status != probe.StatusServing {
		t.Errorf("expected SERVING, got %s", res.Status)
	}
	if res.Attempt != 1 {
		t.Errorf("expected attempt 1, got %d", res.Attempt)
	}
}

func TestProberCheck_NotServing(t *testing.T) {
	addr := startHealthServer(t, grpc_health_v1.HealthCheckResponse_NOT_SERVING)
	cfg := probe.DefaultConfig()
	cfg.Address = addr
	cfg.MaxRetries = 1
	cfg.RetryInterval = 10 * time.Millisecond
	p, err := probe.New(cfg)
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	res := p.Check(context.Background())
	if res.Status != probe.StatusNotServing {
		t.Errorf("expected NOT_SERVING, got %s", res.Status)
	}
	if res.Attempt != 2 {
		t.Errorf("expected 2 attempts after retry, got %d", res.Attempt)
	}
}

func TestProberCheck_InvalidAddress(t *testing.T) {
	cfg := probe.DefaultConfig()
	cfg.Address = "127.0.0.1:1" // nothing listening
	cfg.Timeout = 100 * time.Millisecond
	cfg.MaxRetries = 0
	p, err := probe.New(cfg)
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	res := p.Check(context.Background())
	if res.Err == nil {
		t.Error("expected error for unreachable address")
	}
}

func TestProberNew_InvalidConfig(t *testing.T) {
	cfg := probe.DefaultConfig()
	cfg.Address = ""
	_, err := probe.New(cfg)
	if err == nil {
		t.Error("expected error for empty address")
	}
}

func TestProberCheck_ContextCancelled(t *testing.T) {
	addr := startHealthServer(t, grpc_health_v1.HealthCheckResponse_SERVING)
	cfg := probe.DefaultConfig()
	cfg.Address = addr
	p, err := probe.New(cfg)
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // cancel immediately before calling Check
	res := p.Check(ctx)
	if res.Err == nil {
		t.Error("expected error when context is already cancelled")
	}
}

func TestStatusString(t *testing.T) {
	cases := []struct {
		s    probe.Status
		want string
	}{
		{probe.StatusServing, "SERVING"},
		{probe.StatusNotServing, "NOT_SERVING"},
		{probe.StatusServiceUnknown, "SERVICE_UNKNOWN"},
		{probe.StatusUnknown, "UNKNOWN"},
	}
	for _, tc := range cases {
		if got := tc.s.String(); got != tc.want {
			t.Errorf("Status(%d).String() = %q, want %q", tc.s, got, tc.want)
		}
	}
}
