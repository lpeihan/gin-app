package router

import (
	"gin-app/controllers"
	"gin-app/middlewares"

	"gin-app/utils"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.Use(middlewares.Cors())
	r.Use(gin.LoggerWithConfig(utils.LoggerToFile()))
	r.Use(utils.Recover)

	api := r.Group("/api")

	// user
	user := api.Group("/user")
	{
		user.POST("/register", controllers.Register)
		user.GET("/info", middlewares.Auth(), controllers.GetUserInfo)
	}

	return r
}
