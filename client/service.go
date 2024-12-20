package client

// Service represents a Netbox API service
type Service struct {
	Client *Client
}

// newService creates a new service with the given client
func newService(client *Client) *Service {
	return &Service{
		Client: client,
	}
}

// BuildPath builds a full API path from the given parts
func (s *Service) BuildPath(parts ...string) string {
	path := s.Client.baseURL
	for _, part := range parts {
		path += part + "/"
	}
	return path
}

// NewService creates a new service with the given client
func (c *Client) NewService() *Service {
	return &Service{
		Client: c,
	}
}
