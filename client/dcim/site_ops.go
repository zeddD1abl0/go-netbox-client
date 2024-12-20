package dcim

import (
	"fmt"
	"net/http"

	"github.com/zeddD1abl0/go-netbox-client/client"
)

// ListSites lists all sites matching the input criteria
func (service *Service) ListSites(input *ListSitesInput) ([]Site, error) {
	path := service.BuildPath("api", "dcim", "sites")

	// Build query parameters
	params := map[string]string{}
	if input.Name != "" {
		params["name__ic"] = input.Name
	}
	if input.Region != "" {
		params["region_id"] = input.Region
	}
	if input.Status != "" {
		params["status"] = input.Status
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
		return nil, fmt.Errorf("error listing sites: %w", err)
	}

	// Convert results to []Site
	sites := make([]Site, len(response.Results))
	for i, result := range response.Results {
		resultMap, ok := result.(map[string]any)
		if !ok {
			return nil, fmt.Errorf("unexpected result type at index %d", i)
		}

		// Create a new Site
		var site Site

		// Map basic fields
		if id, ok := resultMap["id"].(float64); ok {
			site.ID = int(id)
		}
		if url, ok := resultMap["url"].(string); ok {
			site.URL = url
		}
		if name, ok := resultMap["name"].(string); ok {
			site.Name = name
		}
		if slug, ok := resultMap["slug"].(string); ok {
			site.Slug = slug
		}
		if description, ok := resultMap["description"].(string); ok {
			site.Description = description
		}
		if created, ok := resultMap["created"].(string); ok {
			site.Created = created
		}
		if lastUpdated, ok := resultMap["last_updated"].(string); ok {
			site.LastUpdated = lastUpdated
		}
		if status, ok := resultMap["status"].(map[string]any); ok {
			if value, ok := status["value"].(string); ok {
				site.Status = &Status{Value: value, Label: value}
			}
		}

		// Map region if present
		if regionMap, ok := resultMap["region"].(map[string]any); ok {
			region := &Region{}
			if regionID, ok := regionMap["id"].(float64); ok {
				region.ID = int(regionID)
			}
			if regionName, ok := regionMap["name"].(string); ok {
				region.Name = regionName
			}
			if regionSlug, ok := regionMap["slug"].(string); ok {
				region.Slug = regionSlug
			}
			site.Region = region
		}

		sites[i] = site
	}

	return sites, nil
}

// GetSite retrieves a single site by ID
func (service *Service) GetSite(id int) (*Site, error) {
	path := service.BuildPath("api", "dcim", "sites", fmt.Sprintf("%d", id))

	var site Site
	resp, err := service.Client.R().
		SetResult(&site).
		Get(path)

	if err != nil {
		return nil, fmt.Errorf("error getting site: %w", err)
	}

	if resp.StatusCode() == http.StatusNotFound {
		return nil, fmt.Errorf("site not found")
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}

	return &site, nil
}

// CreateSite creates a new site
func (service *Service) CreateSite(input *CreateSiteInput) (*Site, error) {
	if err := input.Validate(); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	path := service.BuildPath("api", "dcim", "sites")

	var site Site
	resp, err := service.Client.R().
		SetBody(input).
		SetResult(&site).
		Post(path)

	if err != nil {
		return nil, fmt.Errorf("error creating site: %w", err)
	}

	if resp.StatusCode() != http.StatusCreated {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}

	return &site, nil
}

// UpdateSite updates an existing site
func (service *Service) UpdateSite(input *UpdateSiteInput) (*Site, error) {
	if err := input.Validate(); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	path := service.BuildPath("api", "dcim", "sites", fmt.Sprintf("%d", input.ID))

	var site Site
	resp, err := service.Client.R().
		SetBody(input).
		SetResult(&site).
		Put(path)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() >= 400 {
		return nil, fmt.Errorf("error updating site: %s", resp.Status())
	}

	return &site, nil
}

// PatchSite partially updates an existing site
func (service *Service) PatchSite(input *PatchSiteInput) (*Site, error) {
	if err := input.Validate(); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	path := service.BuildPath("api", "dcim", "sites", fmt.Sprintf("%d", input.ID))

	var site Site
	resp, err := service.Client.R().
		SetBody(input).
		SetResult(&site).
		Patch(path)

	if err != nil {
		return nil, fmt.Errorf("error patching site: %w", err)
	}

	if resp.StatusCode() == http.StatusNotFound {
		return nil, fmt.Errorf("site not found")
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}

	return &site, nil
}

// DeleteSite deletes a site
func (service *Service) DeleteSite(id int) error {
	path := service.BuildPath("api", "dcim", "sites", fmt.Sprintf("%d", id))

	resp, err := service.Client.R().
		Delete(path)

	if err != nil {
		return fmt.Errorf("error deleting site: %w", err)
	}

	if resp.StatusCode() == http.StatusNotFound {
		return fmt.Errorf("site not found")
	}

	if resp.StatusCode() != http.StatusNoContent {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}

	return nil
}
