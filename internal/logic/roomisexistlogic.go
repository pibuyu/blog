package logic

import (
	"blog/rpc/internal/models"
	"context"
	"errors"
	"gorm.io/gorm"

	"blog/rpc/internal/svc"
	"blog/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoomIsExistLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoomIsExistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoomIsExistLogic {
	return &RoomIsExistLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RoomIsExistLogic) RoomIsExist(in *user.RoomIsExistRequest) (*user.RoomIsExistResponse, error) {
	room := new(models.Room)
	err := l.svcCtx.DB.Debug().Where("identity=?", in.Identity).First(room).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errors.New("查找房间信息错误:" + err.Error())
	}
	if room == nil || err == gorm.ErrRecordNotFound {
		return &user.RoomIsExistResponse{
			Exist: false,
			Stock: 0,
		}, nil
	}
	return &user.RoomIsExistResponse{
		Exist: true,
		Stock: room.Stock,
		Price: float32(room.Price),
	}, nil
}
