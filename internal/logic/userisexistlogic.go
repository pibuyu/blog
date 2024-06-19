package logic

import (
	"blog/rpc/internal/models"
	"context"

	"blog/rpc/internal/svc"
	"blog/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserIsExistLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserIsExistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserIsExistLogic {
	return &UserIsExistLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserIsExistLogic) UserIsExist(in *user.UserIsExistRequest) (*user.UserIsExistResponse, error) {
	var count int64
	err := l.svcCtx.DB.Debug().Model(&models.UserBasic{}).Where("identity=?", in.UserIdentity).Count(&count).Error
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return &user.UserIsExistResponse{
			Exist: true,
		}, nil
	}
	return &user.UserIsExistResponse{
		Exist: false,
	}, nil
}
