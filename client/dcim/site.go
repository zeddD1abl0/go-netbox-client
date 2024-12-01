package dcim

import (
	"fmt"
	"net/http"
	"github.com/jordan/go-netbox-client/models"
)

// Valid status values for sites
const (
	SiteStatusActive    = "active"
	SiteStatusPlanned   = "planned"
	SiteStatusStaging   = "staging"
	SiteStatusDecommissioning = "decommissioning"
	SiteStatusRetired   = "retired"
)

// Site represents a Netbox site
type Site struct {
	ID           int              `json:"id"`
	URL          string           `json:"url"`
	Name         string           `json:"name"`
	Slug         string           `json:"slug"`
	Status       *Status          `json:"status"`
	Region       *Region          `json:"region"`
	Description  string           `json:"description"`
	PhysicalAddress string        `json:"physical_address,omitempty"`
	ShippingAddress string        `json:"shipping_address,omitempty"`
	Latitude     *float64         `json:"latitude,omitempty"`
	Longitude    *float64         `json:"longitude,omitempty"`
	Comments     string           `json:"comments,omitempty"`
	Tags         []models.Tag     `json:"tags,omitempty"`
	CustomFields map[string]any   `json:"custom_fields,omitempty"`
	Created      string           `json:"created"`
	LastUpdated  string           `json:"last_updated"`
}

// Status represents a site's status
type Status struct {
	Value string `json:"value"`
	Label string `json:"label"`
}

// Region represents a site's region
type Region struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

// ListSitesInput represents the input for listing sites
type ListSitesInput struct {
	Name     string
	Status   string
	Region   string
	Tag      string
	Limit    int
	Offset   int
}

// CreateSiteInput represents the input for creating a site
type CreateSiteInput struct {
	Name            string            `json:"name"`
	Slug            string            `json:"slug"`
	Status          string            `json:"status,omitempty"`
	Region          int              `json:"region,omitempty"`
	Description     string           `json:"description,omitempty"`
	PhysicalAddress string           `json:"physical_address,omitempty"`
	ShippingAddress string           `json:"shipping_address,omitempty"`
	Latitude        *float64         `json:"latitude,omitempty"`
	Longitude       *float64         `json:"longitude,omitempty"`
	Comments        string           `json:"comments,omitempty"`
	Tags            []models.Tag     `json:"tags,omitempty"`
	CustomFields    map[string]any   `json:"custom_fields,omitempty"`
}

// Validate validates the CreateSiteInput
func (input *CreateSiteInput) Validate() error {
	var errors models.ValidationErrors
	
	if err := models.ValidateRequired("name", input.Name); err != nil {
		errors = append(errors, *err.(*models.ValidationError))
	}
	
	if err := models.ValidateSlug(input.Slug); err != nil {
		errors = append(errors, *err.(*models.ValidationError))
	}

	if input.Status != "" {
		validStatuses := []string{
			SiteStatusActive,
			SiteStatusPlanned,
			SiteStatusStaging,
			SiteStatusDecommissioning,
			SiteStatusRetired,
		}
		isValid := false
		for _, status := range validStatuses {
			if input.Status == status {
				isValid = true
				break
			}
		}
		if !isValid {
			errors = append(errors, models.ValidationError{
				Field:   "status",
				Message: fmt.Sprintf("must be one of: %v", validStatuses),
			})
		}
	}

	if input.Latitude != nil && (*input.Latitude < -90 || *input.Latitude > 90) {
		errors = append(errors, models.ValidationError{
			Field:   "latitude",
			Message: "must be between -90 and 90",
		})
	}

	if input.Longitude != nil && (*input.Longitude < -180 || *input.Longitude > 180) {
		errors = append(errors, models.ValidationError{
			Field:   "longitude",
			Message: "must be between -180 and 180",
		})
	}

	if len(errors) > 0 {
		return errors
	}
	return nil
}

