package models

type PurchaseItems struct {
	ID                int     `json:"id"`
	BookID            int     `json:"book_id"`
	PurchaseHistoryID int     `json:"purchase_history_id"`
	Qty               int     `json:"Qty"`
	TotalPrice        float64 `json:"total_price"` // total price based on how many same book is bought
}

type PurchaseHistories struct {
	ID                int      `json:"id"`
	CustomerID        int      `json:"customer_id"`
	DateOfSale        JSONTime `json:"date_of_sale"`
	Status            string   `json:"status"`
	TotalPricePayment float64  `json:"total_price_payment"`
}
