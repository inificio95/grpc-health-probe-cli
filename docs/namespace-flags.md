# Namespace Flags

The namespace flags allow you to prefix the gRPC health-check service name
with a logical namespace, making it easier to distinguish services across
environments (e.g. `prod`, `staging`, `dev`).

## Flags

| Flag | Type | Default | Description |
|---|---|---|---|
| `--namespace-enabled` | bool | `false` | Enable namespace prefixing |
| `--namespace` | string | `""` | Namespace prefix (e.g. `prod`) |

## Behaviour

When `--namespace-enabled` is `true`, the resolved service name becomes
`<namespace>/<service>`. If the service name is empty (server-level check),
only the namespace is used.

When disabled (default), the service name is passed through unchanged.

## Examples

```bash
# Standard server-level health check — no namespace
grpc-health-probe --addr=localhost:50051

# Service-level check with namespace prefix
grpc-health-probe \
  --addr=localhost:50051 \
  --service=my-service \
  --namespace-enabled \
  --namespace=prod
# Effective service checked: prod/my-service

# Namespace only (server-level check with namespace)
grpc-health-probe \
  --addr=localhost:50051 \
  --namespace-enabled \
  --namespace=staging
```

## Validation

- If `--namespace-enabled` is `true`, `--namespace` must be a non-empty string.
- Namespace values should not contain leading/trailing slashes.
