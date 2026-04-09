# Verbosity Flags

The `grpc-health-probe` CLI supports configurable verbosity levels to control the amount of output produced during a health check.

## Flags

| Flag          | Type   | Default | Description                                      |
|---------------|--------|---------|--------------------------------------------------|
| `--verbosity` | string | `info`  | Log verbosity level (`debug`, `info`, `warn`, `error`) |
| `--quiet`     | bool   | `false` | Suppress all output except the exit code         |

## Verbosity Levels

- **debug** — Prints detailed connection and request information.
- **info** — Default level. Prints health status and key events.
- **warn** — Prints only warnings and errors.
- **error** — Prints only errors.

## Examples

### Default (info)
```bash
grpc-health-probe --addr=localhost:50051
```

### Debug mode
```bash
grpc-health-probe --addr=localhost:50051 --verbosity=debug
```

### Quiet mode (exit code only)
```bash
grpc-health-probe --addr=localhost:50051 --quiet
echo $?
```

### Warn level
```bash
grpc-health-probe --addr=localhost:50051 --verbosity=warn
```

## Notes

- `--quiet` takes precedence over `--verbosity`.
- When `--quiet` is set, no output is written regardless of the verbosity level.
- Invalid verbosity levels will cause the command to exit with an error.
