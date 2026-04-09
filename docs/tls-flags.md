# TLS Flags

The `grpc-health-probe` CLI supports TLS configuration via command-line flags, allowing you to secure connections to gRPC services.

## Available Flags

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--tls` | bool | `false` | Enable TLS for the connection |
| `--tls-insecure-skip-verify` | bool | `false` | Skip server certificate verification (insecure) |
| `--tls-ca-cert` | string | `""` | Path to a CA certificate file to verify the server |
| `--tls-client-cert` | string | `""` | Path to a client certificate file for mutual TLS (mTLS) |
| `--tls-client-key` | string | `""` | Path to a client key file for mutual TLS (mTLS) |
| `--tls-server-name` | string | `""` | Override the server name used for TLS SNI |

## Examples

### Basic TLS (server certificate verification)

```bash
grpc-health-probe --addr=myservice:443 --tls
```

### TLS with a custom CA certificate

```bash
grpc-health-probe --addr=myservice:443 --tls --tls-ca-cert=/etc/ssl/my-ca.crt
```

### Mutual TLS (mTLS)

```bash
grpc-health-probe --addr=myservice:443 \
  --tls \
  --tls-ca-cert=/etc/ssl/ca.crt \
  --tls-client-cert=/etc/ssl/client.crt \
  --tls-client-key=/etc/ssl/client.key
```

### Skip certificate verification (development only)

```bash
grpc-health-probe --addr=localhost:50051 --tls --tls-insecure-skip-verify
```

> **Warning:** Using `--tls-insecure-skip-verify` disables certificate validation and should never be used in production environments.

## Notes

- `--tls-client-cert` and `--tls-client-key` must be provided together for mTLS.
- `--tls-server-name` is useful when the server's certificate CN/SAN does not match the address used to connect.
