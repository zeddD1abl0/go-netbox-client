package extras

import (
	"github.com/zeddD1abl0/go-netbox-client/models"
)

// ListTagsInput represents the input for listing tags
type ListTagsInput struct {
	Name   string
	Slug   string
	Color  string
	Limit  int
	Offset int
}

// CreateTagInput represents the input for creating a tag
type CreateTagInput struct {
	Name        string   `json:"name"`
	Slug        string   `json:"slug"`
	Color       string   `json:"color"`
	Description string   `json:"description,omitempty"`
	ObjectTypes []string `json:"object_types,omitempty"`
}

// Validate validates the CreateTagInput
func (input *CreateTagInput) Validate() error {
	var errors models.ValidationErrors

	if err := models.ValidateRequired("name", input.Name); err != nil {
		errors = append(errors, *err.(*models.ValidationError))
	}

	if err := models.ValidateSlug(input.Slug); err != nil {
		errors = append(errors, *err.(*models.ValidationError))
	}

	if len(errors) > 0 {
		return errors
	}
	return nil
}

// UpdateTagInput represents the input for updating a tag
type UpdateTagInput struct {
	ID          int      `json:"-"`
	Name        string   `json:"name"`
	Slug        string   `json:"slug"`
	Color       string   `json:"color"`
	Description string   `json:"description,omitempty"`
	ObjectTypes []string `json:"object_types,omitempty"`
}

// Validate validates the UpdateTagInput
func (input *UpdateTagInput) Validate() error {
	var errors models.ValidationErrors

	if err := models.ValidateRequired("name", input.Name); err != nil {
		errors = append(errors, *err.(*models.ValidationError))
	}

	if err := models.ValidateSlug(input.Slug); err != nil {
		errors = append(errors, *err.(*models.ValidationError))
	}

	if len(errors) > 0 {
		return errors
	}
	return nil
}

// PatchTagInput represents the input for patching a tag
type PatchTagInput struct {
	ID          int       `json:"-"`
	Name        *string   `json:"name,omitempty"`
	Slug        *string   `json:"slug,omitempty"`
	Color       *string   `json:"color,omitempty"`
	Description *string   `json:"description,omitempty"`
	ObjectTypes *[]string `json:"object_types,omitempty"`
}

// Validate validates the PatchTagInput
func (input *PatchTagInput) Validate() error {
	var errors models.ValidationErrors

	if input.Slug != nil {
		if err := models.ValidateSlug(*input.Slug); err != nil {
			errors = append(errors, *err.(*models.ValidationError))
		}
	}

	if len(errors) > 0 {
		return errors
	}
	return nil
}
