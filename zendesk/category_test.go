package zendesk

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCategoryCrud(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}

	client, err := NewEnvClient()
	require.NoError(t, err)

	category := &Category{
		Name: String("My Category"),
	}

	created, err := client.CreateCategory(category)
	require.NoError(t, err)
	require.NotNil(t, created.ID)
	require.Equal(t, category.Name, created.Name)

	found, err := client.ShowCategory(*created.ID)
	require.NoError(t, err)
	require.Equal(t, category.Name, found.Name)

	input := &Category{
		Position: Int(2),
	}

	updated, err := client.UpdateCategory(*created.ID, input)
	require.NoError(t, err)
	require.Equal(t, input.Position, updated.Position)

	err = client.DeleteCategory(*created.ID)
	require.NoError(t, err)
}
