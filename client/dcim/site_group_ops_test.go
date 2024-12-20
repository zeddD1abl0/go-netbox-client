package dcim

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zeddD1abl0/go-netbox-client/client"
)

func TestListSiteGroups(test *testing.T) {
	tests := []struct {
		name           string
		input          *ListSiteGroupsInput
		expectedPath   string
		expectedParams map[string]string
		mockResponse   string
		mockStatus     int
		expectError    bool
	}{
		{
			name:         "successful list with no filters",
			input:        &ListSiteGroupsInput{},
			expectedPath: "/api/dcim/site-groups",
			mockResponse: `{"count": 2, "results": [{"id": 1, "name": "Group 1"}, {"id": 2, "name": "Group 2"}]}`,
			mockStatus:   http.StatusOK,
		},
		{
			name: "successful list with filters",
			input: &ListSiteGroupsInput{
				Name:   "Test",
				Parent: "parent-group",
				Tag:    "prod",
				Limit:  10,
				Offset: 0,
			},
			expectedPath: "/api/dcim/site-groups",
			expectedParams: map[string]string{
				"name":   "Test",
				"parent": "parent-group",
				"tag":    "prod",
				"limit":  "10",
			},
			mockResponse: `{"count": 1, "results": [{"id": 1, "name": "Test Group"}]}`,
			mockStatus:   http.StatusOK,
		},
		{
			name:         "server error",
			input:        &ListSiteGroupsInput{},
			expectedPath: "/api/dcim/site-groups",
			mockStatus:   http.StatusInternalServerError,
			expectError:  true,
		},
	}

	for _, spec_test := range tests {
		test.Run(spec_test.name, func(test *testing.T) {
			client := client.NewMockClient(test, spec_test.expectedPath, spec_test.mockResponse, spec_test.mockStatus)
			service := NewService(client)

			siteGroups, err := service.ListSiteGroups(spec_test.input)
			if spec_test.expectError {
				assert.Error(test, err)
				return
			}

			assert.NoError(test, err)
			assert.NotNil(test, siteGroups)
		})
	}
}

func TestGetSiteGroup(test *testing.T) {
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
			expectedPath: "/api/dcim/site-groups/1",
			mockResponse: `{"id": 1, "name": "Test Group"}`,
			mockStatus:   http.StatusOK,
		},
		{
			name:         "not found",
			id:           999,
			expectedPath: "/api/dcim/site-groups/999",
			mockStatus:   http.StatusNotFound,
			expectError:  true,
		},
		{
			name:         "server error",
			id:           1,
			expectedPath: "/api/dcim/site-groups/1",
			mockStatus:   http.StatusInternalServerError,
			expectError:  true,
		},
	}

	for _, spec_test := range tests {
		test.Run(spec_test.name, func(test *testing.T) {
			client := client.NewMockClient(test, spec_test.expectedPath, spec_test.mockResponse, spec_test.mockStatus)
			service := NewService(client)

			siteGroup, err := service.GetSiteGroup(spec_test.id)
			if spec_test.expectError {
				assert.Error(test, err)
				return
			}

			assert.NoError(test, err)
			assert.NotNil(test, siteGroup)
		})
	}
}

func TestCreateSiteGroup(test *testing.T) {
	tests := []struct {
		name         string
		input        *CreateSiteGroupInput
		expectedPath string
		mockResponse string
		mockStatus   int
		expectError  bool
	}{
		{
			name: "successful create",
			input: &CreateSiteGroupInput{
				Name: "Test Group",
				Slug: "test-group",
			},
			expectedPath: "/api/dcim/site-groups",
			mockResponse: `{"id": 1, "name": "Test Group", "slug": "test-group"}`,
			mockStatus:   http.StatusCreated,
		},
		{
			name: "validation error",
			input: &CreateSiteGroupInput{
				Name: "", // Required field
				Slug: "test-group",
			},
			expectError: true,
		},
		{
			name: "server error",
			input: &CreateSiteGroupInput{
				Name: "Test Group",
				Slug: "test-group",
			},
			expectedPath: "/api/dcim/site-groups",
			mockStatus:   http.StatusInternalServerError,
			expectError:  true,
		},
	}

	for _, spec_test := range tests {
		test.Run(spec_test.name, func(test *testing.T) {
			client := client.NewMockClient(test, spec_test.expectedPath, spec_test.mockResponse, spec_test.mockStatus)
			service := NewService(client)

			siteGroup, err := service.CreateSiteGroup(spec_test.input)
			if spec_test.expectError {
				assert.Error(test, err)
				return
			}

			assert.NoError(test, err)
			assert.NotNil(test, siteGroup)
		})
	}
}

