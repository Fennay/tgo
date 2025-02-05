package user

type CreateReq struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
	Nickname string `form:"nickname" json:"nickname"`
	Phone    string `form:"phone" json:"phone"`
	Email    string `form:"email" json:"email"`
	Sex      int    `form:"sex" json:"sex"` //
}

type LoginReq struct {
	Username string `form:"username" json:"username" `
	Password string `form:"password" json:"password" binding:"required"`
}
