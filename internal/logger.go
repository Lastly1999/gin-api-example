package internal

import (
	"fmt"
	"github.com/gin-api-example/global"
	"github.com/gin-api-example/pkg"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

func InitZap() {
	cfg := global.GLOBAL_CONFIG.Zap //获取上述的配置文件

	// 判断是否有Director文件夹
	if ok, _ := pkg.PathExists(cfg.Director); !ok {
		fmt.Printf("create %v directory\n", cfg.Director)
		_ = os.Mkdir(cfg.Director, os.ModePerm)
	}

	// zap.LevelEnablerFunc(func(lev zapcore.Level) bool 用来划分不同级别的输出
	// 根据不同的级别输出到不同的日志文件

	// 调试级别
	debugPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.DebugLevel
	})
	// 日志级别
	infoPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.InfoLevel
	})
	// 警告级别
	warnPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.WarnLevel
	})
	// 错误级别
	errorPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev >= zap.ErrorLevel
	})

	cores := [...]zapcore.Core{
		getEncoderCore(fmt.Sprintf("./%s/server_debug.log", cfg.Director), debugPriority),
		getEncoderCore(fmt.Sprintf("./%s/server_info.log", cfg.Director), infoPriority),
		getEncoderCore(fmt.Sprintf("./%s/server_warn.log", cfg.Director), warnPriority),
		getEncoderCore(fmt.Sprintf("./%s/server_error.log", cfg.Director), errorPriority),
	}

	//zapcore.NewTee(cores ...zapcore.Core) zapcore.Core
	//NewTee创建一个Core，将日志条目复制到两个或更多的底层Core中

	logger := zap.New(zapcore.NewTee(cores[:]...))

	//用文件名、行号和zap调用者的函数名注释每条消息
	if cfg.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}

	//把初始化好的logger赋值到全局日志变量中
	global.GLOBAL_LOGGER = logger
}

// getEncoderCore 获取Encoder的zapcore.Core
func getEncoderCore(fileName string, level zapcore.LevelEnabler) (core zapcore.Core) {
	writer := pkg.GetWriteSyncer(fileName) // 使用file-rotatelogs进行日志分割
	return zapcore.NewCore(getEncoder(), writer, level)
}

// getEncoder 获取zapcore.Encoder
func getEncoder() zapcore.Encoder {
	//获取配置文件的输出格式,json or console
	cfg := global.GLOBAL_CONFIG.Zap
	if cfg.Format == "json" {
		return zapcore.NewJSONEncoder(getEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(getEncoderConfig())
}

// getEncoderConfig 获取zapcore.EncoderConfig
func getEncoderConfig() (config zapcore.EncoderConfig) {
	cfg := global.GLOBAL_CONFIG.Zap
	config = zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  cfg.StacktraceKey,              //栈名
		LineEnding:     zapcore.DefaultLineEnding,      //默认的结尾\n
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  //小写字母输出
		EncodeTime:     CustomTimeEncoder,              //自定义时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //编码间隔 s
		EncodeCaller:   cEncodeCaller,                  // 编码调用者
	}
	//根据配置文件重新配置编码颜色和字体
	switch {
	case cfg.EncodeLevel == "LowercaseLevelEncoder": // 小写编码器(默认)
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	case cfg.EncodeLevel == "LowercaseColorLevelEncoder": // 小写编码器带颜色
		config.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	case cfg.EncodeLevel == "CapitalLevelEncoder": // 大写编码器
		config.EncodeLevel = zapcore.CapitalLevelEncoder
	case cfg.EncodeLevel == "CapitalColorLevelEncoder": // 大写编码器带颜色
		config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	default:
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	}
	return config
}

// cEncodeCaller 自定义行号显示
func cEncodeCaller(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + caller.TrimmedPath() + "]")
}

// CustomTimeEncoder 自定义日志输出时间格式
// 输出格式为 prfix + 具体的时间日期
func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	cfg := global.GLOBAL_CONFIG.Zap
	enc.AppendString(t.Format(cfg.Prefix + " 2006-01-02 15:04:05"))
}
