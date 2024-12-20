package integration_tests

import (
	"os"
	"testing"
)

// TestConfig holds configuration for integration tests
type TestConfig struct {
	NetboxURL   string
	NetboxToken string
}

// LoadTestConfig loads the test configuration from environment variables
func LoadTestConfig(t *testing.T) *TestConfig {
	url := os.Getenv("NETBOX_URL")
	if url == "" {
		t.Skip("NETBOX_URL environment variable not set")
	}

	token := os.Getenv("NETBOX_TOKEN")
	if token == "" {
		t.Skip("NETBOX_TOKEN environment variable not set")
	}

	return &TestConfig{
		NetboxURL:   url,
		NetboxToken: token,
	}
}
