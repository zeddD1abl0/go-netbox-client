package integration_tests

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zeddD1abl0/go-netbox-client/client/dcim"
	"github.com/zeddD1abl0/go-netbox-client/models"
)

func locationsToPointers(locations []dcim.Location) []*dcim.Location {
	result := make([]*dcim.Location, len(locations))
	for i := range locations {
		result[i] = &locations[i]
	}
	return result
}

func TestLocationIntegration(t *testing.T) {
	_, service := setupTestClient(t)
	cleanup := newCleanupList(t)
	defer cleanup.runAll()

	t.Run("Complex location hierarchy with relationships", func(t *testing.T) {
		// Create a site
		site, err := service.CreateSite(&dcim.CreateSiteInput{
			Name:        "Test Site",
			Slug:        "test-site",
			Description: "Top level site",
			Status:      "active",
		})
		require.NoError(t, err)
		cleanup.add(func() error {
			return service.DeleteSite(site.ID)
		})
		// Create a region hierarchy
		parentRegion, err := service.CreateRegion(&dcim.CreateRegionInput{
			Name:        "Parent Region",
			Slug:        "parent-region",
			Description: "Top level region",
		})
		require.NoError(t, err)
		cleanup.add(func() error {
			return service.DeleteRegion(parentRegion.ID)
		})

		childRegion, err := service.CreateRegion(&dcim.CreateRegionInput{
			Name:        "Child Region",
			Slug:        "child-region",
			Description: "Child region",
			Parent:      parentRegion.ID,
		})
		require.NoError(t, err)
		cleanup.add(func() error {
			return service.DeleteRegion(childRegion.ID)
		})

		// Create a complex location hierarchy
		locations := make(map[string]*dcim.Location)

		// Create campus location
		campusInput := &dcim.CreateLocationInput{
			Name:        "Main Campus",
			Slug:        "main-campus",
			Description: "Main campus location",
			Site:        site.ID,
			Tags: []models.TagCreate{
				{
					Name:  "campus",
					Slug:  "campus",
					Color: "ff0000",
				},
				{
					Name:  "main",
					Slug:  "main",
					Color: "00ff00",
				},
			},
		}
		campus, err := service.CreateLocation(campusInput)
		require.NoError(t, err)
		locations["campus"] = campus
		cleanup.add(func() error {
			return service.DeleteLocation(campus.ID)
		})

		// Create building locations
		for i := 1; i <= 2; i++ {
			buildingInput := &dcim.CreateLocationInput{
				Name:        fmt.Sprintf("Building %d", i),
				Slug:        fmt.Sprintf("building-%d", i),
				Description: fmt.Sprintf("Building %d in main campus", i),
				Parent:      campus.ID,
				Site:        site.ID,
				Tags: []models.TagCreate{
					{
						Name:  "building",
						Slug:  "building",
						Color: "0000ff",
					},
				},
			}
			building, err := service.CreateLocation(buildingInput)
			require.NoError(t, err)
			locations[fmt.Sprintf("building%d", i)] = building
			cleanup.add(func() error {
				return service.DeleteLocation(building.ID)
			})

			// Create floor locations for each building
			for j := 1; j <= 3; j++ {
				floorInput := &dcim.CreateLocationInput{
					Name:        fmt.Sprintf("Floor %d", j),
					Slug:        fmt.Sprintf("floor-%d", j),
					Description: fmt.Sprintf("Floor %d in Building %d", j, i),
					Parent:      building.ID,
					Site:        site.ID,
					Tags: []models.TagCreate{
						{
							Name:  "floor",
							Slug:  "floor",
							Color: "ff00ff",
						},
					},
				}
				floor, err := service.CreateLocation(floorInput)
				require.NoError(t, err)
				locations[fmt.Sprintf("building%d_floor%d", i, j)] = floor
				cleanup.add(func() error {
					return service.DeleteLocation(floor.ID)
				})
			}
		}

		// Test various filtering and relationship scenarios
		tests := []struct {
			name          string
			input         *dcim.ListLocationsInput
			expectedCount int
			validate      func(t *testing.T, results []*dcim.Location)
		}{
			{
				name: "Filter by parent",
				input: &dcim.ListLocationsInput{
					Parent: fmt.Sprintf("%d", locations["campus"].ID),
				},
				expectedCount: 2,
				validate: func(t *testing.T, results []*dcim.Location) {
					for _, loc := range results {
						assert.Contains(t, loc.Name, "Building")
					}
				},
			},
			{
				name: "Filter by tag",
				input: &dcim.ListLocationsInput{
					Tag: "floor",
				},
				expectedCount: 6,
				validate: func(t *testing.T, results []*dcim.Location) {
					for _, loc := range results {
						assert.Contains(t, loc.Name, "Floor")
					}
				},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				list, err := service.ListLocations(tt.input)
				require.NoError(t, err)
				require.NotNil(t, list)
				assert.Equal(t, tt.expectedCount, len(list))
				if tt.validate != nil {
					tt.validate(t, locationsToPointers(list))
				}
			})
		}

		// Test updating location with new relationships
		updateInput := &dcim.UpdateLocationInput{
			ID:          locations["building1"].ID,
			Name:        "Updated Building 1",
			Slug:        "updated-building-1",
			Description: "Updated building description",
			Site:        site.ID, // Change to parent group
			Parent:      campus.ID,
		}
		updated, err := service.UpdateLocation(updateInput)
		require.NoError(t, err)
		require.NotNil(t, updated)
		assert.Equal(t, updateInput.Name, updated.Name)
		assert.Equal(t, parentRegion.ID, updated.Site.Region.ID)
		assert.Equal(t, site.ID, updated.Site.ID)
	})
}
