package models

import (
	"encoding/json"
	"strings"
)

type Books struct {
	ID       int     `json:"id"`
	BookName string  `json:"book_name"`
	Price    float64 `json:"price"`
	StockQty int     `json:"stock_qty"`
}

type Author struct {
	ID                 int      `json:"id"`
	AuthorName         string   `json:"author_name"`
	SerialAuthorNumber string   `json:"serial_author_number"`
	CreatedAt          JSONTime `json:"created_at"`
}

type Scanner interface {
	Scan(value interface{}) error
}

type TitleType string
type PriceType float64
type AuthorType string

type BooksAuthorLike struct {
	Title      string   `json:"title"`
	AuthorName string   `json:"author_name"`
	CreatedAt  JSONTime `json:"author_registered_at"`
}

type BooksLike struct {
	Title string  `json:"title"`
	Price float64 `json:"price"`
}

// note , sebelumnya add TitleType di dalam struct, kemudian digunakan as method struct MarshalJSON => overflow error!
// solusi pakai type native "string" di dalam struct agar punya `json:"any_name"`!
func (typ *BooksAuthorLike) MarshalJSON() ([]byte, error) {
	// implement MarshalJSON()

	// log.Println(1, reflect.TypeOf(typ))
	// log.Println(2, reflect.TypeOf(*typ))

	if strings.Contains(string(typ.Title), "Plague") {

		typ.Title = "censored title"

		return json.Marshal(*typ) // note ???? harus pake asterisk
	} else {
		return json.Marshal(*typ)
	}

}

// Value implements the driver.Valuer interface for inserting/updating database values
// func (t JSONTime) Value() (driver.Value, error) {
// 	return t.Time, nil
// }

// UnmarshalJSON implements the json.Unmarshaler interface
// func (t *JSONTime) UnmarshalJSON(data []byte) error {
// 	// Parse the JSON date
// 	parsedTime, err := time.Parse("\"2006-01-02\"", string(data))
// 	if err != nil {
// 		return err
// 	}
// 	t.Time = parsedTime
// 	return nil
// }
