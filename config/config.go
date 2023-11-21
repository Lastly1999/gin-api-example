package config

type Mysql struct {
	Host     string `json:"host" yaml:"host"`
	Port     string `json:"port" yaml:"port"`
	User     string `json:"user" yaml:"user"`
	DB       string `json:"db" yaml:"db"`
	Password string `json:"password" yaml:"password"`
}

type Redis struct {
	Host     string `json:"host" yaml:"host"`
	Port     string `json:"port" yaml:"port"`
	Password string `json:"password" yaml:"password"`
	Db       int    `json:"db" yaml:"db"`
}

type Application struct {
	Name      string `json:"name" yaml:"name"`
	Host      string `json:"host" yaml:"host"`
	Port      string `json:"port" yaml:"port"`
	Version   string `json:"version" yaml:"version"`
	ApiPrefix string `json:"apiPrefix" yaml:"apiPrefix"`
	RunMode   string `json:"runMode" yaml:"runMode"`
}

type ZapConfig struct {
	Level         string `mapstructure:"level" json:"level" yaml:"level"`                           // 级别
	Format        string `mapstructure:"format" json:"format" yaml:"format"`                        // 输出
	Prefix        string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`                        // 日志前缀
	Director      string `mapstructure:"director" json:"director"  yaml:"director"`                 // 日志文件夹
	ShowLine      bool   `mapstructure:"show-line" json:"showLine" yaml:"showLine"`                 // 显示行
	EncodeLevel   string `mapstructure:"encode-level" json:"encodeLevel" yaml:"encode-level"`       // 编码级
	StacktraceKey string `mapstructure:"stacktrace-key" json:"stacktraceKey" yaml:"stacktrace-key"` // 栈名
	LogInConsole  bool   `mapstructure:"log-in-console" json:"logInConsole" yaml:"log-in-console"`  // 输出控制台
}

type Miniprogram struct {
	AppId     string `json:"appid" yaml:"appid"`         // 小程序、公众号或者企业微信的appid
	AppSecret string `json:"appsecret" yaml:"appsecret"` // 商户号 appID
	HttpDebug bool   `json:"httpDebug" yaml:"httpDebug"` // 是否开启http请求日志
	Debug     bool   `json:"debug" yaml:"debug"`         // 是否开启debug模式
}

type PowerWechat struct {
	Miniprogram Miniprogram `json:"miniprogram" yaml:"miniprogram"`
}

type AliyunSms struct {
	AccessKeyId     string `json:"accessKeyId" yaml:"accessKeyId"`
	AccessKeySecret string `json:"accessKeySecret" yaml:"accessKeySecret"`
	SignName        string `json:"signName" yaml:"signName"`
	TemplateCode    string `json:"templateCode" yaml:"templateCode"`
	Endpoint        string `json:"endpoint" yaml:"endpoint"`
}

type AliyunOss struct {
	AccessKeyId     string `json:"accessKeyId" yaml:"accessKeyId"`
	AccessKeySecret string `json:"accessKeySecret" yaml:"accessKeySecret"`
	Endpoint        string `json:"endpoint" yaml:"endpoint"`
	Bucket          string `json:"bucket" yaml:"bucket"`
}

type Aliyun struct {
	Sms AliyunSms `json:"sms" yaml:"sms"`
	Oss AliyunOss `json:"oss" yaml:"oss"`
}

type Jwt struct {
	ExpireTime int64  `json:"expireTime" yaml:"expireTime"` // 过期时间
	Secret     string `json:"secret" yaml:"secret"`         // 密钥
}

type WxJwt struct {
	ExpireTime int64  `json:"expireTime" yaml:"expireTime"` // 过期时间
	Secret     string `json:"secret" yaml:"secret"`         // 密钥
}

type Config struct {
	Application Application `json:"application" yaml:"application"`
	Jwt         Jwt         `json:"jwt" yaml:"jwt"`
	WxJwt       WxJwt       `json:"wxJwt" yaml:"wxJwt"`
	Mysql       Mysql       `json:"mysql" yaml:"mysql"`
	Redis       Redis       `json:"redis" yaml:"redis"`
	Zap         ZapConfig   `json:"zap" yaml:"zap"`
	PowerWechat PowerWechat `json:"powerWechat" yaml:"powerWechat"`
	Aliyun      Aliyun      `json:"aliyun" yaml:"aliyun"`
}

var ConfigFileMapping = map[string]string{
	"dev":  "development.yaml",
	"test": "testing.yaml",
	"prod": "production.yaml",
}
