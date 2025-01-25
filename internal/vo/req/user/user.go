package user

type CreateReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Sex      int    `json:"sex"` //
}

type LoginReq struct {
	Username string `json:"username"` //
	Password string `json:"password"`
}
