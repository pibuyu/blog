package svc

import (
	"blog/rpc/internal/config"
	"fmt"
	"github.com/hibiken/asynq"
)

func newAsynqServer(c config.Config) *asynq.Server {

	return asynq.NewServer(
		asynq.RedisClientOpt{Addr: "127.0.0.1:6379"},
		asynq.Config{
			IsFailure: func(err error) bool {
				fmt.Printf("asynq server 执行延迟任务出错 err : %+v \n", err)
				return true
			},
			Concurrency: 20, //max concurrent process job task num
		},
	)
}
