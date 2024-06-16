package logic

import (
	"blog/rpc/internal/svc"
	"context"
	"fmt"
	"github.com/hibiken/asynq"
)

type SettleRecordHandler struct {
	svcCtx *svc.ServiceContext
}

func NewUserRegisterHandler(svcCtx *svc.ServiceContext) *SettleRecordHandler {
	return &SettleRecordHandler{
		svcCtx: svcCtx,
	}
}
func (l *SettleRecordHandler) ProcessTask(ctx context.Context, _ *asynq.Task) error {

	fmt.Printf("这是userRegister延时注册调用的handler \n")

	return nil
}
