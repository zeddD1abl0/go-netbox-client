#!/bin/bash

# Colors for better output
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo -e "${GREEN}Netbox Integration Test Runner${NC}\n"

# Function to validate URL format
validate_url() {
    if [[ ! $1 =~ ^https?:// ]]; then
        echo -e "${YELLOW}Warning: URL should start with http:// or https://${NC}"
        read -p "Continue anyway? (y/n) " -n 1 -r
        echo
        if [[ ! $REPLY =~ ^[Yy]$ ]]; then
            return 1
        fi
    fi
    return 0
}

# Function to mask token input
read_token() {
    if [ -t 0 ]; then  # Check if running interactively
        stty -echo
        read -p "Enter Netbox API Token: " TOKEN
        stty echo
        echo
    else
        read TOKEN
    fi
}

# Get Netbox URL
read -p "Enter Netbox URL (e.g., https://netbox.example.com/api): " URL

# Validate URL
if ! validate_url "$URL"; then
    echo "Exiting..."
    exit 1
fi

# Get Netbox token
echo -e "\nEnter your Netbox API token"
echo "This token should have read/write permissions for integration testing"
read_token

# Remove trailing slash from URL if present
URL=${URL%/}

# Export environment variables
export NETBOX_URL="$URL"
export NETBOX_TOKEN="$TOKEN"

echo -e "\n${GREEN}Configuration:${NC}"
echo "Netbox URL: $URL"
echo "Token: ${TOKEN:0:4}...${TOKEN: -4}"

# Verify configuration
echo -e "\nVerifying configuration..."
if [ -z "$URL" ] || [ -z "$TOKEN" ]; then
    echo "Error: URL and token are required"
    exit 1
fi

# Run tests
echo -e "\n${GREEN}Running integration tests...${NC}\n"

# Create a temporary file for test output
TEMP_OUTPUT=$(mktemp)

# Run tests and capture output
if go test ./integration_tests/... -v 2>&1 | tee "$TEMP_OUTPUT"; then
    echo -e "\n${GREEN}All tests passed successfully!${NC}"
else
    echo -e "\n${YELLOW}Some tests failed. See output above for details.${NC}"
fi

# Clean up
rm "$TEMP_OUTPUT"
