package models

import (
	"gin-app/common/global"

	"gorm.io/gorm"
)

type User struct {
	ID         int    `json:"id" gorm:"primary_key"`
	Name       string `json:"name"`
	Password   string `json:"password"`
	Phone      string `json:"phone"`
	Email      string `json:"email" `
	Avatar     string `json:"avatar"`
	Salt       string `json:"salt"`
	LoginTime  string `json:"login_time"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}

func (table *User) TableName() string {
	return "user"
}

func Register(user *User) *gorm.DB {
	return global.DB.Create(user)
}
