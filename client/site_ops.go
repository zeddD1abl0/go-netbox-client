package client

import (
	"fmt"
	"net/http"
)

// ListSites lists all sites matching the input criteria
func (c *Client) ListSites(input *ListSitesInput) ([]Site, error) {
	path := c.BuildPath("dcim", "sites")

	// Build query parameters
	params := map[string]string{}
	if input.Name != "" {
		params["name__ic"] = input.Name
	}
	if input.Region != "" {
		params["region"] = input.Region
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
	var response Response
	response.Results = make([]any, 0)
	_, err := c.R().
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
		err := convertMapToStruct(resultMap, &site)
		if err != nil {
			return nil, fmt.Errorf("error converting map to struct at index %d: %w", i, err)
		}

		sites[i] = site
	}

	return sites, nil
}

// GetSite retrieves a single site by ID
func (c *Client) GetSite(id int) (*Site, error) {
	path := c.BuildPath("dcim", "sites", fmt.Sprintf("%d", id))

	var site Site
	resp, err := c.R().
		SetResult(&site).
		Get(path)

	if err != nil {
		return nil, fmt.Errorf("error getting site: %w", err)
	}

	if resp.StatusCode() == http.StatusNotFound {
		return nil, fmt.Errorf("site not found")
	}

	return &site, nil
}

// CreateSite creates a new site
func (c *Client) CreateSite(input *CreateSiteInput) (*Site, error) {
	path := c.BuildPath("dcim", "sites")

	var site Site
	resp, err := c.R().
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
func (c *Client) UpdateSite(input *UpdateSiteInput) (*Site, error) {
	path := c.BuildPath("dcim", "sites", fmt.Sprintf("%d", input.ID))

	var site Site
	resp, err := c.R().
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

// PatchSite patches an existing site
func (c *Client) PatchSite(input *PatchSiteInput) (*Site, error) {
	if input.ID == nil {
		return nil, fmt.Errorf("site ID is required")
	}

	path := c.BuildPath("dcim", "sites", fmt.Sprintf("%d", *input.ID))

	var site Site
	resp, err := c.R().
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
func (c *Client) DeleteSite(id int) error {
	path := c.BuildPath("dcim", "sites", fmt.Sprintf("%d", id))

	resp, err := c.R().
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
