package routes

import (
	"github.com/gin-api-example/api"
	"github.com/gin-gonic/gin"
)

func InitLoginRouter(router *gin.RouterGroup) {
	var loginApi api.LoginApi
	router.POST("/login", loginApi.Login)
}
