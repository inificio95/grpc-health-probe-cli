package probe

import "testing"

func TestNamespaceConfig_ApplyNamespace_Disabled(t *testing.T) {
	c := &NamespaceConfig{Enabled: false, Namespace: "prod"}
	got := c.ApplyNamespace("my-service")
	if got != "my-service" {
		t.Errorf("expected 'my-service', got %q", got)
	}
}

func TestNamespaceConfig_ApplyNamespace_Nil(t *testing.T) {
	var c *NamespaceConfig
	got := c.ApplyNamespace("my-service")
	if got != "my-service" {
		t.Errorf("expected 'my-service', got %q", got)
	}
}

func TestNamespaceConfig_ApplyNamespace_Enabled(t *testing.T) {
	c := &NamespaceConfig{Enabled: true, Namespace: "prod"}
	got := c.ApplyNamespace("my-service")
	if got != "prod/my-service" {
		t.Errorf("expected 'prod/my-service', got %q", got)
	}
}

func TestNamespaceConfig_ApplyNamespace_EmptyService(t *testing.T) {
	c := &NamespaceConfig{Enabled: true, Namespace: "staging"}
	got := c.ApplyNamespace("")
	if got != "staging" {
		t.Errorf("expected 'staging', got %q", got)
	}
}

func TestNamespaceConfig_String_Nil(t *testing.T) {
	var c *NamespaceConfig
	if got := c.String(); got != "namespace(nil)" {
		t.Errorf("expected 'namespace(nil)', got %q", got)
	}
}

func TestNamespaceConfig_String_Disabled(t *testing.T) {
	c := &NamespaceConfig{Enabled: false}
	if got := c.String(); got != "namespace(disabled)" {
		t.Errorf("expected 'namespace(disabled)', got %q", got)
	}
}

func TestNamespaceConfig_String_Enabled(t *testing.T) {
	c := &NamespaceConfig{Enabled: true, Namespace: "prod"}
	if got := c.String(); got != "namespace(prod)" {
		t.Errorf("expected 'namespace(prod)', got %q", got)
	}
}
