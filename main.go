package main

import (
	"github.com/gin-api-example/cmd"
	"github.com/gin-api-example/internal"
)

func init() {
	internal.InitConfig()
	internal.InitZap()
	internal.InitRedis()
	internal.InitGorm()
	internal.InitPowerWechat()
	internal.InitAliyunSmsClient()
	internal.InitAliyunoss()
}

func main() {
	cmd.Start()
}
