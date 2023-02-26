package model

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func InitMysql(datasource string) *gorm.DB {
	newLogger := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
		SlowThreshold: time.Second,
		LogLevel:      logger.Warn,
		Colorful:      true,
	})
	DB, err := gorm.Open(mysql.Open(datasource), &gorm.Config{Logger: newLogger})
	if err != nil {
		logrus.Println("init mysql failed... ...", err)
		panic("数据库连接失败")
	}
	logrus.Println("init mysql success... ...", viper.Get("mysql"))
	return DB
}

func InitRedis(addr, password string, db int) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	ping, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Println("redis init failed......err:", err)
		return nil
	}
	log.Println("redis init success......", ping)
	return client
}
