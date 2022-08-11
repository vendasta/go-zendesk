package zendesk

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSectionCrud(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}

	client, err := NewEnvClient()
	require.NoError(t, err)

	category := &Category{
		Name: String("My category"),
	}
	createdCategory, err := client.CreateCategory(category)
	require.NoError(t, err)
	require.NotNil(t, createdCategory.ID)

	section := &Section{
		Name:       String("My Section"),
		CategoryID: createdCategory.ID,
	}

	created, err := client.CreateSection(*createdCategory.ID, section)
	require.NoError(t, err)
	require.NotNil(t, created.ID)
	require.Equal(t, section.Name, created.Name)

	found, err := client.ShowSection(*created.ID)
	require.NoError(t, err)
	require.Equal(t, section.Name, found.Name)

	input := &Section{
		Position: Int(2),
	}

	updated, err := client.UpdateSection(*created.ID, input)
	require.NoError(t, err)
	require.Equal(t, input.Position, updated.Position)

	err = client.DeleteSection(*created.ID)
	require.NoError(t, err)
}
