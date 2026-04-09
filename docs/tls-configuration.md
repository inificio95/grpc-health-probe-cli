# TLS Configuration

`grpc-health-probe-cli` supports TLS-secured gRPC connections via the `--tls*` family of flags.

## Flags

| Flag | Default | Description |
|---|---|---|
| `--tls` | `false` | Enable TLS for the gRPC connection |
| `--tls-ca-cert` | `` | Path to a PEM-encoded CA certificate used to verify the server |
| `--tls-client-cert` | `` | Path to a PEM-encoded client certificate (mTLS) |
| `--tls-client-key` | `` | Path to the private key matching `--tls-client-cert` |
| `--tls-server-name` | `` | Override the server name used for TLS SNI / verification |
| `--tls-insecure-skip-verify` | `false` | Disable server certificate verification (**not for production**) |

## Examples

### Server-side TLS (custom CA)

```bash
grpc-health-probe \
  --addr=myservice.example.com:443 \
  --tls \
  --tls-ca-cert=/etc/certs/ca.pem
```

### Mutual TLS (mTLS)

```bash
grpc-health-probe \
  --addr=myservice.example.com:443 \
  --tls \
  --tls-ca-cert=/etc/certs/ca.pem \
  --tls-client-cert=/etc/certs/client.pem \
  --tls-client-key=/etc/certs/client-key.pem
```

### Skip verification (development only)

```bash
grpc-health-probe \
  --addr=localhost:50051 \
  --tls \
  --tls-insecure-skip-verify
```

> **Warning:** `--tls-insecure-skip-verify` disables all certificate validation and must never be used in production.
