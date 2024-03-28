package user

import (
	"github.com/fennay/tgo/internal/utils/response"
	"github.com/gin-gonic/gin"
)

func PageList(c *gin.Context) {

	response.Ok(c, "", nil)
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
