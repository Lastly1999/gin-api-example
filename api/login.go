package api

import (
	api "github.com/gin-api-example/api/common"
	"github.com/gin-api-example/api/request"
	"github.com/gin-api-example/service"
	"github.com/gin-gonic/gin"
)

type LoginApi struct {
	api.BaseApi
}

func (api *LoginApi) Login(c *gin.Context) {
	var req request.LoginReq
	api.ValidParam(c, &req)
	s := service.LoginService{}
	token, err := s.ValidLogin(req.Username, req.Password)
	if err != nil {
		api.Error(c, err.Error())
	}
	api.Success(c, token)
}
