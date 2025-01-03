package integration_tests

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zeddD1abl0/go-netbox-client/client/dcim"
)

func TestRegionIntegration(t *testing.T) {
	_, service := setupTestClient(t)
	cleanup := newCleanupList(t)
	defer cleanup.runAll()

	t.Run("Hierarchical regions", func(t *testing.T) {
		// Create parent region
		parentInput := &dcim.CreateRegionInput{
			Name:        "Parent Region",
			Slug:        "parent-region",
			Description: "Parent region for testing",
		}
		parent, err := service.CreateRegion(parentInput)
		require.NoError(t, err)
		require.NotNil(t, parent)
		cleanup.add(func() error {
			return service.DeleteRegion(parent.ID)
		})

		// Create child regions
		childRegions := make([]*dcim.Region, 3)
		for i := 0; i < 3; i++ {
			childInput := &dcim.CreateRegionInput{
				Name:        fmt.Sprintf("Child Region %d", i+1),
				Slug:        fmt.Sprintf("child-region-%d", i+1),
				Description: fmt.Sprintf("Child region %d of Parent Region", i+1),
				Parent:      parent.ID,
			}
			child, err := service.CreateRegion(childInput)
			require.NoError(t, err)
			require.NotNil(t, child)
			childRegions[i] = child
			cleanup.add(func() error {
				return service.DeleteRegion(child.ID)
			})
		}

		// List and verify hierarchy
		listInput := &dcim.ListRegionsInput{
			Parent: fmt.Sprintf("%d", parent.ID),
		}
		list, err := service.ListRegions(listInput)
		require.NoError(t, err)
		require.NotNil(t, list)
		//assert.Equal(t, 3, len(list))

		// Update a child region
		updateInput := &dcim.UpdateRegionInput{
			ID:          childRegions[0].ID,
			Name:        "Updated Child Region",
			Slug:        "updated-child-region",
			Description: "Updated child region description",
			Parent:      parent.ID,
		}
		updated, err := service.UpdateRegion(updateInput)
		require.NoError(t, err)
		require.NotNil(t, updated)
		assert.Equal(t, updateInput.Name, updated.Name)
		assert.Equal(t, updateInput.Description, updated.Description)

		// Verify parent-child relationship
		retrieved, err := service.GetRegion(childRegions[0].ID)
		require.NoError(t, err)
		require.NotNil(t, retrieved)
		require.NotNil(t, retrieved.Parent)
		assert.Equal(t, parent.ID, retrieved.Parent.ID)
	})

	t.Run("Region search and filtering", func(t *testing.T) {
		// Create regions with different attributes
		regions := []struct {
			name        string
			slug        string
			description string
			tags        []string
		}{
			{
				name:        "Production Region",
				slug:        "production-region",
				description: "Production environment",
				tags:        []string{"prod", "critical"},
			},
			{
				name:        "Staging Region",
				slug:        "staging-region",
				description: "Staging environment",
				tags:        []string{"staging", "test"},
			},
			{
				name:        "Development Region",
				slug:        "development-region",
				description: "Development environment",
				tags:        []string{"dev", "test"},
			},
		}

		// Create the regions
		for _, r := range regions {
			input := &dcim.CreateRegionInput{
				Name:        r.name,
				Slug:        r.slug,
				Description: r.description,
				//Tags:        make([]models.TagCreate, len(r.tags)),
			}
			// for i, tag := range r.tags {
			// 	input.Tags[i] = models.TagCreate{
			// 		Name:  tag,
			// 		Slug:  tag,
			// 		Color: "0xFF00FF",
			// 	}
			// }
			// pretty, err := json.MarshalIndent(input, "", "  ")
			// fmt.Printf("Creating region: %s\n", pretty)
			created, err := service.CreateRegion(input)
			require.NoError(t, err)
			require.NotNil(t, created)
			cleanup.add(func() error {
				return service.DeleteRegion(created.ID)
			})
		}

		// Test different filter combinations
		tests := []struct {
			name          string
			input         *dcim.ListRegionsInput
			expectedCount int
			expectedName  string
		}{
			{
				name: "Filter by tag",
				input: &dcim.ListRegionsInput{
					Tag: "prod",
				},
				expectedCount: 1,
				expectedName:  "Production Region",
			},
			{
				name: "Filter by name contains",
				input: &dcim.ListRegionsInput{
					Name: "Staging",
				},
				expectedCount: 1,
				expectedName:  "Staging Region",
			},
			{
				name: "Filter by multiple tags",
				input: &dcim.ListRegionsInput{
					Tag: "test",
				},
				expectedCount: 2,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				list, err := service.ListRegions(tt.input)
				require.NoError(t, err)
				require.NotNil(t, list)
				assert.Equal(t, tt.expectedCount, len(list))
				// if tt.expectedName != "" {
				// 	assert.Equal(t, tt.expectedName, list[0].Name)
				// }
			})
		}
	})
}
