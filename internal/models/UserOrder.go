package models

import "gorm.io/gorm"

type UserOder struct {
	*gorm.Model
	Identity     string  `json:"identity"`
	UserIdentity string  `json:"user_identity"`
	RoomIdentity string  `json:"room_identity"`
	Price        float32 `json:"price"`
}

func (table UserOder) TableName() string {
	return "user_order"
}
