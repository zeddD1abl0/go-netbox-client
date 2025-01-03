package integration_tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zeddD1abl0/go-netbox-client/models"
)

func TestTagIntegration(t *testing.T) {
	client := setupTestClient(t)
	cleanup := newCleanupList(t)
	defer cleanup.runAll()

	t.Run("CRUD operations", func(t *testing.T) {
		// Create test data
		testTags := []struct {
			name        string
			slug        string
			color       string
			description string
			objectTypes []string
		}{
			{
				name:        "Production",
				slug:        "production",
				color:       "ff0000",
				description: "Production environment",
				objectTypes: []string{"dcim.site", "dcim.region"},
			},
			{
				name:        "Development",
				slug:        "development",
				color:       "00ff00",
				description: "Development environment",
				objectTypes: []string{"dcim.site", "dcim.region"},
			},
			{
				name:        "Testing",
				slug:        "testing",
				color:       "0000ff",
				description: "Testing environment",
				objectTypes: []string{"dcim.site", "dcim.region"},
			},
		}

		// Create tags
		var createdTags []*models.Tag
		for _, tt := range testTags {
			input := &client.CreateTagInput{
				Name:        tt.name,
				Slug:        tt.slug,
				Color:       tt.color,
				Description: tt.description,
				ObjectTypes: tt.objectTypes,
			}

			tag, err := client.CreateTag(input)
			require.NoError(t, err)
			require.NotNil(t, tag)
			assert.Equal(t, tt.name, tag.Name)
			assert.Equal(t, tt.slug, tag.Slug)
			assert.Equal(t, tt.color, tag.Color)
			assert.Equal(t, tt.description, tag.Description)
			assert.Equal(t, tt.objectTypes, tag.ObjectTypes)

			createdTags = append(createdTags, tag)
			cleanup.add(func() error {
				return client.DeleteTag(tag.ID)
			})
		}

		// List tags
		tags, err := client.ListTags(&extras.ListTagsInput{})
		require.NoError(t, err)
		assert.GreaterOrEqual(t, len(tags), len(testTags))

		// Get individual tags
		for _, ct := range createdTags {
			tag, err := client.GetTag(ct.ID)
			require.NoError(t, err)
			assert.Equal(t, ct.Name, tag.Name)
			assert.Equal(t, ct.Slug, tag.Slug)
			assert.Equal(t, ct.Color, tag.Color)
			assert.Equal(t, ct.Description, tag.Description)
		}

		// Update a tag
		updateInput := &client.UpdateTagInput{
			ID:          createdTags[0].ID,
			Name:        "Updated Production",
			Slug:        "updated-production",
			Color:       "ff0000",
			Description: "Updated production environment",
			ObjectTypes: []string{"dcim.site"},
		}
		updatedTag, err := client.UpdateTag(updateInput)
		require.NoError(t, err)
		assert.Equal(t, updateInput.Name, updatedTag.Name)
		assert.Equal(t, updateInput.Slug, updatedTag.Slug)
		assert.Equal(t, updateInput.Color, updatedTag.Color)
		assert.Equal(t, updateInput.Description, updatedTag.Description)
		assert.Equal(t, updateInput.ObjectTypes, updatedTag.ObjectTypes)

		// Patch a tag
		patchInput := &client.PatchTagInput{
			ID:          createdTags[1].ID,
			Description: strPtr("Patched development environment"),
		}
		patchedTag, err := client.PatchTag(patchInput)
		require.NoError(t, err)
		assert.Equal(t, *patchInput.Description, patchedTag.Description)
	})
}

func strPtr(s string) *string {
	return &s
}
