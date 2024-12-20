package dcim

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zeddD1abl0/go-netbox-client/client"
)

func intPtr(i int) *int {
	return &i
}

func TestListSites(test *testing.T) {
	tests := []struct {
		name           string
		input          *ListSitesInput
		expectedPath   string
		expectedParams map[string]string
		mockResponse   string
		mockStatus     int
		expectError    bool
	}{
		{
			name:         "successful list with no filters",
			input:        &ListSitesInput{},
			expectedPath: "/api/dcim/sites",
			mockResponse: `{"count": 2, "results": [{"id": 1, "name": "Site 1"}, {"id": 2, "name": "Site 2"}]}`,
			mockStatus:   http.StatusOK,
		},
		{
			name: "successful list with filters",
			input: &ListSitesInput{
				BaseListInput: BaseListInput{
					Name:   "Test",
					Tag:    "prod",
					Limit:  10,
					Offset: 0,
				},
				Status: "active",
				Region: "us-west",
			},
			expectedPath: "/api/dcim/sites",
			expectedParams: map[string]string{
				"name":   "Test",
				"status": "active",
				"region": "us-west",
				"tag":    "prod",
				"limit":  "10",
			},
			mockResponse: `{"count": 1, "results": [{"id": 1, "name": "Test Site"}]}`,
			mockStatus:   http.StatusOK,
		},
		{
			name:         "server error",
			input:        &ListSitesInput{},
			expectedPath: "/api/dcim/sites",
			mockStatus:   http.StatusInternalServerError,
			expectError:  true,
		},
	}

	for _, spec_test := range tests {
		test.Run(spec_test.name, func(test *testing.T) {
			client := client.NewMockClient(test, spec_test.expectedPath, spec_test.mockResponse, spec_test.mockStatus)
			service := NewService(client)

			sites, err := service.ListSites(spec_test.input)
			if spec_test.expectError {
				assert.Error(test, err)
				return
			}

			assert.NoError(test, err)
			assert.NotNil(test, sites)
		})
	}
}

func TestGetSite(test *testing.T) {
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
			expectedPath: "/api/dcim/sites/1",
			mockResponse: `{"id": 1, "name": "Test Site"}`,
			mockStatus:   http.StatusOK,
		},
		{
			name:         "not found",
			id:           999,
			expectedPath: "/api/dcim/sites/999",
			mockStatus:   http.StatusNotFound,
			expectError:  true,
		},
		{
			name:         "server error",
			id:           1,
			expectedPath: "/api/dcim/sites/1",
			mockStatus:   http.StatusInternalServerError,
			expectError:  true,
		},
	}

	for _, spec_test := range tests {
		test.Run(spec_test.name, func(test *testing.T) {
			client := client.NewMockClient(test, spec_test.expectedPath, spec_test.mockResponse, spec_test.mockStatus)
			service := NewService(client)

			site, err := service.GetSite(spec_test.id)
			if spec_test.expectError {
				assert.Error(test, err)
				return
			}

			assert.NoError(test, err)
			assert.NotNil(test, site)
			assert.Equal(test, spec_test.id, site.ID)
		})
	}
}

func TestCreateSite(test *testing.T) {
	tests := []struct {
		name         string
		input        *CreateSiteInput
		expectedPath string
		mockResponse string
		mockStatus   int
		expectError  bool
	}{
		{
			name: "successful create",
			input: &CreateSiteInput{
				Name:   "New Site",
				Slug:   "new-site",
				Status: "active",
			},
			expectedPath: "/api/dcim/sites",
			mockResponse: `{"id": 1, "name": "New Site", "slug": "new-site", "status": "active"}`,
			mockStatus:   http.StatusCreated,
		},
		{
			name: "validation error",
			input: &CreateSiteInput{
				Name: "", // Required field
				Slug: "new-site",
			},
			expectError: true,
		},
		{
			name: "server error",
			input: &CreateSiteInput{
				Name: "New Site",
				Slug: "new-site",
			},
			expectedPath: "/api/dcim/sites",
			mockStatus:   http.StatusInternalServerError,
			expectError:  true,
		},
	}

	for _, spec_test := range tests {
		test.Run(spec_test.name, func(test *testing.T) {
			client := client.NewMockClient(test, spec_test.expectedPath, spec_test.mockResponse, spec_test.mockStatus)
			service := NewService(client)

			site, err := service.CreateSite(spec_test.input)
			if spec_test.expectError {
				assert.Error(test, err)
				return
			}

			assert.NoError(test, err)
			assert.NotNil(test, site)
		})
	}
}

