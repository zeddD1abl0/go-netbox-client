package client

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
		resultMap, ok := result.(map[string]any)
		if !ok {
			return nil, fmt.Errorf("unexpected result type at index %d", i)
		}

		// Create a new Region
		var region Region

		// Map basic fields
		if id, ok := resultMap["id"].(float64); ok {
			region.ID = int(id)
		}
		if url, ok := resultMap["url"].(string); ok {
			region.URL = url
		}
		if name, ok := resultMap["name"].(string); ok {
			region.Name = name
		}
		if slug, ok := resultMap["slug"].(string); ok {
			region.Slug = slug
		}
		if description, ok := resultMap["description"].(string); ok {
			region.Description = description
		}
		if created, ok := resultMap["created"].(string); ok {
			region.Created = created
		}
		if lastUpdated, ok := resultMap["last_updated"].(string); ok {
			region.LastUpdated = lastUpdated
		}
		if siteCount, ok := resultMap["site_count"].(float64); ok {
			region.SiteCount = int(siteCount)
		}

		// Map parent if present
		if parentMap, ok := resultMap["parent"].(map[string]any); ok {
			parent := &Region{}
			if parentID, ok := parentMap["id"].(float64); ok {
				parent.ID = int(parentID)
			}
			if parentName, ok := parentMap["name"].(string); ok {
				parent.Name = parentName
			}
			if parentSlug, ok := parentMap["slug"].(string); ok {
				parent.Slug = parentSlug
			}
			region.Parent = parent
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

	if resp.StatusCode() != http.StatusOK && resp.StatusCode() != http.StatusCreated {
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

	if resp.StatusCode() != http.StatusOK && resp.StatusCode() != http.StatusCreated {
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

	if resp.StatusCode() != http.StatusOK && resp.StatusCode() != http.StatusCreated {
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
