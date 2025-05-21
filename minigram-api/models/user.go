package models

import (
	"minigram-api/helpers"
	"time"

	"github.com/asaskevich/govalidator"
)

type User struct {
	Id          uint       `json:"id"`
	Username    string     `json:"username" valid:"required~Username is Required"`
	FullName    string     `json:"full_name" valid:"required~Name is Required"`
	Email       string     `json:"email" valid:"required~Email is Required, email~Invalid format Email"`
	Password    string     `json:"password" valid:"required~Username is Required, minstringlength(6)~Password minimum 6 Character"`
	Avatar      string     `json:"avatar"`
	Token       string     `json:"token"`
	CreatedDate *time.Time `json:"created_date"`
	UpdatedDate *time.Time `json:"updated_date"`
}

func (u *User) BeforeCreate() (res bool, err error) {
	res, err = govalidator.ValidateStruct(u)

	u.Password = helpers.HashPassword(u.Password)
	u.Token, err = helpers.GenerateJWT(u.Username, u.Email)

	return
}
