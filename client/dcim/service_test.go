package dcim

import (
	"net/http"
	"testing"

	"github.com/zeddD1abl0/go-netbox-client/client"
	"github.com/zeddD1abl0/go-netbox-client/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestService_ListSites(t *testing.T) {
	tests := []struct {
		name           string
		input          *ListSitesInput
		mockResponse   interface{}
		expectedSites  []Site
		expectedError  string
		responseStatus int
	}{
		{
			name: "successful list sites",
			input: &ListSitesInput{
				Name:   "test-site",
				Status: "active",
				Limit:  10,
			},
			mockResponse: client.mockPaginatedResponse([]Site{
				{
					ID:   1,
					Name: "test-site",
					Status: &Status{
						Value: "active",
						Label: "Active",
					},
					Tags: []models.Tag{
						{
							ID:   1,
							Name: "test-tag",
						},
					},
				},
			}),
			expectedSites: []Site{
				{
					ID:   1,
					Name: "test-site",
					Status: &Status{
						Value: "active",
						Label: "Active",
					},
					Tags: []models.Tag{
						{
							ID:   1,
							Name: "test-tag",
						},
					},
				},
			},
			responseStatus: http.StatusOK,
		},
		{
			name:           "empty response",
			input:          &ListSitesInput{},
			mockResponse:   client.mockPaginatedResponse([]Site{}),
			expectedSites:  []Site{},
			responseStatus: http.StatusOK,
		},
		{
			name:           "server error",
			input:          &ListSitesInput{},
			mockResponse:   map[string]string{"detail": "Internal server error"},
			expectedError:  "error listing sites",
			responseStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create test server
			ts := client.newTestServer(t, tt.responseStatus, tt.mockResponse)
			defer ts.Close()

			// Create test client
			c := client.newTestClient(t, ts)
			service := &Service{client.NewService(c)}

			// Make request
			sites, err := service.ListSites(tt.input)

			// Check error
			if tt.expectedError != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedError)
				return
			}

			// Check response
			require.NoError(t, err)
			assert.Equal(t, tt.expectedSites, sites)
		})
	}
}

func TestService_GetSite(t *testing.T) {
	tests := []struct {
		name           string
		siteID         int
		mockResponse   interface{}
		expectedSite   *Site
		expectedError  string
		responseStatus int
	}{
		{
			name:   "successful get site",
			siteID: 1,
			mockResponse: Site{
				ID:   1,
				Name: "test-site",
				Status: &Status{
					Value: "active",
					Label: "Active",
				},
			},
			expectedSite: &Site{
				ID:   1,
				Name: "test-site",
				Status: &Status{
					Value: "active",
					Label: "Active",
				},
			},
			responseStatus: http.StatusOK,
		},
		{
			name:           "site not found",
			siteID:         999,
			mockResponse:   map[string]string{"detail": "Not found"},
			expectedError:  "site not found",
			responseStatus: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create test server
			ts := client.newTestServer(t, tt.responseStatus, tt.mockResponse)
			defer ts.Close()

			// Create test client
			c := client.newTestClient(t, ts)
			service := &Service{client.NewService(c)}

			// Make request
			site, err := service.GetSite(tt.siteID)

			// Check error
			if tt.expectedError != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedError)
				return
			}

			// Check response
			require.NoError(t, err)
			assert.Equal(t, tt.expectedSite, site)
		})
	}
}

