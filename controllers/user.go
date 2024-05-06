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
	req := dto.RegisterReq{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ReturnError(ctx, code.CommonError, err.Error())
		return
	}

	if models.GetUserByName(req.Name).ID != 0 {
		response.ReturnError(ctx, code.CommonError, "用户已存在")
		return
	}

	user := models.User{
		Name:       req.Name,
		Password:   req.Password,
		Phone:      req.Phone,
		Email:      req.Email,
		Avatar:     req.Avatar,
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

	res := &dto.UserInfoRes{
		ID:         user.ID,
		Name:       user.Name,
		Phone:      user.Phone,
		Email:      user.Email,
		Avatar:     user.Avatar,
		LoginTime:  user.LoginTime,
		CreateTime: user.CreateTime,
		UpdateTime: user.UpdateTime,
		PostCount:  postCount,
	}

	// jsonBytes, err := json.Marshal(res)
	// if err != nil {
	// 	fmt.Println("序列化失败:", err)
	// 	return
	// }

	// err = global.RDB.Set(ctx, fmt.Sprintf("user:%d", user.ID), string(jsonBytes), 0).Err()
	// if err != nil {
	// 	fmt.Println("存储 JSON 字符串到 Redis 出错:", err)
	// 	return
	// }

	// 返回响应
	response.ReturnOK(ctx, res)
}
