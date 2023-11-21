package service

import (
	"errors"
	"github.com/gin-api-example/global"
	"github.com/gin-api-example/models"
)

type UserService struct{}

// SelectUserByUsername 根据用户名查询用户
func (s *UserService) SelectUserByUsername(username string) (*models.User, error) {
	var user models.User
	user.Username = username
	if err := global.GLOBAL_ORM.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// CreateUser 创建用户
func (s *UserService) CreateUser(user *models.User) error {
	if user == nil {
		return errors.New("user is nil")
	}
	if err := global.GLOBAL_ORM.Create(&user).Error; err != nil {
		return err
	}
	return nil
}
