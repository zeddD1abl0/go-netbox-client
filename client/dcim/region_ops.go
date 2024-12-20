package dcim

import (
	"fmt"
	"net/http"

	"github.com/zeddD1abl0/go-netbox-client/client"
)

// ListRegions lists all regions
func (service *Service) ListRegions(input *ListRegionsInput) ([]Region, error) {
	path := service.BuildPath("dcim", "regions")

	// Build query parameters
	params := map[string]string{}
	if input.Name != "" {
		params["name"] = input.Name
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
		return nil, fmt.Errorf("error listing regions: %w", err)
	}

	// Convert results to []Region
	regions := make([]Region, len(response.Results))
	for i, result := range response.Results {
		region, ok := result.(Region)
		if !ok {
			return nil, fmt.Errorf("unexpected result type at index %d", i)
		}
		regions[i] = region
	}

	return regions, nil
}

// GetRegion retrieves a single region by ID
func (service *Service) GetRegion(id int) (*Region, error) {
	path := service.BuildPath("dcim", "regions", fmt.Sprintf("%d", id))

	var region Region
	resp, err := service.Client.R().
		SetResult(&region).
		Get(path)

	if err != nil {
		return nil, fmt.Errorf("error getting region: %w", err)
	}

	if resp.StatusCode() == http.StatusNotFound {
		return nil, fmt.Errorf("region not found")
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}

	return &region, nil
}

// CreateRegion creates a new region
func (service *Service) CreateRegion(input *CreateRegionInput) (*Region, error) {
	if err := input.Validate(); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	path := service.BuildPath("dcim", "regions")

	var region Region
	resp, err := service.Client.R().
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
func (service *Service) UpdateRegion(input *UpdateRegionInput) (*Region, error) {
	if err := input.Validate(); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	path := service.BuildPath("dcim", "regions", fmt.Sprintf("%d", input.ID))

	var region Region
	resp, err := service.Client.R().
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
func (service *Service) PatchRegion(input *PatchRegionInput) (*Region, error) {
	if err := input.Validate(); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	path := service.BuildPath("dcim", "regions", fmt.Sprintf("%d", input.ID))

	var region Region
	resp, err := service.Client.R().
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
func (service *Service) DeleteRegion(id int) error {
	path := service.BuildPath("dcim", "regions", fmt.Sprintf("%d", id))

	resp, err := service.Client.R().
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
