package service

import (
	"context"
	"github.com/gin-api-example/api/response"
	"github.com/gin-api-example/global"
	models "github.com/gin-api-example/models/jwt"
)

type WxService struct{}

// Wxlogin 微信登录换取session
func (s *WxService) Wxlogin(ctx context.Context, code string) (*response.WxloginRes, error) {
	r, err := global.GLOBAL_POWER_WECHAT_MINI_PROGRAM.Auth.Session(ctx, code)
	if err != nil {
		return nil, err
	}
	// 查询用户信息，如果没有则创建
	wxUserService := WxUserService{}
	wxUser, err := wxUserService.SelectMiniAppWxUserInfoOrCreate(r.OpenID, r.UnionID)
	if err != nil {
		return nil, err
	}
	wxJwtService := WxJwtService{}
	token, err := wxJwtService.GenerateToken(&models.WxJwtUserInfo{
		OpenId:  wxUser.OpenId,
		Unionid: wxUser.Unionid,
	})
	if err != nil {
		return nil, err
	}
	return &response.WxloginRes{
		Token:  token,
		OpenId: wxUser.OpenId,
	}, nil
}
