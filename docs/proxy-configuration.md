# Proxy Configuration

The `grpc-health-probe` CLI supports routing gRPC connections through an HTTP or HTTPS proxy.

## Flags

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--proxy` | bool | `false` | Enable HTTP/HTTPS proxy for gRPC connections |
| `--proxy-url` | string | `""` | Proxy URL (e.g. `http://proxy.example.com:8080`) |

## Usage

```bash
grpc-health-probe \
  --addr=myservice:50051 \
  --proxy \
  --proxy-url=http://proxy.internal:3128
```

## Supported Schemes

Only `http` and `https` proxy URLs are supported. SOCKS5 or other proxy
protocols are not currently supported.

## Validation

- When `--proxy` is `false` (default), the `--proxy-url` flag is ignored.
- When `--proxy` is `true`, `--proxy-url` must be set to a valid `http://` or
  `https://` URL; otherwise the tool will exit with a validation error.

## Examples

### HTTP proxy

```bash
grpc-health-probe --addr=svc:50051 --proxy --proxy-url=http://10.0.0.1:8080
```

### HTTPS proxy

```bash
grpc-health-probe --addr=svc:50051 --proxy --proxy-url=https://secure-proxy.example.com:443
```

### No proxy (default)

```bash
grpc-health-probe --addr=svc:50051
```
