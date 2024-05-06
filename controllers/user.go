package controllers

import (
	"gin-app/common/code"
	"gin-app/common/global"
	"gin-app/common/response"
	"gin-app/dto"
	"gin-app/models"
	"gin-app/utils"
	"time"

	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	json := dto.RegisterReq{}

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
	contextUser, _ := ctx.Get("user")

	user := contextUser.(models.User)
	var postCount int64
	global.DB.Model(models.Post{}).Where("user_id = ?", user.ID).Count(&postCount)

	response.ReturnOK(ctx, &dto.UserInfoRes{
		ID:         user.ID,
		Name:       user.Name,
		Phone:      user.Phone,
		Email:      user.Email,
		Avatar:     user.Avatar,
		LoginTime:  user.LoginTime,
		CreateTime: user.CreateTime,
		UpdateTime: user.UpdateTime,
		PostCount:  postCount,
	})
}
