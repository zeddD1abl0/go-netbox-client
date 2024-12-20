package dcim

import (
	"github.com/zeddD1abl0/go-netbox-client/models"
)

// SiteGroup represents a Netbox site group
type SiteGroup struct {
	ID           int                `json:"id"`
	URL          string             `json:"url"`
	Name         string             `json:"name"`
	Slug         string             `json:"slug"`
	Parent       *SiteGroup         `json:"parent,omitempty"`
	Description  string             `json:"description,omitempty"`
	Tags         []models.TagCreate `json:"tags,omitempty"`
	CustomFields map[string]any     `json:"custom_fields,omitempty"`
	Created      string             `json:"created"`
	LastUpdated  string             `json:"last_updated"`
}

// ListSiteGroupsInput represents the input for listing site groups
type ListSiteGroupsInput struct {
	Name   string
	Parent string
	Tag    string
	Limit  int
	Offset int
}

// CreateSiteGroupInput represents the input for creating a site group
type CreateSiteGroupInput struct {
	Name         string             `json:"name"`
	Slug         string             `json:"slug"`
	Parent       int                `json:"parent,omitempty"`
	Description  string             `json:"description,omitempty"`
	Tags         []models.TagCreate `json:"tags,omitempty"`
	CustomFields map[string]any     `json:"custom_fields,omitempty"`
}

// Validate validates the CreateSiteGroupInput
func (input *CreateSiteGroupInput) Validate() error {
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

// UpdateSiteGroupInput represents the input for updating a site group
type UpdateSiteGroupInput struct {
	ID           int                `json:"-"`
	Name         string             `json:"name"`
	Slug         string             `json:"slug"`
	Parent       int                `json:"parent,omitempty"`
	Description  string             `json:"description,omitempty"`
	Tags         []models.TagCreate `json:"tags,omitempty"`
	CustomFields map[string]any     `json:"custom_fields,omitempty"`
}

// Validate validates the UpdateSiteGroupInput
func (input *UpdateSiteGroupInput) Validate() error {
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

// PatchSiteGroupInput represents the input for patching a site group
type PatchSiteGroupInput struct {
	ID           int                 `json:"-"`
	Name         *string             `json:"name,omitempty"`
	Slug         *string             `json:"slug,omitempty"`
	Parent       *int                `json:"parent,omitempty"`
	Description  *string             `json:"description,omitempty"`
	Tags         *[]models.TagCreate `json:"tags,omitempty"`
	CustomFields map[string]any      `json:"custom_fields,omitempty"`
}

// Validate validates the PatchSiteGroupInput
func (input *PatchSiteGroupInput) Validate() error {
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
