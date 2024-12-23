package models

type Users struct {
	ID              int
	Username        string
	Email           string
	FirstName       string
	LastName        string
	Password        string
	PasswordHistory string
}

type UserRoles struct {
	RoleID int
	UserID int
}

type Roles struct {
	ID       int
	RoleName string
}
