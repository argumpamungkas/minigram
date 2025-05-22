package models

type ReponseInfo struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type ReponseLogin struct {
	Status  int          `json:"status"`
	Message string       `json:"message"`
	User    ResponseUser `json:"user"`
}

type ResponseUser struct {
	Username string  `json:"username"`
	FullName string  `json:"full_name"`
	Email    string  `json:"email"`
	Avatar   *string `json:"avatar"` // menggunakan arterisk karena dapat bernilai null
	Token    string  `json:"token"`
}
