package controllers

import (
	"gin-app/common/code"
	"gin-app/common/response"
	"gin-app/dto"
	"gin-app/models"
	"gin-app/utils"
	"gin-app/vo"
	"time"

	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	json := vo.RegisterRequestJson{}

	if err := ctx.ShouldBindJSON(&json); err != nil {
		response.ReturnError(ctx, code.CommonError, err.Error())
		return
	}

	user := models.User{
		Name:       json.Name,
		Password:   json.Password,
		Phone:      json.Phone,
		Email:      json.Email,
		Avatar:     json.Avatar,
		CreateTime: time.Now().Format(time.DateTime),
		UpdateTime: time.Now().Format(time.DateTime),
		LoginTime:  time.Now().Format(time.DateTime),
	}

	models.Register(&user)
	token, err := utils.GenerateToken(user.ID)

	if err != nil {
		response.ReturnError(ctx, code.CommonError, code.GetCodeMessage(code.CommonError))
		return
	}

	response.ReturnOK(ctx, gin.H{"token": token})
}

func GetUserInfo(ctx *gin.Context) {
	user, _ := ctx.Get("user")

	response.ReturnOK(ctx, dto.ToUserInfoDto(user.(models.User)))
}
