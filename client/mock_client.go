package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/go-resty/resty/v2"
)

// MockClient is a mock implementation of the Netbox client for testing
type MockClient struct {
	t            *testing.T
	expectedPath string
	response     interface{}
	statusCode   int
}

// NewMockClient creates a new mock client for testing
func NewMockClient(t *testing.T, expectedPath string, response interface{}, statusCode int) *Client {
	return &Client{
		httpClient: &mockHTTPClient{
			t:            t,
			expectedPath: expectedPath,
			response:     response,
			statusCode:   statusCode,
		},
		baseURL: "http://mock.test",
		token:   "mock-token",
	}
}

// mockHTTPClient is a mock implementation of HTTPClient for testing
type mockHTTPClient struct {
	t            *testing.T
	expectedPath string
	response     interface{}
	statusCode   int
}

// R returns a new mock request object
func (m *mockHTTPClient) R() *resty.Request {
	client := resty.New()
	return client.R()
}

type mockRequest struct {
	t            *testing.T
	expectedPath string
	response     interface{}
	statusCode   int
	headers      map[string]string
}

func (r *mockRequest) SetResult(v interface{}) *mockRequest {
	return r
}

func (r *mockRequest) SetHeader(key, value string) *mockRequest {
	if r.headers == nil {
		r.headers = make(map[string]string)
	}
	r.headers[key] = value
	return r
}

func (r *mockRequest) SetHeaders(headers map[string]string) *mockRequest {
	r.headers = headers
	return r
}

func (r *mockRequest) SetBody(body interface{}) *mockRequest {
	return r
}

func (r *mockRequest) Get(path string) (*mockResponse, error) {
	if path != r.expectedPath {
		r.t.Errorf("Expected path %s, got %s", r.expectedPath, path)
	}
	return r.createResponse(), nil
}

func (r *mockRequest) Post(path string) (*mockResponse, error) {
	if path != r.expectedPath {
		r.t.Errorf("Expected path %s, got %s", r.expectedPath, path)
	}
	return r.createResponse(), nil
}

func (r *mockRequest) Put(path string) (*mockResponse, error) {
	if path != r.expectedPath {
		r.t.Errorf("Expected path %s, got %s", r.expectedPath, path)
	}
	return r.createResponse(), nil
}

func (r *mockRequest) Delete(path string) (*mockResponse, error) {
	if path != r.expectedPath {
		r.t.Errorf("Expected path %s, got %s", r.expectedPath, path)
	}
	return r.createResponse(), nil
}

func (r *mockRequest) createResponse() *mockResponse {
	var body []byte
	var err error
	if r.response != nil {
		body, err = json.Marshal(r.response)
		if err != nil {
			r.t.Fatalf("Failed to marshal mock response: %v", err)
		}
	}

	return &mockResponse{
		statusCode: r.statusCode,
		body:       body,
	}
}

type mockResponse struct {
	statusCode int
	body       []byte
}

func (r *mockResponse) StatusCode() int {
	return r.statusCode
}

func (r *mockResponse) Body() []byte {
	return r.body
}

func (r *mockResponse) Error() error {
	if r.statusCode >= http.StatusBadRequest {
		return fmt.Errorf("HTTP error: %d", r.statusCode)
	}
	return nil
}
