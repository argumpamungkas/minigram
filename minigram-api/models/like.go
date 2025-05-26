package models

type Like struct {
	GormModel
	PhotoId uint `json:"phpto_id"`
	UserId  uint `json:""`
}
