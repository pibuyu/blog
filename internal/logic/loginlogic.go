package logic

import (
	"blog/rpc/internal/helper"
	"blog/rpc/internal/svc"
	"blog/rpc/types/user"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginRequest) (*user.LoginResponse, error) {
	token, err := helper.GenerateToken(in.Id, in.Identity, in.Name)
	if err != nil {
		return nil, err
	}
	return &user.LoginResponse{
		Code: 200,
		Msg:  "登陆成功",
		Data: token,
	}, nil
}
