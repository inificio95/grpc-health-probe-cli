# Service Configuration

The `ServiceConfig` struct controls which gRPC service health endpoint is probed.

## Fields

| Field | Type | Default | Description |
|-------|------|---------|-------------|
| `Service` | `string` | `""` | Fully-qualified gRPC service name. Empty means server-level health. |
| `WatchMode` | `bool` | `false` | Use streaming Watch RPC instead of a single Check. |

## Usage

### Check overall server health (default)

```go
cfg := probe.DefaultServiceConfig()
// cfg.Service == "" — targets the server-level health endpoint
```

### Check a specific named service

```go
cfg := probe.ServiceConfig{
    Service: "my.package.MyService",
}
```

### Enable watch mode

```go
cfg := probe.ServiceConfig{
    Service:   "my.package.MyService",
    WatchMode: true,
}
```

## CLI flags

```
--service   gRPC service name to check (default: server-level)
--watch     Use streaming Watch RPC instead of unary Check
```

## Notes

- An empty `Service` name is valid and instructs the health server to return
  its overall readiness status.
- Unknown service names will be rejected by the remote server with
  `NOT_FOUND`; this is surfaced as a non-serving result.
