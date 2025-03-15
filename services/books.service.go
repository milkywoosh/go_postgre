package services

// note handle sisi business logic

import (
	"context"
	"io"

	"github.com/milkyway/gin_beginer/models"
	"github.com/milkyway/gin_beginer/repositories"
	xlsx "github.com/milkyway/gin_beginer/utils"
	"github.com/xuri/excelize/v2"
)

// note: layer ini khusus bisnis logic, data received from REPO layer
type BooksService struct {
	BooksRepo repositories.BooksRepo
}

// implement Prototype Interface
func (bs BooksService) Clone() Prototype {
	return &BooksService{
		BooksRepo: bs.BooksRepo,
	}
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

func (bs BooksService) UploadBulkyBooks(ctx context.Context, file_content io.Reader, password string) ([][]string, error) {
	// map[string][][]string
	var opts excelize.Options
	if password == "" {
		password = ""
	}
	opts.Password = password
	data, err := xlsx.OpenReader(file_content, opts)
	if err != nil {
		return nil, err
	}

	// var result map[string][][]string = make(map[string][][]string)
	// result["header"] = [][]string{data[0]}
	// length_data := len(data)
	// result["data"] = data[1:length_data]
	return data, nil
}
