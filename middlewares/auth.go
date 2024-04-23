package middlewares

import (
	"gin-app/common/global"
	"gin-app/common/response"
	"gin-app/models"
	"gin-app/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")

		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			response.ReturnUnauthorized(ctx)
			ctx.Abort()
			return
		}

		tokenString = tokenString[7:]

		token, claims, err := utils.ParseToken(tokenString)

		if err != nil || !token.Valid {
			response.ReturnUnauthorized(ctx)
			ctx.Abort()
			return
		}

		userId := claims.UserId

		var user models.User
		global.DB.First(&user, userId)

		if user.ID == 0 {
			response.ReturnUnauthorized(ctx)
			ctx.Abort()
			return
		}

		ctx.Set("user", user)
		ctx.Next()
	}
}
