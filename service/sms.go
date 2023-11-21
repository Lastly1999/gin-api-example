package service

import (
	aliyunSms "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	"github.com/gin-api-example/global"
)

type SmsService struct{}

// SendVerificationCode 发送验证码
func (s *SmsService) SendVerificationCode(phone string) (*aliyunSms.SendSmsResponse, error) {
	request := aliyunSms.SendSmsRequest{
		PhoneNumbers: &phone,
		SignName:     &global.GLOBAL_CONFIG.Aliyun.Sms.SignName,
		TemplateCode: &global.GLOBAL_CONFIG.Aliyun.Sms.TemplateCode,
	}
	sms, err := global.GLOBAL_ALIYUN_SMS_CLIENT.SendSms(&request)
	if err != nil {
		return nil, err
	}
	return sms, nil
}
