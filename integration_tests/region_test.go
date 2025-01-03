package integration_tests

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zeddD1abl0/go-netbox-client/client"
)

func TestRegionIntegration(t *testing.T) {
	c := setupTestClient(t)
	cleanup := newCleanupList(t)
	defer cleanup.runAll()

	t.Run("Hierarchical regions", func(t *testing.T) {
		// Create parent region
		parentInput := &client.CreateRegionInput{
			Name:        "Parent Region",
			Slug:        "parent-region",
			Description: "Parent region for testing",
		}

		parent, err := c.CreateRegion(parentInput)
		require.NoError(t, err)
		require.NotNil(t, parent)
		cleanup.add(func() error {
			return c.DeleteRegion(parent.ID)
		})

		// Create child regions
		childRegions := make([]*client.Region, 3)
		for i := 0; i < 3; i++ {
			childInput := &client.CreateRegionInput{
				Name:        fmt.Sprintf("Child Region %d", i+1),
				Slug:        fmt.Sprintf("child-region-%d", i+1),
				Description: fmt.Sprintf("Child region %d of Parent Region", i+1),
				Parent:      parent.ID,
			}
			child, err := c.CreateRegion(childInput)
			require.NoError(t, err)
			require.NotNil(t, child)
			childRegions[i] = child
			cleanup.add(func() error {
				return c.DeleteRegion(child.ID)
			})
		}

		// List child regions
		listInput := &client.ListRegionsInput{
			Parent: fmt.Sprintf("%d", parent.ID),
		}
		list, err := c.ListRegions(listInput)
		require.NoError(t, err)
		require.NotNil(t, list)
		assert.Equal(t, 3, len(list))

		// Update a child region
		updateInput := &client.UpdateRegionInput{
			ID:          childRegions[0].ID,
			Name:        "Updated Child Region",
			Slug:        "updated-child-region",
			Description: "Updated child region",
			Parent:      parent.ID,
		}
		updated, err := c.UpdateRegion(updateInput)
		require.NoError(t, err)
		require.NotNil(t, updated)
		assert.Equal(t, updateInput.Name, updated.Name)
		assert.Equal(t, updateInput.Description, updated.Description)

		// Verify parent-child relationship
		retrieved, err := c.GetRegion(childRegions[0].ID)
		require.NoError(t, err)
		require.NotNil(t, retrieved)
		require.NotNil(t, retrieved.Parent)
		assert.Equal(t, parent.ID, retrieved.Parent.ID)
	})

	t.Run("Region filtering", func(t *testing.T) {
		// Test data
		regions := []struct {
			name        string
			slug        string
			description string
			tags        []string
		}{
			{
				name:        "Production Region",
				slug:        "prod-region",
				description: "Production region",
				tags:        []string{"prod"},
			},
			{
				name:        "Staging Region",
				slug:        "staging-region",
				description: "Staging region",
				tags:        []string{"test"},
			},
			{
				name:        "Development Region",
				slug:        "dev-region",
				description: "Development region",
				tags:        []string{"test"},
			},
		}

		// Create the regions
		for _, r := range regions {
			input := &client.CreateRegionInput{
				Name:        r.name,
				Slug:        r.slug,
				Description: r.description,
			}
			created, err := c.CreateRegion(input)
			require.NoError(t, err)
			require.NotNil(t, created)
			cleanup.add(func() error {
				return c.DeleteRegion(created.ID)
			})
		}

		// Test filtering
		tests := []struct {
			name          string
			input         *client.ListRegionsInput
			expectedCount int
			expectedName  string
		}{
			{
				name: "Filter by tag",
				input: &client.ListRegionsInput{
					Tag: "prod",
				},
				expectedCount: 1,
				expectedName:  "Production Region",
			},
			{
				name: "Filter by name contains",
				input: &client.ListRegionsInput{
					Name: "Staging",
				},
				expectedCount: 1,
				expectedName:  "Staging Region",
			},
			{
				name: "Filter by multiple tags",
				input: &client.ListRegionsInput{
					Tag: "test",
				},
				expectedCount: 2,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				list, err := c.ListRegions(tt.input)
				require.NoError(t, err)
				require.NotNil(t, list)
				assert.Equal(t, tt.expectedCount, len(list))
				if tt.expectedName != "" {
					assert.Equal(t, tt.expectedName, list[0].Name)
				}
			})
		}
	})

	t.Run("CRUD operations", func(t *testing.T) {
		// Test data
		testRegions := []struct {
			name        string
			slug        string
			description string
		}{
			{
				name:        "Region 1",
				slug:        "region-1",
				description: "Test region 1",
			},
			{
				name:        "Region 2",
				slug:        "region-2",
				description: "Test region 2",
			},
		}

		// Create regions
		var createdRegions []*client.Region
		for _, tt := range testRegions {
			input := &client.CreateRegionInput{
				Name:        tt.name,
				Slug:        tt.slug,
				Description: tt.description,
			}

			region, err := c.CreateRegion(input)
			require.NoError(t, err)
			require.NotNil(t, region)
			assert.Equal(t, tt.name, region.Name)
			assert.Equal(t, tt.slug, region.Slug)
			assert.Equal(t, tt.description, region.Description)

			createdRegions = append(createdRegions, region)
			cleanup.add(func() error {
				return c.DeleteRegion(region.ID)
			})
		}

		// List regions
		regions, err := c.ListRegions(&client.ListRegionsInput{})
		require.NoError(t, err)
		assert.GreaterOrEqual(t, len(regions), len(testRegions))

		// Get individual regions
		for _, cr := range createdRegions {
			region, err := c.GetRegion(cr.ID)
			require.NoError(t, err)
			assert.Equal(t, cr.Name, region.Name)
			assert.Equal(t, cr.Slug, region.Slug)
			assert.Equal(t, cr.Description, region.Description)
		}

		// Update a region
		updateInput := &client.UpdateRegionInput{
			ID:          createdRegions[0].ID,
			Name:        "Updated Region 1",
			Slug:        "updated-region-1",
			Description: "Updated test region 1",
		}
		updatedRegion, err := c.UpdateRegion(updateInput)
		require.NoError(t, err)
		assert.Equal(t, updateInput.Name, updatedRegion.Name)
		assert.Equal(t, updateInput.Slug, updatedRegion.Slug)
		assert.Equal(t, updateInput.Description, updatedRegion.Description)

		// Patch a region
		patchInput := &client.PatchRegionInput{
			ID:          createdRegions[1].ID,
			Description: strPtr("Patched test region 2"),
		}
		patchedRegion, err := c.PatchRegion(patchInput)
		require.NoError(t, err)
		assert.Equal(t, *patchInput.Description, patchedRegion.Description)
	})
}
