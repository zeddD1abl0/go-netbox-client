package dcim

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zeddD1abl0/go-netbox-client/client"
)

func TestListRegions(test *testing.T) {
	tests := []struct {
		name           string
		input          *ListRegionsInput
		expectedPath   string
		expectedParams map[string]string
		mockResponse   string
		mockStatus     int
		expectError    bool
	}{
		{
			name:         "successful list with no filters",
			input:        &ListRegionsInput{},
			expectedPath: "/api/dcim/regions",
			mockResponse: `{"count": 2, "results": [{"id": 1, "name": "Region 1"}, {"id": 2, "name": "Region 2"}]}`,
			mockStatus:   http.StatusOK,
		},
		{
			name: "successful list with filters",
			input: &ListRegionsInput{
				Name:   "Test",
				Parent: "parent-region",
				Tag:    "prod",
				Limit:  10,
				Offset: 0,
			},
			expectedPath: "/api/dcim/regions",
			expectedParams: map[string]string{
				"name":   "Test",
				"parent": "parent-region",
				"tag":    "prod",
				"limit":  "10",
			},
			mockResponse: `{"count": 1, "results": [{"id": 1, "name": "Test Region"}]}`,
			mockStatus:   http.StatusOK,
		},
		{
			name:         "server error",
			input:        &ListRegionsInput{},
			expectedPath: "/api/dcim/regions",
			mockStatus:   http.StatusInternalServerError,
			expectError:  true,
		},
	}

	for _, spec_test := range tests {
		test.Run(spec_test.name, func(test *testing.T) {
			client := client.NewMockClient(test, spec_test.expectedPath, spec_test.mockResponse, spec_test.mockStatus)
			service := NewService(client)

			regions, err := service.ListRegions(spec_test.input)
			if spec_test.expectError {
				assert.Error(test, err)
				return
			}

			assert.NoError(test, err)
			assert.NotNil(test, regions)
		})
	}
}

func TestGetRegion(test *testing.T) {
	tests := []struct {
		name         string
		id           int
		expectedPath string
		mockResponse string
		mockStatus   int
		expectError  bool
	}{
		{
			name:         "successful get",
			id:           1,
			expectedPath: "/api/dcim/regions/1",
			mockResponse: `{"id": 1, "name": "Test Region"}`,
			mockStatus:   http.StatusOK,
		},
		{
			name:         "not found",
			id:           999,
			expectedPath: "/api/dcim/regions/999",
			mockStatus:   http.StatusNotFound,
			expectError:  true,
		},
		{
			name:         "server error",
			id:           1,
			expectedPath: "/api/dcim/regions/1",
			mockStatus:   http.StatusInternalServerError,
			expectError:  true,
		},
	}

	for _, spec_test := range tests {
		test.Run(spec_test.name, func(test *testing.T) {
			client := client.NewMockClient(test, spec_test.expectedPath, spec_test.mockResponse, spec_test.mockStatus)
			service := NewService(client)

			region, err := service.GetRegion(spec_test.id)
			if spec_test.expectError {
				assert.Error(test, err)
				return
			}

			assert.NoError(test, err)
			assert.NotNil(test, region)
		})
	}
}

func TestCreateRegion(test *testing.T) {
	tests := []struct {
		name         string
		input        *CreateRegionInput
		expectedPath string
		mockResponse string
		mockStatus   int
		expectError  bool
	}{
		{
			name: "successful create",
			input: &CreateRegionInput{
				Name: "Test Region",
				Slug: "test-region",
			},
			expectedPath: "/api/dcim/regions",
			mockResponse: `{"id": 1, "name": "Test Region", "slug": "test-region"}`,
			mockStatus:   http.StatusCreated,
		},
		{
			name: "validation error",
			input: &CreateRegionInput{
				Name: "", // Required field
				Slug: "test-region",
			},
			expectError: true,
		},
		{
			name: "server error",
			input: &CreateRegionInput{
				Name: "Test Region",
				Slug: "test-region",
			},
			expectedPath: "/api/dcim/regions",
			mockStatus:   http.StatusInternalServerError,
			expectError:  true,
		},
	}

	for _, spec_test := range tests {
		test.Run(spec_test.name, func(test *testing.T) {
			client := client.NewMockClient(test, spec_test.expectedPath, spec_test.mockResponse, spec_test.mockStatus)
			service := NewService(client)

			region, err := service.CreateRegion(spec_test.input)
			if spec_test.expectError {
				assert.Error(test, err)
				return
			}

			assert.NoError(test, err)
			assert.NotNil(test, region)
		})
	}
}

