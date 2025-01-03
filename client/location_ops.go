package client

import (
	"fmt"
	"net/http"
)

// ListLocations lists all locations
func (c *Client) ListLocations(input *ListLocationsInput) ([]Location, error) {
	path := c.BuildPath("dcim", "locations")

	// Build query parameters
	params := map[string]string{}
	if input.Name != "" {
		params["name__ic"] = input.Name
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
	var response Response
	response.Results = make([]any, 0)
	_, err := c.R().
		SetQueryParams(params).
		SetResult(&response).
		Get(path)

	if err != nil {
		return nil, fmt.Errorf("error listing locations: %w", err)
	}

	// Convert results to []Location
	locations := make([]Location, len(response.Results))
	for i, result := range response.Results {
		resultMap, ok := result.(map[string]any)
		if !ok {
			return nil, fmt.Errorf("unexpected result type at index %d", i)
		}

		// Create a new Location
		var location Location
		err := convertMapToStruct(resultMap, &location)
		if err != nil {
			return nil, fmt.Errorf("error converting map to struct at index %d: %w", i, err)
		}

		locations[i] = location
	}

	return locations, nil
}

// GetLocation retrieves a single location by ID
func (c *Client) GetLocation(id int) (*Location, error) {
	path := c.BuildPath("dcim", "locations", fmt.Sprintf("%d", id))

	var location Location
	resp, err := c.R().
		SetResult(&location).
		Get(path)

	if err != nil {
		return nil, fmt.Errorf("error getting location: %w", err)
	}

	if resp.StatusCode() == http.StatusNotFound {
		return nil, fmt.Errorf("location not found")
	}

	return &location, nil
}

// CreateLocation creates a new location
func (c *Client) CreateLocation(input *CreateLocationInput) (*Location, error) {
	path := c.BuildPath("dcim", "locations")

	var location Location
	resp, err := c.R().
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
func (c *Client) UpdateLocation(input *UpdateLocationInput) (*Location, error) {
	path := c.BuildPath("dcim", "locations", fmt.Sprintf("%d", input.ID))

	var location Location
	resp, err := c.R().
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

// PatchLocation patches an existing location
func (c *Client) PatchLocation(input *PatchLocationInput) (*Location, error) {
	path := c.BuildPath("dcim", "locations", fmt.Sprintf("%d", input.ID))

	var location Location
	resp, err := c.R().
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

// DeleteLocation deletes a location
func (c *Client) DeleteLocation(id int) error {
	path := c.BuildPath("dcim", "locations", fmt.Sprintf("%d", id))

	resp, err := c.R().
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
