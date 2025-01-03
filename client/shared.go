package client

// Status represents a status value in Netbox
type Status struct {
	Value       string `json:"value"`
	Label       string `json:"label"`
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Color       string `json:"color"`
	Description string `json:"description"`
}