func TestUpdateRegion(test *testing.T) {
	tests := []struct {
		name         string
		input        *UpdateRegionInput
		expectedPath string
		mockResponse string
		mockStatus   int
		expectError  bool
	}{
		{
			name: "successful update",
			input: &UpdateRegionInput{
				ID:   1,
				Name: "Updated Region",
				Slug: "updated-region",
			},
			expectedPath: "/api/dcim/regions/1",
			mockResponse: `{"id": 1, "name": "Updated Region", "slug": "updated-region"}`,
			mockStatus:   http.StatusOK,
		},
		{
			name: "not found",
			input: &UpdateRegionInput{
				ID:   999,
				Name: "Updated Region",
				Slug: "updated-region",
			},
			expectedPath: "/api/dcim/regions/999",
			mockStatus:   http.StatusNotFound,
			expectError:  true,
		},
		{
			name: "validation error",
			input: &UpdateRegionInput{
				ID:   1,
				Name: "", // Required field
				Slug: "updated-region",
			},
			expectError: true,
		},
	}

	for _, spec_test := range tests {
		test.Run(spec_test.name, func(test *testing.T) {
			client := client.NewMockClient(test, spec_test.expectedPath, spec_test.mockResponse, spec_test.mockStatus)
			service := NewService(client)

			region, err := service.UpdateRegion(spec_test.input)
			if spec_test.expectError {
				assert.Error(test, err)
				return
			}

			assert.NoError(test, err)
			assert.NotNil(test, region)
		})
	}
}

func TestPatchRegion(test *testing.T) {
	tests := []struct {
		name         string
		input        *PatchRegionInput
		expectedPath string
		mockResponse string
		mockStatus   int
		expectError  bool
	}{
		{
			name: "successful patch",
			input: &PatchRegionInput{
				ID:   1,
				Name: stringPtr("Patched Region"),
				Slug: stringPtr("patched-region"),
			},
			expectedPath: "/api/dcim/regions/1",
			mockResponse: `{"id": 1, "name": "Patched Region", "slug": "patched-region"}`,
			mockStatus:   http.StatusOK,
		},
		{
			name: "not found",
			input: &PatchRegionInput{
				ID:   999,
				Name: stringPtr("Patched Region"),
			},
			expectedPath: "/api/dcim/regions/999",
			mockStatus:   http.StatusNotFound,
			expectError:  true,
		},
		{
			name:        "validation error",
			input:       &PatchRegionInput{}, // Missing ID
			expectError: true,
		},
	}

	for _, spec_test := range tests {
		test.Run(spec_test.name, func(test *testing.T) {
			client := client.NewMockClient(test, spec_test.expectedPath, spec_test.mockResponse, spec_test.mockStatus)
			service := NewService(client)

			region, err := service.PatchRegion(spec_test.input)
			if spec_test.expectError {
				assert.Error(test, err)
				return
			}

			assert.NoError(test, err)
			assert.NotNil(test, region)
		})
	}
}

func TestDeleteRegion(test *testing.T) {
	tests := []struct {
		name         string
		id           int
		expectedPath string
		mockStatus   int
		expectError  bool
	}{
		{
			name:         "successful delete",
			id:           1,
			expectedPath: "/api/dcim/regions/1",
			mockStatus:   http.StatusNoContent,
		},
		{
			name:         "not found",
			id:           999,
			expectedPath: "/api/dcim/regions/999",
			mockStatus:   http.StatusNotFound,
			expectError:  true,
		},
		{
			name:         "server error",
			id:           1,
			expectedPath: "/api/dcim/regions/1",
			mockStatus:   http.StatusInternalServerError,
			expectError:  true,
		},
	}

	for _, spec_test := range tests {
		test.Run(spec_test.name, func(test *testing.T) {
			client := client.NewMockClient(test, spec_test.expectedPath, "", spec_test.mockStatus)
			service := NewService(client)

			err := service.DeleteRegion(spec_test.id)
			if spec_test.expectError {
				assert.Error(test, err)
				return
			}

			assert.NoError(test, err)
		})
	}
}
