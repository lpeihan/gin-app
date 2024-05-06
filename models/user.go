package models

import (
	"gin-app/common/global"

	"gorm.io/gorm"
)

type User struct {
	ID         uint   `json:"id" gorm:"primary_key"`
	Name       string `json:"name"`
	Password   string `json:"password"`
	Phone      string `json:"phone"`
	Email      string `json:"email" `
	Avatar     string `json:"avatar"`
	Salt       string `json:"salt"`
	LoginTime  string `json:"loginTime"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
}

func (table *User) TableName() string {
	return "user"
}

func GetUserById(id uint) User {
	user := User{}

	global.DB.Where("ID = ?", id).First(&user)
	return user
}

func GetUserByName(name string) User {
	user := User{}

	global.DB.Where("name = ?", name).First(&user)
	return user
}

func Register(user *User) *gorm.DB {
	return global.DB.Create(user)
}
