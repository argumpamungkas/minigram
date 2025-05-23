package models

import "time"

type GormModel struct {
	Id          uint       `gorm:"primaryKey;column:id" json:"id"`
	CreatedDate *time.Time `gorm:"column:created_date" json:"created_date"`
	UpdatedDate *time.Time `gorm:"column:updated_date" json:"updated_date"`
}
