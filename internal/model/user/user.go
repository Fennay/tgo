package user

import (
	commonModel "github.com/fennay/tgo/internal/model"
)

// User model struct
type User struct {
	commonModel.Base
	Username string `json:"username" form:"username"`
	Nickname string `json:"nickname" form:"nickname"`
	Password string `json:"password" form:"password"`
	Phone    string `json:"phone" form:"phone"`
	Email    string `json:"email" form:"email"`
	Sex      uint8  `json:"sex" form:"sex"`
	Status   uint8  `json:"status" form:"status"`
	Sort     int    `json:"sort"`
}

// TableName 设置表名称
func (User) TableName() string {
	return "user"
}
