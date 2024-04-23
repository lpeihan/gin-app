package controllers

import (
	"gin-app/common/code"
	"gin-app/common/response"
	"gin-app/models"

	"github.com/gin-gonic/gin"
)

type CreateCategoryRequestJson struct {
	Name string `json:"name" binding:"required"`
}

func CreateCategory(ctx *gin.Context) {
	json := CreateCategoryRequestJson{}

	if err := ctx.ShouldBindJSON(&json); err != nil {
		response.ReturnError(ctx, code.CommonError, err.Error())
		return
	}

	category := models.Category{
		Name: json.Name,
	}

	models.CreateCategory(&category)

	response.ReturnOK(ctx, category)
}
