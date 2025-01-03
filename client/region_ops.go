package client

import (
	"fmt"
	"net/http"
)

// ListRegions lists all regions
func (c *Client) ListRegions(input *ListRegionsInput) ([]Region, error) {
	path := c.BuildPath("dcim", "regions")

	// Build query parameters
	params := map[string]string{}
	if input.Name != "" {
		params["name__ic"] = input.Name
	}
	if input.Parent != "" {
		params["parent_id"] = input.Parent
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
		return nil, fmt.Errorf("error listing regions: %w", err)
	}

	// Convert results to []Region
	regions := make([]Region, len(response.Results))
	for i, result := range response.Results {
		resultMap, ok := result.(map[string]any)
		if !ok {
			return nil, fmt.Errorf("unexpected result type at index %d", i)
		}

		// Create a new Region
		var region Region
		err := convertMapToStruct(resultMap, &region)
		if err != nil {
			return nil, fmt.Errorf("error converting map to struct at index %d: %w", i, err)
		}

		regions[i] = region
	}

	return regions, nil
}

// GetRegion retrieves a single region by ID
func (c *Client) GetRegion(id int) (*Region, error) {
	path := c.BuildPath("dcim", "regions", fmt.Sprintf("%d", id))

	var region Region
	resp, err := c.R().
		SetResult(&region).
		Get(path)

	if err != nil {
		return nil, fmt.Errorf("error getting region: %w", err)
	}

	if resp.StatusCode() == http.StatusNotFound {
		return nil, fmt.Errorf("region not found")
	}

	return &region, nil
}

// CreateRegion creates a new region
func (c *Client) CreateRegion(input *CreateRegionInput) (*Region, error) {
	path := c.BuildPath("dcim", "regions")

	var region Region
	resp, err := c.R().
		SetBody(input).
		SetResult(&region).
		Post(path)

	if err != nil {
		return nil, fmt.Errorf("error creating region: %w", err)
	}

	if resp.StatusCode() != http.StatusCreated {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}

	return &region, nil
}

// UpdateRegion updates an existing region
func (c *Client) UpdateRegion(input *UpdateRegionInput) (*Region, error) {
	path := c.BuildPath("dcim", "regions", fmt.Sprintf("%d", input.ID))

	var region Region
	resp, err := c.R().
		SetBody(input).
		SetResult(&region).
		Put(path)

	if err != nil {
		return nil, fmt.Errorf("error updating region: %w", err)
	}

	if resp.StatusCode() == http.StatusNotFound {
		return nil, fmt.Errorf("region not found")
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}

	return &region, nil
}

// PatchRegion patches an existing region
func (c *Client) PatchRegion(input *PatchRegionInput) (*Region, error) {
	path := c.BuildPath("dcim", "regions", fmt.Sprintf("%d", input.ID))

	var region Region
	resp, err := c.R().
		SetBody(input).
		SetResult(&region).
		Patch(path)

	if err != nil {
		return nil, fmt.Errorf("error patching region: %w", err)
	}

	if resp.StatusCode() == http.StatusNotFound {
		return nil, fmt.Errorf("region not found")
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}

	return &region, nil
}

// DeleteRegion deletes a region
func (c *Client) DeleteRegion(id int) error {
	path := c.BuildPath("dcim", "regions", fmt.Sprintf("%d", id))

	resp, err := c.R().
		Delete(path)

	if err != nil {
		return fmt.Errorf("error deleting region: %w", err)
	}

	if resp.StatusCode() == http.StatusNotFound {
		return fmt.Errorf("region not found")
	}

	if resp.StatusCode() != http.StatusNoContent {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}

	return nil
}
