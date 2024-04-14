package controllers

import (
	"gin-app/models"
	"time"

	"github.com/gin-gonic/gin"
)

func GetUserList(c *gin.Context) {
	ReturnSuccess(c, gin.H{
		"list": models.GetUserList(),
	})
}

func CreateUser(c *gin.Context) {
	user := models.User{}

	if err := c.ShouldBindJSON(&user); err != nil {
		ReturnError(c, CodeError, err.Error())
		return
	}

	user.CreateTime = time.Now().Format(time.DateTime)
	user.UpdateTime = time.Now().Format(time.DateTime)
	user.LoginTime = time.Now().Format(time.DateTime)

	models.CreateUser(user)

	ReturnSuccess(c, user)
}
