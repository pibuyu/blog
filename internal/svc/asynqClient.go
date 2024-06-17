package svc

import (
	"blog/rpc/internal/config"
	"github.com/hibiken/asynq"
)

// create asynq client.
func newAsynqClient(c config.Config) *asynq.Client {
	//return asynq.NewClient(asynq.RedisClientOpt{Addr: "127.0.0.1:6379"})
	return asynq.NewClient(asynq.RedisClientOpt{Addr: c.RedisConf.Host})
}