func TestUpdateSiteGroup(test *testing.T) {
	tests := []struct {
		name         string
		input        *UpdateSiteGroupInput
		expectedPath string
		mockResponse string
		mockStatus   int
		expectError  bool
	}{
		{
			name: "successful update",
			input: &UpdateSiteGroupInput{
				ID:   1,
				Name: "Updated Group",
				Slug: "updated-group",
			},
			expectedPath: "/api/dcim/site-groups/1",
			mockResponse: `{"id": 1, "name": "Updated Group", "slug": "updated-group"}`,
			mockStatus:   http.StatusOK,
		},
		{
			name: "not found",
			input: &UpdateSiteGroupInput{
				ID:   999,
				Name: "Updated Group",
				Slug: "updated-group",
			},
			expectedPath: "/api/dcim/site-groups/999",
			mockStatus:   http.StatusNotFound,
			expectError:  true,
		},
		{
			name: "validation error",
			input: &UpdateSiteGroupInput{
				ID:   1,
				Name: "", // Required field
				Slug: "updated-group",
			},
			expectError: true,
		},
	}

	for _, spec_test := range tests {
		test.Run(spec_test.name, func(test *testing.T) {
			client := client.NewMockClient(test, spec_test.expectedPath, spec_test.mockResponse, spec_test.mockStatus)
			service := NewService(client)

			siteGroup, err := service.UpdateSiteGroup(spec_test.input)
			if spec_test.expectError {
				assert.Error(test, err)
				return
			}

			assert.NoError(test, err)
			assert.NotNil(test, siteGroup)
		})
	}
}

func TestPatchSiteGroup(test *testing.T) {
	tests := []struct {
		name         string
		input        *PatchSiteGroupInput
		expectedPath string
		mockResponse string
		mockStatus   int
		expectError  bool
	}{
		{
			name: "successful patch",
			input: &PatchSiteGroupInput{
				ID:   1,
				Name: stringPtr("Patched Group"),
				Slug: stringPtr("patched-group"),
			},
			expectedPath: "/api/dcim/site-groups/1",
			mockResponse: `{"id": 1, "name": "Patched Group", "slug": "patched-group"}`,
			mockStatus:   http.StatusOK,
		},
		{
			name: "not found",
			input: &PatchSiteGroupInput{
				ID:   999,
				Name: stringPtr("Patched Group"),
			},
			expectedPath: "/api/dcim/site-groups/999",
			mockStatus:   http.StatusNotFound,
			expectError:  true,
		},
		{
			name:        "validation error",
			input:       &PatchSiteGroupInput{}, // Missing ID
			expectError: true,
		},
	}

	for _, spec_test := range tests {
		test.Run(spec_test.name, func(test *testing.T) {
			client := client.NewMockClient(test, spec_test.expectedPath, spec_test.mockResponse, spec_test.mockStatus)
			service := NewService(client)

			siteGroup, err := service.PatchSiteGroup(spec_test.input)
			if spec_test.expectError {
				assert.Error(test, err)
				return
			}

			assert.NoError(test, err)
			assert.NotNil(test, siteGroup)
		})
	}
}

func TestDeleteSiteGroup(test *testing.T) {
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
			expectedPath: "/api/dcim/site-groups/1",
			mockStatus:   http.StatusNoContent,
		},
		{
			name:         "not found",
			id:           999,
			expectedPath: "/api/dcim/site-groups/999",
			mockStatus:   http.StatusNotFound,
			expectError:  true,
		},
		{
			name:         "server error",
			id:           1,
			expectedPath: "/api/dcim/site-groups/1",
			mockStatus:   http.StatusInternalServerError,
			expectError:  true,
		},
	}

	for _, spec_test := range tests {
		test.Run(spec_test.name, func(test *testing.T) {
			client := client.NewMockClient(test, spec_test.expectedPath, "", spec_test.mockStatus)
			service := NewService(client)

			err := service.DeleteSiteGroup(spec_test.id)
			if spec_test.expectError {
				assert.Error(test, err)
				return
			}

			assert.NoError(test, err)
		})
	}
}
