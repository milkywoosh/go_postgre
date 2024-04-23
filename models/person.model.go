package models

import "time"

// note : paradigma TYPE STRUCT can be based on FIELD or CERTAIN METHOD
//

type Person struct {
	ID         *int       `json:"id_person,omitempty"`
	NamePerson *string    `json:"name_person,omitempty"`
	SchoolID   *int       `json:"school_id,omitempty"`
	CreatedAt  *time.Time `json:"created_at" form:"created_at" time_format:"unix"`
	CreatedBy  *int       `json:"created_by,omitempty"`
}

type CreatePerson struct{}
