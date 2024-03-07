package models

import "time"

// note : paradigma TYPE STRUCT can be based on FIELD or CERTAIN METHOD
//

type Person struct {
	ID         int       `json:"id,omitempty"`
	NamePerson string    `json:"name_person,omitempty"`
	SchoolID   int       `json:"school_id,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	CreatedBy  int       `json:"created_by,omitempty"`
}

type CreatePerson struct{}
