package client

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

const (
	defaultAPIVersion = "4.0"
)

// HTTPClient represents an HTTP client interface
type HTTPClient interface {
	R() *resty.Request
}

// Client represents a Netbox API client
type Client struct {
	httpClient HTTPClient
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

	client.httpClient.(*resty.Client).SetHeader("Authorization", fmt.Sprintf("Token %s", token))
	client.httpClient.(*resty.Client).SetHeader("Accept", "application/json")
	client.httpClient.(*resty.Client).SetHeader("Content-Type", "application/json")

	return client, nil
}

// R returns a new request object
func (c *Client) R() *resty.Request {
	return c.httpClient.R()
}

// Response represents a paginated response from the Netbox API
type Response struct {
	Count    int   `json:"count"`
	Next     *int  `json:"next"`
	Previous *int  `json:"previous"`
	Results  []any `json:"results"`
}
