# Compression Configuration

The compression configuration allows you to enable and configure compression for gRPC health probe requests.

## Overview

Compression can reduce bandwidth usage when probing gRPC services, especially useful in bandwidth-constrained environments or when probing services over the internet.

## Configuration

### CompressionConfig

```go
type CompressionConfig struct {
    Enabled bool             // Whether compression is enabled
    Type    CompressionType  // Compression algorithm to use
}
```

### Compression Types

- `none`: No compression (default)
- `gzip`: GZIP compression

## Usage Examples

### Disable Compression (Default)

```go
config := probe.DefaultCompressionConfig()
// Enabled: false, Type: "none"
```

### Enable GZIP Compression

```go
config := &probe.CompressionConfig{
    Enabled: true,
    Type:    probe.CompressionGzip,
}
```

## CLI Usage

```bash
# Use gzip compression
grpc-health-probe --addr=localhost:9090 --compression=gzip

# Disable compression (default)
grpc-health-probe --addr=localhost:9090 --compression=none
```

## Notes

- Compression is disabled by default to minimize overhead for local health checks
- GZIP is the only supported compression algorithm currently
- Both client and server must support the same compression algorithm
- Compression adds CPU overhead but reduces network bandwidth
