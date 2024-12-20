# Integration Tests

This package contains integration tests that run against a real Netbox server. These tests verify that our client works correctly with an actual Netbox instance.

## Setup

1. Set up environment variables:
   ```bash
   export NETBOX_URL="https://your-netbox-server/api"
   export NETBOX_TOKEN="your-api-token"
   ```

2. Run the tests:
   ```bash
   # Run all integration tests
   go test ./integration_tests/... -v

   # Run specific test
   go test ./integration_tests/... -v -run TestSiteGroupIntegration
   ```

## Test Structure

Each test file focuses on a specific resource type (e.g., site groups, regions, locations) and performs CRUD operations to verify the client works correctly with the real Netbox API.

The tests will be skipped if the required environment variables are not set. This allows the integration tests to be run as part of the normal test suite without failing when Netbox credentials are not available.
