package main

import (
	"blog/rpc/internal/config"
	"blog/rpc/internal/server"
	"blog/rpc/internal/svc"
	"blog/rpc/types/user"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		user.RegisterUserServiceServer(grpcServer, server.NewUserServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	//这些事kq的消费者的启动配置
	//svcCtx := svc.NewServiceContext(c)
	//backgroundCtx := context.Background()
	//serviceGroup := service.NewServiceGroup()
	//defer serviceGroup.Stop()
	//for _, mq := range mqs.AllConsumers(c) {
	//	serviceGroup.Add(mq)
	//}
	//serviceGroup.Start()

	//这个是asynq 延时队列消费者的注册和启动
	//ctxBackground := context.Background()
	//svcContext := svc.NewServiceContext(c)
	//cronJob := logic.NewCronJob(ctxBackground, svcContext)
	//mux := cronJob.Register()
	//if err := svcContext.AsynqServer.Run(mux); err != nil {
	//	logx.WithContext(ctxBackground).Errorf("!!!CronJobErr!!! run err:%+v", err)
	//	os.Exit(1)
	//}

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
