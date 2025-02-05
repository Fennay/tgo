package user

import (
	"github.com/fennay/tgo/internal/model"
	"github.com/fennay/tgo/internal/store/user"
	"github.com/fennay/tgo/internal/utils/ecode"
	"github.com/fennay/tgo/internal/utils/response"
	userVo "github.com/fennay/tgo/internal/vo/req/user"
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
	response.Ok(c, data)
	return
}

func Save(c *gin.Context) {
	userReqParam := userVo.CreateReq{}
	err := c.BindJSON(userReqParam)
	if err != nil {
		response.Fail(c, ecode.BadRequest, nil)
		return
	}
	user.Save(&model.User{
		Username: userReqParam.Username,
		Nickname: "",
		Password: userReqParam.Password,
		Phone:    "",
		Email:    "",
		Sex:      model.OTHER_GENDER,
	})
	response.Ok(c, nil)
	return
}

func Delete(c *gin.Context) {
	id := c.GetInt("id")
	user.Delete(&model.User{
		Base: model.Base{ID: id},
	})
	response.Fail(c, ecode.BadRequest, nil)
	return
}

func Detail(c *gin.Context) {
	id := c.GetInt("id")
	user.Detail(id)
	response.Ok(c, nil)
	return
}
