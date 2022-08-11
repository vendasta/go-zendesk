package zendesk

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPermissionGroupsCrud(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}

	client, err := NewEnvClient()
	require.NoError(t, err)

	permissionGroup := &PermissionGroup{
		Name: String("My Permission Group"),
	}

	created, err := client.CreatePermissionGroup(permissionGroup)
	require.NoError(t, err)
	require.NotNil(t, created.ID)
	require.Equal(t, permissionGroup.Name, created.Name)

	found, err := client.ShowPermissionGroup(*created.ID)
	require.NoError(t, err)
	require.Equal(t, permissionGroup.Name, found.Name)

	input := &PermissionGroup{
		Name: String("My Updated Permission Group"),
	}

	updated, err := client.UpdatePermissionGroup(*created.ID, input)
	require.NoError(t, err)
	require.Equal(t, input.Name, updated.Name)

	err = client.DeletePermissionGroup(*created.ID)
	require.NoError(t, err)
}
