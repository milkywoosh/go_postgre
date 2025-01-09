package repositories

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

// type RoleStructTest struct {
// 	Username string `json:"username"`
// 	RoleName string `json:"rolename"`
// }

func TestFetchRolesByUsername(t *testing.T) {
	// var info_roles []RoleStructTest
	var roles []string = []string{}
	var err error = nil
	info_roles, err := UserRolesTest.FetchRolesByUsername(context.Background(), "luke")
	for i := 0; i < len(info_roles); i++ {
		roles = append(roles, info_roles[i].RoleName)
	}

	require.NoError(t, err, "get roles by usernam no error")
	require.Contains(t, roles, "adminn")
	require.Contains(t, roles, "viewer")
	require.Contains(t, roles, "editor")
}
