package zendesk

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSearchTickets(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}

	client, err := NewEnvClient()
	require.NoError(t, err)

	options := ListOptions{
		Page:      0,
		SortBy:    "",
		SortOrder: "",
	}
	include := []SideLoad{IncludeGroups()}

	status := StatusFilter(Status("OPEN"), Equality)

	_, err = client.SearchTickets("", &options, include, status)
	require.NoError(t, err)
}
