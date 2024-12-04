package dcim

import (
	"fmt"
	"github.com/zeddD1abl0/go-netbox-client/client"
	"github.com/zeddD1abl0/go-netbox-client/models"
	"net/http"
)

// Service handles DCIM endpoints
type Service struct {
	*client.Service
}

// ListSitesInput represents the input for listing sites
type ListSitesInput struct {
	Name     string
	Status   string
	Region   string
	Limit    int
	Offset   int
}

// Site represents a Netbox site
type Site struct {
	ID          int              `json:"id"`
	URL         string           `json:"url"`
	Name        string           `json:"name"`
	Slug        string           `json:"slug"`
	Status      *Status         `json:"status"`
	Region      *Region         `json:"region"`
	Tags        []models.Tag    `json:"tags"`
	CustomFields map[string]any `json:"custom_fields"`
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

// CreateSiteInput represents the input for creating a site
type CreateSiteInput struct {
	Name         string            `json:"name"`
	Slug         string            `json:"slug"`
	Status       string            `json:"status,omitempty"`
	Region       int              `json:"region,omitempty"`
	Description  string           `json:"description,omitempty"`
	Tags         []models.Tag     `json:"tags,omitempty"`
	CustomFields map[string]any   `json:"custom_fields,omitempty"`
}

// Validate validates the CreateSiteInput
func (i *CreateSiteInput) Validate() error {
	var errors models.ValidationErrors
	
	if err := models.ValidateRequired("name", i.Name); err != nil {
		errors = append(errors, *err.(*models.ValidationError))
	}
	
	if err := models.ValidateSlug(i.Slug); err != nil {
		errors = append(errors, *err.(*models.ValidationError))
	}

	if len(errors) > 0 {
		return errors
	}
	return nil
}

// UpdateSiteInput represents the input for updating a site
type UpdateSiteInput struct {
	ID           int               `json:"-"` // Used in URL, not in body
	Name         string            `json:"name"`
	Slug         string            `json:"slug"`
	Status       string            `json:"status,omitempty"`
	Region       int              `json:"region,omitempty"`
	Description  string           `json:"description,omitempty"`
	Tags         []models.Tag     `json:"tags,omitempty"`
	CustomFields map[string]any   `json:"custom_fields,omitempty"`
}

// Validate validates the UpdateSiteInput
func (i *UpdateSiteInput) Validate() error {
	var errors models.ValidationErrors
	
	if err := models.ValidateRequired("name", i.Name); err != nil {
		errors = append(errors, *err.(*models.ValidationError))
	}
	
	if err := models.ValidateSlug(i.Slug); err != nil {
		errors = append(errors, *err.(*models.ValidationError))
	}

	if len(errors) > 0 {
		return errors
	}
	return nil
}

// PatchSiteInput represents the input for patching a site
type PatchSiteInput struct {
	ID           int               `json:"-"` // Used in URL, not in body
	Name         *string           `json:"name,omitempty"`
	Slug         *string           `json:"slug,omitempty"`
	Status       *string           `json:"status,omitempty"`
	Region       *int             `json:"region,omitempty"`
	Description  *string          `json:"description,omitempty"`
	Tags         *[]models.Tag    `json:"tags,omitempty"`
	CustomFields map[string]any   `json:"custom_fields,omitempty"`
}

// Validate validates the PatchSiteInput
func (i *PatchSiteInput) Validate() error {
	var errors models.ValidationErrors
	
	if i.Slug != nil {
		if err := models.ValidateSlug(*i.Slug); err != nil {
			errors = append(errors, *err.(*models.ValidationError))
		}
	}

	if len(errors) > 0 {
		return errors
	}
	return nil
}

// BulkDeleteSitesInput represents the input for bulk deleting sites
type BulkDeleteSitesInput struct {
	IDs []int `json:"ids"`
}

// Validate validates the BulkDeleteSitesInput
func (i *BulkDeleteSitesInput) Validate() error {
	if len(i.IDs) == 0 {
		return &models.ValidationError{
			Field:   "ids",
			Message: "at least one ID must be provided",
		}
	}
	return nil
}

// BulkCreateSitesInput represents the input for bulk creating sites
type BulkCreateSitesInput struct {
	Sites []CreateSiteInput `json:"sites"`
}

// Validate validates the BulkCreateSitesInput
func (i *BulkCreateSitesInput) Validate() error {
	if len(i.Sites) == 0 {
		return &models.ValidationError{
			Field:   "sites",
			Message: "at least one site must be provided",
		}
	}

	var errors models.ValidationErrors
	for idx, site := range i.Sites {
		if err := site.Validate(); err != nil {
			if verr, ok := err.(models.ValidationErrors); ok {
				for _, e := range verr {
					e.Field = fmt.Sprintf("sites[%d].%s", idx, e.Field)
					errors = append(errors, e)
				}
			}
		}
	}

	if len(errors) > 0 {
		return errors
	}
	return nil
}

// ListSites lists all sites matching the input criteria
func (s *Service) ListSites(input *ListSitesInput) ([]Site, error) {
	path := s.buildPath("dcim", "sites")
	
	// Build query parameters
	params := map[string]string{}
	if input.Name != "" {
		params["name"] = input.Name
	}
	if input.Status != "" {
		params["status"] = input.Status
	}
	if input.Region != "" {
		params["region"] = input.Region
	}
	if input.Limit > 0 {
		params["limit"] = fmt.Sprintf("%d", input.Limit)
	}
	if input.Offset > 0 {
		params["offset"] = fmt.Sprintf("%d", input.Offset)
	}

	// Make request
	var response models.PaginatedResponse
	response.Results = make([]any, 0)
	_, err := s.client.httpClient.R().
		SetQueryParams(params).
		SetResult(&response).
		Get(path)
	
	if err != nil {
		return nil, fmt.Errorf("error listing sites: %w", err)
	}

	// Convert results to []Site
	sites := make([]Site, len(response.Results))
	for i, result := range response.Results {
		site, ok := result.(Site)
		if !ok {
			return nil, fmt.Errorf("unexpected result type at index %d", i)
		}
		sites[i] = site
	}

	return sites, nil
}

// GetSite retrieves a single site by ID
func (s *Service) GetSite(id int) (*Site, error) {
	path := s.buildPath("dcim", "sites", fmt.Sprintf("%d", id))
	
	var site Site
	resp, err := s.client.httpClient.R().
		SetResult(&site).
		Get(path)
	
	if err != nil {
		return nil, fmt.Errorf("error getting site: %w", err)
	}

	if resp.StatusCode() == http.StatusNotFound {
		return nil, fmt.Errorf("site not found")
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}

	return &site, nil
}

// CreateSite creates a new site
func (s *Service) CreateSite(input *CreateSiteInput) (*Site, error) {
	if err := input.Validate(); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	path := s.buildPath("dcim", "sites")
	
	var site Site
	resp, err := s.client.httpClient.R().
		SetBody(input).
		SetResult(&site).
		Post(path)
	
	if err != nil {
		return nil, fmt.Errorf("error creating site: %w", err)
	}

	if resp.StatusCode() != http.StatusCreated {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}

	return &site, nil
}

// UpdateSite updates an existing site
func (s *Service) UpdateSite(input *UpdateSiteInput) (*Site, error) {
	if err := input.Validate(); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	path := s.buildPath("dcim", "sites", fmt.Sprintf("%d", input.ID))
	
	var site Site
	resp, err := s.client.httpClient.R().
		SetBody(input).
		SetResult(&site).
		Put(path)
	
	if err != nil {
		return nil, fmt.Errorf("error updating site: %w", err)
	}

	if resp.StatusCode() == http.StatusNotFound {
		return nil, fmt.Errorf("site not found")
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}

	return &site, nil
}

// PatchSite partially updates an existing site
func (s *Service) PatchSite(input *PatchSiteInput) (*Site, error) {
	if err := input.Validate(); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	path := s.buildPath("dcim", "sites", fmt.Sprintf("%d", input.ID))
	
	var site Site
	resp, err := s.client.httpClient.R().
		SetBody(input).
		SetResult(&site).
		Patch(path)
	
	if err != nil {
		return nil, fmt.Errorf("error patching site: %w", err)
	}

	if resp.StatusCode() == http.StatusNotFound {
		return nil, fmt.Errorf("site not found")
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}

	return &site, nil
}

// BulkDeleteSites deletes multiple sites
func (s *Service) BulkDeleteSites(input *BulkDeleteSitesInput) error {
	if err := input.Validate(); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	path := s.buildPath("dcim", "sites", "bulk", "delete")
	
	resp, err := s.client.httpClient.R().
		SetBody(input).
		Post(path)
	
	if err != nil {
		return fmt.Errorf("error bulk deleting sites: %w", err)
	}

	if resp.StatusCode() != http.StatusNoContent {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}

	return nil
}

// BulkCreateSites creates multiple sites
func (s *Service) BulkCreateSites(input *BulkCreateSitesInput) ([]Site, error) {
	if err := input.Validate(); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	path := s.buildPath("dcim", "sites", "bulk", "create")
	
	var sites []Site
	resp, err := s.client.httpClient.R().
		SetBody(input).
		SetResult(&sites).
		Post(path)
	
	if err != nil {
		return nil, fmt.Errorf("error bulk creating sites: %w", err)
	}

	if resp.StatusCode() != http.StatusCreated {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}

	return sites, nil
}

// DeleteSite deletes a site
func (s *Service) DeleteSite(id int) error {
	path := s.buildPath("dcim", "sites", fmt.Sprintf("%d", id))
	
	resp, err := s.client.httpClient.R().
		Delete(path)
	
	if err != nil {
		return fmt.Errorf("error deleting site: %w", err)
	}

	if resp.StatusCode() == http.StatusNotFound {
		return fmt.Errorf("site not found")
	}

	if resp.StatusCode() != http.StatusNoContent {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}

	return nil
}
