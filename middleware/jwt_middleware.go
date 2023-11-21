package middleware

import (
	"strings"

	"github.com/gin-api-example/api/response"
	"github.com/gin-api-example/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader("Authorization")
		if authorizationHeader == "" {
			c.JSON(http.StatusOK, response.JsonResponseData{
				Code:    http.StatusUnauthorized,
				Message: "Authorization参数,请求头为空",
				Data:    nil,
			})
			c.Abort()
			return
		}

		// 检查 Authorization 头部是否以 "Bearer " 开头
		if !strings.HasPrefix(authorizationHeader, "Bearer ") {
			c.JSON(http.StatusOK, response.JsonResponseData{
				Code:    http.StatusUnauthorized,
				Message: "Authorization头部格式错误，应以Bearer开头",
				Data:    nil,
			})
			c.Abort()
			return
		}

		// 提取 token
		tokenString := strings.TrimPrefix(authorizationHeader, "Bearer ")

		jwtService := service.JwtService{}
		claims, err := jwtService.ParseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusOK, response.JsonResponseData{
				Code:    http.StatusForbidden,
				Message: "认证失败，请重新登录！",
				Data:    nil,
			})
			c.Abort()
			return
		}

		c.Set("user", claims)
		c.Next()
	}
}
