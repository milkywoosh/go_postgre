package models

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type Books struct {
	ID       int     `json:"id"`
	BookName string  `json:"book_name"`
	Price    float64 `json:"price"`
	StockQty int     `json:"stock_qty"`
}

type JSONTime struct {
	Time time.Time
}
type Author struct {
	ID                 int      `json:"id"`
	AuthorName         string   `json:"author_name"`
	SerialAuthorNumber string   `json:"serial_author_number"`
	CreatedAt          JSONTime `json:"created_at"`
}

type Marshaler interface {
	MarshalJSON() ([]byte, error)
}

func (t JSONTime) MarshalJSON() ([]byte, error) {
	// Format the time as desired
	// fmt.Println("kapan dipanggil boss?")
	// stamp := fmt.Sprintf("\"%s\"", t.Time.Format("2006-01-02")) // date format golang
	// return []byte(stamp), nil

	tm := JSONTime{
		Time: t.Time, // or you can use a specific time
	}

	return json.Marshal(tm.Time.Format("2006-01-02"))
}

type Scanner interface {
	Scan(value interface{}) error
}

func (t *JSONTime) Scan(value interface{}) error {
	switch v := value.(type) {
	case time.Time:
		t.Time = v
		return nil
	case []byte:
		parsedTime, err := time.Parse(time.RFC3339, string(v))
		if err != nil {
			return err
		}
		t.Time = parsedTime
		return nil
	case string:
		parsedTime, err := time.Parse(time.RFC3339, v)
		if err != nil {
			return err
		}
		t.Time = parsedTime
		return nil
	default:
		return fmt.Errorf("unsupported type: %T", v)
	}
}

type TitleType string
type PriceType float64
type AuthorType string

type BooksAuthorLike struct {
	Title      TitleType `json:"title"`
	AuthorName string    `json:"author_name"`
	CreatedAt  JSONTime  `json:"author_registered_at"`
}

type BooksLike struct {
	Title TitleType `json:"title"`
	Price float64   `json:"price"`
}

func (typ TitleType) MarshalJSON() ([]byte, error) {
	var filtered string
	var non_filtered string

	if strings.Contains(string(typ), "Plague") {
		filtered = "Cencored Title"
		return json.Marshal(filtered)
	} else {
		non_filtered = fmt.Sprintf("add %s", typ)
		return json.Marshal(non_filtered)
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
