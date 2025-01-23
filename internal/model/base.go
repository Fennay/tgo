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

const MALE_GENDER = 1   // 男
const FEMALE_GENDER = 2 // 女
const OTHER_GENDER = 0  // 其他

// DEFAULT_SORT 默认排序
const DEFAULT_SORT = 999999

type ExtendField struct {
	Status uint8 `json:"status"` //
	Sort   int   `json:"sort"`   //
}
