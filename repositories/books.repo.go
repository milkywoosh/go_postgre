package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/milkyway/gin_beginer/models"
)

type BooksRepo struct {
	DB *sql.DB
}

// for testing repo
func NewBooksRepo(arg_db *sql.DB) *BooksRepo {
	return &BooksRepo{
		DB: arg_db,
	}
}

func (br BooksRepo) FetchBookByID(ctx context.Context, book_id int) (models.Books, error) {

	var book_info models.Books
	var row *sql.Row
	var err error = nil

	query_book_get_by_id := `SELECT id, title, stock_qty, price FROM books b where b.id = $1`
	row = br.DB.QueryRowContext(ctx, query_book_get_by_id, book_id)

	if err = row.Scan(
		&book_info.ID,
		&book_info.BookName,
		&book_info.StockQty,
		&book_info.Price,
	); err != nil {
		if err == sql.ErrNoRows {
			return book_info, fmt.Errorf("data is not found")
		}
		return book_info, err
	}

	if err = row.Err(); err != nil {
		return book_info, err
	}

	return book_info, nil
}

func (br BooksRepo) BookIsExistedByID(ctx context.Context, book_id int) (int, error) {

	var book_info models.Books
	var row *sql.Row
	var err error = nil

	query_book_get_by_id := `SELECT id FROM books b where b.id = $1`
	row = br.DB.QueryRowContext(ctx, query_book_get_by_id, book_id)

	// if row kosong

	if err = row.Scan(&book_info.ID); err != nil {
		log.Printf("=========>  %v ==> %v", err, err == sql.ErrNoRows)
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("data is not found")
		}
		return book_info.ID, err
	}

	if err == sql.ErrNoRows {
		return 0, err
	}
	if err = row.Err(); err != nil {
		return book_info.ID, err
	}

	return book_info.ID, nil
}

func (br BooksRepo) FetchBooksLikeName(ctx context.Context, name_like string) ([]models.BooksLike, error) {

	query_books_like := fmt.Sprintf(`SELECT title, price FROM books b WHERE title LIKE %s`, "'%'||$1||'%'") // '%'||$1||'%'
	var rows *sql.Rows
	var err error = nil
	var eachResult models.BooksLike
	var searchResults []models.BooksLike = make([]models.BooksLike, 0)

	rows, err = br.DB.QueryContext(ctx, query_books_like, name_like)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&eachResult.Title, &eachResult.Price); err != nil {
			return nil, err
		}
		searchResults = append(searchResults, eachResult)
	}
	if err = rows.Close(); err != nil {
		return nil, err
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return searchResults, nil
}

func (br BooksRepo) FetchBooksByAuthor(ctx context.Context, author_name_like string) ([]models.BooksAuthorLike, error) {

	query_books_like := fmt.Sprintf(`
	SELECT a.author_name, b.title, DATE(a.created_at) FROM books b 
      INNER JOIN authors a ON a.id = b.author_id 
      WHERE author_name LIKE %s`, "'%'||$1||'%'") // '%'||$1||'%'

	var rows *sql.Rows
	var err error = nil
	var eachResult models.BooksAuthorLike
	var searchResults []models.BooksAuthorLike = make([]models.BooksAuthorLike, 0)

	rows, err = br.DB.QueryContext(ctx, query_books_like, author_name_like)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&eachResult.AuthorName, &eachResult.Title, &eachResult.CreatedAt); err != nil {
			return nil, err
		}

		searchResults = append(searchResults, eachResult)
	}
	if err = rows.Close(); err != nil {
		return nil, err
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return searchResults, nil
}

type ListBulkyBookParams [][]string

func (br BooksRepo) InsertManyBooks(ctx context.Context, arg ListBulkyBookParams) error {
	var tx *sql.Tx
	var err error = nil
	var opts *sql.TxOptions = &sql.TxOptions{
		Isolation: sql.LevelReadCommitted, // Choose an appropriate isolation level
		ReadOnly:  false,                  // Set to true if this is a read-only operation
	}

	tx, err = br.DB.BeginTx(ctx, opts)
	if err != nil {
		return err
	}

	query_insert := "INSERT INTO books (title, price, stock_qty) VALUES ($1, $2, $3)"
	stmt, err := tx.PrepareContext(ctx, query_insert)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer stmt.Close()

	for i, row := range arg {
		if i == 0 {
			continue
		}
		title := row[0]
		price := row[1]
		stock_qty := row[2]

		if _, err := stmt.ExecContext(ctx, title, price, stock_qty); err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to execute statement: %w", err)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (br BooksRepo) InsertManyBooksConcurrent(ctx context.Context, arg ListBulkyBookParams) error {

	// br.DB.BeginTx(ctx)
	// br.DB.ExecContext()

	return fmt.Errorf("not yet implemented")
}
