package integration_tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zeddD1abl0/go-netbox-client/client/dcim"
)

func TestSiteIntegration(t *testing.T) {
	client := setupTestClient(t)
	cleanup := newCleanupList(t)
	defer cleanup.runAll()

	t.Run("CRUD operations", func(t *testing.T) {

		testRegions := &dcim.CreateRegionInput{
			Name:        "AP South East 2",
			Slug:        "ap-southeast-2",
			Description: "AWS AP South East 2",
		}

		region, err := client.DCIM().CreateRegion(testRegions)

		// Create test data
		testSites := []struct {
			name        string
			slug        string
			status      string
			region      string
			description string
		}{
			{
				name:        "Test Site 1",
				slug:        "test-site-1",
				status:      "active",
				region:      region.ID,
				description: "Test site 1",
			},
			{
				name:        "Test Site 2",
				slug:        "test-site-2",
				status:      "planned",
				region:      region.ID,
				description: "Test site 2",
			},
		}

		// Create sites
		var createdSites []*dcim.Site
		for _, tt := range testSites {
			input := &dcim.CreateSiteInput{
				Name:        tt.name,
				Slug:        tt.slug,
				Status:      tt.status,
				Region:      region.ID,
				Description: tt.description,
			}

			site, err := client.DCIM().CreateSite(input)
			require.NoError(t, err)
			require.NotNil(t, site)
			assert.Equal(t, tt.name, site.Name)
			assert.Equal(t, tt.slug, site.Slug)
			assert.Equal(t, tt.status, site.Status)
			assert.Equal(t, tt.description, site.Description)

			createdSites = append(createdSites, site)
			cleanup.add(func() error {
				return client.DCIM().DeleteSite(site.ID)
			})
		}

		// List sites
		sites, err := client.DCIM().ListSites(&dcim.ListSitesInput{})
		require.NoError(t, err)
		assert.GreaterOrEqual(t, len(sites), len(testSites))

		// Get individual sites
		for _, cs := range createdSites {
			site, err := client.DCIM().GetSite(cs.ID)
			require.NoError(t, err)
			assert.Equal(t, cs.Name, site.Name)
			assert.Equal(t, cs.Slug, site.Slug)
			assert.Equal(t, cs.Status, site.Status)
			assert.Equal(t, cs.Description, site.Description)
		}

		// Update a site
		updateInput := &dcim.UpdateSiteInput{
			ID:          createdSites[0].ID,
			Name:        "Updated Site 1",
			Slug:        "updated-site-1",
			Status:      "decommissioning",
			Description: "Updated test site 1",
		}
		updatedSite, err := client.DCIM().UpdateSite(updateInput)
		require.NoError(t, err)
		assert.Equal(t, updateInput.Name, updatedSite.Name)
		assert.Equal(t, updateInput.Slug, updatedSite.Slug)
		assert.Equal(t, updateInput.Status, updatedSite.Status)
		assert.Equal(t, updateInput.Description, updatedSite.Description)

		// Patch a site
		patchInput := &dcim.PatchSiteInput{
			ID:          &createdSites[1].ID,
			Description: strPtr("Patched test site 2"),
		}
		patchedSite, err := client.DCIM().PatchSite(patchInput)
		require.NoError(t, err)
		assert.Equal(t, *patchInput.Description, patchedSite.Description)
	})
}
