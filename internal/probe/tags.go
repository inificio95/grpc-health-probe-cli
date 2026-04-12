package probe

import (
	"fmt"
	"strings"
)

// TagsConfig holds key-value label pairs attached to probe results for
// identification and filtering purposes.
type TagsConfig struct {
	Enabled bool
	Tags    map[string]string
}

// DefaultTagsConfig returns a TagsConfig with tagging disabled.
func DefaultTagsConfig() *TagsConfig {
	return &TagsConfig{
		Enabled: false,
		Tags:    map[string]string{},
	}
}

// Validate checks the TagsConfig for correctness.
func (c *TagsConfig) Validate() error {
	if c == nil {
		return fmt.Errorf("tags config must not be nil")
	}
	if !c.Enabled {
		return nil
	}
	for k, v := range c.Tags {
		if strings.TrimSpace(k) == "" {
			return fmt.Errorf("tag key must not be empty or whitespace")
		}
		if strings.TrimSpace(v) == "" {
			return fmt.Errorf("tag value for key %q must not be empty or whitespace", k)
		}
	}
	return nil
}

// AsLabels returns the tags as a flat slice of "key=value" strings.
func (c *TagsConfig) AsLabels() []string {
	if c == nil || !c.Enabled {
		return nil
	}
	labels := make([]string, 0, len(c.Tags))
	for k, v := range c.Tags {
		labels = append(labels, fmt.Sprintf("%s=%s", k, v))
	}
	return labels
}
