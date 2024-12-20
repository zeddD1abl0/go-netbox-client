package dcim

import (
	"github.com/zeddD1abl0/go-netbox-client/models"
)

// Region represents a Netbox region
type Region struct {
	ID           int            `json:"id"`
	URL          string         `json:"url"`
	Name         string         `json:"name"`
	Slug         string         `json:"slug"`
	Parent       *Region        `json:"parent,omitempty"`
	Description  string         `json:"description,omitempty"`
	Tags         []models.Tag   `json:"tags,omitempty"`
	CustomFields map[string]any `json:"custom_fields,omitempty"`
	Created      string         `json:"created"`
	LastUpdated  string         `json:"last_updated"`
	SiteCount    int            `json:"site_count"`
}

// ListRegionsInput represents the input for listing regions
type ListRegionsInput struct {
	Name   string
	Parent string
	Tag    string
	Limit  int
	Offset int
}

// CreateRegionInput represents the input for creating a region
type CreateRegionInput struct {
	Name         string         `json:"name"`
	Slug         string         `json:"slug"`
	Parent       int            `json:"parent,omitempty"`
	Description  string         `json:"description,omitempty"`
	Tags         []models.Tag   `json:"tags,omitempty"`
	CustomFields map[string]any `json:"custom_fields,omitempty"`
}

// Validate validates the CreateRegionInput
func (input *CreateRegionInput) Validate() error {
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

// UpdateRegionInput represents the input for updating a region
type UpdateRegionInput struct {
	ID           int            `json:"-"`
	Name         string         `json:"name"`
	Slug         string         `json:"slug"`
	Parent       int            `json:"parent,omitempty"`
	Description  string         `json:"description,omitempty"`
	Tags         []models.Tag   `json:"tags,omitempty"`
	CustomFields map[string]any `json:"custom_fields,omitempty"`
}

// Validate validates the UpdateRegionInput
func (input *UpdateRegionInput) Validate() error {
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

// PatchRegionInput represents the input for patching a region
type PatchRegionInput struct {
	ID           int            `json:"-"`
	Name         *string        `json:"name,omitempty"`
	Slug         *string        `json:"slug,omitempty"`
	Parent       *int           `json:"parent,omitempty"`
	Description  *string        `json:"description,omitempty"`
	Tags         *[]models.Tag  `json:"tags,omitempty"`
	CustomFields map[string]any `json:"custom_fields,omitempty"`
}

// Validate validates the PatchRegionInput
func (input *PatchRegionInput) Validate() error {
	if input.ID == 0 {
		return &models.ValidationError{
			Field:   "id",
			Message: "ID is required",
		}
	}

	if input.Slug != nil {
		if err := models.ValidateSlug(*input.Slug); err != nil {
			return err
		}
	}

	return nil
}
