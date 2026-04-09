# Timeout Flags

The `grpc-health-probe` CLI provides several timeout-related flags to control the timing behavior of health check operations.

## Available Flags

### `--timeout`

Sets the timeout for the health check request itself.

**Default:** 5 seconds

**Example:**
```bash
grpc-health-probe --addr=localhost:9090 --timeout=10s
```

### `--dial-timeout`

Sets the timeout for establishing a connection to the gRPC server.

**Default:** 3 seconds

**Example:**
```bash
grpc-health-probe --addr=localhost:9090 --dial-timeout=5s
```

### `--connect-timeout`

Sets the timeout for the entire connection process, including dialing and initial setup.

**Default:** 5 seconds

**Example:**
```bash
grpc-health-probe --addr=localhost:9090 --connect-timeout=8s
```

## Combined Usage

You can combine multiple timeout flags for fine-grained control:

```bash
grpc-health-probe \
  --addr=localhost:9090 \
  --timeout=15s \
  --dial-timeout=5s \
  --connect-timeout=10s
```

## Duration Format

All timeout flags accept Go duration strings:
- `5s` - 5 seconds
- `500ms` - 500 milliseconds
- `1m` - 1 minute
- `1m30s` - 1 minute 30 seconds

## Timeout Hierarchy

1. **Dial Timeout**: Controls how long to wait for the TCP/network connection
2. **Connect Timeout**: Controls the overall connection establishment (includes dial)
3. **Request Timeout**: Controls how long to wait for the health check RPC response

## Best Practices

- Set `dial-timeout` lower than `connect-timeout`
- Set `request-timeout` based on expected service response time
- In production, use conservative timeouts to avoid false negatives
- For slow networks, increase `dial-timeout` and `connect-timeout`
- For slow services, increase `request-timeout`

## See Also

- [Timeout Configuration](timeout-configuration.md) - Internal timeout configuration details
- [Retry Configuration](retry-configuration.md) - Retry behavior with timeouts
