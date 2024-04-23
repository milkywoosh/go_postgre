package models

// note : paradigma TYPE STRUCT can be based on FIELD or CERTAIN METHOD
type Subject struct {
	IdSubject   *int    `json:"id_subject,omitempty"`
	SubjectName *string `json:"subject_name,omitempty"`
	IdPeople    *int    `json:"id_people_subject,omitempty"`
}

type SubjectInfo struct {
	IdSubject   *int    `json:"id_subject"`
	NamePerson  *string `json:"name_person"`
	SubjectName *string `json:"subject_name"`
}
