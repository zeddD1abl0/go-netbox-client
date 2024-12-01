package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// testServer represents a test HTTP server
type testServer struct {
	*httptest.Server
	// ResponseCode is the HTTP response code to return
	ResponseCode int
	// ResponseBody is the response body to return
	ResponseBody interface{}
}

// newTestServer creates a new test server with the given response
func newTestServer(t *testing.T, code int, body interface{}) *testServer {
	ts := &testServer{
		ResponseCode: code,
		ResponseBody: body,
	}

	ts.Server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(ts.ResponseCode)

		if body != nil {
			err := json.NewEncoder(w).Encode(ts.ResponseBody)
			if err != nil {
				t.Fatalf("failed to encode response body: %v", err)
			}
		}
	}))

	return ts
}

// newTestClient creates a new client that uses the test server
func newTestClient(t *testing.T, ts *testServer) *Client {
	client, err := NewClient(ts.URL, "test-token")
	if err != nil {
		t.Fatalf("failed to create test client: %v", err)
	}
	return client
}

// mockPaginatedResponse creates a mock paginated response
func mockPaginatedResponse(items interface{}) map[string]interface{} {
	return map[string]interface{}{
		"count":    len(fmt.Sprint(items)), // This is just for testing
		"next":     nil,
		"previous": nil,
		"results":  items,
	}
}
