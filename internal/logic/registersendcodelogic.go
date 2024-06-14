package logic

import (
	"blog/rpc/internal/helper"
	"blog/rpc/internal/models"
	"blog/rpc/internal/svc"
	"blog/rpc/types/user"
	"context"
	"errors"
	"fmt"

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
	//先查询该邮箱是否注册过
	var count int64
	err := l.svcCtx.DB.Debug().Where("email=?", in.Email).Model(&models.UserBasic{}).Count(&count).Error
	if err != nil {
		return nil, fmt.Errorf("查询是否已注册出现错误:%w", err)
	}
	if count > 0 {
		return nil, errors.New("该邮箱已被注册！")
	}

	//生成6位随机验证码，并发送到该email
	code := helper.GenVerifyCode()
	err = helper.SendVerifyCode(in.Email, code)
	if err != nil {
		return nil, fmt.Errorf("发送邮件错误：%w", err)
	}
	return &user.RegisterSendCodeResponse{
		Success: true,
		Message: "发送验证码成功",
	}, nil
}
