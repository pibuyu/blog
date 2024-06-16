package svc

import (
	"github.com/hibiken/asynq"
)

// create asynq client.
func newAsynqClient() *asynq.Client {
	return asynq.NewClient(asynq.RedisClientOpt{Addr: "127.0.0.1:6379"})
}
