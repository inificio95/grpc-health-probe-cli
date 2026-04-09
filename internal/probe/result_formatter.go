package probe

import (
	"encoding/json"
	"fmt"
	"io"
	"text/tabwriter"
)

// OutputFormat controls how results are rendered.
type OutputFormat string

const (
	OutputFormatText OutputFormat = "text"
	OutputFormatJSON OutputFormat = "json"
)

// jsonResult is the serialisable form of Result.
type jsonResult struct {
	Status   string  `json:"status"`
	Attempt  int     `json:"attempt"`
	DurationMs float64 `json:"duration_ms"`
	Error    string  `json:"error,omitempty"`
}

// FormatResult writes a Result to w in the requested format.
func FormatResult(w io.Writer, r Result, format OutputFormat) error {
	switch format {
	case OutputFormatJSON:
		return formatJSON(w, r)
	default:
		return formatText(w, r)
	}
}

func formatText(w io.Writer, r Result) error {
	tw := tabwriter.NewWriter(w, 0, 0, 2, ' ', 0)
	fmt.Fprintf(tw, "Status:\t%s\n", r.Status)
	fmt.Fprintf(tw, "Attempt:\t%d\n", r.Attempt)
	fmt.Fprintf(tw, "Duration:\t%.2fms\n", float64(r.Duration.Microseconds())/1000)
	if r.Err != nil {
		fmt.Fprintf(tw, "Error:\t%s\n", r.Err)
	}
	return tw.Flush()
}

func formatJSON(w io.Writer, r Result) error {
	jr := jsonResult{
		Status:     r.Status.String(),
		Attempt:    r.Attempt,
		DurationMs: float64(r.Duration.Microseconds()) / 1000,
	}
	if r.Err != nil {
		jr.Error = r.Err.Error()
	}
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	return enc.Encode(jr)
}
