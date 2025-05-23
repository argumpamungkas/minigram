package models

type Like struct {
	GormModel
	PhotoId uint
	UserId  uint
}