// UpdateSiteInput represents the input for updating a site
type UpdateSiteInput struct {
	ID              int               `json:"-"` // Used in URL, not in body
	Name            string            `json:"name"`
	Slug            string            `json:"slug"`
	Status          string            `json:"status,omitempty"`
	Region          int              `json:"region,omitempty"`
	Description     string           `json:"description,omitempty"`
	PhysicalAddress string           `json:"physical_address,omitempty"`
	ShippingAddress string           `json:"shipping_address,omitempty"`
	Latitude        *float64         `json:"latitude,omitempty"`
	Longitude       *float64         `json:"longitude,omitempty"`
	Comments        string           `json:"comments,omitempty"`
	Tags            []models.Tag     `json:"tags,omitempty"`
	CustomFields    map[string]any   `json:"custom_fields,omitempty"`
}

// Validate validates the UpdateSiteInput
func (input *UpdateSiteInput) Validate() error {
	var errors models.ValidationErrors
	
	if err := models.ValidateRequired("name", input.Name); err != nil {
		errors = append(errors, *err.(*models.ValidationError))
	}
	
	if err := models.ValidateSlug(input.Slug); err != nil {
		errors = append(errors, *err.(*models.ValidationError))
	}

	if input.Status != "" {
		validStatuses := []string{
			SiteStatusActive,
			SiteStatusPlanned,
			SiteStatusStaging,
			SiteStatusDecommissioning,
			SiteStatusRetired,
		}
		isValid := false
		for _, status := range validStatuses {
			if input.Status == status {
				isValid = true
				break
			}
		}
		if !isValid {
			errors = append(errors, models.ValidationError{
				Field:   "status",
				Message: fmt.Sprintf("must be one of: %v", validStatuses),
			})
		}
	}

	if input.Latitude != nil && (*input.Latitude < -90 || *input.Latitude > 90) {
		errors = append(errors, models.ValidationError{
			Field:   "latitude",
			Message: "must be between -90 and 90",
		})
	}

	if input.Longitude != nil && (*input.Longitude < -180 || *input.Longitude > 180) {
		errors = append(errors, models.ValidationError{
			Field:   "longitude",
			Message: "must be between -180 and 180",
		})
	}

	if len(errors) > 0 {
		return errors
	}
	return nil
}

// PatchSiteInput represents the input for patching a site
type PatchSiteInput struct {
	ID                *int     `json:"-"`
	Name              *string  `json:"name,omitempty"`
	Slug              *string  `json:"slug,omitempty"`
	Status            *string  `json:"status,omitempty"`
	Region            *int     `json:"region,omitempty"`
	Group             *int     `json:"group,omitempty"`
	Tenant            *int     `json:"tenant,omitempty"`
	Facility          *string  `json:"facility,omitempty"`
	TimeZone          *string  `json:"time_zone,omitempty"`
	Description       *string  `json:"description,omitempty"`
	PhysicalAddress   *string  `json:"physical_address,omitempty"`
	ShippingAddress   *string  `json:"shipping_address,omitempty"`
	Latitude          *float64 `json:"latitude,omitempty"`
	Longitude         *float64 `json:"longitude,omitempty"`
	Comments          *string  `json:"comments,omitempty"`
	AsnsIDs           *[]int   `json:"asns,omitempty"`
	Tags             *[]int   `json:"tags,omitempty"`
	CustomFields     map[string]interface{} `json:"custom_fields,omitempty"`
}

// Validate validates the PatchSiteInput
func (input *PatchSiteInput) Validate() error {
	if input.ID == nil {
		return validation.NewValidationError("id is required")
	}

	if input.Slug != nil {
		if err := validation.ValidateSlug(*input.Slug); err != nil {
			return validation.WrapValidationError(err, "slug")
		}
	}

	if input.Status != nil {
		if err := validation.ValidateStatus(*input.Status); err != nil {
			return validation.WrapValidationError(err, "status")
		}
	}

	return nil
}
