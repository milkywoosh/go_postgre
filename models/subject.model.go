package models

// note : paradigma TYPE STRUCT can be based on FIELD or CERTAIN METHOD
//
type Subject struct {
	IdSubject   uint   `json:"id_subject,omitempty"`
	SubjectName string `json:"subject_name,omitempty"`
	IdPeople    uint   `json:"id_people_subject,omitempty"`
}

type SubjectInfo struct {
	Subject
	Person
}
