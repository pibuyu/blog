package logic

import (
	"blog/rpc/internal/define"
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

	//判断用户输入的验证码是否正确
	inputCode := in.Code
	//向redis存储验证码时，key为define.REDIS_VERI_CODE_PRE+in.Email；但是微服务调用时，
	actualCode, err := l.svcCtx.RDB.Get(l.ctx, define.REDIS_VERI_CODE_PRE+in.Email).Result()
	if err != nil {
		return nil, fmt.Errorf("从redis取验证码出错：%w", err)
	}
	fmt.Printf("用户输入的验证码为：%s,redis中的验证码为%s", inputCode, actualCode)
	if inputCode != actualCode {
		return nil, errors.New("输入的验证码错误")
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
