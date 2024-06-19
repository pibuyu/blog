package logic

import (
	"blog/rpc/internal/helper"
	"blog/rpc/internal/models"
	"blog/rpc/internal/svc"
	"blog/rpc/types/user"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/dtm-labs/client/dtmgrpc"
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
	//userOrder := &models.UserOder{
	//	Identity:     helper.GetUUID(),
	//	UserIdentity: in.UserIdentity,
	//	RoomIdentity: in.RoomIdentity,
	//	Price:        in.Price,
	//}
	//err := l.svcCtx.DB.Debug().Create(userOrder).Error
	//if err != nil {
	//	return nil, errors.New("创建订单失败")
	//}

	//把上面的代码修改为dtm负责的代码
	fmt.Println("新增订单开始了")
	barrier, _ := dtmgrpc.BarrierFromGrpc(l.ctx)

	//datasource := "root:123456@tcp(127.0.0.1:3306)/dtm_barrier?charset=utf8mb4&parseTime=True&loc=Local"
	//
	//db, _ := gorm.Open(mysql.Open(datasource), &gorm.Config{})
	//sqlDB, _ := db.DB()
	userOrder := &models.UserOder{
		Identity:     helper.GetUUID(),
		UserIdentity: in.UserIdentity,
		RoomIdentity: in.RoomIdentity,
		Price:        in.Price,
	}
	barrier.CallWithDB(l.svcCtx.DtmBarrierDB, func(tx *sql.Tx) error {

		err := l.svcCtx.DB.Debug().Create(userOrder).Error
		if err != nil {
			return errors.New("创建订单失败")
		}
		return nil
	})
	return &user.CreateOrderResponse{
		Success:  true,
		Identity: userOrder.Identity,
		Msg:      "创建订单成功",
	}, nil
}
