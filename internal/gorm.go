package internal

import (
	"fmt"
	"github.com/gin-api-example/global"
	"github.com/gin-api-example/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

func InitGorm() {
	conf := global.GLOBAL_CONFIG.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.User,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.DB,
	)
	// 禁用表名复数
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NowFunc: func() time.Time {
			tmp := time.Now().Local().Format("2006-01-02 15:04:05")
			now, _ := time.ParseInLocation("2006-01-02 15:04:05", tmp, time.Local)
			return now
		},
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})
	if err != nil {
		global.GLOBAL_LOGGER.Error("Failed to connect to MySQL: " + err.Error())
		return
	}
	global.GLOBAL_ORM = db
	global.GLOBAL_ORM.Debug()
	global.GLOBAL_ORM.AutoMigrate(&models.User{}, &models.UserWx{}, &models.UserInfo{})
	global.GLOBAL_LOGGER.Info("Connected to MySQL successfully. DSN: " + dsn)
}
