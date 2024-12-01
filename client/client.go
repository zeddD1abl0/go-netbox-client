package client

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

const (
	defaultAPIVersion = "4.0"
)

// Client represents a Netbox API client
type Client struct {
	httpClient *resty.Client
	baseURL    string
	token      string
}

// NewClient creates a new Netbox client
func NewClient(baseURL, token string) (*Client, error) {
	if baseURL == "" {
		return nil, fmt.Errorf("baseURL cannot be empty")
	}
	if token == "" {
		return nil, fmt.Errorf("token cannot be empty")
	}

	client := &Client{
		httpClient: resty.New(),
		baseURL:    baseURL,
		token:      token,
	}

	client.httpClient.SetHeader("Authorization", fmt.Sprintf("Token %s", token))
	client.httpClient.SetHeader("Accept", "application/json")
	client.httpClient.SetHeader("Content-Type", "application/json")

	return client, nil
}
