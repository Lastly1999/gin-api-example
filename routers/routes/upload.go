package routes

import (
	"github.com/gin-api-example/api"
	"github.com/gin-gonic/gin"
)

func InitUploadRouter(router *gin.RouterGroup) {
	var uploadApi api.UploadApi
	uploadRouter := router.Group("/upload")
	uploadRouter.POST("/files", uploadApi.Upload)
}
