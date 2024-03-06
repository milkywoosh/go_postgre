package models

// note : paradigma TYPE STRUCT can be based on FIELD or CERTAIN METHOD
//

type Person struct {
	ID       uint      `json:"id,omitempty"`
	Name     string    `json:"name,omitempty"`
	SchoolID uint      `json:"school_id,omitempty"`
	Subjects []Subject `json:"subjects,omitempty"`
	School   *School   `json:"school,omitempty"`
}

type CreatePerson struct{}
