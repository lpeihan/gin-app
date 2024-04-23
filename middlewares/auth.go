package middlewares

import (
	"gin-app/controllers"
	"gin-app/models"
	"gin-app/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")

		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			controllers.ResponseUnauthorized(ctx)
			ctx.Abort()
			return
		}

		tokenString = tokenString[7:]

		token, claims, err := utils.ParseToken(tokenString)

		if err != nil || !token.Valid {
			controllers.ResponseUnauthorized(ctx)
			ctx.Abort()
			return
		}

		userId := claims.UserId

		var user models.User
		utils.DB.First(&user, userId)

		if user.ID == 0 {
			controllers.ResponseUnauthorized(ctx)
			ctx.Abort()
			return
		}

		ctx.Set("user", user)
		ctx.Next()
	}
}
