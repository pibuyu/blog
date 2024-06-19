package svc

import (
	"blog/rpc/internal/config"
	"blog/rpc/internal/models"
	"database/sql"
	"github.com/hibiken/asynq"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-queue/kq"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config

	DB  *gorm.DB
	RDB *redis.Client

	KqPusherClient *kq.Pusher

	AsynqClient *asynq.Client
	AsynqServer *asynq.Server

	DtmBarrierDB *sql.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	//gorm.DB转化为sql.DB
	dtmBarrierGormDB, _ := gorm.Open(mysql.Open(c.DtmBarrierDB.Datasource), &gorm.Config{})
	dtmBarrierSqlDB, _ := dtmBarrierGormDB.DB()
	return &ServiceContext{
		Config:         c,
		DB:             models.InitDB(c.Mysql.Datasource),
		RDB:            models.InitRedisConnection(c),
		KqPusherClient: kq.NewPusher(c.KqPusherConf.Brokers, c.KqPusherConf.Topic),
		AsynqClient:    newAsynqClient(c),
		AsynqServer:    newAsynqServer(c),
		DtmBarrierDB:   dtmBarrierSqlDB,
	}
}
