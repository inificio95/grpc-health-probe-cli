# Keepalive Flags

The `grpc-health-probe` CLI supports configuring gRPC keepalive parameters to maintain long-lived connections.

## Flags

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--keepalive-time` | duration | `0` (disabled) | Interval between keepalive pings |
| `--keepalive-timeout` | duration | `20s` | Time to wait for a keepalive ping ack |
| `--keepalive-permit-without-stream` | bool | `false` | Allow pings even without active RPCs |

## Behavior

Keepalive is **disabled by default**. It is automatically enabled when `--keepalive-time` is set to a non-zero duration.

## Examples

### Enable keepalive with default timeout

```bash
grpc-health-probe --addr=localhost:50051 --keepalive-time=30s
```

### Enable keepalive with custom timeout

```bash
grpc-health-probe --addr=localhost:50051 \
  --keepalive-time=30s \
  --keepalive-timeout=10s
```

### Enable keepalive for idle connections

```bash
grpc-health-probe --addr=localhost:50051 \
  --keepalive-time=60s \
  --keepalive-permit-without-stream=true
```

## Notes

- The `--keepalive-time` value must be at least `1s` when enabled.
- Setting `--keepalive-permit-without-stream` is useful when probing services that may be idle.
- Keepalive settings interact with server-side keepalive enforcement policies.
