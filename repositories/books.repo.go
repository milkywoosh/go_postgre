package repositories

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

type BooksRepo struct {
	DB *sql.DB
}

func (br BooksRepo) FetchBookByID(ctx *gin.Context, book_id string) (string, error) {
	var book_name string
	var rows *sql.Row

	query_book_get_by_id := `SELECT title FROM books b where b.id = $1`

	rows = br.DB.QueryRowContext(ctx, query_book_get_by_id, book_id)
	if err := rows.Scan(&book_name); err != nil {
		return "", err
	}

	return book_name, nil
}
