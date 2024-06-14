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

type UserRegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserRegisterLogic) UserRegister(in *user.UserRegisterRequest) (*user.UserRegisterResponse, error) {

	//先查询该邮箱是否注册过
	var count int64
	err := l.svcCtx.DB.Debug().Where("email=?", in.Email).Model(&models.UserBasic{}).Count(&count).Error
	if err != nil {
		return nil, fmt.Errorf("查询是否已注册出现错误:%w", err)
	}
	if count > 0 {
		return nil, errors.New("该邮箱已被注册！")
	}

	//进行注册
	ub := &models.UserBasic{
		Identity: helper.GetUUID(),
		Name:     in.Username,
		Password: helper.Md5(in.Password),
		Email:    in.Email,
	}
	err = l.svcCtx.DB.Debug().Create(ub).Error
	if err != nil {
		return nil, fmt.Errorf("插入用户出错：%w", err)
	}

	return &user.UserRegisterResponse{
		Result: true,
		Msg:    "用户注册成功",
	}, nil
}
