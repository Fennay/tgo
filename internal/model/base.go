package model

import "time"

// Base 基础Model
type Base struct {
	ID        int       `json:"id" form:"id"`               // 自增id
	CreatedAt time.Time `json:"createdAt" form:"createdAt"` // 创建时间
	UpdatedAt time.Time `json:"updatedAt" form:"updatedAt"` // 更新时间
}

// OPEN_STATUS 开启状态
const OPEN_STATUS = 1
const CLOSE_STATUS = 0

// DEAULT_SORT 默认排序
const DEAULT_SORT = 999999

type Status struct {
	Status int `json:"status"` //
	Sort   int `json:"sort"`   //
}
