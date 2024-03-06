package models

type Subject struct {
	IdSubject   uint   `json:"id_subject,omitempty"`
	SubjectName string `json:"subject_name,omitempty"`
	IdPeople    uint   `json:"id_people_subject,omitempty"`
}
