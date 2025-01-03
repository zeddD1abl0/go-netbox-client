package client

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
)

const (
	defaultAPIVersion    = "4.0"
	defaultTimeout       = 30 // seconds
	defaultRetryCount    = 3
	defaultRetryWaitTime = 5 // seconds
)

// ClientOption represents a function that can configure a Client
type ClientOption func(*Client)

// WithTimeout sets the client timeout
func WithTimeout(timeout int) ClientOption {
	return func(c *Client) {
		c.httpClient.(*resty.Client).SetTimeout(time.Duration(timeout) * time.Second)
	}
}

// WithRetry sets the retry configuration
func WithRetry(count int, waitTime int) ClientOption {
	return func(c *Client) {
		c.httpClient.(*resty.Client).
			SetRetryCount(count).
			SetRetryWaitTime(time.Duration(waitTime) * time.Second).
			AddRetryCondition(func(r *resty.Response, err error) bool {
				return err != nil || r.StatusCode() >= 500
			})
	}
}

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
func NewClient(baseURL, token string, opts ...ClientOption) (*Client, error) {
	if baseURL == "" {
		return nil, fmt.Errorf("baseURL cannot be empty")
	}
	if token == "" {
		return nil, fmt.Errorf("token cannot be empty")
	}

	// Ensure baseURL ends with /api
	if !strings.HasSuffix(baseURL, "/api") {
		baseURL = strings.TrimSuffix(baseURL, "/") + "/api"
	}

	// Create HTTP client
	httpClient := resty.New().
		SetTimeout(time.Duration(defaultTimeout)*time.Second).
		SetRetryCount(defaultRetryCount).
		SetRetryWaitTime(time.Duration(defaultRetryWaitTime)*time.Second).
		SetAuthToken(token).
		SetHeader("Accept", "application/json")

	// Create client
	c := &Client{
		httpClient: httpClient,
		baseURL:    baseURL,
		token:      token,
	}

	// Apply options
	for _, opt := range opts {
		opt(c)
	}

	return c, nil
}

// R returns a new request object
func (c *Client) R() *resty.Request {
	return c.httpClient.R()
}

// BuildPath builds a full API path from the given parts
func (c *Client) BuildPath(parts ...string) string {
	path := c.baseURL
	for _, part := range parts {
		path += part + "/"
	}
	return path
}

// Response represents a paginated response from the Netbox API
type Response struct {
	Count    int   `json:"count"`
	Next     *int  `json:"next"`
	Previous *int  `json:"previous"`
	Results  []any `json:"results"`
}
