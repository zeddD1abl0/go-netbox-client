package dcim

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/jordan/go-netbox-client/client"
)

// ListLocations lists all locations
func (s *Service) ListLocations(input *ListLocationsInput) ([]Location, error) {
	params := url.Values{}
	if input != nil {
		if input.Name != "" {
			params.Add("name", input.Name)
		}
		if input.Site != "" {
			params.Add("site", input.Site)
		}
		if input.Parent != "" {
			params.Add("parent", input.Parent)
		}
		if input.Tag != "" {
			params.Add("tag", input.Tag)
		}
		if input.Limit != 0 {
			params.Add("limit", strconv.Itoa(input.Limit))
		}
		if input.Offset != 0 {
			params.Add("offset", strconv.Itoa(input.Offset))
		}
	}

	req, err := s.client.NewRequest("GET", "dcim/locations/", params)
	if err != nil {
		return nil, err
	}

	locations := new(client.Response[[]Location])
	err = s.client.Do(req, locations)
	if err != nil {
		return nil, err
	}

	return locations.Results, nil
}

// GetLocation gets a location by ID
func (s *Service) GetLocation(id int) (*Location, error) {
	req, err := s.client.NewRequest("GET", fmt.Sprintf("dcim/locations/%d/", id), nil)
	if err != nil {
		return nil, err
	}

	location := new(Location)
	err = s.client.Do(req, location)
	if err != nil {
		return nil, err
	}

	return location, nil
}

// CreateLocation creates a new location
func (s *Service) CreateLocation(input *CreateLocationInput) (*Location, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("POST", "dcim/locations/", nil)
	if err != nil {
		return nil, err
	}

	req.Body = input
	location := new(Location)
	err = s.client.Do(req, location)
	if err != nil {
		return nil, err
	}

	return location, nil
}

// UpdateLocation updates a location
func (s *Service) UpdateLocation(input *UpdateLocationInput) (*Location, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("PUT", fmt.Sprintf("dcim/locations/%d/", input.ID), nil)
	if err != nil {
		return nil, err
	}

	req.Body = input
	location := new(Location)
	err = s.client.Do(req, location)
	if err != nil {
		return nil, err
	}

	return location, nil
}

// PutLocation creates or updates a location
func (s *Service) PutLocation(input *UpdateLocationInput) (*Location, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("PUT", fmt.Sprintf("dcim/locations/%d/", input.ID), nil)
	if err != nil {
		return nil, err
	}

	req.Body = input
	location := new(Location)
	err = s.client.Do(req, location)
	if err != nil {
		return nil, err
	}

	return location, nil
}

// PatchLocation patches a location
func (s *Service) PatchLocation(input *PatchLocationInput) (*Location, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("PATCH", fmt.Sprintf("dcim/locations/%d/", input.ID), nil)
	if err != nil {
		return nil, err
	}

	req.Body = input
	location := new(Location)
	err = s.client.Do(req, location)
	if err != nil {
		return nil, err
	}

	return location, nil
}

// DeleteLocation deletes a location
func (s *Service) DeleteLocation(id int) error {
	req, err := s.client.NewRequest("DELETE", fmt.Sprintf("dcim/locations/%d/", id), nil)
	if err != nil {
		return err
	}

	return s.client.Do(req, nil)
}
