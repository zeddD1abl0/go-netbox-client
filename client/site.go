package client

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/zeddD1abl0/go-netbox-client/models"
)

// Valid status values for sites
const (
	SiteStatusActive          = "active"
	SiteStatusPlanned         = "planned"
	SiteStatusStaging         = "staging"
	SiteStatusDecommissioning = "decommissioning"
	SiteStatusRetired         = "retired"
)

// Site represents a Netbox site
type Site struct {
	ID              int                `json:"id"`
	URL             string             `json:"url"`
	Name            string             `json:"name"`
	Slug            string             `json:"slug"`
	Status          *Status            `json:"status"`
	Region          *Region            `json:"region"`
	Description     string             `json:"description"`
	PhysicalAddress string             `json:"physical_address,omitempty"`
	ShippingAddress string             `json:"shipping_address,omitempty"`
	Latitude        *float64           `json:"latitude,omitempty"`
	Longitude       *float64           `json:"longitude,omitempty"`
	Comments        string             `json:"comments,omitempty"`
	Tags            []models.TagCreate `json:"tags,omitempty"`
	CustomFields    map[string]any     `json:"custom_fields,omitempty"`
	Created         string             `json:"created"`
	LastUpdated     string             `json:"last_updated"`
}

// CreateSiteInput represents the input for creating a site
type CreateSiteInput struct {
	Name            string             `json:"name"`
	Slug            string             `json:"slug"`
	Status          string             `json:"status,omitempty"`
	Region          int                `json:"region,omitempty"`
	Description     string             `json:"description,omitempty"`
	PhysicalAddress string             `json:"physical_address,omitempty"`
	ShippingAddress string             `json:"shipping_address,omitempty"`
	Latitude        *float64           `json:"latitude,omitempty"`
	Longitude       *float64           `json:"longitude,omitempty"`
	Comments        string             `json:"comments,omitempty"`
	Tags            []models.TagCreate `json:"tags,omitempty"`
	CustomFields    map[string]any     `json:"custom_fields,omitempty"`
}

// Validate validates the CreateSiteInput
func (input *CreateSiteInput) Validate() error {
	return validation.ValidateStruct(input,
		validation.Field(&input.Name, validation.Required),
		validation.Field(&input.Slug, validation.Required),
		validation.Field(&input.Status, validation.By(func(value any) error {
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
				return validation.NewError("status", "must be one of: active, planned, staging, decommissioning, retired")
			}
			return nil
		})),
		validation.Field(&input.Latitude, validation.By(func(value any) error {
			if input.Latitude != nil && (*input.Latitude < -90 || *input.Latitude > 90) {
				return validation.NewError("latitude", "must be between -90 and 90")
			}
			return nil
		})),
		validation.Field(&input.Longitude, validation.By(func(value any) error {
			if input.Longitude != nil && (*input.Longitude < -180 || *input.Longitude > 180) {
				return validation.NewError("longitude", "must be between -180 and 180")
			}
			return nil
		})),
	)
}

// UpdateSiteInput represents the input for updating a site
type UpdateSiteInput struct {
	ID              int                `json:"-"` // Used in URL, not in body
	Name            string             `json:"name"`
	Slug            string             `json:"slug"`
	Status          string             `json:"status,omitempty"`
	Region          int                `json:"region,omitempty"`
	Description     string             `json:"description,omitempty"`
	PhysicalAddress string             `json:"physical_address,omitempty"`
	ShippingAddress string             `json:"shipping_address,omitempty"`
	Latitude        *float64           `json:"latitude,omitempty"`
	Longitude       *float64           `json:"longitude,omitempty"`
	Comments        string             `json:"comments,omitempty"`
	Tags            []models.TagCreate `json:"tags,omitempty"`
	CustomFields    map[string]any     `json:"custom_fields,omitempty"`
}

// Validate validates the UpdateSiteInput
func (input *UpdateSiteInput) Validate() error {
	return validation.ValidateStruct(input,
		validation.Field(&input.Name, validation.Required),
		validation.Field(&input.Slug, validation.Required),
		validation.Field(&input.Status, validation.By(func(value any) error {
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
				return validation.NewError("status", "must be one of: active, planned, staging, decommissioning, retired")
			}
			return nil
		})),
		validation.Field(&input.Latitude, validation.By(func(value any) error {
			if input.Latitude != nil && (*input.Latitude < -90 || *input.Latitude > 90) {
				return validation.NewError("latitude", "must be between -90 and 90")
			}
			return nil
		})),
		validation.Field(&input.Longitude, validation.By(func(value any) error {
			if input.Longitude != nil && (*input.Longitude < -180 || *input.Longitude > 180) {
				return validation.NewError("longitude", "must be between -180 and 180")
			}
			return nil
		})),
	)
}

// PatchSiteInput represents the input for patching a site
type PatchSiteInput struct {
	ID              *int                   `json:"-"`
	Name            *string                `json:"name,omitempty"`
	Slug            *string                `json:"slug,omitempty"`
	Status          *string                `json:"status,omitempty"`
	Region          *int                   `json:"region,omitempty"`
	Group           *int                   `json:"group,omitempty"`
	Tenant          *int                   `json:"tenant,omitempty"`
	Facility        *string                `json:"facility,omitempty"`
	TimeZone        *string                `json:"time_zone,omitempty"`
	Description     *string                `json:"description,omitempty"`
	PhysicalAddress *string                `json:"physical_address,omitempty"`
	ShippingAddress *string                `json:"shipping_address,omitempty"`
	Latitude        *float64               `json:"latitude,omitempty"`
	Longitude       *float64               `json:"longitude,omitempty"`
	Comments        *string                `json:"comments,omitempty"`
	AsnsIDs         *[]int                 `json:"asns,omitempty"`
	Tags            *[]int                 `json:"tags,omitempty"`
	CustomFields    map[string]interface{} `json:"custom_fields,omitempty"`
}

// Validate validates the PatchSiteInput
func (input *PatchSiteInput) Validate() error {
	if input.ID == nil {
		return validation.NewError("id", "is required")
	}

	if input.Slug != nil {
		if err := validation.Validate(*input.Slug, validation.Match(regexp.MustCompile(`^[a-zA-Z0-9_-]+$`))); err != nil {
			return validation.NewError("slug", "must be a valid slug")
		}
	}

	if input.Status != nil {
		validStatuses := []string{
			SiteStatusActive,
			SiteStatusPlanned,
			SiteStatusStaging,
			SiteStatusDecommissioning,
			SiteStatusRetired,
		}
		isValid := false
		for _, status := range validStatuses {
			if *input.Status == status {
				isValid = true
				break
			}
		}
		if !isValid {
			return validation.NewError("status", "must be one of: active, planned, staging, decommissioning, retired")
		}
	}

	if input.Latitude != nil && (*input.Latitude < -90 || *input.Latitude > 90) {
		return validation.NewError("latitude", "must be between -90 and 90")
	}

	if input.Longitude != nil && (*input.Longitude < -180 || *input.Longitude > 180) {
		return validation.NewError("longitude", "must be between -180 and 180")
	}

	return nil
}
