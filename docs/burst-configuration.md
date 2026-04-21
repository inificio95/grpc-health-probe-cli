# Burst Configuration

The burst detection feature prevents probe storms by tracking how many health
checks are issued within a sliding time window and suppressing further attempts
once the configured maximum is exceeded.

## Flags

| Flag | Type | Default | Description |
|---|---|---|---|
| `--burst-enabled` | bool | `false` | Enable burst detection |
| `--burst-max` | int | `5` | Max probes allowed in the burst window |
| `--burst-interval` | duration | `500ms` | Sliding window size for burst counting |
| `--burst-cooldown` | duration | `2s` | Suppression period after burst is detected |

## Behaviour

When `--burst-enabled` is set:

1. Each probe attempt is recorded with a timestamp.
2. Attempts older than `--burst-interval` are discarded from the window.
3. If the number of recent attempts exceeds `--burst-max`, the probe is
   suppressed and the cooldown timer is started.
4. While the cooldown timer is active all subsequent probes are suppressed.
5. After the cooldown expires the window is reset and probes are allowed again.

## Example

```bash
grpc-health-probe \
  --addr=localhost:50051 \
  --burst-enabled \
  --burst-max=3 \
  --burst-interval=200ms \
  --burst-cooldown=5s
```

With the above settings, if more than 3 probes arrive within any 200 ms window
the tool will suppress probes for the next 5 seconds.

## Validation

- `--burst-max` must be greater than zero.
- `--burst-interval` must be a positive duration.
- `--burst-cooldown` must be a positive duration.
