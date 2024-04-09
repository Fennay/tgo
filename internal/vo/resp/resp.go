package resp

type IDResp struct {
	ID int `json:"id" form:"id"` // ID
}

// PageListResp 分页返回内容
type PageListResp struct {
	Pagination Pagination    `json:"pagination"` // 分页数据信息
	List       []interface{} `json:"list"`       // 数据内容
	Query      []interface{} `json:"query"`      // 查询参数
}

type Pagination struct {
	Page     int   `json:"page"`     // 当前页码
	PageSize int   `json:"pageSize"` // 每页数量
	Total    int64 `json:"total"`    // 总数量
}
