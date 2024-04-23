package controllers

import (
	"gin-app/models"
	"time"

	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	user := models.User{}

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ResponseError(ctx, CodeError, err.Error())
		return
	}

	user.CreateTime = time.Now().Format(time.DateTime)
	user.UpdateTime = time.Now().Format(time.DateTime)
	user.LoginTime = time.Now().Format(time.DateTime)

	models.Register(user)

	ResponseSuccess(ctx, user)
}
