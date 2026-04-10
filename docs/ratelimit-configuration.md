# Rate Limit Configuration

The `RateLimitConfig` controls how many health probe requests are allowed within a sliding time window. This is useful when probing services that enforce server-side rate limits.

## Fields

| Field        | Type            | Default      | Description                                      |
|--------------|-----------------|--------------|--------------------------------------------------|
| `Enabled`    | `bool`          | `false`      | Whether rate limiting is active.                 |
| `MaxRequests`| `int`           | `10`         | Maximum number of requests per window.           |
| `WindowSize` | `time.Duration` | `1s`         | Duration of the sliding window.                  |

## Default Configuration

```go
cfg := probe.DefaultRateLimitConfig()
// Enabled:     false
// MaxRequests: 10
// WindowSize:  1s
```

## Enabling Rate Limiting

```go
cfg := &probe.RateLimitConfig{
    Enabled:     true,
    MaxRequests: 5,
    WindowSize:  time.Second,
}
if err := cfg.Validate(); err != nil {
    log.Fatal(err)
}
```

## Checking Allowance

Before issuing a probe request, call `Allow()` to check whether the request is permitted:

```go
if !cfg.Allow() {
    fmt.Println("rate limit exceeded, skipping probe")
    return
}
```

Use `Remaining()` to inspect how many requests are left in the current window:

```go
fmt.Printf("%d requests remaining in window\n", cfg.Remaining())
```

## Validation Rules

- Config must not be nil.
- When disabled, no further validation is performed.
- `MaxRequests` must be greater than zero.
- `WindowSize` must be greater than zero.
