package config

import "testing"

func TestNacosNamespaceIDForPublicNamespace(t *testing.T) {
	if got := nacosNamespaceID(); got != "" {
		t.Fatalf("nacosNamespaceID() = %q, want empty string for public namespace", got)
	}
}

func TestParseServicePort(t *testing.T) {
	got, err := parseServicePort("8088")
	if err != nil {
		t.Fatalf("parseServicePort returned error: %v", err)
	}
	if got != 8088 {
		t.Fatalf("parseServicePort() = %d, want 8088", got)
	}
}

func TestParseServicePortRejectsInvalidValue(t *testing.T) {
	if _, err := parseServicePort("bad-port"); err == nil {
		t.Fatal("parseServicePort() error = nil, want non-nil")
	}
}
