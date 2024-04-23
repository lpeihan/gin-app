package main

import (
	"gin-app/router"
	"gin-app/utils"
)

func main() {
	utils.InitMysql()

	r := router.Router()

	r.Run("0.0.0.0:7002")
}
