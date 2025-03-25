package models

type Customers struct {
	CustID       int             `json:"id"`
	CustName     string          `json:"cust_name"`
	Email        string          `json:"email"`
	Phone        string          `json:"phone"`
	CreatedAt    YYYMMDDCustTime `json:"created_at"`
	CustRegistry string          `json:"cust_registry"`
}
