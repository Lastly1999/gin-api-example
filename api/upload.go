package api

import (
	api "github.com/gin-api-example/api/common"
	"github.com/gin-api-example/service"
	"github.com/gin-gonic/gin"
)

type UploadApi struct {
	api.BaseApi
}

func (api *UploadApi) Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		api.Error(c, err.Error())
	}
	uploadService := service.UploadService{}
	r, err := uploadService.UploadFile(file)
	if err != nil {
		api.Error(c, err.Error())
	}
	api.Success(c, r)
}
