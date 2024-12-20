package dcim

import (
	"net/http"
	"testing"

	"github.com/zeddD1abl0/go-netbox-client/client"
	"github.com/stretchr/testify/assert"
)

func TestListLocations(test *testing.T) {
	tests := []struct {
		name         string
		input        *ListLocationsInput
		expectedPath string
		mockResponse string
		mockStatus   int
		expectError  bool
	}{
		{
			name:         "successful list with no filters",
			input:        &ListLocationsInput{},
			expectedPath: "/api/dcim/locations",
			mockResponse: `{"count": 2, "results": [{"id": 1, "name": "Location 1"}, {"id": 2, "name": "Location 2"}]}`,
			mockStatus:   http.StatusOK,
		},
		{
			name: "successful list with filters",
			input: &ListLocationsInput{
				Name:   "Test",
				Parent: "parent-location",
				Tag:    "prod",
				Limit:  10,
				Offset: 0,
			},
			expectedPath: "/api/dcim/locations",
			mockResponse: `{"count": 1, "results": [{"id": 1, "name": "Test Location"}]}`,
			mockStatus:   http.StatusOK,
		},
		{
			name:         "server error",
			input:        &ListLocationsInput{},
			expectedPath: "/api/dcim/locations",
			mockStatus:   http.StatusInternalServerError,
			expectError:  true,
		},
	}

	for _, spec_test := range tests {
		test.Run(spec_test.name, func(test *testing.T) {
			mockClient := client.NewMockClient(test, spec_test.expectedPath, spec_test.mockResponse, spec_test.mockStatus)
			service := NewService(mockClient)

			locations, err := service.ListLocations(spec_test.input)
			if spec_test.expectError {
				assert.Error(test, err)
				return
			}

			assert.NoError(test, err)
			assert.NotNil(test, locations)
		})
	}
}

func TestGetLocation(test *testing.T) {
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
			expectedPath: "/api/dcim/locations/1",
			mockResponse: `{"id": 1, "name": "Test Location"}`,
			mockStatus:   http.StatusOK,
		},
		{
			name:         "not found",
			id:           999,
			expectedPath: "/api/dcim/locations/999",
			mockStatus:   http.StatusNotFound,
			expectError:  true,
		},
		{
			name:         "server error",
			id:           1,
			expectedPath: "/api/dcim/locations/1",
			mockStatus:   http.StatusInternalServerError,
			expectError:  true,
		},
	}

	for _, spec_test := range tests {
		test.Run(spec_test.name, func(test *testing.T) {
			mockClient := client.NewMockClient(test, spec_test.expectedPath, spec_test.mockResponse, spec_test.mockStatus)
			service := NewService(mockClient)

			location, err := service.GetLocation(spec_test.id)
			if spec_test.expectError {
				assert.Error(test, err)
				return
			}

			assert.NoError(test, err)
			assert.NotNil(test, location)
		})
	}
}

func TestCreateLocation(test *testing.T) {
	tests := []struct {
		name         string
		input        *CreateLocationInput
		expectedPath string
		mockResponse string
		mockStatus   int
		expectError  bool
	}{
		{
			name: "successful create",
			input: &CreateLocationInput{
				Name: "New Location",
				Slug: "new-location",
			},
			expectedPath: "/api/dcim/locations",
			mockResponse: `{"id": 1, "name": "New Location", "slug": "new-location"}`,
			mockStatus:   http.StatusCreated,
		},
		{
			name: "validation error",
			input: &CreateLocationInput{
				Name: "", // Required field
				Slug: "new-location",
			},
			expectError: true,
		},
		{
			name: "server error",
			input: &CreateLocationInput{
				Name: "New Location",
				Slug: "new-location",
			},
			expectedPath: "/api/dcim/locations",
			mockStatus:   http.StatusInternalServerError,
			expectError:  true,
		},
	}

	for _, spec_test := range tests {
		test.Run(spec_test.name, func(test *testing.T) {
			mockClient := client.NewMockClient(test, spec_test.expectedPath, spec_test.mockResponse, spec_test.mockStatus)
			service := NewService(mockClient)

			location, err := service.CreateLocation(spec_test.input)
			if spec_test.expectError {
				assert.Error(test, err)
				return
			}

			assert.NoError(test, err)
			assert.NotNil(test, location)
		})
	}
}

func TestUpdateLocation(test *testing.T) {
	tests := []struct {
		name         string
		input        *UpdateLocationInput
		expectedPath string
		mockResponse string
		mockStatus   int
		expectError  bool
	}{
		{
			name: "successful update",
			input: &UpdateLocationInput{
				ID:   1,
				Name: "Updated Location",
				Slug: "updated-location",
			},
			expectedPath: "/api/dcim/locations/1",
			mockResponse: `{"id": 1, "name": "Updated Location", "slug": "updated-location"}`,
			mockStatus:   http.StatusOK,
		},
		{
			name: "validation error",
			input: &UpdateLocationInput{
				ID:   1,
				Name: "", // Required field
			},
			expectError: true,
		},
		{
			name: "not found",
			input: &UpdateLocationInput{
				ID:   999,
				Name: "Updated Location",
			},
			expectedPath: "/api/dcim/locations/999",
			mockStatus:   http.StatusNotFound,
			expectError:  true,
		},
	}

	for _, spec_test := range tests {
		test.Run(spec_test.name, func(test *testing.T) {
			mockClient := client.NewMockClient(test, spec_test.expectedPath, spec_test.mockResponse, spec_test.mockStatus)
			service := NewService(mockClient)

			location, err := service.UpdateLocation(spec_test.input)
			if spec_test.expectError {
				assert.Error(test, err)
				return
			}

			assert.NoError(test, err)
			assert.NotNil(test, location)
		})
	}
}

func TestDeleteLocation(test *testing.T) {
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
			expectedPath: "/api/dcim/locations/1",
			mockStatus:   http.StatusNoContent,
		},
		{
			name:         "not found",
			id:           999,
			expectedPath: "/api/dcim/locations/999",
			mockStatus:   http.StatusNotFound,
			expectError:  true,
		},
		{
			name:         "server error",
			id:           1,
			expectedPath: "/api/dcim/locations/1",
			mockStatus:   http.StatusInternalServerError,
			expectError:  true,
		},
	}

	for _, spec_test := range tests {
		test.Run(spec_test.name, func(test *testing.T) {
			mockClient := client.NewMockClient(test, spec_test.expectedPath, "", spec_test.mockStatus)
			service := NewService(mockClient)

			err := service.DeleteLocation(spec_test.id)
			if spec_test.expectError {
				assert.Error(test, err)
				return
			}

			assert.NoError(test, err)
		})
	}
}
