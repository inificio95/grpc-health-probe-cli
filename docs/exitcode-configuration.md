# Exit Code Configuration

The `grpc-health-probe-cli` maps probe results to specific OS exit codes, enabling shell scripts and CI pipelines to react to health status without parsing output.

## Default Exit Codes

| Situation             | Exit Code |
|-----------------------|-----------|
| Service is SERVING    | `0`       |
| Service is NOT_SERVING| `1`       |
| Status is UNKNOWN     | `2`       |
| Connection failure    | `3`       |
| Probe timed out       | `4`       |
| Configuration error   | `5`       |

## Flags

| Flag                      | Default | Description                                      |
|---------------------------|---------|--------------------------------------------------|
| `--exit-code-enabled`     | `true`  | Enable result-to-exit-code mapping               |
| `--exit-code-not-serving` | `1`     | Exit code when the service reports NOT_SERVING   |
| `--exit-code-unknown`     | `2`     | Exit code when the service reports UNKNOWN       |
| `--exit-code-timeout`     | `4`     | Exit code when the probe request times out       |

## Examples

```bash
# Default behaviour — exit 1 when not healthy
grpc-health-probe --addr=localhost:50051
echo $?  # 0 = SERVING, 1 = NOT_SERVING

# Use custom exit codes for integration with a specific monitoring system
grpc-health-probe --addr=localhost:50051 \
  --exit-code-not-serving=2 \
  --exit-code-unknown=3

# Disable exit-code mapping — always exits 0 regardless of status
grpc-health-probe --addr=localhost:50051 --exit-code-enabled=false
```

## Disabling Exit Code Mapping

When `--exit-code-enabled=false` the tool always exits with code `0`. This is useful when you only care about the printed output and do not want the process to signal failure to a parent process.
