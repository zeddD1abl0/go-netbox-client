package dcim

import (
	"fmt"
	"net/http"
)

// ListSiteGroups lists all site groups matching the input criteria
func (service *Service) ListSiteGroups(input *ListSiteGroupsInput) ([]SiteGroup, error) {
	path := service.buildPath("dcim", "site-groups")
	
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
		Count    int         `json:"count"`
		Next     *string     `json:"next"`
		Previous *string     `json:"previous"`
		Results  []SiteGroup `json:"results"`
	}

	resp, err := service.client.httpClient.R().
		SetQueryParams(params).
		SetResult(&response).
		Get(path)
	
	if err != nil {
		return nil, fmt.Errorf("error listing site groups: %w", err)
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}

	return response.Results, nil
}

// GetSiteGroup retrieves a single site group by ID
func (service *Service) GetSiteGroup(id int) (*SiteGroup, error) {
	path := service.buildPath("dcim", "site-groups", fmt.Sprintf("%d", id))
	
	var siteGroup SiteGroup
	resp, err := service.client.httpClient.R().
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

	path := service.buildPath("dcim", "site-groups")
	
	var siteGroup SiteGroup
	resp, err := service.client.httpClient.R().
		SetBody(input).
		SetResult(&siteGroup).
		Post(path)
	
	if err != nil {
		return nil, fmt.Errorf("error creating site group: %w", err)
	}

	if resp.StatusCode() != http.StatusCreated {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}

	return &siteGroup, nil
}

// UpdateSiteGroup updates an existing site group
func (service *Service) UpdateSiteGroup(input *UpdateSiteGroupInput) (*SiteGroup, error) {
	if err := input.Validate(); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	path := service.buildPath("dcim", "site-groups", fmt.Sprintf("%d", input.ID))
	
	var siteGroup SiteGroup
	resp, err := service.client.httpClient.R().
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

	path := service.buildPath("dcim", "site-groups", fmt.Sprintf("%d", input.ID))
	
	var siteGroup SiteGroup
	resp, err := service.client.httpClient.R().
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

	req, err := service.client.NewRequest("PUT", fmt.Sprintf("dcim/site-groups/%d/", input.ID), nil)
	if err != nil {
		return nil, err
	}

	req.Body = input
	siteGroup := new(SiteGroup)
	err = service.client.Do(req, siteGroup)
	if err != nil {
		return nil, err
	}

	return siteGroup, nil
}

// DeleteSiteGroup deletes a site group
func (service *Service) DeleteSiteGroup(id int) error {
	path := service.buildPath("dcim", "site-groups", fmt.Sprintf("%d", id))
	
	resp, err := service.client.httpClient.R().
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
