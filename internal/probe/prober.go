package probe

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/health/grpc_health_v1"
)

// Status represents the health check result.
type Status int

const (
	StatusUnknown    Status = iota
	StatusServing
	StatusNotServing
	StatusServiceUnknown
)

func (s Status) String() string {
	switch s {
	case StatusServing:
		return "SERVING"
	case StatusNotServing:
		return "NOT_SERVING"
	case StatusServiceUnknown:
		return "SERVICE_UNKNOWN"
	default:
		return "UNKNOWN"
	}
}

// Result holds the outcome of a single health probe attempt.
type Result struct {
	Status   Status
	Attempt  int
	Duration time.Duration
	Err      error
}

// IsHealthy reports whether the result indicates the service is serving.
func (r Result) IsHealthy() bool {
	return r.Err == nil && r.Status == StatusServing
}

// Prober performs gRPC health checks against a target.
type Prober struct {
	cfg Config
}

// New creates a new Prober with the given Config.
func New(cfg Config) (*Prober, error) {
	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}
	return &Prober{cfg: cfg}, nil
}

// Check performs the health check with retry logic defined in the config.
func (p *Prober) Check(ctx context.Context) Result {
	var last Result
	for attempt := 1; attempt <= p.cfg.MaxRetries+1; attempt++ {
		last = p.doCheck(ctx, attempt)
		if last.IsHealthy() {
			return last
		}
		if attempt <= p.cfg.MaxRetries {
			select {
			case <-ctx.Done():
				last.Err = ctx.Err()
				return last
			case <-time.After(p.cfg.RetryInterval):
			}
		}
	}
	return last
}

func (p *Prober) doCheck(ctx context.Context, attempt int) Result {
	start := time.Now()
	tctx, cancel := context.WithTimeout(ctx, p.cfg.Timeout)
	defer cancel()

	conn, err := grpc.DialContext(tctx, p.cfg.Address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		return Result{Status: StatusUnknown, Attempt: attempt, Duration: time.Since(start), Err: fmt.Errorf("dial: %w", err)}
	}
	defer conn.Close()

	client := grpc_health_v1.NewHealthClient(conn)
	resp, err := client.Check(tctx, &grpc_health_v1.HealthCheckRequest{Service: p.cfg.Service})
	if err != nil {
		return Result{Status: StatusUnknown, Attempt: attempt, Duration: time.Since(start), Err: fmt.Errorf("health check: %w", err)}
	}

	return Result{
		Status:   grpcStatusToStatus(resp.GetStatus()),
		Attempt:  attempt,
		Duration: time.Since(start),
	}
}

func grpcStatusToStatus(s grpc_health_v1.HealthCheckResponse_ServingStatus) Status {
	switch s {
	case grpc_health_v1.HealthCheckResponse_SERVING:
		return StatusServing
	case grpc_health_v1.HealthCheckResponse_NOT_SERVING:
		return StatusNotServing
	case grpc_health_v1.HealthCheckResponse_SERVICE_UNKNOWN:
		return StatusServiceUnknown
	default:
		return StatusUnknown
	}
}
