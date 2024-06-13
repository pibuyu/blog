package logic

import (
	"blog/rpc/internal/models"
	"context"

	"blog/rpc/internal/svc"
	"blog/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserRepositoryByUserIdentityLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserRepositoryByUserIdentityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserRepositoryByUserIdentityLogic {
	return &GetUserRepositoryByUserIdentityLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserRepositoryByUserIdentityLogic) GetUserRepositoryByUserIdentity(in *user.GetUserRepositoryByUserIdentityRequest) (*user.GetUserRepositoryByUserIdentityResponse, error) {
	// todo: add your logic here and delete this line
	ups := make([]models.UserRepository, 0)
	err := l.svcCtx.DB.Debug().Where("user_identity = ?", in.UserIdentity).Find(&ups).Error
	if err != nil {
		return nil, err
	}

	//将[]models.UserRepository转换为[]user.Repositories
	repositories := make([]*user.UserRepository, len(ups))
	for i, up := range ups {
		repositories[i] = &user.UserRepository{
			Identity:           up.Identity,
			UserIdentity:       up.UserIdentity,
			ParentId:           up.ParentId,
			RepositoryIdentity: up.RepositoryIdentity,
			Ext:                up.Ext,
			Name:               up.Name,
		}
	}
	return &user.GetUserRepositoryByUserIdentityResponse{
		Repositories: repositories,
	}, nil
}
