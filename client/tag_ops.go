package client

import (
	"fmt"
	"net/http"

	"github.com/zeddD1abl0/go-netbox-client/models"
)

// ListTags lists all tags
func (c *Client) ListTags(input *ListTagsInput) ([]models.Tag, error) {
	path := c.BuildPath("extras", "tags")

	// Build query parameters
	params := map[string]string{}
	if input.Name != "" {
		params["name__ic"] = input.Name
	}
	if input.Slug != "" {
		params["slug__ic"] = input.Slug
	}
	if input.Color != "" {
		params["color__ic"] = input.Color
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
		return nil, fmt.Errorf("error listing tags: %w", err)
	}

	// Convert results to []Tag
	tags := make([]models.Tag, len(response.Results))
	for i, result := range response.Results {
		resultMap, ok := result.(map[string]any)
		if !ok {
			return nil, fmt.Errorf("unexpected result type at index %d", i)
		}

		// Create a new Tag
		var tag models.Tag
		err := convertMapToStruct(resultMap, &tag)
		if err != nil {
			return nil, fmt.Errorf("error converting map to struct at index %d: %w", i, err)
		}

		tags[i] = tag
	}

	return tags, nil
}

// GetTag retrieves a single tag by ID
func (c *Client) GetTag(id int) (*models.Tag, error) {
	path := c.BuildPath("extras", "tags", fmt.Sprintf("%d", id))

	var tag models.Tag
	resp, err := c.R().
		SetResult(&tag).
		Get(path)

	if err != nil {
		return nil, fmt.Errorf("error getting tag: %w", err)
	}

	if resp.StatusCode() == http.StatusNotFound {
		return nil, fmt.Errorf("tag not found")
	}

	return &tag, nil
}

// CreateTag creates a new tag
func (c *Client) CreateTag(input *CreateTagInput) (*models.Tag, error) {
	path := c.BuildPath("extras", "tags")

	var tag models.Tag
	resp, err := c.R().
		SetBody(input).
		SetResult(&tag).
		Post(path)

	if err != nil {
		return nil, fmt.Errorf("error creating tag: %w", err)
	}

	if resp.StatusCode() != http.StatusCreated {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}

	return &tag, nil
}

// UpdateTag updates an existing tag
func (c *Client) UpdateTag(input *UpdateTagInput) (*models.Tag, error) {
	path := c.BuildPath("extras", "tags", fmt.Sprintf("%d", input.ID))

	var tag models.Tag
	resp, err := c.R().
		SetBody(input).
		SetResult(&tag).
		Put(path)

	if err != nil {
		return nil, fmt.Errorf("error updating tag: %w", err)
	}

	if resp.StatusCode() == http.StatusNotFound {
		return nil, fmt.Errorf("tag not found")
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}

	return &tag, nil
}

// PatchTag patches an existing tag
func (c *Client) PatchTag(input *PatchTagInput) (*models.Tag, error) {
	path := c.BuildPath("extras", "tags", fmt.Sprintf("%d", input.ID))

	var tag models.Tag
	resp, err := c.R().
		SetBody(input).
		SetResult(&tag).
		Patch(path)

	if err != nil {
		return nil, fmt.Errorf("error patching tag: %w", err)
	}

	if resp.StatusCode() == http.StatusNotFound {
		return nil, fmt.Errorf("tag not found")
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}

	return &tag, nil
}

// DeleteTag deletes a tag
func (c *Client) DeleteTag(id int) error {
	path := c.BuildPath("extras", "tags", fmt.Sprintf("%d", id))

	resp, err := c.R().
		Delete(path)

	if err != nil {
		return fmt.Errorf("error deleting tag: %w", err)
	}

	if resp.StatusCode() == http.StatusNotFound {
		return fmt.Errorf("tag not found")
	}

	if resp.StatusCode() != http.StatusNoContent {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}

	return nil
}
