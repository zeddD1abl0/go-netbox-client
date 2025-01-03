package extras

import (
	"github.com/zeddD1abl0/go-netbox-client/client"
)

// Service handles extras endpoints
type Service struct {
	*client.Service
}

// NewService creates a new extras service
func NewService(client *client.Client) *Service {
	return &Service{
		Service: client.NewService(),
	}
}
