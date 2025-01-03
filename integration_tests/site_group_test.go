package integration_tests

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zeddD1abl0/go-netbox-client/client/dcim"
	"github.com/zeddD1abl0/go-netbox-client/models"
)

func TestSiteGroupIntegration(t *testing.T) {
	client := setupTestClient(t)
	cleanup := newCleanupList(t)
	defer cleanup.runAll()

	t.Run("Hierarchical site groups", func(t *testing.T) {
		// Create parent site group
		parentInput := &dcim.CreateSiteGroupInput{
			Name:        "Parent Site Group",
			Slug:        "parent-site-group",
			Description: "Parent site group for testing",
		}
		parent, err := client.DCIM().CreateSiteGroup(parentInput)
		require.NoError(t, err)
		require.NotNil(t, parent)
		cleanup.add(func() error {
			return client.DCIM().DeleteSiteGroup(parent.ID)
		})

		// Create child site groups
		childGroups := make([]*dcim.SiteGroup, 3)
		for i := 0; i < 3; i++ {
			childInput := &dcim.CreateSiteGroupInput{
				Name:        fmt.Sprintf("Child Site Group %d", i+1),
				Slug:        fmt.Sprintf("child-site-group-%d", i+1),
				Description: fmt.Sprintf("Child site group %d of Parent Site Group", i+1),
				Parent:      parent.ID,
			}
			child, err := client.DCIM().CreateSiteGroup(childInput)
			require.NoError(t, err)
			require.NotNil(t, child)
			childGroups[i] = child
			cleanup.add(func() error {
				return client.DCIM().DeleteSiteGroup(child.ID)
			})
		}

		// List and verify hierarchy
		fmt.Println(parent.ID)
		listInput := &dcim.ListSiteGroupsInput{
			Parent: fmt.Sprintf("%d", parent.ID),
		}
		list, err := client.DCIM().ListSiteGroups(listInput)
		require.NoError(t, err)
		require.NotNil(t, list)
		assert.Equal(t, 3, len(list))

		// Update a child site group
		updateInput := &dcim.UpdateSiteGroupInput{
			ID:          childGroups[0].ID,
			Name:        "Updated Child Site Group",
			Slug:        "updated-child-site-group",
			Description: "Updated child site group description",
			Parent:      parent.ID,
		}
		updated, err := client.DCIM().UpdateSiteGroup(updateInput)
		require.NoError(t, err)
		require.NotNil(t, updated)
		assert.Equal(t, updateInput.Name, updated.Name)
		assert.Equal(t, updateInput.Description, updated.Description)

		// Verify parent-child relationship
		retrieved, err := client.DCIM().GetSiteGroup(childGroups[0].ID)
		require.NoError(t, err)
		require.NotNil(t, retrieved)
		require.NotNil(t, retrieved.Parent)
		assert.Equal(t, parent.ID, retrieved.Parent.ID)
	})

	t.Run("Site group search and filtering", func(t *testing.T) {
		// Create site groups with different attributes
		groups := []struct {
			name        string
			slug        string
			description string
			tags        []string
		}{
			{
				name:        "Production Site Group",
				slug:        "production-site-group",
				description: "Production environment",
				tags:        []string{"prod", "critical"},
			},
			{
				name:        "Staging Site Group",
				slug:        "staging-site-group",
				description: "Staging environment",
				tags:        []string{"staging", "test"},
			},
			{
				name:        "Development Site Group",
				slug:        "development-site-group",
				description: "Development environment",
				tags:        []string{"dev", "test"},
			},
		}

		// Create the site groups
		for _, g := range groups {
			input := &dcim.CreateSiteGroupInput{
				Name:        g.name,
				Slug:        g.slug,
				Description: g.description,
				Tags:        make([]models.TagCreate, len(g.tags)),
			}
			for i, tag := range g.tags {
				input.Tags[i] = models.TagCreate{
					Name:  tag,
					Slug:  tag,
					Color: "0xFF00FF",
				}
			}
			created, err := client.DCIM().CreateSiteGroup(input)
			require.NoError(t, err)
			require.NotNil(t, created)
			cleanup.add(func() error {
				return client.DCIM().DeleteSiteGroup(created.ID)
			})
		}

		// Test different filter combinations
		tests := []struct {
			name          string
			input         *dcim.ListSiteGroupsInput
			expectedCount int
			expectedName  string
		}{
			{
				name: "Filter by tag",
				input: &dcim.ListSiteGroupsInput{
					Tag: "prod",
				},
				expectedCount: 1,
				expectedName:  "Production Site Group",
			},
			{
				name: "Filter by name contains",
				input: &dcim.ListSiteGroupsInput{
					Name: "Staging",
				},
				expectedCount: 1,
				expectedName:  "Staging Site Group",
			},
			{
				name: "Filter by multiple tags",
				input: &dcim.ListSiteGroupsInput{
					Tag: "test",
				},
				expectedCount: 2,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				list, err := client.DCIM().ListSiteGroups(tt.input)
				require.NoError(t, err)
				require.NotNil(t, list)
				assert.Equal(t, tt.expectedCount, len(list))
				if tt.expectedName != "" {
					assert.Equal(t, tt.expectedName, list[0].Name)
				}
			})
		}
	})
}
