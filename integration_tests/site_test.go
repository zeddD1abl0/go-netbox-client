package integration_tests

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zeddD1abl0/go-netbox-client/client/dcim"
	"github.com/zeddD1abl0/go-netbox-client/models"
)

func TestSiteIntegration(t *testing.T) {
	_, service := setupTestClient(t)
	cleanup := newCleanupList(t)
	defer cleanup.runAll()

	t.Run("Site with region and tags", func(t *testing.T) {
		// Create a region first
		region, err := service.CreateRegion(&dcim.CreateRegionInput{
			Name:        "Test Region",
			Slug:        "test-region",
			Description: "Region for site testing",
		})
		require.NoError(t, err)
		require.NotNil(t, region)
		cleanup.add(func() error {
			return service.DeleteRegion(region.ID)
		})

		// Create a site with various attributes
		siteInput := &dcim.CreateSiteInput{
			Name:            "Main Data Center",
			Slug:            "main-dc",
			Status:          dcim.SiteStatusActive,
			Region:          region.ID,
			Description:     "Primary data center facility",
			PhysicalAddress: "123 Server Street, Rack City",
			Latitude:        float64Ptr(37.7749),
			Longitude:       float64Ptr(-122.4194),
			Tags: []models.TagCreate{
				{
					Name:  "datacenter",
					Slug:  "datacenter",
					Color: "0xFF00FF",
				},
				{
					Name:  "production",
					Slug:  "production",
					Color: "0xFF00FF",
				},
			},
		}

		site, err := service.CreateSite(siteInput)
		require.NoError(t, err)
		require.NotNil(t, site)
		cleanup.add(func() error {
			return service.DeleteSite(site.ID)
		})

		// Verify the created site
		assert.Equal(t, siteInput.Name, site.Name)
		assert.Equal(t, siteInput.Status, site.Status.Value)
		assert.Equal(t, region.ID, site.Region.ID)
		assert.Equal(t, siteInput.PhysicalAddress, site.PhysicalAddress)
		assert.Equal(t, siteInput.Latitude, site.Latitude)
		assert.Equal(t, siteInput.Longitude, site.Longitude)
		assert.Len(t, site.Tags, 2)

		// Update the site
		updateInput := &dcim.UpdateSiteInput{
			ID:              site.ID,
			Name:            "Updated Data Center",
			Slug:            "updated-dc",
			Status:          dcim.SiteStatusPlanned,
			Region:          region.ID,
			Description:     "Updated description",
			PhysicalAddress: "456 Updated Street",
		}

		updated, err := service.UpdateSite(updateInput)
		require.NoError(t, err)
		require.NotNil(t, updated)
		assert.Equal(t, updateInput.Name, updated.Name)
		assert.Equal(t, updateInput.Status, updated.Status.Value)
		assert.Equal(t, updateInput.PhysicalAddress, updated.PhysicalAddress)
	})

	t.Run("Site lifecycle and status transitions", func(t *testing.T) {
		// Create sites with different statuses
		statuses := []string{
			dcim.SiteStatusPlanned,
			dcim.SiteStatusStaging,
			dcim.SiteStatusActive,
			dcim.SiteStatusDecommissioning,
			dcim.SiteStatusRetired,
		}

		for i, status := range statuses {
			input := &dcim.CreateSiteInput{
				Name:        fmt.Sprintf("Site %d", i+1),
				Slug:        fmt.Sprintf("site-%d", i+1),
				Status:      status,
				Description: fmt.Sprintf("Site with status: %s", status),
			}
			site, err := service.CreateSite(input)
			require.NoError(t, err)
			require.NotNil(t, site)
			cleanup.add(func() error {
				return service.DeleteSite(site.ID)
			})

			// Verify status
			assert.Equal(t, status, site.Status.Value)
		}

		// Test listing with status filter
		listInput := &dcim.ListSitesInput{
			Status: dcim.SiteStatusActive,
		}
		list, err := service.ListSites(listInput)
		require.NoError(t, err)
		require.NotNil(t, list)

		// Verify all returned sites have active status
		for _, site := range list {
			assert.Equal(t, dcim.SiteStatusActive, site.Status.Value)
		}
	})

	t.Run("Site search and filtering", func(t *testing.T) {
		// Create sites with different attributes for testing filters
		sites := []struct {
			name       string
			slug       string
			status     string
			tags       []string
			facilities []string
		}{
			{
				name:       "Production Site",
				slug:       "prod-site",
				status:     dcim.SiteStatusActive,
				tags:       []string{"prod", "critical"},
				facilities: []string{"datacenter"},
			},
			{
				name:       "Staging Site",
				slug:       "staging-site",
				status:     dcim.SiteStatusStaging,
				tags:       []string{"staging", "test"},
				facilities: []string{"lab"},
			},
			{
				name:       "Development Site",
				slug:       "dev-site",
				status:     dcim.SiteStatusActive,
				tags:       []string{"dev", "test"},
				facilities: []string{"office"},
			},
		}

		// Create the sites
		for _, s := range sites {
			input := &dcim.CreateSiteInput{
				Name:   s.name,
				Slug:   s.slug,
				Status: s.status,
				Tags:   make([]models.TagCreate, len(s.tags)),
			}
			for i, tag := range s.tags {
				input.Tags[i] = models.TagCreate{
					Name:  tag,
					Slug:  tag,
					Color: "0xFF00FF",
				}
			}
			created, err := service.CreateSite(input)
			require.NoError(t, err)
			require.NotNil(t, created)
			cleanup.add(func() error {
				return service.DeleteSite(created.ID)
			})
		}

		// Test different filter combinations
		tests := []struct {
			name          string
			input         *dcim.ListSitesInput
			expectedCount int
			expectedName  string
		}{
			{
				name: "Filter by status",
				input: &dcim.ListSitesInput{
					Status: dcim.SiteStatusActive,
				},
				expectedCount: 2,
			},
			{
				name: "Filter by name contains",
				input: &dcim.ListSitesInput{
					BaseListInput: dcim.BaseListInput{
						Name: "Site",
					},
				},
				expectedCount: 1,
				expectedName:  "Staging Site",
			},
			{
				name: "Filter by tag",
				input: &dcim.ListSitesInput{
					BaseListInput: dcim.BaseListInput{
						Tag: "prod",
					},
				},
				expectedCount: 1,
				expectedName:  "Production Site",
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				list, err := service.ListSites(tt.input)
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

func float64Ptr(v float64) *float64 {
	return &v
}
