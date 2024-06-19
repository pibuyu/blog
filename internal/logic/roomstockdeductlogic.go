package logic

import (
	"blog/rpc/internal/models"
	"context"
	"errors"
	"gorm.io/gorm"
	"strconv"

	"blog/rpc/internal/svc"
	"blog/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoomStockDeductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoomStockDeductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoomStockDeductLogic {
	return &RoomStockDeductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RoomStockDeductLogic) RoomStockDeduct(in *user.RoomStockDeductRequest) (*user.RoomStockDeductResponse, error) {
	//要先判断库存够不够减的
	room := &models.Room{}
	err := l.svcCtx.DB.Where("identity=?", in.RoomIdentity).First(room).Error
	if err != nil {
		return nil, errors.New("查询房间信息出错")
	}
	if room.Stock <= in.DeductNumber {
		return nil, errors.New("房间余量不足,仅剩" + strconv.Itoa(int(room.Stock)) + "间")
	}
	err = l.svcCtx.DB.Model(&models.Room{}).Where("identity=?", in.RoomIdentity).Update("stock", gorm.Expr("stock - ?", in.DeductNumber)).Error
	if err != nil {
		return nil, errors.New("库存减少出错：" + err.Error())
	}
	return &user.RoomStockDeductResponse{
		Success: true,
		Msg:     "库存减少成功，room_identity:" + in.RoomIdentity + ",减少量为:" + string(in.DeductNumber),
	}, nil
}
