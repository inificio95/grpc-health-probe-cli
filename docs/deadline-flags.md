# Deadline Flags

The deadline flags allow you to set an overall time limit for the entire probe
operation, independent of per-request timeouts.

## Flags

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--deadline` | bool | `false` | Enable an overall deadline for the probe |
| `--deadline-duration` | duration | `10s` | Duration of the overall deadline |

## Usage

```bash
# Run with a 30-second overall deadline
grpc-health-probe --addr=localhost:50051 \
  --deadline \
  --deadline-duration=30s
```

## Relationship to Timeout Flags

- **Per-request timeout** (`--request-timeout`): limits each individual RPC.
- **Dial timeout** (`--dial-timeout`): limits the initial connection attempt.
- **Overall deadline** (`--deadline-duration`): caps the total wall-clock time
  across all retries and dial attempts combined.

When both retry logic and a deadline are active, the probe stops retrying as
soon as the deadline is exceeded, even if attempts remain.

## Validation

- `--deadline-duration` must be a positive duration when `--deadline` is set.
- Passing `--deadline-duration` without `--deadline` has no effect.
