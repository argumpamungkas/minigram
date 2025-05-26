package models

import (
	"time"

	"github.com/asaskevich/govalidator"
)

type Posting struct {
	GormModel
	UserId  uint   `json:"user_id"`
	Caption string `json:"caption" form:"caption"`
	Photo   string `json:"photo" form:"photo" valid:"required~Photo is required"`
	// CreatedDate  *time.Time `json:"created_date"`
	// UpdatedDate  *time.Time `json:"updated_date"`
}

func (p *Posting) BeforeCreate() (res bool, err error) {
	res, err = govalidator.ValidateStruct(p)
	if err != nil {
		return
	}

	currentTime := time.Now()
	p.CreatedDate = &currentTime

	return
}
