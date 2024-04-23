package controllers

import (
	"gin-app/common/code"
	"gin-app/common/response"
	"gin-app/dto"
	"gin-app/models"
	"gin-app/utils"
	"time"

	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	user := models.User{}

	if err := ctx.ShouldBindJSON(&user); err != nil {
		response.ReturnError(ctx, code.CommonError, err.Error())
		return
	}

	user.CreateTime = time.Now().Format(time.DateTime)
	user.UpdateTime = time.Now().Format(time.DateTime)
	user.LoginTime = time.Now().Format(time.DateTime)

	models.Register(&user)
	token, err := utils.GenerateToken(user.ID)

	if err != nil {
		response.ReturnError(ctx, code.OK, "系统异常")
		return
	}

	response.ReturnOK(ctx, gin.H{"token": token})
}

func GetUserInfo(ctx *gin.Context) {
	user, _ := ctx.Get("user")

	response.ReturnOK(ctx, dto.ToUserInfoDto(user.(models.User)))
}
