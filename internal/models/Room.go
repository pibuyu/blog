package models

import "gorm.io/gorm"

type Room struct {
	*gorm.Model
	Stock int32   `json:"stock"`
	Price float64 `json:"price"`
	Desc  string  `json:"desc"`
}

func (table Room) TableName() string {
	return "room"
}
