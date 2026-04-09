# grpc-health-probe-cli

A lightweight CLI tool for checking gRPC service health endpoints with configurable retry and timeout logic.

---

## Installation

```bash
go install github.com/yourusername/grpc-health-probe-cli@latest
```

Or download a pre-built binary from the [Releases](https://github.com/yourusername/grpc-health-probe-cli/releases) page.

---

## Usage

```bash
grpc-health-probe-cli --addr=localhost:50051
```

### Options

| Flag | Default | Description |
|------|---------|-------------|
| `--addr` | `localhost:50051` | gRPC server address |
| `--service` | `""` | Service name to check (empty checks server health) |
| `--timeout` | `5s` | Timeout per request |
| `--retries` | `3` | Number of retry attempts |
| `--retry-delay` | `1s` | Delay between retries |
| `--tls` | `false` | Enable TLS |
| `--verbose` | `false` | Print detailed output for each attempt |

### Example

```bash
# Check a specific service with retries and custom timeout
grpc-health-probe-cli \
  --addr=myservice.example.com:443 \
  --service=myapp.UserService \
  --timeout=3s \
  --retries=5 \
  --tls
```

Exit codes:
- `0` — Service is healthy
- `1` — Service is unhealthy or unreachable
- `2` — Invalid arguments or configuration error

---

## Requirements

- Go 1.21+
- Target service must implement the [gRPC Health Checking Protocol](https://github.com/grpc/grpc/blob/master/doc/health-checking.md)

---

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

---

## License

[MIT](LICENSE)
