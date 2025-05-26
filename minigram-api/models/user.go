package models

import (
	"minigram-api/helpers"
	"time"

	"github.com/asaskevich/govalidator"
)

type User struct {
	GormModel
	Username string `gorm:"column:username" json:"username" valid:"required~Username is Required"`
	FullName string `gorm:"column:full_name" json:"full_name" valid:"required~Name is Required"`
	Email    string `gorm:"column:email" json:"email" valid:"required~Email is Required, email~Invalid format Email"`
	Password string `gorm:"column:password" json:"password" valid:"required~Password is Required, minstringlength(8)~Password minimum 8 characters"`
	Avatar   string `gorm:"column:avatar" json:"avatar"`
	// CreatedDate *time.Time `gorm:"column:created_date" json:"created_date"`
	// UpdatedDate *time.Time `gorm:"column:updated_date" json:"updated_date"`
}

func (u *User) BeforeCreate() (res bool, err error) {
	res, err = govalidator.ValidateStruct(u)
	if err != nil {
		return
	}

	currentTime := time.Now()
	u.CreatedDate = &currentTime
	u.Password = helpers.HashPassword(u.Password)
	// u.Token, err = helpers.GenerateJWT(u.Username, u.Email)

	return
}
