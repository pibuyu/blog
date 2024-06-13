package models

import (
	"blog/rpc/internal/config"
	"blog/rpc/internal/define"
	"fmt"
	_ "github.com/go-sql-driver/mysql"

	"github.com/redis/go-redis/v9"
	"xorm.io/xorm"
)

func InitDB(datasource string) *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", datasource)
	if err != nil {
		fmt.Print("连接数据库失败" + err.Error())
		return nil
	}
	engine.ShowSQL(true)
	///*
	//	//当然也可以表明的映射规则和字段的映射规则不同
	//	engine.SetTableMapper(core.SnakeMapper{})
	//	engine.SetColumnMapper(core.SameMapper{})
	//*/
	////也支持在setMapper的同时在表名/字段名前面加上前缀
	//engine.SetTableMapper(core.NewPrefixMapper(core.SnakeMapper{}, "xorm_"))
	//engine.Sync(new(User))
	return engine
}

// InitRedisConnection 返回redis连接
func InitRedisConnection(c config.Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     c.Redis.Host,
		Password: define.REDIS_CONN_PWD, // no password set
		DB:       0,                     // use default DB
	})
}
