package models

import (
	"time"
)

type User struct {
	Id          uint       `json:"id"`
	Username    string     `json:"username" valid:"required~Username is Required"`
	FullName    string     `json:"full_name" valid:"required~Name is Required"`
	Password    string     `json:"password" valid:"required~Username is Required, minstringlength(6)~Password minimum 6 Character"`
	Avatar      string     `json:"avatar"`
	Token       string     `json:"token"`
	CreatedDate *time.Time `json:"created_date"`
	UpdatedDate *time.Time `json:"updated_date"`
}

// func (u *User) BeforeCreate() string {
// 	log.Println("CALL GAK user Before create")
// 	u.Password = helpers.HashPassword(u.Password)
// 	return u.Password
// }
