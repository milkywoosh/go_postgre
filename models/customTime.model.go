package models

import (
	"encoding/json"
	"time"
)

type YYYMMDDCustTime time.Time

func (y YYYMMDDCustTime) MarshalJSON() ([]byte, error) {
	formattedTime := time.Time(y).Format("2006-01-02")
	result, err := json.Marshal(formattedTime)
	if err != nil {
		return nil, err
	}
	return result, nil
}
