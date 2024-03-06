package models

import "time"

// note : paradigma TYPE STRUCT can be based on FIELD or CERTAIN METHOD
//
type School struct {
	ID          uint
	NameSchool  string
	Address     string
	CreatedAt   *time.Time
	EmailSchool string
}
