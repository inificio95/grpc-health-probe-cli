package main

import (
	"bytes"
	"context"
	"net"
	"strings"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func startTestServer(t *testing.T, status grpc_health_v1.HealthCheckResponse_ServingStatus) string {
	t.Helper()

	lis, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("failed to listen: %v", err)
	}

	srv := grpc.NewServer()
	healthSrv := health.NewServer()
	healthSrv.SetServingStatus("", status)
	grpc_health_v1.RegisterHealthServer(srv, healthSrv)

	go func() {
		_ = srv.Serve(lis)
	}()
	t.Cleanup(srv.GracefulStop)

	return lis.Addr().String()
}

func TestRun_Serving(t *testing.T) {
	addr := startTestServer(t, grpc_health_v1.HealthCheckResponse_SERVING)

	cmd := newRootCmd()
	buf := &bytes.Buffer{}
	cmd.SetOut(buf)
	cmd.SetArgs([]string{"--addr", addr, "--format", "text"})

	err := cmd.ExecuteContext(context.Background())
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if !strings.Contains(buf.String(), "SERVING") {
		t.Errorf("expected output to contain SERVING, got: %s", buf.String())
	}
}

func TestRun_NotServing(t *testing.T) {
	addr := startTestServer(t, grpc_health_v1.HealthCheckResponse_NOT_SERVING)

	cmd := newRootCmd()
	cmd.SetOut(&bytes.Buffer{})
	cmd.SetArgs([]string{"--addr", addr})

	err := cmd.ExecuteContext(context.Background())
	if err == nil {
		t.Fatal("expected error for unhealthy service, got nil")
	}
}

func TestRun_MissingAddr(t *testing.T) {
	cmd := newRootCmd()
	cmd.SetOut(&bytes.Buffer{})
	cmd.SetErr(&bytes.Buffer{})
	cmd.SetArgs([]string{})

	err := cmd.ExecuteContext(context.Background())
	if err == nil {
		t.Fatal("expected error when addr flag is missing")
	}
}

func TestRun_JSONFormat(t *testing.T) {
	addr := startTestServer(t, grpc_health_v1.HealthCheckResponse_SERVING)

	cmd := newRootCmd()
	buf := &bytes.Buffer{}
	cmd.SetOut(buf)
	cmd.SetArgs([]string{"--addr", addr, "--format", "json"})

	_ = cmd.ExecuteContext(context.Background())

	if !strings.Contains(buf.String(), `"healthy"`) {
		t.Errorf("expected JSON output with healthy field, got: %s", buf.String())
	}
}
