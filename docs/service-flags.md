# Service Flags

The `--service` flag controls which gRPC health check service is queried.

## Usage

```
grpc-health-probe --addr=localhost:50051 --service=mypackage.MyService
```

## Flags

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--service` | string | `""` (empty) | gRPC health service name to check |

## Server-Level vs Service-Level Checks

### Server-Level (default)

When `--service` is omitted or set to an empty string, the probe performs a
**server-level** health check. This queries the overall health of the gRPC
server rather than a specific service.

```bash
grpc-health-probe --addr=localhost:50051
```

### Service-Level

When `--service` is provided, the probe checks the health of the named service
registered with the server's health service.

```bash
grpc-health-probe --addr=localhost:50051 --service=mypackage.MyService
```

## Exit Codes

- `0` — Service (or server) is `SERVING`
- `1` — Service is `NOT_SERVING`, `UNKNOWN`, or the check failed

## Examples

```bash
# Check server-level health
grpc-health-probe --addr=localhost:50051

# Check a specific service
grpc-health-probe --addr=localhost:50051 --service=grpc.health.v1.Health

# JSON output for a named service
grpc-health-probe --addr=localhost:50051 --service=orders.OrderService --format=json
```
