# Authentication Configuration

The `grpc-health-probe-cli` supports configurable authentication for gRPC health check calls.

## Auth Types

### None (default)

No authentication is applied. This is the default behaviour.

```
grpc-health-probe --addr=localhost:50051
```

### Bearer Token

Attaches an `Authorization: Bearer <token>` header to every RPC call.

```
grpc-health-probe --addr=localhost:50051 \
  --auth-type=bearer \
  --auth-token=my-secret-token
```

### Basic Auth

Attaches an `Authorization: Basic <base64(user:pass)>` header to every RPC call.

```
grpc-health-probe --addr=localhost:50051 \
  --auth-type=basic \
  --auth-username=admin \
  --auth-password=secret
```

## Notes

- `RequireTransportSecurity` is set to `false` for both credential types, allowing use over plaintext connections. For production use, combine with TLS (see `tls-configuration.md`).
- Auth headers are sent on every RPC, including retries.
