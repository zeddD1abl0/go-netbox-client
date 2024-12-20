package dcim

import (
	"github.com/zeddD1abl0/go-netbox-client/client"
)

// Service handles DCIM endpoints
type Service struct {
	*client.Service
}

// NewService creates a new DCIM service
func NewService(client *client.Client) *Service {
	return &Service{
		Service: client.NewService(),
	}
}
