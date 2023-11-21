package internal

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram"
	"github.com/gin-api-example/global"
	"strconv"
)

func InitPowerWechat() {
	pwrWechatConf := global.GLOBAL_CONFIG.PowerWechat
	// 1. 初始化小程序应用实例
	app, err := miniProgram.NewMiniProgram(&miniProgram.UserConfig{
		AppID:     pwrWechatConf.Miniprogram.AppId,     // 小程序、公众号或者企业微信的appid
		Secret:    pwrWechatConf.Miniprogram.AppSecret, // 商户号 appID
		HttpDebug: pwrWechatConf.Miniprogram.HttpDebug,
		Debug:     pwrWechatConf.Miniprogram.Debug,
	})
	if err != nil {
		global.GLOBAL_LOGGER.Error("Failed to initialize Power Wechat error:" + err.Error())
		return
	}
	global.GLOBAL_POWER_WECHAT_MINI_PROGRAM = app
	global.GLOBAL_LOGGER.Info("Power Wechat Mini Program initialized successfully. AppID: " +
		pwrWechatConf.Miniprogram.AppId +
		", HttpDebug: " + strconv.FormatBool(pwrWechatConf.Miniprogram.HttpDebug) +
		", Debug: " + strconv.FormatBool(pwrWechatConf.Miniprogram.Debug))
}
