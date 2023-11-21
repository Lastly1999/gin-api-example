package routes

import (
	"github.com/gin-api-example/api"
	"github.com/gin-gonic/gin"
)

func InitWxRouter(router *gin.RouterGroup) {
	var wxApi api.WxApi
	wxRouter := router.Group("/wx")
	wxRouter.POST("/login", wxApi.Sign)
}
