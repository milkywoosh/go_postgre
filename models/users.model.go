package models

type Users struct {
	ID              int    `json:"id"`
	Username        string `json:"username"`
	Email           string `json:"email"`
	FirstName       string `json:"firstname"`
	LastName        string `json:"lastname"`
	Password        string `json:"password"`
	PasswordHistory string `json:"passwordhistory"`
}

type UserRoles struct {
	RoleID int `json:"roleid"`
	UserID int `json:"userid"`
}

type Roles struct {
	ID       int    `json:"id"`
	RoleName string `json:"rolename"`
}
