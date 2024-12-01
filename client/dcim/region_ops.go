package dcim

import (
	"fmt"
	"net/http"
)

// ListRegions lists all regions matching the input criteria
func (service *Service) ListRegions(input *ListRegionsInput) ([]Region, error) {
	path := service.buildPath("dcim", "regions")
	
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
	var response struct {
		Count    int      `json:"count"`
		Next     *string  `json:"next"`
		Previous *string  `json:"previous"`
		Results  []Region `json:"results"`
	}

	resp, err := service.client.httpClient.R().
		SetQueryParams(params).
		SetResult(&response).
		Get(path)
	
	if err != nil {
		return nil, fmt.Errorf("error listing regions: %w", err)
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}

	return response.Results, nil
}

// GetRegion retrieves a single region by ID
func (service *Service) GetRegion(id int) (*Region, error) {
	path := service.buildPath("dcim", "regions", fmt.Sprintf("%d", id))
	
	var region Region
	resp, err := service.client.httpClient.R().
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

	path := service.buildPath("dcim", "regions")
	
	var region Region
	resp, err := service.client.httpClient.R().
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

	path := service.buildPath("dcim", "regions", fmt.Sprintf("%d", input.ID))
	
	var region Region
	resp, err := service.client.httpClient.R().
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

// PatchRegion partially updates an existing region
func (service *Service) PatchRegion(input *PatchRegionInput) (*Region, error) {
	if err := input.Validate(); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	path := service.buildPath("dcim", "regions", fmt.Sprintf("%d", input.ID))
	
	var region Region
	resp, err := service.client.httpClient.R().
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

// PutRegion creates or updates a region
func (service *Service) PutRegion(input *UpdateRegionInput) (*Region, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	req, err := service.client.NewRequest("PUT", fmt.Sprintf("dcim/regions/%d/", input.ID), nil)
	if err != nil {
		return nil, err
	}

	req.Body = input
	region := new(Region)
	err = service.client.Do(req, region)
	if err != nil {
		return nil, err
	}

	return region, nil
}

// DeleteRegion deletes a region
func (service *Service) DeleteRegion(id int) error {
	path := service.buildPath("dcim", "regions", fmt.Sprintf("%d", id))
	
	resp, err := service.client.httpClient.R().
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
