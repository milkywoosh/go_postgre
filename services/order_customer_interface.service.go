package services

import "github.com/milkyway/gin_beginer/models"

// get from UI
type OrderBookRequest struct {
	ID    string `json:"id" binding:"required"` // string or int ??
	Title string `json:"title" binding:"required"`
	Qty   int    `json:"qty" binding:"required"`
}

type OrderCustomer interface {
	BuyBooks(models.Books, []OrderBookRequest) bool
	SetDiscount(float64) float64
}
