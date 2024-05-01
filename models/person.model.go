package models

import "time"

// note : paradigma TYPE STRUCT can be based on FIELD DB or CERTAIN METHOD
//

type Person struct {
	ID            int        `json:"id_person"`
	FirstName     *string    `json:"first_name"`
	LastName      *string    `json:"last_name"`
	Address       *string    `json:"address"`
	PhoneNumber   *string    `json:"phone_number"`
	Job           *string    `json:"job"`
	InstagramName *string    `json:"instagram_uname"`
	FBName        *string    `json:"facebook_uname"`
	PostalCode    *string    `json:"postal_code"`
	CreatedAt     *time.Time `json:"created_at"`
	IDUser        *int       `json:"id_user"`
}

type CreatePerson struct{}
