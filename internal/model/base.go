package model

import "time"

// Base 基础Model
type Base struct {
	ID        int       `json:"id" form:"id"`               // 自增id
	CreatedAt time.Time `json:"createdAt" form:"createdAt"` // 创建时间
	UpdatedAt time.Time `json:"updatedAt" form:"updatedAt"` // 更新时间
}
