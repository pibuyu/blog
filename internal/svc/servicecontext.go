package svc

import (
	"blog/rpc/internal/config"
	"blog/rpc/internal/models"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config

	DB *gorm.DB
	//RDB    *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		DB:     models.InitDB(c.Mysql.Datasource),
		//RDB:    models.InitRedisConnection(c),
	}
}
