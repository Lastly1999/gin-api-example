package internal

import (
	"context"
	"fmt"
	"github.com/gin-api-example/global"
	"github.com/redis/go-redis/v9"
)

func InitRedis() {
	addr := fmt.Sprintf("%s:%s", global.GLOBAL_CONFIG.Redis.Host, global.GLOBAL_CONFIG.Redis.Port)
	password := global.GLOBAL_CONFIG.Redis.Password
	db := global.GLOBAL_CONFIG.Redis.Db
	global.GLOBAL_REDIS_CLIENT = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	ctx := context.Background()
	ping := global.GLOBAL_REDIS_CLIENT.Ping(ctx)
	if ping.Err() != nil {
		global.GLOBAL_LOGGER.Error("Failed to connect to Redis: " + ping.Err().Error())
	} else {
		global.GLOBAL_LOGGER.Info("Connected to Redis successfully. Address: " + addr + ", DB: " + fmt.Sprint(db))
	}
}
