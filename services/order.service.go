package services

import (
	"github.com/milkyway/gin_beginer/repositories"
)

type OrdersService struct {
	BooksRepo     repositories.BooksRepo
	CustomersRepo repositories.CustomersRepo
}

func NewOrdersService(booksRepo repositories.BooksRepo, custRepo repositories.CustomersRepo) OrdersService {
	return OrdersService{
		BooksRepo:     booksRepo,
		CustomersRepo: custRepo,
	}
}

// get from UI
type OrderBookRequest struct {
	ID    string `json:"id" binding:"required"` // string or int ??
	Title string `json:"title" binding:"required"`
	Qty   int    `json:"qty" binding:"required"`
}

type OrderCustomer interface {
	BuyBooks([]OrderBookRequest) bool
	SetDiscount(float64) float64
}

func (o OrdersService) BuyBooks([]OrderBookRequest) bool {
	return true
}
func (o OrdersService) SetDiscount(float64) float64 {
	panic("not implemented yet!")
}
