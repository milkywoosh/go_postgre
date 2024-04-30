package models

type Users struct {
	IDUser   int     `json:"id_user"`
	UserName *string `json:"username"`
	FullName *string `json:"fullname"`
	Password *string `json:"password"`
}
