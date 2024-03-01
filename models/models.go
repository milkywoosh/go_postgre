package models

import "time"

// type People struct {
// 	ID       uint   `gorm:"primaryKey"`
// 	Name     string `gorm:"type:varchar;not null"`
// 	SchoolID uint   `gorm: "not null"`
// 	// School   *School // belongs to relationship
// }

// type School struct {
// 	ID         uint   `gorm:"primaryKey"`
// 	NameSchool string `gorm:"type:varchar;not null"`
// 	Address    string `gorm:"type:varchar;not null"`
// 	CreatedAt  *time.Time
// 	Email      string `gorm:"type:varchar"`
// }

// note !!!
// penggunaan `gorm` atau `json` diperlukan jika akan penyesuaian untuk res API

// note !!
// perlu dibuat interface repository

// type People struct {
// 	ID       uint    `json:"id,omitempty"`
// 	Name     string  `json:"name,omitempty"`
// 	SchoolID uint    `json:"school_id,omitempty"`
// 	School   *School // belongs to relationship
// }

type School struct {
	ID         uint
	NameSchool string
	Address    string
	CreatedAt  *time.Time
	Email      string
}

type People struct {
	ID       uint   `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	SchoolID uint   `json:"school_id,omitempty"`
	Subjects []Subject
	School   *School
}

type Subject struct {
	IdSubject uint   `json:"id_subject,omitempty"`
	Subject   string `json:"subject,omitempty"`
	IdPeople  uint   `json:"id_people,omitempty"`
}
