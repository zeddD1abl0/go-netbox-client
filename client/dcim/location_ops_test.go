package dcim

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zeddD1abl0/go-netbox-client/client"
)

func TestListLocations(t *testing.T) {
	client, mux, teardown := client.NewClientForTesting()
	defer teardown()

	mux.HandleFunc("/api/dcim/locations/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		resp := client.Response[[]Location]{
			Results: []Location{
				{
					ID:   1,
					Name: "Test Location",
					Slug: "test-location",
				},
			},
		}
		json.NewEncoder(w).Encode(resp)
	})

	service := NewService(client)
	locations, err := service.ListLocations(&ListLocationsInput{
		Name: "Test Location",
	})

	assert.NoError(t, err)
	assert.Equal(t, 1, len(locations))
	assert.Equal(t, "Test Location", locations[0].Name)
}

func TestGetLocation(t *testing.T) {
	client, mux, teardown := client.NewClientForTesting()
	defer teardown()

	mux.HandleFunc("/api/dcim/locations/1/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		location := Location{
			ID:   1,
			Name: "Test Location",
			Slug: "test-location",
		}
		json.NewEncoder(w).Encode(location)
	})

	service := NewService(client)
	location, err := service.GetLocation(1)

	assert.NoError(t, err)
	assert.Equal(t, "Test Location", location.Name)
}

func TestCreateLocation(t *testing.T) {
	client, mux, teardown := client.NewClientForTesting()
	defer teardown()

	mux.HandleFunc("/api/dcim/locations/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)

		var input CreateLocationInput
		json.NewDecoder(r.Body).Decode(&input)
		assert.Equal(t, "Test Location", input.Name)

		location := Location{
			ID:   1,
			Name: input.Name,
			Slug: input.Slug,
		}
		json.NewEncoder(w).Encode(location)
	})

	service := NewService(client)
	input := &CreateLocationInput{
		Name: "Test Location",
		Slug: "test-location",
		Site: 1,
	}
	location, err := service.CreateLocation(input)

	assert.NoError(t, err)
	assert.Equal(t, "Test Location", location.Name)
}

func TestUpdateLocation(t *testing.T) {
	client, mux, teardown := client.NewClientForTesting()
	defer teardown()

	mux.HandleFunc("/api/dcim/locations/1/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "PUT", r.Method)

		var input UpdateLocationInput
		json.NewDecoder(r.Body).Decode(&input)
		assert.Equal(t, "Updated Location", input.Name)

		location := Location{
			ID:   1,
			Name: input.Name,
			Slug: input.Slug,
		}
		json.NewEncoder(w).Encode(location)
	})

	service := NewService(client)
	input := &UpdateLocationInput{
		ID:   1,
		Name: "Updated Location",
		Slug: "updated-location",
		Site: 1,
	}
	location, err := service.UpdateLocation(input)

	assert.NoError(t, err)
	assert.Equal(t, "Updated Location", location.Name)
}

func TestPatchLocation(t *testing.T) {
	client, mux, teardown := client.NewClientForTesting()
	defer teardown()

	mux.HandleFunc("/api/dcim/locations/1/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "PATCH", r.Method)

		var input PatchLocationInput
		json.NewDecoder(r.Body).Decode(&input)
		assert.Equal(t, "Patched Location", *input.Name)

		location := Location{
			ID:   1,
			Name: *input.Name,
			Slug: *input.Slug,
		}
		json.NewEncoder(w).Encode(location)
	})

	service := NewService(client)
	name := "Patched Location"
	slug := "patched-location"
	input := &PatchLocationInput{
		ID:   1,
		Name: &name,
		Slug: &slug,
	}
	location, err := service.PatchLocation(input)

	assert.NoError(t, err)
	assert.Equal(t, "Patched Location", location.Name)
}

func TestDeleteLocation(t *testing.T) {
	client, mux, teardown := client.NewClientForTesting()
	defer teardown()

	mux.HandleFunc("/api/dcim/locations/1/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "DELETE", r.Method)
		w.WriteHeader(http.StatusNoContent)
	})

	service := NewService(client)
	err := service.DeleteLocation(1)

	assert.NoError(t, err)
}
