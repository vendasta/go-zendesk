package zendesk

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestArticleCrud(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}

	client, err := NewEnvClient()
	require.NoError(t, err)

	category := &Category{
		Name: String("My Category"),
	}

	createCategory, err := client.CreateCategory(category)
	require.NoError(t, err)
	require.NotNil(t, createCategory.ID)

	section := &Section{
		Name: String("My Section"),
	}

	createdSection, err := client.CreateSection(*createCategory.ID, section)
	require.NoError(t, err)
	require.NotNil(t, createdSection.ID)

	permissionGroup := &PermissionGroup{
		Name: String("My permission group"),
	}
	createdPermissionGroup, err := client.CreatePermissionGroup(permissionGroup)
	require.NoError(t, err)
	require.NotNil(t, createdPermissionGroup.ID)

	article := &Article{
		Body:              String("My very good article."),
		Title:             String("A good article"),
		PermissionGroupID: createdPermissionGroup.ID,
	}

	created, err := client.CreateArticle(*createdSection.ID, article)
	require.NoError(t, err)
	require.NotNil(t, created.ID)
	require.Equal(t, article.Title, created.Title)

	found, err := client.ShowArticle(*created.ID)
	require.NoError(t, err)
	require.Equal(t, article.Title, found.Title)

	input := &Article{
		Promoted: Bool(true),
	}

	updated, err := client.UpdateArticle(*created.ID, input)
	require.NoError(t, err)
	require.Equal(t, input.Promoted, updated.Promoted)

	err = client.DeleteArticle(*created.ID)
	require.NoError(t, err)

	err = client.DeleteSection(*createdSection.ID)
	require.NoError(t, err)

	err = client.DeleteCategory(*createCategory.ID)
	require.NoError(t, err)

	err = client.DeletePermissionGroup(*createdPermissionGroup.ID)
	require.NoError(t, err)
}
