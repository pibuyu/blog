package logic

import (
	"context"
	"errors"

	"blog/rpc/internal/svc"
	"blog/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserRepositoryByIdentityLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteUserRepositoryByIdentityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserRepositoryByIdentityLogic {
	return &DeleteUserRepositoryByIdentityLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteUserRepositoryByIdentityLogic) DeleteUserRepositoryByIdentity(in *user.DeleteUserRepositoryByIdentityRequest) (*user.DeleteUserRepositoryByIdentityResponse, error) {

	result := l.svcCtx.DB.Table("user_repository").Debug().Where("identity = ?", in.Identity).Delete(&user.UserRepository{})
	//根据delete返回的tx对象的RowsAffected字段判断是否找到了该条记录
	if result.RowsAffected == 0 {
		return nil, errors.New("不存在该条记录！")
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &user.DeleteUserRepositoryByIdentityResponse{
		Result: true,
	}, nil
}