func TestService_CreateSite(t *testing.T) {
	tests := []struct {
		name           string
		input          *CreateSiteInput
		mockResponse   interface{}
		expectedSite   *Site
		expectedError  string
		responseStatus int
	}{
		{
			name: "successful create",
			input: &CreateSiteInput{
				Name:   "new-site",
				Slug:   "new-site",
				Status: "active",
			},
			mockResponse: Site{
				ID:     1,
				Name:   "new-site",
				Slug:   "new-site",
				Status: &Status{Value: "active"},
			},
			expectedSite: &Site{
				ID:     1,
				Name:   "new-site",
				Slug:   "new-site",
				Status: &Status{Value: "active"},
			},
			responseStatus: http.StatusCreated,
		},
		{
			name: "validation error",
			input: &CreateSiteInput{
				Name: "new-site",
				// Missing required field 'slug'
			},
			mockResponse:   map[string]string{"slug": ["This field is required."]},
			expectedError:  "unexpected status code: 400",
			responseStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := client.newTestServer(t, tt.responseStatus, tt.mockResponse)
			defer ts.Close()

			c := client.newTestClient(t, ts)
			service := &Service{client.NewService(c)}

			site, err := service.CreateSite(tt.input)

			if tt.expectedError != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedError)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.expectedSite, site)
		})
	}
}

func TestService_UpdateSite(t *testing.T) {
	tests := []struct {
		name           string
		input          *UpdateSiteInput
		mockResponse   interface{}
		expectedSite   *Site
		expectedError  string
		responseStatus int
	}{
		{
			name: "successful update",
			input: &UpdateSiteInput{
				ID:     1,
				Name:   "updated-site",
				Slug:   "updated-site",
				Status: "planned",
			},
			mockResponse: Site{
				ID:     1,
				Name:   "updated-site",
				Slug:   "updated-site",
				Status: &Status{Value: "planned"},
			},
			expectedSite: &Site{
				ID:     1,
				Name:   "updated-site",
				Slug:   "updated-site",
				Status: &Status{Value: "planned"},
			},
			responseStatus: http.StatusOK,
		},
		{
			name: "site not found",
			input: &UpdateSiteInput{
				ID:   999,
				Name: "not-found",
				Slug: "not-found",
			},
			mockResponse:   map[string]string{"detail": "Not found."},
			expectedError:  "site not found",
			responseStatus: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := client.newTestServer(t, tt.responseStatus, tt.mockResponse)
			defer ts.Close()

			c := client.newTestClient(t, ts)
			service := &Service{client.NewService(c)}

			site, err := service.UpdateSite(tt.input)

			if tt.expectedError != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedError)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.expectedSite, site)
		})
	}
}

func TestService_DeleteSite(t *testing.T) {
	tests := []struct {
		name           string
		siteID         int
		mockResponse   interface{}
		expectedError  string
		responseStatus int
	}{
		{
			name:           "successful delete",
			siteID:         1,
			responseStatus: http.StatusNoContent,
		},
		{
			name:           "site not found",
			siteID:         999,
			mockResponse:   map[string]string{"detail": "Not found."},
			expectedError:  "site not found",
			responseStatus: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := client.newTestServer(t, tt.responseStatus, tt.mockResponse)
			defer ts.Close()

			c := client.newTestClient(t, ts)
			service := &Service{client.NewService(c)}

			err := service.DeleteSite(tt.siteID)

			if tt.expectedError != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedError)
				return
			}

			require.NoError(t, err)
		})
	}
}

func TestService_PatchSite(t *testing.T) {
	tests := []struct {
		name           string
		input          *PatchSiteInput
		mockResponse   interface{}
		expectedSite   *Site
		expectedError  string
		responseStatus int
	}{
		{
			name: "successful patch",
			input: &PatchSiteInput{
				ID:     1,
				Status: stringPtr("planned"),
			},
			mockResponse: Site{
				ID:     1,
				Name:   "existing-site",
				Slug:   "existing-site",
				Status: &Status{Value: "planned"},
			},
			expectedSite: &Site{
				ID:     1,
				Name:   "existing-site",
				Slug:   "existing-site",
				Status: &Status{Value: "planned"},
			},
			responseStatus: http.StatusOK,
		},
		{
			name: "invalid slug",
			input: &PatchSiteInput{
				ID:   1,
				Slug: stringPtr("invalid slug"),
			},
			expectedError:  "validation failed: slug: must contain only alphanumeric characters, hyphens, and underscores",
			responseStatus: http.StatusBadRequest,
		},
		{
			name: "site not found",
			input: &PatchSiteInput{
				ID:     999,
				Status: stringPtr("planned"),
			},
			mockResponse:   map[string]string{"detail": "Not found."},
			expectedError:  "site not found",
			responseStatus: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := client.newTestServer(t, tt.responseStatus, tt.mockResponse)
			defer ts.Close()

			c := client.newTestClient(t, ts)
			service := &Service{client.NewService(c)}

			site, err := service.PatchSite(tt.input)

			if tt.expectedError != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedError)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.expectedSite, site)
		})
	}
}

