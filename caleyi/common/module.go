package common

import (
	"fmt"
	"github.com/qinuoyun/shopping/caleyi/utils/ci"
	"gopkg.in/ini.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

func InitModule() {
	// 读取.ini 里面的数据库配置
	config, err := ini.Load("./config.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	ip := config.Section("mysql").Key("ip").String()
	port := config.Section("mysql").Key("port").String()
	user := config.Section("mysql").Key("user").String()
	password := config.Section("mysql").Key("password").String()
	database := config.Section("mysql").Key("database").String()

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
