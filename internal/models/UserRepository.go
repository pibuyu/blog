package models

import "gorm.io/gorm"

type UserRepository struct {
	*gorm.Model
	Identity           string
	UserIdentity       string `json:"user_identity"`
	ParentId           int64
	RepositoryIdentity string
	Ext                string
	Name               string
}

func (table UserRepository) TableName() string {
	return "user_repository"
}
