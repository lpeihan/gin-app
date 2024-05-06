package main

import (
	"gin-app/router"
	"gin-app/utils"
)

func main() {
	utils.InitConfig()
	utils.InitMysql()
	utils.InitRedis()

	r := router.Router()

	r.Run("0.0.0.0:7002")
}
