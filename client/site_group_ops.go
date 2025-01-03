package client

import (
	"fmt"
	"net/http"
)

// ListSiteGroups lists all site groups
func (c *Client) ListSiteGroups(input *ListSiteGroupsInput) ([]SiteGroup, error) {
	path := c.BuildPath("dcim", "site-groups")

	// Build query parameters
	params := map[string]string{}
	if input.Name != "" {
		params["name__ic"] = input.Name
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
		err := convertMapToStruct(resultMap, &siteGroup)
		if err != nil {
			return nil, fmt.Errorf("error converting map to struct at index %d: %w", i, err)
		}

		siteGroups[i] = siteGroup
	}

	return siteGroups, nil
}

// GetSiteGroup retrieves a single site group by ID
func (c *Client) GetSiteGroup(id int) (*SiteGroup, error) {
	path := c.BuildPath("dcim", "site-groups", fmt.Sprintf("%d", id))

	var siteGroup SiteGroup
	resp, err := c.R().
		SetResult(&siteGroup).
		Get(path)

	if err != nil {
		return nil, fmt.Errorf("error getting site group: %w", err)
	}

	if resp.StatusCode() == http.StatusNotFound {
		return nil, fmt.Errorf("site group not found")
	}

	return &siteGroup, nil
}

// CreateSiteGroup creates a new site group
func (c *Client) CreateSiteGroup(input *CreateSiteGroupInput) (*SiteGroup, error) {
	path := c.BuildPath("dcim", "site-groups")

	var siteGroup SiteGroup
	resp, err := c.R().
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
func (c *Client) UpdateSiteGroup(input *UpdateSiteGroupInput) (*SiteGroup, error) {
	path := c.BuildPath("dcim", "site-groups", fmt.Sprintf("%d", input.ID))

	var siteGroup SiteGroup
	resp, err := c.R().
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

// PatchSiteGroup patches an existing site group
func (c *Client) PatchSiteGroup(input *PatchSiteGroupInput) (*SiteGroup, error) {
	path := c.BuildPath("dcim", "site-groups", fmt.Sprintf("%d", input.ID))

	var siteGroup SiteGroup
	resp, err := c.R().
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

// DeleteSiteGroup deletes a site group
func (c *Client) DeleteSiteGroup(id int) error {
	path := c.BuildPath("dcim", "site-groups", fmt.Sprintf("%d", id))

	resp, err := c.R().
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
