package pkg

import (
	"github.com/gin-api-example/global"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

func GetWriteSyncer(file string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   file, // 日志文件的位置
		MaxSize:    10,   // 在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxBackups: 200,  // 保留旧文件的最大个数
		MaxAge:     1,    // 保留旧文件的最大天数
		Compress:   true, // 是否压缩/归档旧文件
	}

	if global.GLOBAL_CONFIG.Zap.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberJackLogger))
	}
	return zapcore.AddSync(lumberJackLogger)
}
