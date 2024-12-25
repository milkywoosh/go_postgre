package models

type Books struct {
	ID       int     `json:"id"`
	BookName string  `json:"book_name"`
	Price    float64 `json:"price"`
	StockQty int     `json:"stock_qty"`
}
