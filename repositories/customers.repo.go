package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/milkyway/gin_beginer/models"
)

// implement interface => see UML :
// https://app.diagrams.net/#G1Zu9rVt9MbPV6yb-oRYg80yfitPyo514e#%7B%22pageId%22%3A%2221v3kVgxeSxbH-oSVx_g%22%7D

type CustomersRepo struct {
	DB *sql.DB
}

func NewCustomersRepo(arg_db *sql.DB) CustomersRepo {
	return CustomersRepo{
		DB: arg_db,
	}
}

func (cr CustomersRepo) FetchByRegistry(ctx context.Context, cust_reg string) (models.Customers, error) {

	var cust_info models.Customers
	var row *sql.Row
	var err error = nil

	query_get_by_registry := `SELECT c.id, c.cust_name, c.email, c.phone, c.created_at, c.cust_registry FROM customers c where c.cust_registry = $1`

	// ERROR : cust_regb ==> doenst existed padahal existed

	row = cr.DB.QueryRowContext(ctx, query_get_by_registry, cust_reg)

	// if err = row.Scan(&cust_info.CustID, &cust_info.CustName, &cust_info.Email); err != nil {
	if err = row.Scan(
		&cust_info.CustID,
		&cust_info.CustName,
		&cust_info.Email,
		&cust_info.Phone,
		&cust_info.CreatedAt,
		&cust_info.CustRegistry,
	); err != nil {
		if err == sql.ErrNoRows {
			return cust_info, fmt.Errorf("data is not found")
		}
		return cust_info, err
	}

	if err = row.Err(); err != nil {
		return cust_info, err
	}

	return cust_info, nil
}
