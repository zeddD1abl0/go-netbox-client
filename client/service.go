package client

// Service represents a Netbox API service
type Service struct {
	client *Client
}

// newService creates a new service with the given client
func newService(client *Client) *Service {
	return &Service{
		client: client,
	}
}

// buildPath builds a full API path from the given parts
func (s *Service) buildPath(parts ...string) string {
	path := s.client.baseURL
	for _, part := range parts {
		path += "/" + part
	}
	return path
}
