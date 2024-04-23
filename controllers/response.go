package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Code int
type ResponseJson struct {
	Code    Code        `json:"code"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data"`
}

const (
	CodeOK           Code = 200
	CodeUnauthorized Code = 401
	CodeError        Code = 0
)

func ResponseSuccess(ctx *gin.Context, data interface{}) {
	json := &ResponseJson{Code: CodeOK, Message: "success", Data: data}

	ctx.JSON(http.StatusOK, json)
}

func ResponseError(ctx *gin.Context, code Code, message string) {
	json := &ResponseJson{Code: code, Message: message, Data: gin.H{}}

	ctx.JSON(http.StatusOK, json)
}

func ResponseUnauthorized(ctx *gin.Context) {
	json := &ResponseJson{Code: CodeUnauthorized, Message: "权限不足", Data: gin.H{}}

	ctx.JSON(http.StatusOK, json)
}
