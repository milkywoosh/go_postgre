package models

import "time"

type School struct {
	ID          uint
	NameSchool  string
	Address     string
	CreatedAt   *time.Time
	EmailSchool string
}
