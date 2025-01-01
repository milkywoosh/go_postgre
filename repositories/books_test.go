package repositories

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFetchBookByID(t *testing.T) {
	// arg := 2

	data, err := BooksRepoTest.FetchBookByID(context.Background(), 2)

	require.NoError(t, err, "fetch get by id not error")
	require.NotEmpty(t, data, "data fetch not empty")

}
