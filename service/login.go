package service

import (
	"errors"
	"github.com/gin-api-example/api/response"
	models "github.com/gin-api-example/models/jwt"
	"github.com/gin-api-example/pkg"
	"strconv"
)

type LoginService struct{}

// ValidLogin 验证登录
func (s *LoginService) ValidLogin(username string, password string) (*response.LoginRes, error) {
	userService := UserService{}
	user, err := userService.SelectUserByUsername(username)
	if err != nil {
		return nil, errors.New("用户不存在！")
	}
	if !pkg.ComparePasswords(user.Password, password) {
		return nil, errors.New("账号或密码错误，请重新输入！")
	}
	jwtService := JwtService{}
	token, err := jwtService.GenerateToken(&models.JwtUserInfo{
		UserId:   strconv.FormatInt(user.ID, 10),
		UserName: user.Username,
	})
	if err != nil {
		return nil, err
	}
	return &response.LoginRes{
		Token: token,
	}, nil
}
