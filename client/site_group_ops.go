package client

import (
	"fmt"
	"net/http"

	"github.com/zeddD1abl0/go-netbox-client/client"
)

// ListSiteGroups lists all site groups matching the input criteria
func (service *Service) ListSiteGroups(input *ListSiteGroupsInput) ([]SiteGroup, error) {
	path := service.BuildPath("dcim", "site-groups")

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
		return nil, fmt.Errorf("error listing site groups: %w", err)
	}

	// Convert results to []SiteGroup
	siteGroups := make([]SiteGroup, len(response.Results))
	for i, result := range response.Results {
		resultMap, ok := result.(map[string]any)
		if !ok {
			return nil, fmt.Errorf("unexpected result type at index %d", i)
		}

		// Create a new SiteGroup
		var siteGroup SiteGroup

		// Map basic fields
		if id, ok := resultMap["id"].(float64); ok {
			siteGroup.ID = int(id)
		}
		if url, ok := resultMap["url"].(string); ok {
			siteGroup.URL = url
		}
		if name, ok := resultMap["name"].(string); ok {
			siteGroup.Name = name
		}
		if slug, ok := resultMap["slug"].(string); ok {
			siteGroup.Slug = slug
		}
		if description, ok := resultMap["description"].(string); ok {
			siteGroup.Description = description
		}
		if created, ok := resultMap["created"].(string); ok {
			siteGroup.Created = created
		}
		if lastUpdated, ok := resultMap["last_updated"].(string); ok {
			siteGroup.LastUpdated = lastUpdated
		}

		// Map parent if present
		if parentMap, ok := resultMap["parent"].(map[string]any); ok {
			parent := &SiteGroup{}
			if parentID, ok := parentMap["id"].(float64); ok {
				parent.ID = int(parentID)
			}
			if parentName, ok := parentMap["name"].(string); ok {
				parent.Name = parentName
			}
			if parentSlug, ok := parentMap["slug"].(string); ok {
				parent.Slug = parentSlug
			}
			siteGroup.Parent = parent
		}

		siteGroups[i] = siteGroup
	}

	return siteGroups, nil
}

// GetSiteGroup retrieves a single site group by ID
func (service *Service) GetSiteGroup(id int) (*SiteGroup, error) {
	path := service.BuildPath("dcim", "site-groups", fmt.Sprintf("%d", id))

	var siteGroup SiteGroup
	resp, err := service.Client.R().
		SetResult(&siteGroup).
		Get(path)

	if err != nil {
		return nil, fmt.Errorf("error getting site group: %w", err)
	}

	if resp.StatusCode() == http.StatusNotFound {
		return nil, fmt.Errorf("site group not found")
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}

	return &siteGroup, nil
}

// CreateSiteGroup creates a new site group
func (service *Service) CreateSiteGroup(input *CreateSiteGroupInput) (*SiteGroup, error) {
	if err := input.Validate(); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	path := service.BuildPath("dcim", "site-groups")

	var siteGroup SiteGroup
	resp, err := service.Client.R().
		SetBody(input).
		SetResult(&siteGroup).
		Post(path)

	if err != nil {
		return nil, fmt.Errorf("error creating site group: %w", err)
	}

	if resp.StatusCode() != http.StatusCreated && resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}

	return &siteGroup, nil
}

// UpdateSiteGroup updates an existing site group
func (service *Service) UpdateSiteGroup(input *UpdateSiteGroupInput) (*SiteGroup, error) {
	if err := input.Validate(); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	path := service.BuildPath("dcim", "site-groups", fmt.Sprintf("%d", input.ID))

	var siteGroup SiteGroup
	resp, err := service.Client.R().
		SetBody(input).
		SetResult(&siteGroup).
		Put(path)

	if err != nil {
		return nil, fmt.Errorf("error updating site group: %w", err)
	}

	if resp.StatusCode() == http.StatusNotFound {
		return nil, fmt.Errorf("site group not found")
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}

	return &siteGroup, nil
}

// PatchSiteGroup partially updates an existing site group
func (service *Service) PatchSiteGroup(input *PatchSiteGroupInput) (*SiteGroup, error) {
	if err := input.Validate(); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	path := service.BuildPath("dcim", "site-groups", fmt.Sprintf("%d", input.ID))

	var siteGroup SiteGroup
	resp, err := service.Client.R().
		SetBody(input).
		SetResult(&siteGroup).
		Patch(path)

	if err != nil {
		return nil, fmt.Errorf("error patching site group: %w", err)
	}

	if resp.StatusCode() == http.StatusNotFound {
		return nil, fmt.Errorf("site group not found")
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}

	return &siteGroup, nil
}

// PutSiteGroup creates or updates a site group
func (service *Service) PutSiteGroup(input *UpdateSiteGroupInput) (*SiteGroup, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	path := service.BuildPath("dcim", "site-groups", fmt.Sprintf("%d", input.ID))

	var siteGroup SiteGroup
	resp, err := service.Client.R().
		SetBody(input).
		SetResult(&siteGroup).
		Put(path)

	if err != nil {
		return nil, fmt.Errorf("error updating site group: %w", err)
	}

	if resp.StatusCode() == http.StatusNotFound {
		return nil, fmt.Errorf("site group not found")
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}

	return &siteGroup, nil
}

// DeleteSiteGroup deletes a site group
func (service *Service) DeleteSiteGroup(id int) error {
	path := service.BuildPath("dcim", "site-groups", fmt.Sprintf("%d", id))

	resp, err := service.Client.R().
		Delete(path)

	if err != nil {
		return fmt.Errorf("error deleting site group: %w", err)
	}

	if resp.StatusCode() == http.StatusNotFound {
		return fmt.Errorf("site group not found")
	}

	if resp.StatusCode() != http.StatusNoContent {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}

	return nil
}
