package api

import (
	"github.com/gin-api-example/api/response"
	"github.com/gin-api-example/global"
	"github.com/gin-api-example/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseApi struct{}

var (
	SUCCESS_CODE = 1
	FAIL_CODE    = -1
)

func (receiver *BaseApi) Error(c *gin.Context, message string) {
	if c.IsAborted() {
		return // 阻止执行后续处理
	}
	c.JSON(http.StatusOK, response.JsonResponseData{
		Code:    FAIL_CODE,
		Message: message,
		Data:    nil,
	})
	global.GLOBAL_LOGGER.Error(message)
	c.Abort()
}

func (receiver *BaseApi) Success(c *gin.Context, data interface{}) {
	if c.IsAborted() {
		return // 阻止执行后续处理
	}
	c.JSON(http.StatusOK, response.JsonResponseData{
		Code:    SUCCESS_CODE,
		Message: "ok",
		Data:    data,
	})
}

func (receiver *BaseApi) ValidParam(c *gin.Context, u interface{}) {
	if err := c.ShouldBind(&u); err != nil {
		receiver.Error(c, pkg.GetValidMsg(err, u))
		c.Abort() // 停止请求的继续处理
	}
}
