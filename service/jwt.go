package service

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-api-example/global"
	models "github.com/gin-api-example/models/jwt"
	"time"
)

type JwtService struct{}

type Claims struct {
	models.JwtUserInfo
	jwt.StandardClaims
}

// GenerateToken 生成JWT令牌
func (s *JwtService) GenerateToken(info *models.JwtUserInfo) (string, error) {
	claims := Claims{
		JwtUserInfo: *info,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(global.GLOBAL_CONFIG.Jwt.ExpireTime) * time.Minute).Unix(), // 令牌过期时间
			IssuedAt:  time.Now().Unix(),                                                                       // 令牌颁发时间
			Subject:   "jwt-auth",                                                                              // 令牌主题
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(global.GLOBAL_CONFIG.Jwt.Secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// ParseToken 解析JWT令牌
func (s *JwtService) ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(global.GLOBAL_CONFIG.Jwt.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, jwt.ErrSignatureInvalid
}

// RefreshToken 刷新JWT令牌
func (s *JwtService) RefreshToken(tokenString string) (string, error) {
	claims, err := s.ParseToken(tokenString)
	if err != nil {
		return "", err
	}
	claims.ExpiresAt = time.Now().Add(time.Duration(global.GLOBAL_CONFIG.Jwt.ExpireTime) * time.Minute).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString([]byte(global.GLOBAL_CONFIG.Jwt.Secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
