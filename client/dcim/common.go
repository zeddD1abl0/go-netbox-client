package dcim

import (
	"github.com/zeddD1abl0/go-netbox-client/models"
)

// Status represents a status value and label
type Status struct {
	Value string `json:"value"`
	Label string `json:"label"`
}

// Region represents a region with its basic properties
// type Region struct {
// 	ID   int    `json:"id"`
// 	Name string `json:"name"`
// 	Slug string `json:"slug"`
// }

// BaseResource contains common fields shared across DCIM resources
type BaseResource struct {
	ID           int            `json:"id"`
	URL          string         `json:"url"`
	Name         string         `json:"name"`
	Slug         string         `json:"slug"`
	Description  string         `json:"description,omitempty"`
	Tags         []models.Tag   `json:"tags,omitempty"`
	CustomFields map[string]any `json:"custom_fields,omitempty"`
	Created      string         `json:"created"`
	LastUpdated  string         `json:"last_updated"`
}

// BaseListInput contains common fields for list operations
type BaseListInput struct {
	Name   string
	Tag    string
	Limit  int
	Offset int
}

// ListSitesInput represents the input for listing sites
type ListSitesInput struct {
	BaseListInput
	Status string
	Region string
}

// BaseCreateInput contains common fields for create operations
type BaseCreateInput struct {
	Name         string         `json:"name"`
	Slug         string         `json:"slug"`
	Description  string         `json:"description,omitempty"`
	Tags         []models.Tag   `json:"tags,omitempty"`
	CustomFields map[string]any `json:"custom_fields,omitempty"`
}

// BaseUpdateInput contains common fields for update operations
type BaseUpdateInput struct {
	ID           int            `json:"-"`
	Name         string         `json:"name"`
	Slug         string         `json:"slug"`
	Description  string         `json:"description,omitempty"`
	Tags         []models.Tag   `json:"tags,omitempty"`
	CustomFields map[string]any `json:"custom_fields,omitempty"`
}

// BasePatchInput contains common fields for patch operations
type BasePatchInput struct {
	ID           int            `json:"-"`
	Name         *string        `json:"name,omitempty"`
	Slug         *string        `json:"slug,omitempty"`
	Description  *string        `json:"description,omitempty"`
	Tags         *[]models.Tag  `json:"tags,omitempty"`
	CustomFields map[string]any `json:"custom_fields,omitempty"`
}
