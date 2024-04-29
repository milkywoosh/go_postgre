package models

import "time"

// note : paradigma TYPE STRUCT can be based on FIELD or CERTAIN METHOD
//

type Person struct {
	ID            int        `json:"id_person,omitempty"`
	FirstName     string     `json:"first_name,omitempty"`
	LastName      string     `json:"last_name,omitempty"`
	Address       string     `json:"address,omitempty"`
	PhoneNumber   string     `json:"phone_number,omitempty"`
	Job           string     `json:"job,omitempty"`
	InstagramName string     `json:"instagram_uname,omitempty"`
	FBName        string     `json:"facebook_uname,omitempty"`
	PostalCode    string     `json:"postal_code,omitempty"`
	CreatedAt     *time.Time `json:"created_at,omitempty"`
	IDUser        int        `json:"id_user,omitempty"`
}

type CreatePerson struct{}
