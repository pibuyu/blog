package mqs

import (
	"blog/rpc/internal/config"
	"blog/rpc/internal/svc"
	"context"
	"github.com/zeromicro/go-zero/core/service"
)

func AllConsumers(c config.Config) []service.Service {
	svcCtx := svc.NewServiceContext(c)
	ctx := context.Background()
	var services []service.Service
	services = append(services, Consumers(c, ctx, svcCtx)...)
	return services
}
