package logic

import (
	"blog/rpc/internal/models"
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"

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
	userBasic := models.UserBasic{}
	result := l.svcCtx.DB.Debug().Where("id = ?", in.UserId).First(&userBasic)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, errors.New("未查询到该用户")
		}
		return nil, fmt.Errorf("查询过程出错：%w", result.Error)

	}

	return &user.GetUserInfoResponse{
		Username: userBasic.Name,
		Password: userBasic.Password,
	}, nil
}
