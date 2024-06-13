package models

import (
	"gorm.io/gorm"
)

type UserBasic struct {
	*gorm.Model
	Identity string
	Name     string
	Password string
	Email    string

	RepositoryInfo UserRepository `json:"repository_info" gorm:"foreignKey:Identity"`
}

func (table UserBasic) TableName() string {
	return "user_basic"
}
