package repositories

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/milkyway/gin_beginer/models"
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

func (br BooksRepo) FetchBooksLikeName(name_like string, ctx *gin.Context) ([]models.BooksLike, error) {

	query_books_like := fmt.Sprintf(`SELECT title, price FROM books b WHERE title LIKE %s`, "'%'||$1||'%'") // '%'||$1||'%'
	fmt.Println(query_books_like)
	var rows *sql.Rows
	var err error
	var eachResult models.BooksLike
	var searchResults []models.BooksLike = make([]models.BooksLike, 0)

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

func (br BooksRepo) FetchBooksByAuthor(author_name_like string, ctx *gin.Context) ([]models.BooksAuthorLike, error) {

	query_books_like := fmt.Sprintf(`
	SELECT a.author_name, b.title, DATE(a.created_at) FROM books b 
      INNER JOIN authors a ON a.id = b.author_id 
      WHERE author_name LIKE %s`, "'%'||$1||'%'") // '%'||$1||'%'

	var rows *sql.Rows
	var err error
	var eachResult models.BooksAuthorLike
	var searchResults []models.BooksAuthorLike = make([]models.BooksAuthorLike, 0)

	rows, err = br.DB.QueryContext(ctx, query_books_like, author_name_like)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		if err = rows.Scan(&eachResult.AuthorName, &eachResult.Title, &eachResult.CreatedAt); err != nil {
			return nil, err
		}

		searchResults = append(searchResults, eachResult)
	}

	return searchResults, nil
}
