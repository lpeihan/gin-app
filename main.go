package main

import (
	"gin-app/models"
	"gin-app/router"
	"gin-app/utils"
)

func main() {
	utils.InitMysql()

	// AutoMigrate
	utils.DB.AutoMigrate(&models.User{})

	r := router.Router()

	r.Run("0.0.0.0:7002")
}
