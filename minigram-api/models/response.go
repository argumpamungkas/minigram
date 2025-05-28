package models

type ReponseInfo struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type ReponseLogin struct {
	ReponseInfo
	Data ResponseUser `json:"data"`
}

type ResponseUser struct {
	Username string  `json:"username"`
	FullName string  `json:"full_name"`
	Email    string  `json:"email"`
	Avatar   *string `json:"avatar"` // menggunakan arterisk karena dapat bernilai null
	Token    string  `json:"token"`
}

type ReponsePostingAll struct {
	ReponseInfo
	Data []ResponsePosting `json:"data"`
}

type ReponsePostingById struct {
	ReponseInfo
	Data ResponsePosting `json:"data"`
}

type ResponsePosting struct {
	UserId   uint    `json:"user_id"`
	Username string  `json:"username"`
	Avatar   *string `json:"avatar"`
	Caption  *string `json:"caption"`
	Photo    string  `json:"photo"`
	Likes    int     `json:"likes"`
	Comments int     `json:"comments"`
}
