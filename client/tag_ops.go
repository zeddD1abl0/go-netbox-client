package client

import (
	"fmt"

	"github.com/zeddD1abl0/go-netbox-client/models"
)

// ListTags lists all tags
func (service *Service) ListTags(input *ListTagsInput) ([]models.Tag, error) {
	path := service.BuildPath("extras", "tags")

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
	_, err := service.Client.R().
		SetQueryParams(params).
		SetResult(&response).
		Get(path)

	if err != nil {
		return nil, fmt.Errorf("error listing tags: %w", err)
	}

	// Convert response to tags
	tags := make([]models.Tag, len(response.Results))
	for i, result := range response.Results {
		resultMap := result.(map[string]any)
		var tag models.Tag

		if id, ok := resultMap["id"].(float64); ok {
			tag.ID = int(id)
		}
		if name, ok := resultMap["name"].(string); ok {
			tag.Name = name
		}
		if slug, ok := resultMap["slug"].(string); ok {
			tag.Slug = slug
		}
		if color, ok := resultMap["color"].(string); ok {
			tag.Color = color
		}
		if description, ok := resultMap["description"].(string); ok {
			tag.Description = description
		}
		if objectTypes, ok := resultMap["object_types"].([]any); ok {
			tag.ObjectTypes = make([]string, len(objectTypes))
			for j, ot := range objectTypes {
				if otStr, ok := ot.(string); ok {
					tag.ObjectTypes[j] = otStr
				}
			}
		}
		tags[i] = tag
	}

	return tags, nil
}

// GetTag retrieves a single tag by ID
func (service *Service) GetTag(id int) (*models.Tag, error) {
	path := service.BuildPath("extras", "tags", fmt.Sprintf("%d", id))

	var tag models.Tag
	resp, err := service.Client.R().
		SetResult(&tag).
		Get(path)

	if err != nil {
		return nil, fmt.Errorf("error getting tag: %w", err)
	}

	if resp.StatusCode() == 404 {
		return nil, fmt.Errorf("tag not found")
	}

	return &tag, nil
}

// CreateTag creates a new tag
func (service *Service) CreateTag(input *CreateTagInput) (*models.Tag, error) {
	if err := input.Validate(); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	path := service.BuildPath("extras", "tags")

	var tag models.Tag
	resp, err := service.Client.R().
		SetBody(input).
		SetResult(&tag).
		Post(path)

	if err != nil {
		return nil, fmt.Errorf("error creating tag: %w", err)
	}

	if resp.StatusCode() >= 400 {
		return nil, fmt.Errorf("error creating tag: %s", resp.String())
	}

	return &tag, nil
}

// UpdateTag updates an existing tag
func (service *Service) UpdateTag(input *UpdateTagInput) (*models.Tag, error) {
	if err := input.Validate(); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	path := service.BuildPath("extras", "tags", fmt.Sprintf("%d", input.ID))

	var tag models.Tag
	resp, err := service.Client.R().
		SetBody(input).
		SetResult(&tag).
		Put(path)

	if err != nil {
		return nil, fmt.Errorf("error updating tag: %w", err)
	}

	if resp.StatusCode() == 404 {
		return nil, fmt.Errorf("tag not found")
	}

	return &tag, nil
}

// DeleteTag deletes a tag
func (service *Service) DeleteTag(id int) error {
	path := service.BuildPath("extras", "tags", fmt.Sprintf("%d", id))

	resp, err := service.Client.R().Delete(path)
	if err != nil {
		return fmt.Errorf("error deleting tag: %w", err)
	}

	if resp.StatusCode() == 404 {
		return fmt.Errorf("tag not found")
	}

	if resp.StatusCode() >= 400 {
		return fmt.Errorf("error deleting tag: %s", resp.String())
	}

	return nil
}

// PatchTag patches a tag
func (service *Service) PatchTag(input *PatchTagInput) (*models.Tag, error) {
	if err := input.Validate(); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	path := service.BuildPath("extras", "tags", fmt.Sprintf("%d", input.ID))

	var tag models.Tag
	resp, err := service.Client.R().
		SetBody(input).
		SetResult(&tag).
		Patch(path)

	if err != nil {
		return nil, fmt.Errorf("error patching tag: %w", err)
	}

	if resp.StatusCode() == 404 {
		return nil, fmt.Errorf("tag not found")
	}

	return &tag, nil
}
