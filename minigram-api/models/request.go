package models

import (
	"time"

	"github.com/asaskevich/govalidator"
)

type RequestLogin struct {
	Username string `json:"username" valid:"required~Username is Required"`
	Password string `json:"password" valid:"required~Password is Required"`
}

type RequestPosting struct {
	GormModel
	UserId  uint   `json:"user_id"`
	Caption string `json:"caption"`
	Photo   string `json:"photo" form:"photo" valid:"required~Photo is required"`
}

func (p *RequestPosting) BeforeCreatePosting() (res bool, err error) {
	res, err = govalidator.ValidateStruct(p)
	if err != nil {
		return
	}

	currentTime := time.Now()
	p.CreatedDate = &currentTime

	return
}
