package router

import (
	"gin-app/controllers"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")

	user := api.Group("/user")
	{
		user.GET("/info", controllers.GetUserInfo)
	}

	return r
}
