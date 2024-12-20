package models

// PaginatedResponse represents the standard Netbox paginated response
type PaginatedResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []any   `json:"results"`
}

// Tag represents a Netbox tag
type Tag struct {
	ID          int    `json:"id"`
	URL         string `json:"url"`
	Display     string `json:"display"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Color       string `json:"color"`
	Description string `json:"description"`
}

// Tags being created on the fly have no description
type TagCreate struct {
	Name  string `json:"name"`
	Slug  string `json:"slug"`
	Color string `json:"color"`
}

// CustomField represents a custom field in Netbox
type CustomField struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Value any    `json:"value"`
	Type  string `json:"type"`
}
