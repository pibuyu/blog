package svc

import (
	"blog/rpc/internal/config"
	"blog/rpc/internal/models"
	"github.com/hibiken/asynq"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-queue/kq"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config

	DB  *gorm.DB
	RDB *redis.Client

	KqPusherClient *kq.Pusher

	AsynqClient *asynq.Client
	AsynqServer *asynq.Server
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		DB:             models.InitDB(c.Mysql.Datasource),
		RDB:            models.InitRedisConnection(c),
		KqPusherClient: kq.NewPusher(c.KqPusherConf.Brokers, c.KqPusherConf.Topic),
		AsynqClient:    newAsynqClient(),
		AsynqServer:    newAsynqServer(c),
	}
}
