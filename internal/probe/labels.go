package probe

import (
	"errors"
	"fmt"
	"strings"
)

// LabelsConfig holds key-value label pairs to attach to probe results.
type LabelsConfig struct {
	Enabled bool
	Labels  map[string]string
}

// DefaultLabelsConfig returns a LabelsConfig with labels disabled.
func DefaultLabelsConfig() *LabelsConfig {
	return &LabelsConfig{
		Enabled: false,
		Labels:  map[string]string{},
	}
}

// Validate checks that the LabelsConfig is valid.
func (c *LabelsConfig) Validate() error {
	if c == nil {
		return errors.New("labels config must not be nil")
	}
	if !c.Enabled {
		return nil
	}
	for k, v := range c.Labels {
		if strings.TrimSpace(k) == "" {
			return errors.New("label key must not be empty or whitespace")
		}
		if strings.TrimSpace(v) == "" {
			return fmt.Errorf("label value for key %q must not be empty or whitespace", k)
		}
	}
	return nil
}

// AsMap returns a copy of the labels map, or an empty map if disabled.
func (c *LabelsConfig) AsMap() map[string]string {
	out := make(map[string]string)
	if c == nil || !c.Enabled {
		return out
	}
	for k, v := range c.Labels {
		out[k] = v
	}
	return out
}
