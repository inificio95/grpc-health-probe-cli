# Metadata Flags

The `grpc-health-probe` CLI supports attaching arbitrary gRPC metadata (headers)
to every health-check request via the `--header` / `-H` flag.

## Usage

```
--header, -H  string  Metadata header in key=value format (repeatable)
```

## Examples

### Single header

```bash
grpc-health-probe --addr=localhost:50051 \
  -H "x-request-id=abc123"
```

### Multiple headers

```bash
grpc-health-probe --addr=localhost:50051 \
  -H "authorization=Bearer mytoken" \
  -H "x-trace-id=xyz789"
```

## Notes

- The flag may be specified multiple times; each occurrence adds one entry.
- Entries that do not contain an `=` separator are silently ignored.
- Leading and trailing whitespace is trimmed from both the key and the value.
- Metadata is forwarded on every retry attempt.
- This flag is independent of the `--auth-*` flags; both can be used together.
