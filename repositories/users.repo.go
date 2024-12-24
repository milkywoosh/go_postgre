package repositories

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
)

type UsersRepo struct {
	DB *sql.DB
}

func (ur UsersRepo) FetchUsernamePassword(ctx *gin.Context, username string) (string, string, error) {

	// tampungan hash password fetch from DB
	var username_val string
	var hash_password string
	var err error

	query_get_user := `SELECT 
		u.username, 
		u.password
	FROM users u WHERE u.username = $1 LIMIT 1`

	var row *sql.Row = ur.DB.QueryRowContext(ctx, query_get_user, username)

	// scan: tampungan data fetch from DB
	err = row.Scan(&username_val, &hash_password)
	if err != nil {
		log.Fatal("err fetchUsernamePassword", err.Error())
		return "", "", err
	}

	return username_val, hash_password, nil
}
