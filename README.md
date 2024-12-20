# Go NetBox Client

A Go client library for interacting with NetBox API v4.0, designed for use with Terraform providers and general NetBox automation.

[![Go Report Card](https://goreportcard.com/badge/github.com/zeddD1abl0/go-netbox-client)](https://goreportcard.com/report/github.com/zeddD1abl0/go-netbox-client)
[![GoDoc](https://godoc.org/github.com/zeddD1abl0/go-netbox-client?status.svg)](https://godoc.org/github.com/zeddD1abl0/go-netbox-client)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## Features

- Full support for NetBox API v4.0
- Comprehensive test coverage
- Clean and idiomatic Go code
- Easy integration with Terraform providers
- Type-safe API interactions
- Detailed error handling

## Installation

To use this package in your project:

```bash
go get github.com/zeddD1abl0/go-netbox-client
```

## Quick Start

Here's a simple example to get you started:

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/zeddD1abl0/go-netbox-client/client"
)

func main() {
    // Create a new client
    netboxClient, err := client.NewClient(
        "https://netbox.example.com/api",
        "your-api-token",
    )
    if err != nil {
        log.Fatal(err)
    }

    // Example: List all sites
    sites, err := netboxClient.DCIM.ListSites()
    if err != nil {
        log.Fatal(err)
    }

    for _, site := range sites {
        fmt.Printf("Site: %s\n", site.Name)
    }
}
```

## Documentation

For detailed documentation and examples, please refer to the [GoDoc](https://godoc.org/github.com/zeddD1abl0/go-netbox-client).

### Available Modules

- DCIM (Data Center Infrastructure Management)
  - Sites
  - Racks
  - Devices
  - Interfaces
  - Locations
  - Regions
  - Site Groups

More modules will be added as development continues.

## Development

### Requirements

- Go 1.21 or higher
- Access to a NetBox v4.0 instance for integration testing

### Setting Up Development Environment

1. Clone the repository:
   ```bash
   git clone https://github.com/zeddD1abl0/go-netbox-client.git
   cd go-netbox-client
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

### Running Tests

Run all tests:
```bash
go test ./... -v
```

Run tests with race condition detection:
```bash
go test -race ./...
```

### Integration Tests

This project includes a comprehensive suite of integration tests that verify the client's functionality against a real Netbox server. These tests cover various scenarios including:

- Site Groups management
- Region hierarchies
- Location hierarchies with complex relationships
- Cross-resource relationships and filtering

### Running Integration Tests

There are two ways to run the integration tests:

1. Using the setup script (recommended):
   ```bash
   ./run_integration_tests.sh
   ```
   This script will:
   - Prompt for your Netbox URL and API token
   - Validate the configuration
   - Run all integration tests
   - Display results with proper formatting

2. Manually setting environment variables:
   ```bash
   export NETBOX_URL="https://your-netbox-server/api"
   export NETBOX_TOKEN="your-api-token"
   go test ./integration_tests/... -v
   ```

### Test Requirements

- A running Netbox instance
- API token with read/write permissions
- Network access to the Netbox server

### Test Categories

1. Site Groups (`integration_tests/site_group_test.go`)
   - CRUD operations
   - Filtering and search
   - Parent-child relationships

2. Regions (`integration_tests/region_test.go`)
   - Hierarchical relationships
   - Tag-based filtering
   - Complex search scenarios

3. Locations (`integration_tests/location_test.go`)
   - Complex hierarchies (campus/building/floor)
   - Cross-resource relationships
   - Multiple filtering scenarios

For more details about the integration tests, see the [integration_tests/README.md](integration_tests/README.md) file.

### Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

Please ensure your PR includes:
- A clear description of the changes
- Updated tests if needed
- Updated documentation if needed

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Thanks to the NetBox team for providing a great API
- Inspired by various Go API clients and best practices from the community
