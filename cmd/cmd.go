package cmd

import (
	"github.com/gin-api-example/global"
	"github.com/gin-api-example/routers"
)

func Start() {
	if err := routers.InitAppRouter().Run(":" + global.GLOBAL_CONFIG.Application.Port); err != nil {
		global.GLOBAL_LOGGER.Error("app run error")
	}
	return
}
