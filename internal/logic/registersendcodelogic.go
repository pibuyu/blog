package logic

import (
	"blog/rpc/internal/define"
	"blog/rpc/internal/helper"
	"blog/rpc/internal/svc"
	"blog/rpc/types/user"
	"context"
	"fmt"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterSendCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterSendCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterSendCodeLogic {
	return &RegisterSendCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterSendCodeLogic) RegisterSendCode(in *user.RegisterSendCodeRequest) (*user.RegisterSendCodeResponse, error) {

	//不应该在这里判断邮箱是否重复，这只负责发送验证码和向redis存储验证码
	//生成6位随机验证码，并发送到该email
	code := helper.GenVerifyCode()
	err := l.svcCtx.RDB.Set(l.ctx, define.REDIS_VERI_CODE_PRE+in.Email, code, 10*time.Minute).Err()
	if err != nil {
		return nil, fmt.Errorf("redis存储验证码出错：%w", err)
	}
	err = helper.SendVerifyCode(in.Email, code)
	if err != nil {
		return nil, fmt.Errorf("发送邮件错误：%w", err)
	}
	return &user.RegisterSendCodeResponse{
		Success: true,
		Message: code,
	}, nil
}
