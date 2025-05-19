package models

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	FullName string `json:"full_name"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"`
	Token    string `json:"token"`
}
