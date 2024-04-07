package models

type Teacher struct {
	IdTeacher   int    `json:"id_teacher,omitempty"`
	NameTeacher string `json:"name_teacher,omitempty"`
	IdSubject   int    `json:"id_subject,omitempty"`
	Email       string `json:"data_email,omitempty"`
	IdPerson    int    `json:"id_person,omitempty"`
}
