package probe

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var namespacePattern = regexp.MustCompile(`^[a-zA-Z0-9_-]+$`)

// NamespaceConfig holds configuration for grouping probes under a logical namespace.
type NamespaceConfig struct {
	Enabled   bool
	Namespace string
	Separator string
}

// DefaultNamespaceConfig returns a NamespaceConfig with sensible defaults.
func DefaultNamespaceConfig() *NamespaceConfig {
	return &NamespaceConfig{
		Enabled:   false,
		Namespace: "",
		Separator: ".",
	}
}

// Validate checks that the NamespaceConfig is valid.
func (c *NamespaceConfig) Validate() error {
	if c == nil {
		return errors.New("namespace config must not be nil")
	}
	if !c.Enabled {
		return nil
	}
	if strings.TrimSpace(c.Namespace) == "" {
		return errors.New("namespace must not be empty when enabled")
	}
	if !namespacePattern.MatchString(c.Namespace) {
		return fmt.Errorf("namespace %q contains invalid characters; only alphanumeric, dash, and underscore are allowed", c.Namespace)
	}
	if c.Separator == "" {
		return errors.New("namespace separator must not be empty")
	}
	return nil
}

// Qualify prefixes the given name with the namespace when enabled.
func (c *NamespaceConfig) Qualify(name string) string {
	if c == nil || !c.Enabled || c.Namespace == "" {
		return name
	}
	return c.Namespace + c.Separator + name
}
