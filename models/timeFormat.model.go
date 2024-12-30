package models

import (
	"encoding/json"
	"fmt"
	"time"
)

type JSONTime struct {
	Time time.Time
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
