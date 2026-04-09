# Watch Flags

The watch flags control continuous polling of a gRPC health endpoint at a
configurable interval.

## Flags

| Flag              | Type     | Default | Description                                      |
|-------------------|----------|---------|--------------------------------------------------|
| `--watch`         | bool     | false   | Enable watch mode to poll the endpoint repeatedly |
| `--watch-interval`| duration | 10s     | Interval between health checks in watch mode      |

## Usage

```bash
# Poll every 5 seconds
grpc-health-probe --addr=localhost:50051 --watch --watch-interval=5s

# Use default 10s interval
grpc-health-probe --addr=localhost:50051 --watch
```

## Behavior

- When `--watch` is disabled (default), the probe runs once and exits.
- When `--watch` is enabled, the probe polls the endpoint at the specified
  interval until interrupted (e.g. `Ctrl+C`).
- The `--watch-interval` must be greater than zero when watch mode is enabled.
- Exit codes reflect the **last** observed health status.

## Validation

- `--watch-interval` must be a positive duration when `--watch` is `true`.
- If `--watch` is `false`, the interval value is ignored.

## Exit Codes

| Code | Meaning                        |
|------|--------------------------------|
| 0    | Service is SERVING             |
| 1    | Service is NOT_SERVING or unknown |
| 2    | Connection or probe error      |
