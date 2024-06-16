package mqs

import (
	"blog/rpc/internal/logic"
	"blog/rpc/internal/models"
	"blog/rpc/internal/svc"
	"blog/rpc/types/user"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
)

type PaymentSuccess struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPaymentSuccess(ctx context.Context, svcCtx *svc.ServiceContext) *PaymentSuccess {
	return &PaymentSuccess{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

//比方说， 我们把这个kq改造为用户注册的消息队列

// Consume 实现接口Consume，就可以接受消息并处理逻辑了
func (l *PaymentSuccess) Consume(key, val string) error {
	logx.Infof("consumer receive message,PaymentSuccess key :%s , val :%s", key, val)
	return nil
}

// 接受用户注册的基本信息，在这里完成注册,就是调用rpc服务，然后api里就用这个queue的服务
func (l *PaymentSuccess) execService(message models.UserRegisterMessage) (bool, error) {
	registerLogic := logic.NewUserRegisterLogic(l.ctx, l.svcCtx)
	//构造注册用户的请求
	registerReq := &user.UserRegisterRequest{
		Username: message.Name,
		Password: message.Password,
		Email:    message.Email,
		Code:     message.Code,
	}
	response, err := registerLogic.UserRegister(registerReq)
	if err != nil {
		return false, err
	}
	if !response.GetResult() {
		return false, fmt.Errorf("注册用户出错：%s", response.GetMsg())
	}
	return true, nil
}
