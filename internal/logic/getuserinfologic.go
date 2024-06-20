package logic

import (
	"blog/rpc/internal/models"
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"reflect"

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
	fmt.Printf("传递过来的req参数in为：%s,user_id值为：%s,类型为：%s", in, in.GetIdentity(), reflect.TypeOf(in.Identity))
	userBasic := models.UserBasic{}
	result := l.svcCtx.DB.Debug().Where("identity = ?", in.Identity).First(&userBasic)
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
