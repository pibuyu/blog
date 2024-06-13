package logic

import (
	"blog/rpc/internal/models"
	"context"
	"errors"
	"fmt"

	"blog/rpc/internal/svc"
	"blog/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *user.GetUserInfoRequest) (*user.GetUserInfoResponse, error) {

	//这里应该是去查数据库，这里先不用数据库，写死个人信息直接返回
	//return &user.GetUserInfoResponse{
	//	Username: "hhf",
	//	Password: "123456",
	//}, nil
	userBasic := new(models.UserBasic)
	has, err := l.svcCtx.Engine.Where("id = ?", in.UserId).Get(userBasic)
	if err != nil {
		return nil, fmt.Errorf("查询过程出错：%w", err)
	}
	if !has {
		return nil, errors.New("查询的用户不存在")
	}
	return &user.GetUserInfoResponse{
		Username: userBasic.Name,
		Password: userBasic.Password,
	}, nil
}
