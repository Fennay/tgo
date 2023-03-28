package req

// IDReq id参数
type IDReq struct {
	ID int `json:"id" form:"id" binding:"required"`
}

// PageListReq PageListReq
type PageListReq struct {
	PageSize int `json:"pageSize" form:"pageSize"` // 每页数量
	Page     int `json:"page" form:"page"`         // 页码
}
