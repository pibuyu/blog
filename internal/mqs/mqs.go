package mqs

import (
	"blog/rpc/internal/config"
	"blog/rpc/internal/svc"
	"context"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/service"
)

func Consumers(c config.Config, ctx context.Context, svcContext *svc.ServiceContext) []service.Service {
	return []service.Service{
		//监听消费流水状态变更
		kq.MustNewQueue(c.KqConsumerConf, NewPaymentSuccess(ctx, svcContext)),
	}

}
