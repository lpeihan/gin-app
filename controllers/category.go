package controllers

import (
	"gin-app/common/code"
	"gin-app/common/response"
	"gin-app/dto"
	"gin-app/models"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateCategory(ctx *gin.Context) {
	req := dto.CreateCategoryReq{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ReturnError(ctx, code.CommonError, err.Error())
		return
	}

	if models.GetCategoryByName(req.Name).ID != 0 {
		response.ReturnError(ctx, code.CommonError, "分类已存在")
		return
	}

	category := models.Category{
		Name:       req.Name,
		CreateTime: time.Now().Format(time.DateTime),
		UpdateTime: time.Now().Format(time.DateTime),
	}

	models.CreateCategory(&category)

	response.ReturnOK(ctx, category)
}
