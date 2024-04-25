package controllers

import (
	"gin-app/common/code"
	"gin-app/common/response"
	"gin-app/models"
	"gin-app/vo"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateCategory(ctx *gin.Context) {
	data := vo.CreateCategoryRequestJson{}

	if err := ctx.ShouldBindJSON(&data); err != nil {
		response.ReturnError(ctx, code.CommonError, err.Error())
		return
	}

	if models.GetCategoryByName(data.Name).ID != 0 {
		response.ReturnError(ctx, code.CommonError, "分类已存在")
		return
	}

	category := models.Category{
		Name:       data.Name,
		CreateTime: time.Now().Format(time.DateTime),
		UpdateTime: time.Now().Format(time.DateTime),
	}

	models.CreateCategory(&category)

	response.ReturnOK(ctx, category)
}
