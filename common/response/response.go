package response

import (
	"gin-app/common/code"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseJson struct {
	Code    code.Code   `json:"code"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data"`
}

func ReturnOK(ctx *gin.Context, data interface{}) {
	json := &ResponseJson{Code: code.OK, Message: code.GetCodeMessage(code.OK), Data: data}

	ctx.JSON(http.StatusOK, json)
}

func ReturnList(ctx *gin.Context, list interface{}, total int64) {
	json := &ResponseJson{Code: code.OK, Message: code.GetCodeMessage(code.OK), Data: gin.H{
		"list":  list,
		"total": total,
	}}

	ctx.JSON(http.StatusOK, json)
}

func ReturnError(ctx *gin.Context, code code.Code, message string) {
	json := &ResponseJson{Code: code, Message: message, Data: gin.H{}}

	ctx.JSON(http.StatusOK, json)
}

func ReturnUnauthorized(ctx *gin.Context) {
	json := &ResponseJson{Code: code.Unauthorized, Message: code.GetCodeMessage(code.Unauthorized), Data: gin.H{}}

	ctx.JSON(http.StatusOK, json)
}
