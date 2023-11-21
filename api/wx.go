package api

import (
	api "github.com/gin-api-example/api/common"
	"github.com/gin-api-example/api/request"
	"github.com/gin-api-example/service"
	"github.com/gin-gonic/gin"
)

type WxApi struct {
	api.BaseApi
}

// Sign 微信登录换取session
func (api *WxApi) Sign(ctx *gin.Context) {
	var req request.WxSignReq
	api.ValidParam(ctx, &req)
	wxService := service.WxService{}
	wxlogin, err := wxService.Wxlogin(ctx.Request.Context(), req.Code)
	if err != nil {
		api.Error(ctx, err.Error())
	}
	api.Success(ctx, wxlogin)
}
