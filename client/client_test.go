package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	tests := []struct {
		name    string
		baseURL string
		token   string
		wantErr bool
		errMsg  string
	}{
		{
			name:    "valid client creation",
			baseURL: "http://127.0.0.1:8000/api/",
			token:   "42f1da103a3052ae9d2dfb76c93bfaa9e950adc5",
			wantErr: false,
		},
		{
			name:    "empty baseURL",
			baseURL: "",
			token:   "42f1da103a3052ae9d2dfb76c93bfaa9e950adc5",
			wantErr: true,
			errMsg:  "baseURL cannot be empty",
		},
		{
			name:    "empty token",
			baseURL: "http://127.0.0.1:8000/api/",
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
