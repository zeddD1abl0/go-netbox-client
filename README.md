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
