package logic

import (
	"blog/rpc/internal/helper"
	"blog/rpc/internal/models"
	"context"
	"errors"

	"blog/rpc/internal/svc"
	"blog/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderLogic {
	return &CreateOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateOrderLogic) CreateOrder(in *user.CreateOrderRequest) (*user.CreateOrderResponse, error) {
	//api里先确认用户和房间信息、房间库存都满足条件，然后开始创建订单
	//创建订单
	userOrder := &models.UserOder{
		Identity:     helper.GetUUID(),
		UserIdentity: in.UserIdentity,
		RoomIdentity: in.RoomIdentity,
		Price:        in.Price,
	}
	err := l.svcCtx.DB.Create(userOrder).Error
	if err != nil {
		return nil, errors.New("创建订单失败")
	}
	return &user.CreateOrderResponse{
		Success:  true,
		Identity: userOrder.Identity,
		Msg:      "创建订单成功",
	}, nil
}
