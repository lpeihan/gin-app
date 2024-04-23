package controllers

import (
	"gin-app/models"
	"gin-app/utils"
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

	models.Register(&user)
	token, err := utils.GenerateToken(user.ID)

	if err != nil {
		ResponseError(ctx, CodeError, "系统异常")
		return
	}

	ResponseSuccess(ctx, gin.H{"token": token})
}

func GetUserInfo(ctx *gin.Context) {
	user, _ := ctx.Get("user")

	ResponseSuccess(ctx, user)
}
