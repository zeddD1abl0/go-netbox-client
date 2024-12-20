package integration_tests

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zeddD1abl0/go-netbox-client/client"
	"github.com/zeddD1abl0/go-netbox-client/client/dcim"
)

// setupTestClient creates a new client for testing
func setupTestClient(t *testing.T) (*client.Client, *dcim.Service) {
	cfg := LoadTestConfig(t)

	c, err := client.NewClient(
		cfg.NetboxURL,
		cfg.NetboxToken,
		client.WithTimeout(60),
		client.WithRetry(3, 5),
	)
	require.NoError(t, err)
	require.NotNil(t, c)

	return c, dcim.NewService(c)
}

// cleanupResource is a generic cleanup function that takes a delete function
type cleanupFunc func() error

// createCleanupList creates a list of cleanup functions to be run in reverse order
type cleanupList struct {
	t       *testing.T
	cleanup []cleanupFunc
}

func newCleanupList(t *testing.T) *cleanupList {
	return &cleanupList{t: t}
}

func (c *cleanupList) add(fn cleanupFunc) {
	c.cleanup = append(c.cleanup, fn)
}

func (c *cleanupList) runAll() {
	// Run cleanup in reverse order
	for i := len(c.cleanup) - 1; i >= 0; i-- {
		if err := c.cleanup[i](); err != nil {
			c.t.Logf("Cleanup error: %v", err)
		}
	}
}
