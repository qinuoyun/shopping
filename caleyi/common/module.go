package common

import (
	"fmt"
	"github.com/qinuoyun/shopping/caleyi/utils/ci"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

func InitModule() {
	//读取配置文件
	ip := ci.C("mysql.ip")
	port := ci.C("mysql.port")
	user := ci.C("mysql.user")
	password := ci.C("mysql.password")
	database := ci.C("mysql.database")

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Error,
			Colorful:      true,
		},
	)

	options := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "pre_", // 表前缀
			SingularTable: true,   // 禁用表名复数
		},
		Logger: newLogger,
	}
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", user, password, ip, port, database)
	_DB, err := gorm.Open(mysql.Open(dsn), options)
	if err != nil {
		log.Fatal(err)
	}

	// 打开 DB 的 Debug 日志
	_DB.Debug()

	// 获取所有模块
	moduleMap := ci.GetModules()

	// 循环遍历并打印
	for _, value := range moduleMap {
		if err := _DB.AutoMigrate(value); err != nil {
			log.Fatal(err)
		}
	}

	// 将 _DB 设置到 ci 包中
	ci.SetDB(_DB)
}
