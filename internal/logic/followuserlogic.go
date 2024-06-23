package logic

import (
	"blog/rpc/internal/models"
	"context"
	"errors"

	"blog/rpc/internal/svc"
	"blog/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFollowUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowUserLogic {
	return &FollowUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FollowUserLogic) FollowUser(in *user.FollowUserRequest) (*user.FollowUserResponse, error) {
	// todo: add your logic here and delete this line
	//1.先去看一下这个identity是否已经关注过对方：未关注：新增关注；已关注：取消关注
	var count int64
	err := l.svcCtx.DB.Debug().Model(&models.Follow{}).Where("to_user_identity = ? and identity = ?", in.ToUserIdentity, in.Identity).Count(&count).Error
	if err != nil {
		return nil, errors.New("查询是否已关注出错：" + err.Error())
	}
	follow := &models.Follow{
		ToUserIdentity: in.ToUserIdentity,
		Identity:       in.Identity,
	}
	if count == 0 {
		//未关注，新增关注项
		err := l.svcCtx.DB.Create(follow).Error
		if err != nil {
			return nil, errors.New("关注出错:" + err.Error())
		}
		return &user.FollowUserResponse{
			Success: true,
			Msg:     "关注用户成功",
		}, nil
	} else {
		//已关注，取消关注，即删除关注项
		err := l.svcCtx.DB.Where("to_user_identity = ? and identity = ?", in.ToUserIdentity, in.Identity).Delete(follow).Error
		if err != nil {
			if err != nil {
				return nil, errors.New("取消关注出错:" + err.Error())
			}
		}
		return &user.FollowUserResponse{
			Success: true,
			Msg:     "取消关注成功",
		}, nil
	}
}
