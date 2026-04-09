# Output Configuration

The `OutputConfig` struct controls where and how probe results are written.

## Fields

| Field       | Type        | Default      | Description                                      |
|-------------|-------------|--------------|--------------------------------------------------|
| `Writer`    | `io.Writer` | `os.Stdout`  | Destination for formatted probe output           |
| `ErrWriter` | `io.Writer` | `os.Stderr`  | Destination for error and diagnostic messages    |
| `Format`    | `string`    | `"text"`     | Output format: `"text"` or `"json"`              |
| `Verbose`   | `bool`      | `false`      | Emit additional error detail to `ErrWriter`      |

## Usage

```go
cfg := probe.DefaultOutputConfig()
cfg.Format = "json"
cfg.Verbose = true

if err := cfg.Validate(); err != nil {
    log.Fatal(err)
}

result := probe.Result{ /* ... */ }
cfg.Write(result)
```

## Text Format Example

```
status=SERVING address=localhost:50051 duration=4ms
```

## JSON Format Example

```json
{"status":"SERVING","address":"localhost:50051","duration_ms":4}
```

## Validation Rules

- `Writer` and `ErrWriter` must not be `nil`.
- `Format` must be one of `"text"` or `"json"`.
