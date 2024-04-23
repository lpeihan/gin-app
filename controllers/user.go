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

type RegisterRequestJson struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
	Phone    string `json:"phone"`
	Email    string `json:"email" `
	Avatar   string `json:"avatar"`
}

func Register(ctx *gin.Context) {
	json := RegisterRequestJson{}

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
		response.ReturnError(ctx, code.OK, "系统异常")
		return
	}

	response.ReturnOK(ctx, gin.H{"token": token})
}

func GetUserInfo(ctx *gin.Context) {
	user, _ := ctx.Get("user")

	response.ReturnOK(ctx, dto.ToUserInfoDto(user.(models.User)))
}
