package repositories

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFetchPasswordByUsername(t *testing.T) {
	var username string
	var pw string
	var err error = nil
	username, pw, err = UsersRepoTest.FetchPasswordByUsername(context.Background(), "benten")

	require.NoError(t, err, "error check FetchPasswordByUsername")
	require.Equal(t, "benten", username)
	require.Equal(t, "$2a$10$LFuc3at2oqRp6mH3oHooIObTPWEAcoIP3e7EYJeAC3jCCylNSSoBK", pw)

}
