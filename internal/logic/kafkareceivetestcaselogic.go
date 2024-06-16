package logic

import (
	"blog/rpc/internal/svc"
	"blog/rpc/types/user"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type KafkaReceiveTestcaseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewKafkaReceiveTestcaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *KafkaReceiveTestcaseLogic {
	return &KafkaReceiveTestcaseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *KafkaReceiveTestcaseLogic) KafkaReceiveTestcase(in *user.Empty) (*user.KafkaReceiveTestcaseResponse, error) {
	// todo: add your logic here and delete this line

	return &user.KafkaReceiveTestcaseResponse{}, nil
}