func TestUpdateSite(test *testing.T) {
	tests := []struct {
		name         string
		input        *UpdateSiteInput
		expectedPath string
		mockResponse string
		mockStatus   int
		expectError  bool
	}{
		{
			name: "successful update",
			input: &UpdateSiteInput{
				ID:   1,
				Name: "Updated Site",
				Slug: "updated-site",
			},
			expectedPath: "/api/dcim/sites/1",
			mockResponse: `{"id": 1, "name": "Updated Site", "slug": "updated-site"}`,
			mockStatus:   http.StatusOK,
		},
		{
			name: "not found",
			input: &UpdateSiteInput{
				ID:   999,
				Name: "Updated Site",
			},
			expectedPath: "/api/dcim/sites/999",
			mockStatus:   http.StatusNotFound,
			expectError:  true,
		},
		{
			name: "validation error",
			input: &UpdateSiteInput{
				ID:   1,
				Name: "", // Required field
			},
			expectError: true,
		},
	}

	for _, spec_test := range tests {
		test.Run(spec_test.name, func(test *testing.T) {
			client := client.NewMockClient(test, spec_test.expectedPath, spec_test.mockResponse, spec_test.mockStatus)
			service := NewService(client)

			site, err := service.UpdateSite(spec_test.input)
			if spec_test.expectError {
				assert.Error(test, err)
				return
			}

			assert.NoError(test, err)
			assert.NotNil(test, site)
		})
	}
}

func TestPatchSite(test *testing.T) {
	tests := []struct {
		name         string
		input        *PatchSiteInput
		expectedPath string
		mockResponse string
		mockStatus   int
		expectError  bool
	}{
		{
			name: "successful patch",
			input: &PatchSiteInput{
				ID:   intPtr(1),
				Name: stringPtr("Patched Site"),
			},
			expectedPath: "/api/dcim/sites/1",
			mockResponse: `{"id": 1, "name": "Patched Site"}`,
			mockStatus:   http.StatusOK,
		},
		{
			name: "not found",
			input: &PatchSiteInput{
				ID:   intPtr(999),
				Name: stringPtr("Patched Site"),
			},
			expectedPath: "/api/dcim/sites/999",
			mockStatus:   http.StatusNotFound,
			expectError:  true,
		},
		{
			name:        "validation error",
			input:       &PatchSiteInput{}, // Missing ID
			expectError: true,
		},
	}

	for _, spec_test := range tests {
		test.Run(spec_test.name, func(test *testing.T) {
			client := client.NewMockClient(test, spec_test.expectedPath, spec_test.mockResponse, spec_test.mockStatus)
			service := NewService(client)

			site, err := service.PatchSite(spec_test.input)
			if spec_test.expectError {
				assert.Error(test, err)
				return
			}

			assert.NoError(test, err)
			assert.NotNil(test, site)
		})
	}
}

func TestDeleteSite(test *testing.T) {
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
			expectedPath: "/api/dcim/sites/1",
			mockStatus:   http.StatusNoContent,
		},
		{
			name:         "not found",
			id:           999,
			expectedPath: "/api/dcim/sites/999",
			mockStatus:   http.StatusNotFound,
			expectError:  true,
		},
		{
			name:         "server error",
			id:           1,
			expectedPath: "/api/dcim/sites/1",
			mockStatus:   http.StatusInternalServerError,
			expectError:  true,
		},
	}

	for _, spec_test := range tests {
		test.Run(spec_test.name, func(test *testing.T) {
			client := client.NewMockClient(test, spec_test.expectedPath, "", spec_test.mockStatus)
			service := NewService(client)

			err := service.DeleteSite(spec_test.id)
			if spec_test.expectError {
				assert.Error(test, err)
				return
			}

			assert.NoError(test, err)
		})
	}
}
