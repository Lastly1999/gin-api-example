package global

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-api-example/config"
	redis "github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var GLOBAL_REDIS_CLIENT *redis.Client

var GLOBAL_CONFIG *config.Config

var GLOBAL_LOGGER *zap.Logger

var GLOBAL_ORM *gorm.DB

var GLOBAL_POWER_WECHAT_MINI_PROGRAM *miniProgram.MiniProgram

var GLOBAL_ALIYUN_SMS_CLIENT *dysmsapi20170525.Client

var GLOBAL_ALIYUN_OSS_CLIENT *oss.Client
