package integration_tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zeddD1abl0/go-netbox-client/client"
)

func TestSiteIntegration(t *testing.T) {
	c := setupTestClient(t)
	cleanup := newCleanupList(t)
	defer cleanup.runAll()

	t.Run("CRUD operations", func(t *testing.T) {
		// Create a region first
		testRegion := &client.CreateRegionInput{
			Name:        "AP South East 2",
			Slug:        "ap-southeast-2",
			Description: "AWS AP South East 2",
		}

		region, err := c.CreateRegion(testRegion)
		require.NoError(t, err)
		cleanup.add(func() error {
			return c.DeleteRegion(region.ID)
		})

		// Create test data
		testSites := []struct {
			name        string
			slug        string
			status      string
			region      int
			description string
		}{
			{
				name:        "Test Site 1",
				slug:        "test-site-1",
				status:      client.SiteStatusActive,
				region:      region.ID,
				description: "Test site 1",
			},
			{
				name:        "Test Site 2",
				slug:        "test-site-2",
				status:      client.SiteStatusPlanned,
				region:      region.ID,
				description: "Test site 2",
			},
		}

		// Create sites
		var createdSites []*client.Site
		for _, tt := range testSites {
			input := &client.CreateSiteInput{
				Name:        tt.name,
				Slug:        tt.slug,
				Status:      tt.status,
				Region:      tt.region,
				Description: tt.description,
			}

			site, err := c.CreateSite(input)
			require.NoError(t, err)
			require.NotNil(t, site)
			assert.Equal(t, tt.name, site.Name)
			assert.Equal(t, tt.slug, site.Slug)
			assert.Equal(t, tt.status, site.Status.Value)
			assert.Equal(t, tt.description, site.Description)

			createdSites = append(createdSites, site)
			cleanup.add(func() error {
				return c.DeleteSite(site.ID)
			})
		}

		// List sites
		sites, err := c.ListSites(&client.ListSitesInput{})
		require.NoError(t, err)
		assert.GreaterOrEqual(t, len(sites), len(testSites))

		// Get each created site
		for _, cs := range createdSites {
			site, err := c.GetSite(cs.ID)
			require.NoError(t, err)
			assert.Equal(t, cs.Name, site.Name)
			assert.Equal(t, cs.Status.Value, site.Status.Value)
		}

		// Update a site
		updateInput := &client.UpdateSiteInput{
			ID:          createdSites[0].ID,
			Name:        "Updated Test Site 1",
			Slug:        "updated-test-site-1",
			Status:      client.SiteStatusDecommissioning,
			Description: "Updated test site 1",
			Region:      region.ID,
		}
		updatedSite, err := c.UpdateSite(updateInput)
		require.NoError(t, err)
		assert.Equal(t, updateInput.Name, updatedSite.Name)
		assert.Equal(t, updateInput.Status, updatedSite.Status.Value)

		// Patch a site
		patchInput := &client.PatchSiteInput{
			ID:          &createdSites[1].ID,
			Description: strPtr("Patched test site 2"),
		}
		patchedSite, err := c.PatchSite(patchInput)
		require.NoError(t, err)
		assert.Equal(t, *patchInput.Description, patchedSite.Description)
	})
}
