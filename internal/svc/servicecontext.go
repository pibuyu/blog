package svc

import (
	"blog/rpc/internal/config"
	"blog/rpc/internal/models"
	"xorm.io/xorm"
)

type ServiceContext struct {
	Config config.Config

	Engine *xorm.Engine
	//RDB    *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Engine: models.InitDB(c.Mysql.Datasource),
		//RDB:    models.InitRedisConnection(c),
	}
}
