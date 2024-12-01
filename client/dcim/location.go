package dcim

import (
	"github.com/jordan/go-netbox-client/models"
)

// Location represents a Netbox location
type Location struct {
	ID           int              `json:"id"`
	URL          string           `json:"url"`
	Name         string           `json:"name"`
	Slug         string           `json:"slug"`
	Site         Site             `json:"site"`
	Parent       *Location        `json:"parent,omitempty"`
	Description  string           `json:"description,omitempty"`
	Tags         []models.Tag     `json:"tags,omitempty"`
	CustomFields map[string]any   `json:"custom_fields,omitempty"`
	Created      string           `json:"created"`
	LastUpdated  string           `json:"last_updated"`
	RackCount    int             `json:"rack_count"`
	DeviceCount  int             `json:"device_count"`
}

// ListLocationsInput represents the input for listing locations
type ListLocationsInput struct {
	Name     string
	Site     string
	Parent   string
	Tag      string
	Limit    int
	Offset   int
}

// CreateLocationInput represents the input for creating a location
type CreateLocationInput struct {
	Name         string            `json:"name"`
	Slug         string            `json:"slug"`
	Site         int              `json:"site"`
	Parent       int              `json:"parent,omitempty"`
	Description  string           `json:"description,omitempty"`
	Tags         []models.Tag     `json:"tags,omitempty"`
	CustomFields map[string]any   `json:"custom_fields,omitempty"`
}

// Validate validates the CreateLocationInput
func (input *CreateLocationInput) Validate() error {
	var errors models.ValidationErrors
	
	if err := models.ValidateRequired("name", input.Name); err != nil {
		errors = append(errors, *err.(*models.ValidationError))
	}
	
	if err := models.ValidateSlug(input.Slug); err != nil {
		errors = append(errors, *err.(*models.ValidationError))
	}

	if input.Site == 0 {
		errors = append(errors, models.ValidationError{
			Name:    "site",
			Message: "Site is required",
		})
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

// UpdateLocationInput represents the input for updating a location
type UpdateLocationInput struct {
	ID           int               `json:"-"`
	Name         string            `json:"name"`
	Slug         string            `json:"slug"`
	Site         int              `json:"site"`
	Parent       int              `json:"parent,omitempty"`
	Description  string           `json:"description,omitempty"`
	Tags         []models.Tag     `json:"tags,omitempty"`
	CustomFields map[string]any   `json:"custom_fields,omitempty"`
}

// Validate validates the UpdateLocationInput
func (input *UpdateLocationInput) Validate() error {
	var errors models.ValidationErrors
	
	if err := models.ValidateRequired("name", input.Name); err != nil {
		errors = append(errors, *err.(*models.ValidationError))
	}
	
	if err := models.ValidateSlug(input.Slug); err != nil {
		errors = append(errors, *err.(*models.ValidationError))
	}

	if input.Site == 0 {
		errors = append(errors, models.ValidationError{
			Name:    "site",
			Message: "Site is required",
		})
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

// PatchLocationInput represents the input for patching a location
type PatchLocationInput struct {
	ID           int               `json:"-"`
	Name         *string           `json:"name,omitempty"`
	Slug         *string           `json:"slug,omitempty"`
	Site         *int             `json:"site,omitempty"`
	Parent       *int             `json:"parent,omitempty"`
	Description  *string          `json:"description,omitempty"`
	Tags         *[]models.Tag    `json:"tags,omitempty"`
	CustomFields map[string]any   `json:"custom_fields,omitempty"`
}

// Validate validates the PatchLocationInput
func (input *PatchLocationInput) Validate() error {
	if input.ID == 0 {
		return &models.ValidationError{
			Name:    "id",
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
