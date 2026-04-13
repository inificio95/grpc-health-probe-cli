# Timeout Budget Configuration

The timeout budget feature enforces a **shared deadline** across all retry attempts, preventing the total probe time from exceeding a configured limit regardless of per-attempt timeouts or retry counts.

## Flags

| Flag | Type | Default | Description |
|---|---|---|---|
| `--budget-enabled` | bool | `false` | Enable the shared timeout budget |
| `--budget-total` | duration | `30s` | Total time budget for all attempts combined |
| `--budget-reserve` | duration | `500ms` | Buffer subtracted from the deadline before it is considered exhausted |

## How It Works

When `--budget-enabled` is set, a `Budget` is created at probe start. Before each retry attempt the prober checks `Budget.Remaining()`. If the usable time (total minus reserve buffer) is zero the probe stops immediately and reports a timeout error.

The reserve buffer ensures the probe has enough time to cleanly record and report the result before the hard deadline is reached.

## Example

```bash
# Allow up to 15 seconds total, with a 250 ms safety buffer
grpc-health-probe \
  --addr=localhost:50051 \
  --retry-max=10 \
  --budget-enabled \
  --budget-total=15s \
  --budget-reserve=250ms
```

## Validation Rules

- `TotalBudget` must be positive when the budget is enabled.
- `ReserveBuffer` must be non-negative and strictly less than `TotalBudget`.

## Interaction with Retry

The timeout budget works alongside `--retry-max` and per-attempt `--timeout`. Whichever limit is reached first (retry count or budget exhaustion) terminates the probe loop.
