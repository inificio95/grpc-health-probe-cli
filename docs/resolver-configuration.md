# DNS Resolver Configuration

The resolver feature allows `grpc-health-probe` to pre-resolve the target
hostname via DNS before establishing a gRPC connection. This is useful when
you need to pin a connection to a specific IP address or prefer a particular
address family.

## Flags

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--resolve` | bool | `false` | Enable DNS pre-resolution of the target host |
| `--resolve-prefer-ipv6` | bool | `false` | Prefer AAAA records over A records |
| `--resolve-dns-server` | string | `""` | Custom DNS server (`host:port`, e.g. `8.8.8.8:53`) |

## Behaviour

- When `--resolve` is **disabled** (the default) the address is passed
  directly to the gRPC dialer and normal OS/gRPC resolution applies.
- When `--resolve` is **enabled** the CLI resolves the hostname once at
  startup and rewrites the target address to the resolved IP before dialing.
- If the resolved IP is an IPv6 address it is automatically wrapped in
  brackets so that the port separator is unambiguous (e.g. `[::1]:50051`).
- If the host portion of the target address is already an IP literal,
  resolution is skipped and the address is used as-is.

## Examples

```bash
# Resolve the hostname before dialing (IPv4 preferred)
grpc-health-probe --addr myservice.internal:50051 --resolve

# Prefer an IPv6 address
grpc-health-probe --addr myservice.internal:50051 --resolve --resolve-prefer-ipv6

# Use a custom DNS server
grpc-health-probe --addr myservice.internal:50051 \
  --resolve \
  --resolve-dns-server 8.8.8.8:53
```

## Notes

- `--resolve-prefer-ipv6` and `--resolve-dns-server` have no effect unless
  `--resolve` is also set.
- The custom resolver flag expects a `host:port` value. Omitting the port
  will produce a validation error.
