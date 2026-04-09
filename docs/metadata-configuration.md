# Metadata Configuration

The `grpc-health-probe-cli` supports sending custom gRPC metadata headers with each health check request. This is useful for services that require authentication tokens, trace IDs, or other context propagated via metadata.

## Usage

Pass one or more `--header` flags in `key=value` format:

```bash
grpc-health-probe \
  --addr=localhost:50051 \
  --header="authorization=Bearer my-token" \
  --header="x-request-id=abc123"
```

## Format

Each header must follow the `key=value` format:

| Part  | Description                        |
|-------|------------------------------------|
| `key` | The gRPC metadata key (non-empty)  |
| `value` | The associated metadata value    |

Duplicate keys are supported — multiple values will be sent for the same key.

## Validation Rules

- Each entry must contain exactly one `=` separator.
- The key portion must not be empty or whitespace-only.
- An empty list of headers is valid (no metadata will be sent).

## Example: Service with Auth

```bash
grpc-health-probe \
  --addr=my-service:443 \
  --tls \
  --header="authorization=Bearer eyJhbGci..." \
  --format=json
```

This is particularly useful in environments where gRPC interceptors enforce metadata-based authentication on all incoming RPCs, including health checks.
