package integration_tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zeddD1abl0/go-netbox-client/client/dcim"
)

func TestRegionIntegration(t *testing.T) {
	client := setupTestClient(t)
	cleanup := newCleanupList(t)
	defer cleanup.runAll()

	t.Run("Hierarchical regions", func(t *testing.T) {
		// Create parent region
		parentInput := &dcim.CreateRegionInput{
			Name:        "Parent Region",
			Slug:        "parent-region",
			Description: "Parent region for testing",
		}
		parent, err := client.DCIM().CreateRegion(parentInput)
		require.NoError(t, err)
		require.NotNil(t, parent)
		cleanup.add(func() error {
			return client.DCIM().DeleteRegion(parent.ID)
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
			child, err := client.DCIM().CreateRegion(childInput)
			require.NoError(t, err)
			require.NotNil(t, child)
			childRegions[i] = child
			cleanup.add(func() error {
				return client.DCIM().DeleteRegion(child.ID)
			})
		}

		// List and verify hierarchy
		listInput := &dcim.ListRegionsInput{
			Parent: fmt.Sprintf("%d", parent.ID),
		}
		list, err := client.DCIM().ListRegions(listInput)
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
		updated, err := client.DCIM().UpdateRegion(updateInput)
		require.NoError(t, err)
		require.NotNil(t, updated)
		assert.Equal(t, updateInput.Name, updated.Name)
		assert.Equal(t, updateInput.Description, updated.Description)

		// Verify parent-child relationship
		retrieved, err := client.DCIM().GetRegion(childRegions[0].ID)
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
			created, err := client.DCIM().CreateRegion(input)
			require.NoError(t, err)
			require.NotNil(t, created)
			cleanup.add(func() error {
				return client.DCIM().DeleteRegion(created.ID)
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
				list, err := client.DCIM().ListRegions(tt.input)
				require.NoError(t, err)
				require.NotNil(t, list)
				assert.Equal(t, tt.expectedCount, len(list))
				// if tt.expectedName != "" {
				// 	assert.Equal(t, tt.expectedName, list[0].Name)
				// }
			})
		}
	})

	t.Run("CRUD operations", func(t *testing.T) {
		// Create test data
		testRegions := []struct {
			name        string
			slug        string
			description string
		}{
			{
				name:        "Test Region 1",
				slug:        "test-region-1",
				description: "Test region 1",
			},
			{
				name:        "Test Region 2",
				slug:        "test-region-2",
				description: "Test region 2",
			},
		}

		// Create regions
		var createdRegions []*dcim.Region
		for _, tt := range testRegions {
			input := &dcim.CreateRegionInput{
				Name:        tt.name,
				Slug:        tt.slug,
				Description: tt.description,
			}

			region, err := client.DCIM().CreateRegion(input)
			require.NoError(t, err)
			require.NotNil(t, region)
			assert.Equal(t, tt.name, region.Name)
			assert.Equal(t, tt.slug, region.Slug)
			assert.Equal(t, tt.description, region.Description)

			createdRegions = append(createdRegions, region)
			cleanup.add(func() error {
				return client.DCIM().DeleteRegion(region.ID)
			})
		}

		// List regions
		regions, err := client.DCIM().ListRegions(&dcim.ListRegionsInput{})
		require.NoError(t, err)
		assert.GreaterOrEqual(t, len(regions), len(testRegions))

		// Get individual regions
		for _, cr := range createdRegions {
			region, err := client.DCIM().GetRegion(cr.ID)
			require.NoError(t, err)
			assert.Equal(t, cr.Name, region.Name)
			assert.Equal(t, cr.Slug, region.Slug)
			assert.Equal(t, cr.Description, region.Description)
		}

		// Update a region
		updateInput := &dcim.UpdateRegionInput{
			ID:          createdRegions[0].ID,
			Name:        "Updated Region 1",
			Slug:        "updated-region-1",
			Description: "Updated test region 1",
		}
		updatedRegion, err := client.DCIM().UpdateRegion(updateInput)
		require.NoError(t, err)
		assert.Equal(t, updateInput.Name, updatedRegion.Name)
		assert.Equal(t, updateInput.Slug, updatedRegion.Slug)
		assert.Equal(t, updateInput.Description, updatedRegion.Description)

		// Patch a region
		patchInput := &dcim.PatchRegionInput{
			ID:          createdRegions[1].ID,
			Description: strPtr("Patched test region 2"),
		}
		patchedRegion, err := client.DCIM().PatchRegion(patchInput)
		require.NoError(t, err)
		assert.Equal(t, *patchInput.Description, patchedRegion.Description)
	})
}
