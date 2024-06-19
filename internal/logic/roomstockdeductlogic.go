package logic

import (
	"blog/rpc/internal/models"
	"blog/rpc/internal/svc"
	"blog/rpc/types/user"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/dtm-labs/client/dtmgrpc"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"strconv"
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
	//room := &models.Room{}
	//err := l.svcCtx.DB.Where("identity=?", in.RoomIdentity).First(room).Error
	//if err != nil {
	//	return nil, errors.New("查询房间信息出错")
	//}
	//if room.Stock <= in.DeductNumber {
	//	return nil, errors.New("房间余量不足,仅剩" + strconv.Itoa(int(room.Stock)) + "间")
	//}
	//err = l.svcCtx.DB.Debug().Model(&models.Room{}).Where("identity=?", in.RoomIdentity).Update("stock", gorm.Expr("stock - ?", in.DeductNumber)).Error
	//if err != nil {
	//	return nil, errors.New("库存减少出错：" + err.Error())
	//}

	//把上面操作数据库的代码改造成dtm的代码
	fmt.Println("dtmgrpc减少库存开始了")
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)

	//datasource := "root:123456@tcp(127.0.0.1:3306)/dtm_barrier?charset=utf8mb4&parseTime=True&loc=Local"
	//
	//db, _ := gorm.Open(mysql.Open(datasource), &gorm.Config{})
	//sqlDB, _ := db.DB()
	if err = barrier.CallWithDB(l.svcCtx.DtmBarrierDB, func(tx *sql.Tx) error {

		//在这里执行业务逻辑:查询库存是否足够；减去对应数量的库存
		//1.查询库存是否足够
		room := &models.Room{}
		err := l.svcCtx.DB.Debug().Where("identity=?", in.RoomIdentity).First(room).Error
		if err != nil {
			return errors.New("查询房间信息出错")
		}
		if room.Stock <= in.DeductNumber {
			return errors.New("房间余量不足,仅剩" + strconv.Itoa(int(room.Stock)) + "间")
		}
		err = l.svcCtx.DB.Debug().Model(&models.Room{}).Where("identity=?", in.RoomIdentity).Update("stock", gorm.Expr("stock - ?", in.DeductNumber)).Error
		if err != nil {
			return errors.New("库存减少出错：" + err.Error())
		}
		return nil
	}); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &user.RoomStockDeductResponse{
		Success: true,
		Msg:     "库存减少成功，room_identity:" + in.RoomIdentity + ",减少量为:" + string(in.DeductNumber),
	}, nil
}
