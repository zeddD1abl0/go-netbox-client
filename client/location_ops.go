package client

import (
	"fmt"
	"net/http"

	"github.com/zeddD1abl0/go-netbox-client/client"
)

// ListLocations lists all locations
func (service *Service) ListLocations(input *ListLocationsInput) ([]Location, error) {
	path := service.BuildPath("dcim", "locations")

	// Build query parameters
	params := map[string]string{}
	if input.Name != "" {
		params["name"] = input.Name
	}
	if input.Site != "" {
		params["site"] = input.Site
	}
	if input.Parent != "" {
		params["parent"] = input.Parent
	}
	if input.Tag != "" {
		params["tag"] = input.Tag
	}
	if input.Limit > 0 {
		params["limit"] = fmt.Sprintf("%d", input.Limit)
	}
	if input.Offset > 0 {
		params["offset"] = fmt.Sprintf("%d", input.Offset)
	}

	// Make request
	var response client.Response
	response.Results = make([]any, 0)
	_, err := service.Client.R().
		SetQueryParams(params).
		SetResult(&response).
		Get(path)

	if err != nil {
		return nil, fmt.Errorf("error listing locations: %w", err)
	}

	// Convert results to []Location
	locations := make([]Location, len(response.Results))
	for i, result := range response.Results {
		location, ok := result.(Location)
		if !ok {
			return nil, fmt.Errorf("unexpected result type at index %d", i)
		}
		locations[i] = location
	}

	return locations, nil
}

// GetLocation retrieves a single location by ID
func (service *Service) GetLocation(id int) (*Location, error) {
	path := service.BuildPath("dcim", "locations", fmt.Sprintf("%d", id))

	var location Location
	resp, err := service.Client.R().
		SetResult(&location).
		Get(path)

	if err != nil {
		return nil, fmt.Errorf("error getting location: %w", err)
	}

	if resp.StatusCode() == http.StatusNotFound {
		return nil, fmt.Errorf("location not found")
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}

	return &location, nil
}

// CreateLocation creates a new location
func (service *Service) CreateLocation(input *CreateLocationInput) (*Location, error) {
	if err := input.Validate(); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	path := service.BuildPath("dcim", "locations")

	var location Location
	resp, err := service.Client.R().
		SetBody(input).
		SetResult(&location).
		Post(path)

	if err != nil {
		return nil, fmt.Errorf("error creating location: %w", err)
	}

	if resp.StatusCode() != http.StatusCreated {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}

	return &location, nil
}

// UpdateLocation updates an existing location
func (service *Service) UpdateLocation(input *UpdateLocationInput) (*Location, error) {
	if err := input.Validate(); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	path := service.BuildPath("dcim", "locations", fmt.Sprintf("%d", input.ID))

	var location Location
	resp, err := service.Client.R().
		SetBody(input).
		SetResult(&location).
		Put(path)

	if err != nil {
		return nil, fmt.Errorf("error updating location: %w", err)
	}

	if resp.StatusCode() == http.StatusNotFound {
		return nil, fmt.Errorf("location not found")
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}

	return &location, nil
}

// DeleteLocation deletes a location
func (service *Service) DeleteLocation(id int) error {
	path := service.BuildPath("dcim", "locations", fmt.Sprintf("%d", id))

	resp, err := service.Client.R().
		Delete(path)

	if err != nil {
		return fmt.Errorf("error deleting location: %w", err)
	}

	if resp.StatusCode() == http.StatusNotFound {
		return fmt.Errorf("location not found")
	}

	if resp.StatusCode() != http.StatusNoContent {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}

	return nil
}

// PutLocation creates or updates a location
func (service *Service) PutLocation(input *UpdateLocationInput) (*Location, error) {
	if err := input.Validate(); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	path := service.BuildPath("dcim", "locations", fmt.Sprintf("%d", input.ID))

	var location Location
	resp, err := service.Client.R().
		SetBody(input).
		SetResult(&location).
		Put(path)

	if err != nil {
		return nil, fmt.Errorf("error updating location: %w", err)
	}

	if resp.StatusCode() == http.StatusNotFound {
		return nil, fmt.Errorf("location not found")
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}

	return &location, nil
}

// PatchLocation patches a location
func (service *Service) PatchLocation(input *PatchLocationInput) (*Location, error) {
	if err := input.Validate(); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	path := service.BuildPath("dcim", "locations", fmt.Sprintf("%d", input.ID))

	var location Location
	resp, err := service.Client.R().
		SetBody(input).
		SetResult(&location).
		Patch(path)

	if err != nil {
		return nil, fmt.Errorf("error patching location: %w", err)
	}

	if resp.StatusCode() == http.StatusNotFound {
		return nil, fmt.Errorf("location not found")
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}

	return &location, nil
}
