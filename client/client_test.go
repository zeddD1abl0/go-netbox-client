package client

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	tests := []struct {
		name     string
		baseURL  string
		token    string
		wantErr  bool
		errMsg   string
	}{
		{
			name:    "valid client creation",
			baseURL: "https://netbox.example.com/api",
			token:   "abc123",
			wantErr: false,
		},
		{
			name:    "empty baseURL",
			baseURL: "",
			token:   "abc123",
			wantErr: true,
			errMsg:  "baseURL cannot be empty",
		},
		{
			name:    "empty token",
			baseURL: "https://netbox.example.com/api",
			token:   "",
			wantErr: true,
			errMsg:  "token cannot be empty",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, err := NewClient(tt.baseURL, tt.token)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tt.errMsg, err.Error())
				assert.Nil(t, client)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, client)
				assert.Equal(t, tt.baseURL, client.baseURL)
				assert.Equal(t, tt.token, client.token)
			}
		})
	}
}
