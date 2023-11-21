package routers

import (
	"github.com/gin-api-example/api/response"
	"github.com/gin-api-example/global"
	"github.com/gin-api-example/middleware"
	"github.com/gin-api-example/routers/routes"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitAppRouter() *gin.Engine {
	gin.SetMode(global.GLOBAL_CONFIG.Application.RunMode)
	app := gin.Default()

	// ---------------------不受jwt认证影响的路由---------------------
	publicRoutes := app.Group(global.GLOBAL_CONFIG.Application.ApiPrefix)
	// 登录路由
	routes.InitLoginRouter(publicRoutes)
	// 小程序登录路由
	routes.InitWxRouter(publicRoutes)
	// ---------------------------结束----------------------------

	// ---------------------受jwt认证影响的路由---------------------
	protectedRoutes := app.Group(global.GLOBAL_CONFIG.Application.ApiPrefix)
	protectedRoutes.Use(middleware.JWTMiddleware())
	// 上传路由
	routes.InitUploadRouter(protectedRoutes)
	// ---------------------------结束----------------------------
	app.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, response.JsonResponseData{
			Code:    http.StatusNotFound,
			Message: c.Request.RequestURI,
			Data:    nil,
		})
	})
	return app
}
