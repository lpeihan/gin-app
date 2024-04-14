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

	user.Name = c.PostForm("name")
	user.Password = c.PostForm("password")
	user.Email = c.PostForm("email")
	user.Phone = c.PostForm("phone")

	user.CreateTime = time.Now().Format(time.DateTime)
	user.UpdateTime = time.Now().Format(time.DateTime)
	user.LoginTime = time.Now().Format(time.DateTime)

	models.CreateUser(user)

	ReturnSuccess(c, user)
}
