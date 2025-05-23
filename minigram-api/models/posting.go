package models

import "time"

type Posting struct {
	Id           int        `json:"id"`
	UserId       int        `json:"user_id"`
	Caption      string     `json:"caption"`
	CountLike    int        `json:"count_like"`
	CountComment int        `json:"count_comement"`
	CreatedDate  *time.Time `json:"created_date"`
	UpdatedDate  *time.Time `json:"updated_date"`
}
