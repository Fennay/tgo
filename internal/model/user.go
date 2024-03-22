package model

// User model struct
type User struct {
	Base
	Username string `json:"username" form:"username"`
	Nickname string `json:"nickname" form:"nickname"`
	Password string `json:"password" form:"password"`
	Phone    string `json:"phone" form:"phone"`
	Email    string `json:"email" form:"email"`
	Sex      uint8  `json:"sex" form:"sex"`
	ExtendField
}

// TableName 设置表名称
func (User) TableName() string {
	return "user"
}

// AdminUserSession 管理员用户 session 结构
type AdminUserSession struct {
	UserID      int    `json:"userId"`   //
	Username    string `json:"username"` //
	Nickname    string `json:"nickname"` //
	Phone       string `json:"phone"`    //
	Email       string `json:"email"`    //
	Sex         uint8  `json:"sex"`      //
	ExtendField        //
}
