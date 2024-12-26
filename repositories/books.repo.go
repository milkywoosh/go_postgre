package repositories

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
)

type BooksRepo struct {
	DB *sql.DB
}

func (br BooksRepo) FetchBookByID(book_id string, ctx *gin.Context) (string, error) {
	var book_name string
	var rows *sql.Row

	query_book_get_by_id := `SELECT title FROM books b where b.id = $1`

	rows = br.DB.QueryRowContext(ctx, query_book_get_by_id, book_id)
	if err := rows.Scan(&book_name); err != nil {
		return "", err
	}

	return book_name, nil
}

type BooksLike struct {
	Title string `json:"title"`
	Price string `json:"price"`
}

func (br BooksRepo) FetchBooksLikeName(name_like string, ctx *gin.Context) ([]BooksLike, error) {

	query_books_like := fmt.Sprintf(`SELECT title, price FROM books b WHERE title LIKE %s`, "'%'||$1||'%'") // '%'||$1||'%'
	fmt.Println(query_books_like)
	var rows *sql.Rows
	var err error
	var eachResult BooksLike
	var searchResults []BooksLike

	rows, err = br.DB.QueryContext(ctx, query_books_like, name_like)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		if err = rows.Scan(&eachResult.Title, &eachResult.Price); err != nil {
			return nil, err
		}
		searchResults = append(searchResults, eachResult)
	}

	return searchResults, nil
}
