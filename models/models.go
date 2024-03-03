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
	ID          uint
	NameSchool  string
	Address     string
	CreatedAt   *time.Time
	EmailSchool string
}

type People struct {
	ID       uint      `json:"id,omitempty"`
	Name     string    `json:"name,omitempty"`
	SchoolID uint      `json:"school_id,omitempty"`
	Subjects []Subject `json:"subjects,omitempty"`
	School   *School   `json:"school,omitempty"`
}

type Subject struct {
	IdSubject   uint   `json:"id_subject,omitempty"`
	SubjectName string `json:"subject_name,omitempty"`
	IdPeople    uint   `json:"id_people_subject,omitempty"`
}
type Teacher struct {
	IdTeacher    uint   `json:"id_teacher,omitempty"`
	NameTeacher  string `json:"name_teacher,omitempty"`
	IdSubject    uint   `json:"id_subject_teacher,omitempty"`
	EmailTeacher string `json:"email,omitempty"`
	IdPeople     uint   `json:"id_people_teacher,omitempty"`
}

// NOTE: saat buat model untuk JOIN TABLE dipastikan TIDAK ADA [FIELD atau JSONFIELD] YANG SAMA pada tiap STRUCT
type CompleteData struct {
	Teacher
	Subject
	People
}
