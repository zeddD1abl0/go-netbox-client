package integration_tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zeddD1abl0/go-netbox-client/client/extras"
	"github.com/zeddD1abl0/go-netbox-client/models"
)

func TestTagIntegration(t *testing.T) {
	_, service := setupTestClient(t)
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
			input := &extras.CreateTagInput{
				Name:        tt.name,
				Slug:        tt.slug,
				Color:       tt.color,
				Description: tt.description,
				ObjectTypes: tt.objectTypes,
			}
			created, err := service.CreateTag(input)
			require.NoError(t, err)
			require.NotNil(t, created)
			assert.Equal(t, tt.name, created.Name)
			assert.Equal(t, tt.slug, created.Slug)
			assert.Equal(t, tt.color, created.Color)
			assert.Equal(t, tt.description, created.Description)
			createdTags = append(createdTags, created)
		}

		// List and verify tags
		listInput := &extras.ListTagsInput{
			Name: "prod",
		}
		list, err := service.ListTags(listInput)
		require.NoError(t, err)
		require.NotNil(t, list)
		assert.Equal(t, 1, len(list))
		assert.Equal(t, "Production", list[0].Name)

		// Update a tag
		updateInput := &extras.UpdateTagInput{
			ID:          createdTags[0].ID,
			Name:        "Production-Updated",
			Slug:        "production-updated",
			Color:       "ff00ff",
			Description: "Updated production environment",
			ObjectTypes: []string{"dcim.site"},
		}
		updated, err := service.UpdateTag(updateInput)
		require.NoError(t, err)
		require.NotNil(t, updated)
		assert.Equal(t, updateInput.Name, updated.Name)
		assert.Equal(t, updateInput.Slug, updated.Slug)
		assert.Equal(t, updateInput.Color, updated.Color)
		assert.Equal(t, updateInput.Description, updated.Description)

		// Get and verify the updated tag
		retrieved, err := service.GetTag(createdTags[0].ID)
		require.NoError(t, err)
		require.NotNil(t, retrieved)
		assert.Equal(t, updateInput.Name, retrieved.Name)
		assert.Equal(t, updateInput.Slug, retrieved.Slug)
		assert.Equal(t, updateInput.Color, retrieved.Color)
		assert.Equal(t, updateInput.Description, retrieved.Description)

		// Patch a tag
		newName := "Production-Patched"
		patchInput := &extras.PatchTagInput{
			ID:   createdTags[0].ID,
			Name: &newName,
		}
		patched, err := service.PatchTag(patchInput)
		require.NoError(t, err)
		require.NotNil(t, patched)
		assert.Equal(t, newName, patched.Name)
		assert.Equal(t, updateInput.Slug, patched.Slug)   // Should remain unchanged
		assert.Equal(t, updateInput.Color, patched.Color) // Should remain unchanged

		// Clean up - delete all created tags
		for _, tag := range createdTags {
			err := service.DeleteTag(tag.ID)
			assert.NoError(t, err)

			// Verify deletion
			_, err = service.GetTag(tag.ID)
			assert.Error(t, err)
			assert.Contains(t, err.Error(), "tag not found")
		}
	})

	t.Run("Error cases", func(t *testing.T) {
		// Test creating a tag with invalid input
		input := &extras.CreateTagInput{
			Name:  "", // Invalid: empty name
			Slug:  "test",
			Color: "ff0000",
		}
		_, err := service.CreateTag(input)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "validation failed")

		// Test getting a non-existent tag
		_, err = service.GetTag(999999)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "tag not found")

		// Test updating a non-existent tag
		updateInput := &extras.UpdateTagInput{
			ID:    999999,
			Name:  "Test",
			Slug:  "test",
			Color: "ff0000",
		}
		_, err = service.UpdateTag(updateInput)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "tag not found")

		// Test deleting a non-existent tag
		err = service.DeleteTag(999999)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "tag not found")
	})
}
