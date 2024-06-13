package models

import (
	"blog/rpc/internal/config"
	"blog/rpc/internal/define"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/redis/go-redis/v9"
)

func InitDB(datasource string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(datasource), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

// InitRedisConnection 返回redis连接
func InitRedisConnection(c config.Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     c.Redis.Host,
		Password: define.REDIS_CONN_PWD, // no password set
		DB:       0,                     // use default DB
	})
}
