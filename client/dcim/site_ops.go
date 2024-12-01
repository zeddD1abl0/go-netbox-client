package dcim

import (
	"fmt"
	"net/http"
)

// ListSites lists all sites matching the input criteria
func (service *Service) ListSites(input *ListSitesInput) ([]Site, error) {
	path := service.buildPath("dcim", "sites")
	
	// Build query parameters
	params := map[string]string{}
	if input.Name != "" {
		params["name"] = input.Name
	}
	if input.Status != "" {
		params["status"] = input.Status
	}
	if input.Region != "" {
		params["region"] = input.Region
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
		Count    int    `json:"count"`
		Next     *string `json:"next"`
		Previous *string `json:"previous"`
		Results  []Site  `json:"results"`
	}

	resp, err := service.client.httpClient.R().
		SetQueryParams(params).
		SetResult(&response).
		Get(path)
	
	if err != nil {
		return nil, fmt.Errorf("error listing sites: %w", err)
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}

	return response.Results, nil
}

// GetSite retrieves a single site by ID
func (service *Service) GetSite(id int) (*Site, error) {
	path := service.buildPath("dcim", "sites", fmt.Sprintf("%d", id))
	
	var site Site
	resp, err := service.client.httpClient.R().
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

	path := service.buildPath("dcim", "sites")
	
	var site Site
	resp, err := service.client.httpClient.R().
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

	path := service.buildPath("dcim", "sites", fmt.Sprintf("%d", input.ID))
	
	var site Site
	resp, err := service.client.httpClient.R().
		SetBody(input).
		SetResult(&site).
		Put(path)
	
	if err != nil {
		return nil, fmt.Errorf("error updating site: %w", err)
	}

	if resp.StatusCode() == http.StatusNotFound {
		return nil, fmt.Errorf("site not found")
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}

	return &site, nil
}

// PatchSite partially updates an existing site
func (service *Service) PatchSite(input *PatchSiteInput) (*Site, error) {
	if err := input.Validate(); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	path := service.buildPath("dcim", "sites", fmt.Sprintf("%d", input.ID))
	
	var site Site
	resp, err := service.client.httpClient.R().
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

// PutSite creates or updates a site
func (service *Service) PutSite(input *UpdateSiteInput) (*Site, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	req, err := service.client.NewRequest("PUT", fmt.Sprintf("dcim/sites/%d/", input.ID), nil)
	if err != nil {
		return nil, err
	}

	req.Body = input
	site := new(Site)
	err = service.client.Do(req, site)
	if err != nil {
		return nil, err
	}

	return site, nil
}

// DeleteSite deletes a site
func (service *Service) DeleteSite(id int) error {
	path := service.buildPath("dcim", "sites", fmt.Sprintf("%d", id))
	
	resp, err := service.client.httpClient.R().
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
