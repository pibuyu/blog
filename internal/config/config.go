package config

import (
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf

	Mysql struct {
		Datasource string
	}

	//Redis struct {
	//	Host string
	//}
	// Redis redis.RedisConf `yaml:"redis"`
}
