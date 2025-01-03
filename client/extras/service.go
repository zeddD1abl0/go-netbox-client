package extras

import (
	"github.com/zeddD1abl0/go-netbox-client/client"
)

// Service provides access to the NetBox extras endpoints
type Service struct {
	*client.Client
}

// NewService creates a new extras service
func NewService(client *client.Client) *Service {
	return &Service{Client: client}
}

// BuildPath builds a path relative to the extras API root
func (service *Service) BuildPath(parts ...string) string {
	return service.Client.BuildPath(append([]string{"extras"}, parts...)...)
}
