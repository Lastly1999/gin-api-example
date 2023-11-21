package service

import (
	"github.com/gin-api-example/constant"
	"github.com/gin-api-example/global"
	"github.com/gin-api-example/models"
)

type WxUserService struct{}

// SelectMiniAppWxUserInfoOrCreate 根据openId查询用户，如果没有则创建
func (s *WxUserService) SelectMiniAppWxUserInfoOrCreate(openId string, unionid string) (*models.UserWx, error) {
	var user models.UserWx
	err := global.GLOBAL_ORM.FirstOrCreate(&user, models.UserWx{OpenId: openId, Unionid: unionid, Type: constant.USER_WX_MINI_APP}).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
