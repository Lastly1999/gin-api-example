package api

import (
	api "github.com/gin-api-example/api/common"
	"github.com/gin-api-example/global"
	"github.com/gin-gonic/gin"
)

type UserApi struct {
	api.BaseApi
}

func (api *UserApi) Test(ctx *gin.Context) {
	global.GLOBAL_LOGGER.Info("asdkaskdaskdk")
	ctx.JSON(200, gin.H{
		"test": 1,
	})
}
