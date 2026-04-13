package probe

import "fmt"

// ApplyNamespace returns the service name with the namespace prefix applied,
// if the NamespaceConfig is enabled. Otherwise it returns the original name.
func (c *NamespaceConfig) ApplyNamespace(service string) string {
	if c == nil || !c.Enabled {
		return service
	}
	if service == "" {
		return c.Namespace
	}
	return fmt.Sprintf("%s/%s", c.Namespace, service)
}

// String returns a human-readable representation of the NamespaceConfig.
func (c *NamespaceConfig) String() string {
	if c == nil {
		return "namespace(nil)"
	}
	if !c.Enabled {
		return "namespace(disabled)"
	}
	return fmt.Sprintf("namespace(%s)", c.Namespace)
}
