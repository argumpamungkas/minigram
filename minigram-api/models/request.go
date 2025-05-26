package models

type RequestLogin struct {
	Username string `json:"username" valid:"required~Username is Required"`
	Password string `json:"password" valid:"required~Password is Required"`
}
