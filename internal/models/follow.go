package models

import "gorm.io/gorm"

type Follow struct {
	*gorm.Model
	ToUserIdentity string `json:"to_user_identity"`
	Identity       string `json:"identity"`
}

func (table Follow) TableName() string {
	return "follow"
}
