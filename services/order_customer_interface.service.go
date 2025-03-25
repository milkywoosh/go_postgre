package services

import "github.com/milkyway/gin_beginer/models"

// get from UI
type BookOrderInfo struct {
	ID    string // string or int ??
	Title string
	Qty   int
}

type OrderCustomer interface {
	BuyBooks(models.Books, []BookOrderInfo) bool
}
