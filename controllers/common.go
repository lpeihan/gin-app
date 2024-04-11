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
	CodeOK            Code = 200
	CodeNotAuthorized Code = 401
)

func ReturnSuccess(ctx *gin.Context, data interface{}) {
	json := &ResponseJson{Code: CodeOK, Message: "success", Data: data}

	ctx.JSON(http.StatusOK, json)
}

func ReturnError(ctx *gin.Context, code int, message string) {
	json := &ResponseJson{Code: CodeOK, Message: message}

	ctx.JSON(http.StatusOK, json)
}
