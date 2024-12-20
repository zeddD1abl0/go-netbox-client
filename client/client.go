package client

import (
	"fmt"
	"strings"
	"time"
	"github.com/go-resty/resty/v2"
)

const (
	defaultAPIVersion = "4.0"
	defaultTimeout    = 30 // seconds
	defaultRetryCount = 3
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

	// Ensure baseURL has a trailing slash
	if !strings.HasSuffix(baseURL, "/") {
		baseURL = baseURL + "/"
	}

	restyClient := resty.New()
	client := &Client{
		httpClient: restyClient,
		baseURL:    baseURL,
		token:      token,
	}

	// Set default headers
	restyClient.SetHeader("Authorization", fmt.Sprintf("Token %s", token))
	restyClient.SetHeader("Accept", "application/json")
	restyClient.SetHeader("Content-Type", "application/json")

	// Set default timeout
	restyClient.SetTimeout(defaultTimeout * time.Second)

	// Set default retry configuration
	restyClient.
		SetRetryCount(defaultRetryCount).
		SetRetryWaitTime(defaultRetryWaitTime * time.Second).
		AddRetryCondition(func(r *resty.Response, err error) bool {
			return err != nil || r.StatusCode() >= 500
		})

	// Apply any custom options
	for _, opt := range opts {
		opt(client)
	}

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
