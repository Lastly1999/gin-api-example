package service

import (
	"fmt"
	"github.com/gin-api-example/pkg"
	"testing"
)

func TestUserHashPassword(t *testing.T) {
	password, _ := pkg.HashPassword("123456")
	fmt.Print(password)
}
