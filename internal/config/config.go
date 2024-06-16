package config

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf

	Mysql struct {
		Datasource string
	}

	RedisConf struct {
		Host string
	}

	KqPusherConf struct {
		Brokers []string
		Topic   string
	}

	KqConsumerConf kq.KqConf
}
