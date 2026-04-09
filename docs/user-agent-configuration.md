# User-Agent Configuration

The `grpc-health-probe-cli` tool sends a custom `user-agent` header with every gRPC request. This helps server operators identify traffic originating from health probes.

## Default Behaviour

By default, the user-agent string is formatted as:

```
grpc-health-probe-cli/0.1.0 (linux/amd64)
```

The OS and architecture fields are populated automatically at runtime.

## CLI Flags

| Flag | Default | Description |
|------|---------|-------------|
| `--user-agent-name` | `grpc-health-probe-cli` | Application name in the user-agent header |
| `--user-agent-version` | `0.1.0` | Application version in the user-agent header |
| `--no-user-agent` | `false` | Disable the custom user-agent header entirely |

## Examples

### Use a custom application name and version

```bash
grpc-health-probe \
  --addr=localhost:50051 \
  --user-agent-name=my-service \
  --user-agent-version=1.2.3
```

### Disable the user-agent header

```bash
grpc-health-probe \
  --addr=localhost:50051 \
  --no-user-agent
```

## Validation Rules

- When the user-agent is **enabled** (default), both `AppName` and `AppVersion` must be non-empty strings.
- When the user-agent is **disabled**, no validation is performed on name or version fields.
