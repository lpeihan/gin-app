package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseJson struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data"`
}

const CodeOK = 200

func ReturnSuccess(ctx *gin.Context, data interface{}) {
	json := &ResponseJson{Code: CodeOK, Message: "success", Data: data}

	ctx.JSON(http.StatusOK, json)
}

func ReturnError(ctx *gin.Context, code int, message string) {
	json := &ResponseJson{Code: code, Message: message}

	ctx.JSON(200, json)
}
