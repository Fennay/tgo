package user

import (
	"github.com/fennay/tgo/internal/store/user"
	"github.com/fennay/tgo/internal/utils/response"
	"github.com/fennay/tgo/internal/vo/resp"
	"github.com/gin-gonic/gin"
)

func PageList(c *gin.Context) {
	page := c.GetInt("page")
	pageSize := c.GetInt("pageSize")
	list, pagination := user.PageList(page, pageSize)

	data := &resp.PageListResp{
		Pagination: pagination,
		List:       list,
		Query: map[string]interface{}{
			"page":     page,
			"pageSize": pageSize,
		},
	}
	response.Ok(c, "", data)
	return
}

func Save(c *gin.Context) {
	response.Ok(c, "", nil)
	return
}

func Delete(c *gin.Context) {

	response.Fail(c, 400, "", nil)
	return
}

func Detail(c *gin.Context) {
	response.Ok(c, "", nil)
	return
}
