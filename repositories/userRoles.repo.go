package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type UserRolesRepo struct {
	DB *sql.DB
}

func NewUserRolesRepo(arg_db *sql.DB) UserRolesRepo {
	return UserRolesRepo{
		DB: arg_db,
	}
}

type RoleStruct struct {
	Username string `json:"username"`
	RoleName string `json:"rolename"`
}

func (ur UserRolesRepo) FetchRoleByID(user_id int, ctx *gin.Context) ([]RoleStruct, error) {

	get_by_id_query := `
		SELECT r.role_name, u.username from user_roles ur
		INNER JOIN users u on u.id = ur.user_id
		INNER JOIN roles r on r.id = ur.role_id
		where u.id = $1 
	`
	var EachRoleData RoleStruct
	var AllRolesData []RoleStruct

	rows, err := ur.DB.QueryContext(ctx, get_by_id_query, user_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(
			&EachRoleData.RoleName,
			&EachRoleData.Username,
		); err != nil {
			// Check for a scan error.
			// Query rows will be closed with defer.

			return nil, err
		}

		AllRolesData = append(AllRolesData, EachRoleData)

	}

	rerr := rows.Close()
	if rerr != nil {
		return nil, err
	}

	// Rows.Err will report the last error encountered by Rows.Scan.
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return AllRolesData, nil

}
func (ur UserRolesRepo) FetchRolesByUsername(ctx context.Context, username string) ([]RoleStruct, error) {

	get_by_id_query := `
		SELECT u.username, r.role_name from user_roles ur
		INNER JOIN users u on u.id = ur.user_id
		INNER JOIN roles r on r.id = ur.role_id
		where u.username = $1 
	`
	var EachRoleData RoleStruct
	var AllRolesData []RoleStruct

	rows, err := ur.DB.QueryContext(ctx, get_by_id_query, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(
			&EachRoleData.Username,
			&EachRoleData.RoleName,
		); err != nil {
			// Check for a scan error.
			// Query rows will be closed with defer.

			return nil, err
		}

		AllRolesData = append(AllRolesData, EachRoleData)
	}

	rerr := rows.Close()
	if rerr != nil {
		return nil, err
	}

	// Rows.Err will report the last error encountered by Rows.Scan.
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return AllRolesData, nil

}

func (ur UserRolesRepo) DeleteUserRole(ctx context.Context, username, rolename string) error {

	var tx *sql.Tx
	var err error = nil
	tx, err = ur.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	var user_id int
	err = tx.QueryRow("select id from users u where u.username = $1", username).Scan(&user_id)

	switch {
	case err == sql.ErrNoRows:
		log.Printf("no username %s", username)
		return err
	case err != nil:
		log.Printf("err %v", err)
		return err
	case !(err != nil):
		log.Printf("err nil check username %v", err)
	default:
		log.Printf("err check def username %v", err)
		return err
	}

	var role_id int
	err = tx.QueryRow("select id from roles r where r.role_name = $1", rolename).Scan(&role_id)

	switch {
	case err == sql.ErrNoRows:
		log.Printf("no role_name %s", rolename)
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			log.Printf("unable to rollback: %v", rollbackErr)
			return rollbackErr
		}
		return err
	case err != nil:
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			log.Printf("unable to rollback: %v", rollbackErr)
			return rollbackErr
		}
		return err
	case !(err != nil):
		log.Printf("err nil check role_name %v", err)
	default:
		log.Printf("err check def role_name %v", err)
		return err
	}

	delete_user_roles := "DELETE FROM user_roles WHERE role_id=$1 AND user_id=$2"

	_, err = tx.ExecContext(ctx, delete_user_roles, role_id, user_id)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			log.Printf("unable to rollback: %v", rollbackErr)
			return rollbackErr
		}
		return err
	}

	err = tx.Commit()
	if err != nil {
		// note : err commit terjadi ketika sudah rollback tapi belum di return ??
		return err
	}

	return nil
}

func (ur UserRolesRepo) AddNewRole(ctx context.Context, username, rolename string) error {

	tx, err := ur.DB.BeginTx(ctx, nil)
	if err != nil {
		log.Printf("err initial trx %v", err.Error())
		return err
	}

	var user_id int
	err = tx.QueryRow("select id from users u where u.username = $1", username).Scan(&user_id)

	switch {
	case err == sql.ErrNoRows:
		err_msg := fmt.Sprintf("%s: failed query fetch by parameter => %s", err.Error(), username)
		err = errors.New(err_msg)
		log.Printf("no username %s", username)
		return err
	case err != nil:
		log.Printf("err %v", err)
		return err
	case !(err != nil):
		log.Printf("err nil check username %v", err)
	default:
		log.Printf("err check def username %v", err)
		return err
	}

	var role_id int
	err = tx.QueryRow("select id from roles r where r.role_name = $1", rolename).Scan(&role_id)

	switch {
	case err == sql.ErrNoRows:
		log.Printf("no role_name %s", rolename)
		err_msg := fmt.Sprintf("%s: failed query fetch by parameter => %s", err.Error(), rolename)
		err = errors.New(err_msg)
		return err
	case err != nil:
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			log.Fatalf("unable to rollback: %v", rollbackErr)
			return rollbackErr
		}
		log.Fatalf("gett err after rollback %v", err.Error())
		return err
	case !(err != nil):
		log.Printf("err nil check role_name %v", err)
	default:
		log.Printf("err check def role_name %v", err)
		return err
	}

	insert_user_roles := "INSERT INTO user_roles (role_id, user_id) VALUES ($1, $2)"

	// var resultInsertUserRoles sql.Result // note : return <nil>, error kalo value di pake => fmt.Println()

	_, err = tx.ExecContext(ctx, insert_user_roles, role_id, user_id)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			log.Fatalf("unable to rollback: %v", rollbackErr)
			return rollbackErr
		}
		log.Println("rollbackErr: ", err.Error())
		return err
	}

	err = tx.Commit()
	if err != nil {
		// note : err commit terjadi ketika sudah rollback tapi belum di return ??
		return err
	}

	return nil
}