func TestService_BulkDeleteSites(t *testing.T) {
	tests := []struct {
		name           string
		input          *BulkDeleteSitesInput
		mockResponse   interface{}
		expectedError  string
		responseStatus int
	}{
		{
			name: "successful bulk delete",
			input: &BulkDeleteSitesInput{
				IDs: []int{1, 2, 3},
			},
			responseStatus: http.StatusNoContent,
		},
		{
			name: "empty IDs",
			input: &BulkDeleteSitesInput{
				IDs: []int{},
			},
			expectedError: "validation failed: ids: at least one ID must be provided",
		},
		{
			name: "partial failure",
			input: &BulkDeleteSitesInput{
				IDs: []int{1, 999},
			},
			mockResponse:   map[string]string{"detail": "Some sites could not be deleted."},
			expectedError:  "unexpected status code: 400",
			responseStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := client.newTestServer(t, tt.responseStatus, tt.mockResponse)
			defer ts.Close()

			c := client.newTestClient(t, ts)
			service := &Service{client.NewService(c)}

			err := service.BulkDeleteSites(tt.input)

			if tt.expectedError != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedError)
				return
			}

			require.NoError(t, err)
		})
	}
}

func TestService_BulkCreateSites(t *testing.T) {
	tests := []struct {
		name           string
		input          *BulkCreateSitesInput
		mockResponse   interface{}
		expectedSites  []Site
		expectedError  string
		responseStatus int
	}{
		{
			name: "successful bulk create",
			input: &BulkCreateSitesInput{
				Sites: []CreateSiteInput{
					{
						Name: "site-1",
						Slug: "site-1",
					},
					{
						Name: "site-2",
						Slug: "site-2",
					},
				},
			},
			mockResponse: []Site{
				{
					ID:   1,
					Name: "site-1",
					Slug: "site-1",
				},
				{
					ID:   2,
					Name: "site-2",
					Slug: "site-2",
				},
			},
			expectedSites: []Site{
				{
					ID:   1,
					Name: "site-1",
					Slug: "site-1",
				},
				{
					ID:   2,
					Name: "site-2",
					Slug: "site-2",
				},
			},
			responseStatus: http.StatusCreated,
		},
		{
			name: "empty sites",
			input: &BulkCreateSitesInput{
				Sites: []CreateSiteInput{},
			},
			expectedError: "validation failed: sites: at least one site must be provided",
		},
		{
			name: "invalid site in batch",
			input: &BulkCreateSitesInput{
				Sites: []CreateSiteInput{
					{
						Name: "site-1",
						Slug: "site-1",
					},
					{
						Name: "site-2",
						Slug: "invalid slug",
					},
				},
			},
			expectedError: "validation failed: sites[1].slug: must contain only alphanumeric characters, hyphens, and underscores",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := client.newTestServer(t, tt.responseStatus, tt.mockResponse)
			defer ts.Close()

			c := client.newTestClient(t, ts)
			service := &Service{client.NewService(c)}

			sites, err := service.BulkCreateSites(tt.input)

			if tt.expectedError != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedError)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.expectedSites, sites)
		})
	}
}

// Helper function to create string pointers
func stringPtr(s string) *string {
	return &s
}
