package internal

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/gin-api-example/global"
	"go.uber.org/zap"
)

/**
 * 使用AK&SK初始化账号Client
 * @param accessKeyId
 * @param accessKeySecret
 * @return Client
 * @throws Exception
 */
func InitAliyunSmsClient() {
	aliyunSmsConf := global.GLOBAL_CONFIG.Aliyun.Sms
	config := &openapi.Config{
		// 必填，您的 AccessKey ID
		AccessKeyId: &aliyunSmsConf.AccessKeyId,
		// 必填，您的 AccessKey Secret
		AccessKeySecret: &aliyunSmsConf.AccessKeySecret,
	}
	config.Endpoint = tea.String(aliyunSmsConf.Endpoint)
	global.GLOBAL_ALIYUN_SMS_CLIENT = &dysmsapi20170525.Client{}
	client, err := dysmsapi20170525.NewClient(config)
	if err != nil {
		global.GLOBAL_LOGGER.Error("Failed to initialize Aliyun SMS client", zap.Any("error", err))
		return
	}
	global.GLOBAL_ALIYUN_SMS_CLIENT = client
	global.GLOBAL_LOGGER.Info("Aliyun SMS client initialized successfully. AccessKeyId: " + aliyunSmsConf.AccessKeyId + ", Endpoint: " + aliyunSmsConf.Endpoint)

}
