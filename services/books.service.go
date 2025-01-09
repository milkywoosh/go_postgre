package services

// note handle sisi business logic

import (
	"context"

	"github.com/milkyway/gin_beginer/models"
	"github.com/milkyway/gin_beginer/repositories"
)

// note: layer ini khusus bisnis logic, data received from REPO layer
type BooksService struct {
	BooksRepo repositories.BooksRepo
}

func (bs BooksService) GetBookInfo(ctx context.Context, book_id int) (models.Books, error) {
	var book_info_row models.Books
	var err error = nil

	book_info_row, err = bs.BooksRepo.FetchBookByID(ctx, book_id)
	// note: handle error di sisi service
	if err != nil {
		return book_info_row, err
	}
	return book_info_row, err
}

func (bs BooksService) SearchLikeName(ctx context.Context, name_like string) ([]models.BooksLike, error) {
	var search_info_rows []models.BooksLike
	var err error = nil

	search_info_rows, err = bs.BooksRepo.FetchBooksLikeName(ctx, name_like)
	if err != nil {
		return nil, err
	}
	return search_info_rows, nil
}

func (bs BooksService) SearchBookByAuthor(ctx context.Context, author_name_like string) ([]models.BooksAuthorLike, error) {
	var err error = nil
	var searchResults []models.BooksAuthorLike

	searchResults, err = bs.BooksRepo.FetchBooksByAuthor(ctx, author_name_like)
	if err != nil {
		return nil, err
	}

	return searchResults, nil
}
