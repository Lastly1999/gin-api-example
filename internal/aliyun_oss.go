package internal

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-api-example/global"
)

func InitAliyunoss() {
	ossConf := global.GLOBAL_CONFIG.Aliyun.Oss
	client, err := oss.New(ossConf.Endpoint, ossConf.AccessKeyId, ossConf.AccessKeySecret)
	if err != nil {
		global.GLOBAL_LOGGER.Error("oss client init failed, err:" + err.Error())
		return
	}
	global.GLOBAL_ALIYUN_OSS_CLIENT = client
	global.GLOBAL_LOGGER.Info("oss client init success")
}
