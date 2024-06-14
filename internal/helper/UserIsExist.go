package helper

import (
	"blog/rpc/internal/models"
	"blog/rpc/internal/svc"
	"context"
)

// ToolLogic 临时结构体变量，取数据库连接对象用的
type ToolLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// UserIsExist 查询这个email是否被注册过
// UserIsExist首字母大写，意味着方法是exported公开的，但是所属helperLogic首字母小写，类型是private的，所以需要将但是所属helperLogic首字母大写
func (l *ToolLogic) UserIsExist(email string) (bool, error) {
	ub := models.UserBasic{
		Email: email,
	}
	result := l.svcCtx.DB.Where("user_basic", ub)
	if result.Error != nil {
		return false, result.Error
	}
	if result.RowsAffected == 1 {
		return true, nil
	}
	return false, nil
}
