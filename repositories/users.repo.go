package repositories

import (
	"context"
	"database/sql"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/milkyway/gin_beginer/models"
)

type UsersRepo struct {
	DB *sql.DB
}

// for testing repo
func NewUsersRepo(arg_db *sql.DB) *UsersRepo {
	return &UsersRepo{
		DB: arg_db,
	}
}

func (ur UsersRepo) FetchUsernamePassword(ctx context.Context, username string) (string, string, error) {
	return "data1", "data2", errors.New("Error fetch")
}

func (ur UsersRepo) FetchPasswordByUsername(ctx context.Context, username string) (string, string, error) {
	// tampungan hash password fetch frsom DB
	var username_val string
	var hash_password string
	var err error = nil

	query_get_user := `SELECT 
		u.username, 
		u.password
	FROM users u WHERE u.username = $1 LIMIT 1`

	var row *sql.Row = ur.DB.QueryRowContext(ctx, query_get_user, username)

	// scan: tampungan data fetch from DB
	err = row.Scan(&username_val, &hash_password)
	// if err != nil {
	// 	log.Fatal("err fetchPasswordByUsername", err.Error())
	// 	return "", "", err
	// }

	return username_val, hash_password, err

}

func (ur UsersRepo) InsertNewUser(ctx context.Context, UsersModel models.Users, hash_pass string) error {
	var tx *sql.Tx
	var err error = nil

	tx, err = ur.DB.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return err
	}

	insertUsersQry := `insert into users(username, email, password, firstname, lastname, password_history_count) values($1, $2, $3, $4, $5, $6)`
	// test case duplicate entry username : success
	_, err = tx.ExecContext(
		ctx,
		insertUsersQry,
		&UsersModel.Username,
		&UsersModel.Email,
		hash_pass,
		&UsersModel.FirstName,
		&UsersModel.LastName,
		0,
	)

	if err != nil {
		// rollback
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return rollbackErr
		}
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil

}

func (ur UsersRepo) FetchUserByID(ctx *gin.Context, user_id int) (string, error) {
	get_by_id_query := `select username from users u where u.id = $1`

	var username string
	var err error = nil
	var rows *sql.Row = ur.DB.QueryRowContext(ctx, get_by_id_query, user_id)
	if err = rows.Scan(&username); err != nil {
		return "", err
	}
	return username, nil
}
