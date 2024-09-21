package models

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	Password string `json:"password"`
	Code     string `json:"code"`
}
