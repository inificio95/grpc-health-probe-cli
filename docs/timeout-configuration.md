# Timeout Configuration

The `grpc-health-probe-cli` supports fine-grained timeout control via `TimeoutConfig`.

## Fields

| Field            | Default | Description                                                  |
|------------------|---------|--------------------------------------------------------------|
| `dial_timeout`   | `5s`    | Maximum time to establish a gRPC connection to the target.   |
| `request_timeout`| `10s`   | Maximum time for a single health-check RPC to complete.      |

## CLI Flags

```
--dial-timeout duration      Timeout for establishing the gRPC connection (default 5s)
--request-timeout duration   Timeout for the health check RPC call (default 10s)
```

## Behaviour

- If the connection cannot be established within `dial_timeout`, the probe fails immediately.
- If the health RPC does not respond within `request_timeout`, the attempt is considered failed.
- When combined with [retry logic](./retry-configuration.md), each individual attempt is subject
  to its own `request_timeout`; `dial_timeout` applies once per probe invocation.

## Example

```bash
grpc-health-probe --addr=localhost:50051 \
  --dial-timeout=3s \
  --request-timeout=2s \
  --max-retries=3
```

In this example the tool will:
1. Attempt to connect within **3 seconds**.
2. Allow each health-check RPC up to **2 seconds**.
3. Retry up to **3 times** before reporting failure.
