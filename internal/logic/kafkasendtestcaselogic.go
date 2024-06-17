package logic

import (
	"blog/common"
	"blog/rpc/internal/svc"
	"blog/rpc/types/user"
	"context"
	"encoding/json"
	"github.com/hibiken/asynq"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type KafkaSendTestcaseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewKafkaSendTestcaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *KafkaSendTestcaseLogic {
	return &KafkaSendTestcaseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// KafkaSendTestcase asynq发送延时消息
func (l *KafkaSendTestcaseLogic) KafkaSendTestcase(in *user.KafkaSendTestcaseRequest) (*user.KafkaSendTestcaseResponse, error) {
	rpcReq := &user.UserRegisterRequest{
		Username: in.Username,
		Password: in.Password,
		Email:    in.Email,
		Code:     in.Code,
	}
	payload, _ := json.Marshal(rpcReq)

	//创建异步任务
	//_, err := l.svcCtx.AsynqClient.Enqueue(asynq.NewTask(models.USER_REGISTER_JOB, payload),
	//	asynq.ProcessAt(time.Now().Add(1*time.Second)))
	_, err := l.svcCtx.AsynqClient.Enqueue(asynq.NewTask(common.USER_REGISTER_JOB, payload),
		asynq.ProcessAt(time.Now().Add(1*time.Second)))
	if err != nil {
		logx.WithContext(l.ctx).Errorf("创建异步任务出错:%+v", err)
	}
	return &user.KafkaSendTestcaseResponse{
		Result: true,
		Msg:    "创建异步任务成功",
	}, nil
}
