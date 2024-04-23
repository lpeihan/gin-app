package utils

import (
	"gin-app/common/global"
	"gin-app/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMysql() {
	dsn := "hello_user:Test!123@tcp(47.120.71.181:3306)/hello?charset=utf8mb4&parseTime=True&loc=Local"

	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	DB.AutoMigrate(&models.User{})

	global.DB = DB

	if err != nil {
		panic(err)
	}
}
