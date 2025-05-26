package models

type Comment struct {
	GormModel
	UserId  int    `json:"user_id"`
	PhotoId int    `json:"photo_id"`
	Comment string `json:"comment"`
	// CreatedDate *time.Time `json:"created_date"`
	// UpdatedDate *time.Time `json:"updated_date"`
}
