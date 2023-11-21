package internal

import (
	"flag"
	"fmt"
	config "github.com/gin-api-example/config"
	"github.com/gin-api-example/global"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func InitConfig() {

	var conf string
	var confEnv string
	flag.StringVar(&confEnv, "c", "", "choose config file.")
	flag.Parse()
	if confEnv == "" {
		conf = config.ConfigFileMapping["dev"]
	} else {
		conf = config.ConfigFileMapping[confEnv]
	}
	cfg := &config.Config{}
	v := viper.New()
	v.AutomaticEnv()
	fmt.Println("load api service config/" + conf + " mode -> " + confEnv)
	v.SetConfigFile("config/" + conf)
	v.SetConfigType("yaml")

	// set gin mode
	if confEnv == "prod" {
		cfg.Application.RunMode = gin.ReleaseMode
	} else if confEnv == "test" {
		cfg.Application.RunMode = gin.TestMode
	} else if confEnv == "dev" {
		cfg.Application.RunMode = gin.DebugMode
	}

	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error reading config file: %s \n", err))
	}

	if err := v.Unmarshal(&cfg); err != nil {
		panic(fmt.Errorf(err.Error()))
	}

	// Override the global config with the loaded config
	global.GLOBAL_CONFIG = cfg
}
