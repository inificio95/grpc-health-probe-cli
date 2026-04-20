# Throttle Configuration

The throttle feature limits the rate at which the probe sends health-check requests to the target gRPC service. This is useful in watch mode or high-frequency polling scenarios where overwhelming the target must be avoided.

## Flags

| Flag | Type | Default | Description |
|---|---|---|---|
| `--throttle` | bool | `false` | Enable request throttling |
| `--throttle-min-interval` | duration | `500ms` | Minimum time between consecutive probe calls |
| `--throttle-burst` | int | `1` | Number of burst requests allowed before throttling applies |

## How it works

When throttling is enabled the `Throttler` tracks the last probe invocation time. If a new probe call arrives before `min-interval` has elapsed, the call blocks until the interval has passed.

The `burst` value allows a small number of back-to-back calls without delay, consuming tokens from an internal bucket. Once the bucket is empty, the interval enforcement takes over.

## Example

```bash
grpc-health-probe \
  --addr=localhost:50051 \
  --throttle \
  --throttle-min-interval=200ms \
  --throttle-burst=2
```

## Validation rules

- When enabled, `min-interval` must be a positive duration.
- When enabled, `burst` must be at least `1`.
- Throttling is silently skipped when disabled (default).
