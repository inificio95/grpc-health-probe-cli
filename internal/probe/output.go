package probe

import (
	"fmt"
	"io"
	"os"
)

// OutputConfig controls where and how probe results are written.
type OutputConfig struct {
	// Writer is the destination for output. Defaults to os.Stdout.
	Writer io.Writer
	// ErrWriter is the destination for error output. Defaults to os.Stderr.
	ErrWriter io.Writer
	// Format is the output format: "text" or "json".
	Format string
	// Verbose enables additional diagnostic output.
	Verbose bool
}

// DefaultOutputConfig returns an OutputConfig with sensible defaults.
func DefaultOutputConfig() OutputConfig {
	return OutputConfig{
		Writer:    os.Stdout,
		ErrWriter: os.Stderr,
		Format:    "text",
		Verbose:   false,
	}
}

// Validate checks that the OutputConfig is valid.
func (c OutputConfig) Validate() error {
	if c.Writer == nil {
		return fmt.Errorf("output: Writer must not be nil")
	}
	if c.ErrWriter == nil {
		return fmt.Errorf("output: ErrWriter must not be nil")
	}
	if c.Format != "text" && c.Format != "json" {
		return fmt.Errorf("output: unsupported format %q, must be \"text\" or \"json\"", c.Format)
	}
	return nil
}

// Write formats the given Result and writes it to the configured Writer.
// Any formatting or write error is written to ErrWriter.
func (c OutputConfig) Write(r Result) error {
	formatted, err := FormatResult(r, c.Format)
	if err != nil {
		fmt.Fprintf(c.ErrWriter, "output: format error: %v\n", err)
		return err
	}
	_, err = fmt.Fprintln(c.Writer, formatted)
	if err != nil {
		fmt.Fprintf(c.ErrWriter, "output: write error: %v\n", err)
		return err
	}
	if c.Verbose && r.Error != nil {
		fmt.Fprintf(c.ErrWriter, "[verbose] error detail: %v\n", r.Error)
	}
	return nil
}
